#!/usr/bin/env python3
"""
Cloudflare Pages 部署示例脚本
展示如何使用 Email + API Key 方式部署
"""

import requests
import json

# ===== 配置 =====
EMAIL = "your@email.com"
API_KEY = "your_global_api_key"
PROJECT_NAME = "my-test-project"
WORKER_FILE = "./example-worker.js"
BASE_URL = "http://localhost:8000"
# ================

def get_account_id():
    """自动获取 Account ID"""
    headers = {
        "X-Auth-Email": EMAIL,
        "X-Auth-Key": API_KEY,
    }
    r = requests.get("https://api.cloudflare.com/client/v4/accounts", headers=headers)
    data = r.json()
    if not data.get("success"):
        raise Exception(f"获取账户失败: {data.get('errors', [])}")
    return data["result"][0]["id"]

def create_project():
    """创建 Pages 项目"""
    account_id = get_account_id()
    print(f"Account ID: {account_id}")
    input("按回车键继续创建项目...")
    
    url = f"https://api.cloudflare.com/client/v4/accounts/{account_id}/pages/projects"
    
    payload = {
        "name": PROJECT_NAME,
        "production_branch": "main"
    }
    
    headers = {
        "X-Auth-Email": EMAIL,
        "X-Auth-Key": API_KEY,
        "Content-Type": "application/json"
    }
    
    response = requests.post(url, headers=headers, json=payload)
    print(f"\n创建项目响应:")
    print(json.dumps(response.json(), indent=2))

def deploy():
    """部署到 Pages"""
    account_id = get_account_id()
    print(f"Account ID: {account_id}")
    
    url = f"https://api.cloudflare.com/client/v4/accounts/{account_id}/pages/projects/{PROJECT_NAME}/deployments"
    
    headers = {
        "X-Auth-Email": EMAIL,
        "X-Auth-Key": API_KEY,
    }
    
    # 读取 Worker 文件
    with open(WORKER_FILE, 'r', encoding='utf-8') as f:
        worker_content = f.read()
    
    # 使用 multipart/form-data
    files = {
        'manifest': (None, '{}'),
        'branch': (None, 'main'),
        '_worker.js': ('_worker.js', worker_content, 'text/javascript'),
    }
    
    print(f"\n🚀 正在部署到 {PROJECT_NAME}...")
    
    response = requests.post(url, headers=headers, files=files)
    
    if response.status_code in [200, 201]:
        result = response.json()
        if result.get('success'):
            deployment = result['result']
            print(f"\n✅ 部署成功!")
            print(f"🌐 URL: {deployment['url']}")
            print(f"📋 ID: {deployment['id']}")
        else:
            print(f"\n❌ 失败: {result}")
    else:
        print(f"\n❌ 请求失败 ({response.status_code})")
        print(response.text)

def main():
    print("=" * 60)
    print("Cloudflare Pages 部署示例")
    print("=" * 60)
    print(f"\nEmail: {EMAIL}")
    print(f"Project: {PROJECT_NAME}")
    print(f"Worker File: {WORKER_FILE}")
    print("\n提示：请确保已正确设置 EMAIL 和 API_KEY")
    print("=" * 60)
    
    choice = input("\n请选择操作：\n1. 创建项目\n2. 部署\n3. 创建并部署\n\n请输入 (1/2/3): ")
    
    try:
        if choice == "1":
            create_project()
        elif choice == "2":
            deploy()
        elif choice == "3":
            create_project()
            print("\n" + "=" * 60)
            deploy()
        else:
            print("无效的选择")
    except Exception as e:
        print(f"\n❌ 错误: {e}")

if __name__ == "__main__":
    main()
