#!/usr/bin/env python3
"""
测试脚本 - 验证 Cloudflare Pages 部署 API
使用 Email + API Key 认证方式
"""

import requests
import json
import sys

# 配置
BASE_URL = "http://localhost:8000"

# 示例认证信息（需要替换为实际的）
EMAIL = "your@email.com"
API_KEY = "your_api_key_here"

def test_auth():
    """测试认证"""
    print("🔐 测试认证...")
    response = requests.post(f"{BASE_URL}/api/auth", json={
        "email": EMAIL,
        "api_key": API_KEY
    })
    print(f"状态码: {response.status_code}")
    print(f"响应: {response.json()}")
    return response.status_code == 200

def test_get_accounts():
    """测试获取账户列表"""
    print("\n📋 测试获取账户列表...")
    response = requests.get(f"{BASE_URL}/api/accounts")
    print(f"状态码: {response.status_code}")
    data = response.json()
    print(f"响应: {json.dumps(data, indent=2)}")
    return response.status_code == 200

def test_pages_list():
    """测试获取 Pages 项目列表"""
    print("\n📄 测试获取 Pages 项目...")
    response = requests.get(f"{BASE_URL}/api/pages")
    print(f"状态码: {response.status_code}")
    data = response.json()
    print(f"响应: {json.dumps(data, indent=2)}")
    return response.status_code == 200

def test_create_pages_project():
    """测试创建 Pages 项目"""
    print("\n✨ 测试创建 Pages 项目...")
    response = requests.post(f"{BASE_URL}/api/pages", json={
        "name": "test-project-123",
        "production_branch": "main"
    })
    print(f"状态码: {response.status_code}")
    data = response.json()
    print(f"响应: {json.dumps(data, indent=2)}")
    return response.status_code in [200, 201]

def main():
    print("=" * 60)
    print("Cloudflare Manager API 测试")
    print("=" * 60)
    
    # 测试根路由
    print("\n🏠 测试根路由...")
    response = requests.get(f"{BASE_URL}/")
    print(f"状态码: {response.status_code}")
    print(f"响应: {response.json()}")
    
    print("\n" + "=" * 60)
    print("注意：以下测试需要有效的 Email 和 API Key")
    print("请修改脚本中的 EMAIL 和 API_KEY 变量")
    print("=" * 60)
    
    # 如果有真实的认证信息，继续测试
    if EMAIL != "your@email.com":
        test_auth()
        test_get_accounts()
        test_pages_list()
        # test_create_pages_project()  # 取消注释以测试创建项目
    
    print("\n✅ 测试完成！")

if __name__ == "__main__":
    main()
