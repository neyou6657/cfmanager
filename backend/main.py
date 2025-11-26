from fastapi import FastAPI, HTTPException, Depends
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from typing import List, Optional, Dict, Any
import subprocess
import json
import os
from datetime import datetime

app = FastAPI(
    title="Cloudflare Manager API",
    description="Multi-account Cloudflare management platform",
    version="1.0.0"
)

# CORS设置
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# 配置
CFM_PATH = "/app/cloudflare-manager/cfm"

# Pydantic模型
class Account(BaseModel):
    name: str
    api_token: str
    email: Optional[str] = None

class DNSRecord(BaseModel):
    type: str
    name: str
    content: str
    ttl: int = 1
    proxied: bool = False
    priority: Optional[int] = 10

class WorkerDeploy(BaseModel):
    name: str
    script: str

class WorkerRoute(BaseModel):
    pattern: str
    worker_name: str

class KVNamespace(BaseModel):
    title: str

class KVPair(BaseModel):
    key: str
    value: str

class R2Bucket(BaseModel):
    name: str
    location: str = "auto"

# 辅助函数
def run_cfm_command(args: List[str]) -> Dict[str, Any]:
    """运行CFM命令并返回结果"""
    try:
        result = subprocess.run(
            [CFM_PATH] + args,
            capture_output=True,
            text=True,
            timeout=30
        )
        
        if result.returncode != 0:
            raise HTTPException(
                status_code=400,
                detail=f"Command failed: {result.stderr}"
            )
        
        return {
            "success": True,
            "output": result.stdout,
            "stderr": result.stderr
        }
    except subprocess.TimeoutExpired:
        raise HTTPException(status_code=408, detail="Command timeout")
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

# 根路由
@app.get("/")
async def root():
    return {
        "name": "Cloudflare Manager API",
        "version": "1.0.0",
        "status": "running",
        "timestamp": datetime.now().isoformat()
    }

@app.get("/health")
async def health():
    return {"status": "healthy"}

# 账号管理
@app.get("/api/accounts")
async def list_accounts():
    """列出所有账号"""
    result = run_cfm_command(["account", "list"])
    return result

@app.post("/api/accounts")
async def add_account(account: Account):
    """添加新账号"""
    args = ["account", "add", account.name, "--token", account.api_token]
    if account.email:
        args.extend(["--email", account.email])
    result = run_cfm_command(args)
    return result

@app.post("/api/accounts/{name}/switch")
async def switch_account(name: str):
    """切换账号"""
    result = run_cfm_command(["account", "switch", name])
    return result

@app.delete("/api/accounts/{name}")
async def delete_account(name: str):
    """删除账号"""
    result = run_cfm_command(["account", "remove", name])
    return result

@app.get("/api/accounts/current")
async def get_current_account():
    """获取当前账号信息"""
    result = run_cfm_command(["account", "info"])
    return result

# Zone管理
@app.get("/api/zones")
async def list_zones():
    """列出所有Zone"""
    result = run_cfm_command(["zone", "list"])
    return result

@app.post("/api/zones")
async def create_zone(domain: Dict[str, str]):
    """创建新Zone"""
    result = run_cfm_command(["zone", "create", domain["domain"]])
    return result

@app.get("/api/zones/{zone}")
async def get_zone_info(zone: str):
    """获取Zone信息"""
    result = run_cfm_command(["zone", "info", zone])
    return result

@app.delete("/api/zones/{zone}")
async def delete_zone(zone: str):
    """删除Zone"""
    result = run_cfm_command(["zone", "delete", zone])
    return result

@app.post("/api/zones/{zone}/purge")
async def purge_cache(zone: str):
    """清除缓存"""
    result = run_cfm_command(["zone", "purge", zone, "--everything"])
    return result

# DNS记录管理
@app.get("/api/zones/{zone}/dns")
async def list_dns_records(zone: str, type: Optional[str] = None):
    """列出DNS记录"""
    args = ["dns", "list", zone]
    if type:
        args.extend(["--type", type])
    result = run_cfm_command(args)
    return result

@app.post("/api/zones/{zone}/dns")
async def create_dns_record(zone: str, record: DNSRecord):
    """创建DNS记录"""
    args = ["dns", "create", zone, record.type, record.name, record.content]
    args.extend(["--ttl", str(record.ttl)])
    if record.proxied:
        args.append("--proxied")
    if record.priority and record.type in ["MX", "SRV"]:
        args.extend(["--priority", str(record.priority)])
    
    result = run_cfm_command(args)
    return result

@app.delete("/api/zones/{zone}/dns/{record_id}")
async def delete_dns_record(zone: str, record_id: str):
    """删除DNS记录"""
    result = run_cfm_command(["dns", "delete", zone, record_id])
    return result

