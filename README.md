# Cloudflare Manager

一个基于 Web 的多账户 Cloudflare 资源管理平台，支持管理 Zones、DNS、Workers、Pages、KV、R2 等。

## 特性

- 🔐 **简单认证**：使用 Email + API Key 认证（无需 SDK）
- 📄 **Pages 管理**：创建项目、部署 Worker 脚本
- 🌐 **Zone 管理**：域名和 DNS 记录管理
- ⚙️ **Workers**：部署和管理 Cloudflare Workers
- 💾 **存储服务**：KV、R2 桶管理
- 🎨 **现代化 UI**：基于 React + Ant Design

## 快速开始

### 1. 使用 Docker（推荐）

```bash
# 构建镜像
docker build -t cloudflare-manager .

# 运行容器（可选：设置环境变量）
docker run -p 7860:7860 \
  -e CF_EMAIL="your@email.com" \
  -e CF_API_KEY="your_api_key" \
  cloudflare-manager
```

访问 http://localhost:7860

### 2. 本地开发

#### 后端

```bash
cd backend
pip install -r requirements.txt
python main.py
```

后端将运行在 http://localhost:8000

#### 前端

```bash
cd frontend
npm install
npm start
```

前端将运行在 http://localhost:3000

## 认证方式

### 获取 Cloudflare API Key

1. 登录 [Cloudflare Dashboard](https://dash.cloudflare.com/)
2. 进入 **My Profile** > **API Tokens** > **API Keys**
3. 查看或创建 **Global API Key**
4. 在应用的 "账户管理" 页面输入：
   - Email：您的 Cloudflare 账户邮箱
   - API Key：您的 Global API Key

## API 使用示例

### Python 示例

```python
import requests

# 1. 设置认证
response = requests.post("http://localhost:8000/api/auth", json={
    "email": "your@email.com",
    "api_key": "your_api_key"
})

# 2. 获取账户信息
response = requests.get("http://localhost:8000/api/accounts")
account_id = response.json()["result"][0]["id"]

# 3. 创建 Pages 项目
response = requests.post("http://localhost:8000/api/pages", json={
    "name": "my-project",
    "production_branch": "main"
})

# 4. 部署 Worker 到 Pages
files = {
    'branch': (None, 'main'),
    'worker_file': ('_worker.js', open('_worker.js', 'rb'), 'text/javascript')
}
response = requests.post(
    f"http://localhost:8000/api/pages/my-project/deployments",
    files=files
)
```

### cURL 示例

```bash
# 设置认证
curl -X POST http://localhost:8000/api/auth \
  -H "Content-Type: application/json" \
  -d '{"email":"your@email.com","api_key":"your_api_key"}'

# 获取 Pages 项目列表
curl http://localhost:8000/api/pages

# 创建 Pages 项目
curl -X POST http://localhost:8000/api/pages \
  -H "Content-Type: application/json" \
  -d '{"name":"my-project","production_branch":"main"}'

# 部署到 Pages
curl -X POST http://localhost:8000/api/pages/my-project/deployments \
  -F "branch=main" \
  -F "worker_file=@_worker.js"
```

## API 端点

### 认证
- `POST /api/auth` - 设置认证信息
- `GET /api/accounts` - 获取账户列表
- `GET /api/accounts/current` - 获取当前账户信息

### Pages
- `GET /api/pages` - 列出所有项目
- `POST /api/pages` - 创建项目
- `GET /api/pages/{project_name}` - 获取项目信息
- `DELETE /api/pages/{project_name}` - 删除项目
- `GET /api/pages/{project_name}/deployments` - 列出部署
- `POST /api/pages/{project_name}/deployments` - 部署项目

### Zones
- `GET /api/zones` - 列出所有 Zone
- `POST /api/zones` - 创建 Zone
- `GET /api/zones/{zone_id}` - 获取 Zone 信息
- `DELETE /api/zones/{zone_id}` - 删除 Zone
- `POST /api/zones/{zone_id}/purge` - 清除缓存

### DNS
- `GET /api/zones/{zone_id}/dns` - 列出 DNS 记录
- `POST /api/zones/{zone_id}/dns` - 创建 DNS 记录
- `DELETE /api/zones/{zone_id}/dns/{record_id}` - 删除 DNS 记录

### Workers
- `GET /api/workers` - 列出所有 Worker
- `POST /api/workers` - 部署 Worker
- `DELETE /api/workers/{name}` - 删除 Worker

### KV
- `GET /api/kv/namespaces` - 列出 KV 命名空间
- `POST /api/kv/namespaces` - 创建命名空间
- `DELETE /api/kv/namespaces/{namespace_id}` - 删除命名空间
- `GET /api/kv/namespaces/{namespace_id}/keys` - 列出键
- `GET /api/kv/namespaces/{namespace_id}/keys/{key}` - 获取值
- `PUT /api/kv/namespaces/{namespace_id}/keys` - 设置值
- `DELETE /api/kv/namespaces/{namespace_id}/keys/{key}` - 删除键

### R2
- `GET /api/r2/buckets` - 列出所有桶
- `POST /api/r2/buckets` - 创建桶
- `GET /api/r2/buckets/{bucket_name}` - 获取桶信息
- `DELETE /api/r2/buckets/{bucket_name}` - 删除桶

## 技术栈

- **后端**：FastAPI + Python 3.11 + Requests
- **前端**：React + Ant Design
- **部署**：Docker

## 开发说明

### 后端架构

后端直接使用 `requests` 库调用 Cloudflare REST API，通过 `X-Auth-Email` 和 `X-Auth-Key` 请求头进行认证：

```python
headers = {
    "X-Auth-Email": email,
    "X-Auth-Key": api_key,
    "Content-Type": "application/json"
}
response = requests.get(
    "https://api.cloudflare.com/client/v4/accounts",
    headers=headers
)
```

### API 文档

启动后端后访问：
- Swagger UI: http://localhost:8000/docs
- ReDoc: http://localhost:8000/redoc

## 故障排除

### 认证失败

确保：
1. Email 正确
2. API Key 是 **Global API Key**（不是 API Token）
3. API Key 有足够的权限

### 部署失败

Pages 部署需要：
1. 项目已创建
2. Worker 文件是有效的 JavaScript
3. 文件名为 `_worker.js`

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
