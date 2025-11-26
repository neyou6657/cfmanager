from fastapi import FastAPI, HTTPException, UploadFile, File, Form
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from typing import List, Optional, Dict, Any
import requests
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

# Cloudflare API Base URL
CF_API_BASE = "https://api.cloudflare.com/client/v4"

# 从环境变量读取认证信息
EMAIL = os.getenv("CF_EMAIL", "")
API_KEY = os.getenv("CF_API_KEY", "")

# Pydantic模型
class Account(BaseModel):
    email: str
    api_key: str

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

class PagesProject(BaseModel):
    name: str
    production_branch: str = "main"

# 全局认证信息存储
current_auth = {"email": EMAIL, "api_key": API_KEY}

def get_headers() -> Dict[str, str]:
    """获取认证头"""
    if not current_auth.get("email") or not current_auth.get("api_key"):
        raise HTTPException(status_code=401, detail="未设置认证信息，请先设置 Email 和 API Key")
    return {
        "X-Auth-Email": current_auth["email"],
        "X-Auth-Key": current_auth["api_key"],
        "Content-Type": "application/json"
    }

def make_request(method: str, endpoint: str, **kwargs) -> Dict[str, Any]:
    """统一的API请求处理"""
    url = f"{CF_API_BASE}{endpoint}"
    headers = get_headers()
    
    try:
        response = requests.request(method, url, headers=headers, **kwargs)
        data = response.json()
        
        if not data.get("success", False):
            errors = data.get("errors", [{"message": "Unknown error"}])
            raise HTTPException(
                status_code=response.status_code,
                detail=errors[0].get("message", "API request failed")
            )
        
        return data
    except requests.exceptions.RequestException as e:
        raise HTTPException(status_code=500, detail=str(e))

def get_account_id() -> str:
    """获取账户ID"""
    data = make_request("GET", "/accounts")
    if not data.get("result") or len(data["result"]) == 0:
        raise HTTPException(status_code=404, detail="未找到账户")
    return data["result"][0]["id"]

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
@app.post("/api/auth")
async def set_auth(account: Account):
    """设置认证信息"""
    current_auth["email"] = account.email
    current_auth["api_key"] = account.api_key
    
    # 验证认证信息
    try:
        data = make_request("GET", "/accounts")
        return {
            "success": True,
            "message": "认证信息已设置",
            "account": data["result"][0] if data.get("result") else None
        }
    except HTTPException as e:
        current_auth["email"] = ""
        current_auth["api_key"] = ""
        raise HTTPException(status_code=401, detail="认证失败：" + str(e.detail))

@app.get("/api/accounts")
async def list_accounts():
    """列出所有账号"""
    data = make_request("GET", "/accounts")
    return data

@app.get("/api/accounts/current")
async def get_current_account():
    """获取当前账号信息"""
    account_id = get_account_id()
    data = make_request("GET", f"/accounts/{account_id}")
    return data

# Zone管理
@app.get("/api/zones")
async def list_zones():
    """列出所有Zone"""
    data = make_request("GET", "/zones")
    return data

@app.post("/api/zones")
async def create_zone(domain: Dict[str, str]):
    """创建新Zone"""
    account_id = get_account_id()
    payload = {
        "name": domain["domain"],
        "account": {"id": account_id}
    }
    data = make_request("POST", "/zones", json=payload)
    return data

@app.get("/api/zones/{zone_id}")
async def get_zone_info(zone_id: str):
    """获取Zone信息"""
    data = make_request("GET", f"/zones/{zone_id}")
    return data

@app.delete("/api/zones/{zone_id}")
async def delete_zone(zone_id: str):
    """删除Zone"""
    data = make_request("DELETE", f"/zones/{zone_id}")
    return data

@app.post("/api/zones/{zone_id}/purge")
async def purge_cache(zone_id: str):
    """清除缓存"""
    payload = {"purge_everything": True}
    data = make_request("POST", f"/zones/{zone_id}/purge_cache", json=payload)
    return data

# DNS记录管理
@app.get("/api/zones/{zone_id}/dns")
async def list_dns_records(zone_id: str, type: Optional[str] = None):
    """列出DNS记录"""
    params = {}
    if type:
        params["type"] = type
    data = make_request("GET", f"/zones/{zone_id}/dns_records", params=params)
    return data