@app.get("/api/zones/{zone}/dns/export")
async def export_dns_records(zone: str):
    """导出DNS记录"""
    result = run_cfm_command(["dns", "export", zone])
    return result

# Worker管理
@app.post("/api/workers")
async def deploy_worker(worker: WorkerDeploy):
    """部署Worker"""
    # 保存脚本到临时文件
    script_path = f"/tmp/{worker.name}.js"
    with open(script_path, 'w') as f:
        f.write(worker.script)
    
    result = run_cfm_command(["worker", "deploy", worker.name, script_path])
    
    # 清理临时文件
    os.remove(script_path)
    return result

@app.delete("/api/workers/{name}")
async def delete_worker(name: str):
    """删除Worker"""
    result = run_cfm_command(["worker", "delete", name])
    return result

# Worker路由
@app.get("/api/zones/{zone}/routes")
async def list_worker_routes(zone: str):
    """列出Worker路由"""
    result = run_cfm_command(["worker", "route", "list", zone])
    return result

@app.post("/api/zones/{zone}/routes")
async def create_worker_route(zone: str, route: WorkerRoute):
    """创建Worker路由"""
    result = run_cfm_command([
        "worker", "route", "create",
        zone, route.pattern, route.worker_name
    ])
    return result

@app.delete("/api/zones/{zone}/routes/{route_id}")
async def delete_worker_route(zone: str, route_id: str):
    """删除Worker路由"""
    result = run_cfm_command(["worker", "route", "delete", zone, route_id])
    return result

# Pages管理
@app.get("/api/pages")
async def list_pages():
    """列出所有Pages项目"""
    result = run_cfm_command(["pages", "list"])
    return result

@app.get("/api/pages/{project}")
async def get_pages_info(project: str):
    """获取Pages项目信息"""
    result = run_cfm_command(["pages", "info", project])
    return result

@app.delete("/api/pages/{project}")
async def delete_pages_project(project: str):
    """删除Pages项目"""
    result = run_cfm_command(["pages", "delete", project])
    return result

@app.get("/api/pages/{project}/deployments")
async def list_pages_deployments(project: str):
    """列出Pages部署"""
    result = run_cfm_command(["pages", "deployment", "list", project])
    return result

# KV管理
@app.get("/api/kv/namespaces")
async def list_kv_namespaces():
    """列出KV命名空间"""
    result = run_cfm_command(["kv", "namespace", "list"])
    return result

@app.post("/api/kv/namespaces")
async def create_kv_namespace(namespace: KVNamespace):
    """创建KV命名空间"""
    result = run_cfm_command(["kv", "namespace", "create", namespace.title])
    return result

@app.delete("/api/kv/namespaces/{namespace_id}")
async def delete_kv_namespace(namespace_id: str):
    """删除KV命名空间"""
    result = run_cfm_command(["kv", "namespace", "delete", namespace_id])
    return result

@app.get("/api/kv/namespaces/{namespace_id}/keys")
async def list_kv_keys(namespace_id: str, prefix: Optional[str] = None):
    """列出KV键"""
    args = ["kv", "key", "list", namespace_id]
    if prefix:
        args.extend(["--prefix", prefix])
    result = run_cfm_command(args)
    return result

@app.get("/api/kv/namespaces/{namespace_id}/keys/{key}")
async def get_kv_value(namespace_id: str, key: str):
    """获取KV值"""
    result = run_cfm_command(["kv", "key", "get", namespace_id, key])
    return result

@app.put("/api/kv/namespaces/{namespace_id}/keys")
async def put_kv_value(namespace_id: str, pair: KVPair):
    """设置KV值"""
    result = run_cfm_command([
        "kv", "key", "put",
        namespace_id, pair.key, pair.value
    ])
    return result

@app.delete("/api/kv/namespaces/{namespace_id}/keys/{key}")
async def delete_kv_key(namespace_id: str, key: str):
    """删除KV键"""
    result = run_cfm_command(["kv", "key", "delete", namespace_id, key])
    return result

# R2管理
@app.get("/api/r2/buckets")
async def list_r2_buckets():
    """列出R2桶"""
    result = run_cfm_command(["r2", "list"])
    return result

@app.post("/api/r2/buckets")
async def create_r2_bucket(bucket: R2Bucket):
    """创建R2桶"""
    result = run_cfm_command([
        "r2", "create", bucket.name,
        "--location", bucket.location
    ])
    return result

@app.get("/api/r2/buckets/{bucket_name}")
async def get_r2_bucket_info(bucket_name: str):
    """获取R2桶信息"""
    result = run_cfm_command(["r2", "info", bucket_name])
    return result

@app.delete("/api/r2/buckets/{bucket_name}")
async def delete_r2_bucket(bucket_name: str):
    """删除R2桶"""
    result = run_cfm_command(["r2", "delete", bucket_name])
    return result

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