@app.post("/api/zones/{zone_id}/dns")
async def create_dns_record(zone_id: str, record: DNSRecord):
    """创建DNS记录"""
    payload = {
        "type": record.type,
        "name": record.name,
        "content": record.content,
        "ttl": record.ttl,
        "proxied": record.proxied
    }
    if record.priority and record.type in ["MX", "SRV"]:
        payload["priority"] = record.priority
    
    data = make_request("POST", f"/zones/{zone_id}/dns_records", json=payload)
    return data

@app.delete("/api/zones/{zone_id}/dns/{record_id}")
async def delete_dns_record(zone_id: str, record_id: str):
    """删除DNS记录"""
    data = make_request("DELETE", f"/zones/{zone_id}/dns_records/{record_id}")
    return data

@app.get("/api/zones/{zone_id}/dns/export")
async def export_dns_records(zone_id: str):
    """导出DNS记录"""
    data = make_request("GET", f"/zones/{zone_id}/dns_records/export")
    return data

# Worker管理
@app.get("/api/workers")
async def list_workers():
    """列出所有Worker"""
    account_id = get_account_id()
    data = make_request("GET", f"/accounts/{account_id}/workers/scripts")
    return data

@app.post("/api/workers")
async def deploy_worker(worker: WorkerDeploy):
    """部署Worker"""
    account_id = get_account_id()
    headers = {
        "X-Auth-Email": current_auth["email"],
        "X-Auth-Key": current_auth["api_key"],
        "Content-Type": "application/javascript"
    }
    
    url = f"{CF_API_BASE}/accounts/{account_id}/workers/scripts/{worker.name}"
    response = requests.put(url, headers=headers, data=worker.script.encode('utf-8'))
    data = response.json()
    
    if not data.get("success", False):
        errors = data.get("errors", [{"message": "Unknown error"}])
        raise HTTPException(
            status_code=response.status_code,
            detail=errors[0].get("message", "Worker deployment failed")
        )
    
    return data

@app.delete("/api/workers/{name}")
async def delete_worker(name: str):
    """删除Worker"""
    account_id = get_account_id()
    data = make_request("DELETE", f"/accounts/{account_id}/workers/scripts/{name}")
    return data

# Worker路由
@app.get("/api/zones/{zone_id}/routes")
async def list_worker_routes(zone_id: str):
    """列出Worker路由"""
    data = make_request("GET", f"/zones/{zone_id}/workers/routes")
    return data

@app.post("/api/zones/{zone_id}/routes")
async def create_worker_route(zone_id: str, route: WorkerRoute):
    """创建Worker路由"""
    payload = {
        "pattern": route.pattern,
        "script": route.worker_name
    }
    data = make_request("POST", f"/zones/{zone_id}/workers/routes", json=payload)
    return data

@app.delete("/api/zones/{zone_id}/routes/{route_id}")
async def delete_worker_route(zone_id: str, route_id: str):
    """删除Worker路由"""
    data = make_request("DELETE", f"/zones/{zone_id}/workers/routes/{route_id}")
    return data

# Pages管理
@app.get("/api/pages")
async def list_pages():
    """列出所有Pages项目"""
    account_id = get_account_id()
    data = make_request("GET", f"/accounts/{account_id}/pages/projects")
    return data

@app.post("/api/pages")
async def create_pages_project(project: PagesProject):
    """创建Pages项目"""
    account_id = get_account_id()
    payload = {
        "name": project.name,
        "production_branch": project.production_branch
    }
    data = make_request("POST", f"/accounts/{account_id}/pages/projects", json=payload)
    return data

@app.get("/api/pages/{project_name}")
async def get_pages_info(project_name: str):
    """获取Pages项目信息"""
    account_id = get_account_id()
    data = make_request("GET", f"/accounts/{account_id}/pages/projects/{project_name}")
    return data

@app.delete("/api/pages/{project_name}")
async def delete_pages_project(project_name: str):
    """删除Pages项目"""
    account_id = get_account_id()
    data = make_request("DELETE", f"/accounts/{account_id}/pages/projects/{project_name}")
    return data

@app.get("/api/pages/{project_name}/deployments")
async def list_pages_deployments(project_name: str):
    """列出Pages部署"""
    account_id = get_account_id()
    data = make_request("GET", f"/accounts/{account_id}/pages/projects/{project_name}/deployments")
    return data

@app.post("/api/pages/{project_name}/deployments")
async def deploy_pages(
    project_name: str,
    branch: str = Form("main"),
    worker_file: UploadFile = File(...)
):
    """部署Pages项目"""
    account_id = get_account_id()
    
    # 读取Worker文件内容
    worker_content = await worker_file.read()
    
    # 使用multipart/form-data上传
    files = {
        'manifest': (None, '{}'),
        'branch': (None, branch),
        '_worker.js': ('_worker.js', worker_content, 'text/javascript'),
    }
    
    headers = {
        "X-Auth-Email": current_auth["email"],
        "X-Auth-Key": current_auth["api_key"],
    }
    
    url = f"{CF_API_BASE}/accounts/{account_id}/pages/projects/{project_name}/deployments"
    response = requests.post(url, headers=headers, files=files)
    data = response.json()
    
    if not data.get("success", False):
        errors = data.get("errors", [{"message": "Unknown error"}])
        raise HTTPException(
            status_code=response.status_code,
            detail=errors[0].get("message", "Pages deployment failed")
        )
    
    return data

# KV管理
@app.get("/api/kv/namespaces")
async def list_kv_namespaces():
    """列出KV命名空间"""
    account_id = get_account_id()
    data = make_request("GET", f"/accounts/{account_id}/storage/kv/namespaces")
    return data

@app.post("/api/kv/namespaces")
async def create_kv_namespace(namespace: KVNamespace):
    """创建KV命名空间"""
    account_id = get_account_id()
    payload = {"title": namespace.title}
    data = make_request("POST", f"/accounts/{account_id}/storage/kv/namespaces", json=payload)
    return data

@app.delete("/api/kv/namespaces/{namespace_id}")
async def delete_kv_namespace(namespace_id: str):
    """删除KV命名空间"""
    account_id = get_account_id()
    data = make_request("DELETE", f"/accounts/{account_id}/storage/kv/namespaces/{namespace_id}")
    return data

@app.get("/api/kv/namespaces/{namespace_id}/keys")
async def list_kv_keys(namespace_id: str, prefix: Optional[str] = None):
    """列出KV键"""
    account_id = get_account_id()
    params = {}
    if prefix:
        params["prefix"] = prefix
    data = make_request("GET", f"/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/keys", params=params)
    return data

@app.get("/api/kv/namespaces/{namespace_id}/keys/{key}")
async def get_kv_value(namespace_id: str, key: str):
    """获取KV值"""
    account_id = get_account_id()
    headers = {
        "X-Auth-Email": current_auth["email"],
        "X-Auth-Key": current_auth["api_key"],
    }
    url = f"{CF_API_BASE}/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key}"
    response = requests.get(url, headers=headers)
    
    if response.status_code == 200:
        return {"success": True, "result": response.text}
    else:
        raise HTTPException(status_code=response.status_code, detail="Key not found")

@app.put("/api/kv/namespaces/{namespace_id}/keys")
async def put_kv_value(namespace_id: str, pair: KVPair):
    """设置KV值"""
    account_id = get_account_id()
    headers = {
        "X-Auth-Email": current_auth["email"],
        "X-Auth-Key": current_auth["api_key"],
    }
    url = f"{CF_API_BASE}/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{pair.key}"
    response = requests.put(url, headers=headers, data=pair.value.encode('utf-8'))
    data = response.json()
    
    if not data.get("success", False):
        raise HTTPException(status_code=response.status_code, detail="Failed to put KV value")
    
    return data

@app.delete("/api/kv/namespaces/{namespace_id}/keys/{key}")
async def delete_kv_key(namespace_id: str, key: str):
    """删除KV键"""
    account_id = get_account_id()
    data = make_request("DELETE", f"/accounts/{account_id}/storage/kv/namespaces/{namespace_id}/values/{key}")
    return data

# R2管理
@app.get("/api/r2/buckets")
async def list_r2_buckets():
    """列出R2桶"""
    account_id = get_account_id()
    data = make_request("GET", f"/accounts/{account_id}/r2/buckets")
    return data

@app.post("/api/r2/buckets")
async def create_r2_bucket(bucket: R2Bucket):
    """创建R2桶"""
    account_id = get_account_id()
    payload = {
        "name": bucket.name,
        "locationHint": bucket.location
    }
    data = make_request("POST", f"/accounts/{account_id}/r2/buckets", json=payload)
    return data

@app.get("/api/r2/buckets/{bucket_name}")
async def get_r2_bucket_info(bucket_name: str):
    """获取R2桶信息"""
    account_id = get_account_id()
    data = make_request("GET", f"/accounts/{account_id}/r2/buckets/{bucket_name}")
    return data

@app.delete("/api/r2/buckets/{bucket_name}")
async def delete_r2_bucket(bucket_name: str):
    """删除R2桶"""
    account_id = get_account_id()
    data = make_request("DELETE", f"/accounts/{account_id}/r2/buckets/{bucket_name}")
    return data

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
