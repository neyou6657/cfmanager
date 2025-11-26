# 更改说明

## 概述

将 Cloudflare Manager 从基于 CLI 工具的方式改为**直接使用 HTTP API + Email/API Key 认证**的方式，完全移除了对 SDK 和 Go CLI 的依赖。

## 主要更改

### 1. 后端重写 (`backend/main.py`)

**之前**：
- 使用 Go 编写的 `cfm` CLI 工具
- 通过 `subprocess` 调用命令行
- 需要管理账户配置文件

**现在**：
- 直接使用 Python `requests` 库
- 使用 `X-Auth-Email` 和 `X-Auth-Key` 请求头进行认证
- 所有 API 调用都是标准 HTTP 请求
- 新增 `/api/auth` 端点用于设置认证信息

### 2. 认证方式简化

**之前**：
- 复杂的多账户管理
- 需要维护配置文件
- 账户切换机制

**现在**：
- 简单的 Email + API Key
- 内存中存储认证信息
- 一次设置，所有 API 调用共享

### 3. Pages 功能完整实现

新增完整的 Pages 管理功能：
- ✅ 创建项目 (`POST /api/pages`)
- ✅ 列出项目 (`GET /api/pages`)
- ✅ 删除项目 (`DELETE /api/pages/{project_name}`)
- ✅ 上传并部署 Worker (`POST /api/pages/{project_name}/deployments`)

### 4. 前端界面更新

#### Accounts 页面 (`frontend/src/pages/Accounts.js`)
- 改为登录界面
- 输入 Email 和 API Key
- 显示认证状态和账户信息

#### Pages 页面 (`frontend/src/pages/Pages.js`)
- 完整的项目管理界面
- 文件上传功能（支持 Worker 脚本）
- 部署状态展示

### 5. 部署简化

**之前**：
```dockerfile
# 需要 Go 编译器
FROM golang:1.21-alpine AS go-builder
# 编译 cfm 工具
RUN go build -o cfm
```

**现在**：
```dockerfile
# 只需要 Python
FROM python:3.11-slim
# 安装 requests
RUN pip install -r requirements.txt
```

## API 使用示例

### 设置认证

```bash
curl -X POST http://localhost:8000/api/auth \
  -H "Content-Type: application/json" \
  -d '{
    "email": "your@email.com",
    "api_key": "your_global_api_key"
  }'
```

### 创建 Pages 项目

```bash
curl -X POST http://localhost:8000/api/pages \
  -H "Content-Type: application/json" \
  -d '{
    "name": "my-project",
    "production_branch": "main"
  }'
```

### 部署到 Pages

```bash
curl -X POST http://localhost:8000/api/pages/my-project/deployments \
  -F "branch=main" \
  -F "worker_file=@_worker.js"
```

## 技术细节

### 认证实现

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

### 文件上传实现

```python
files = {
    'manifest': (None, '{}'),
    'branch': (None, 'main'),
    '_worker.js': ('_worker.js', worker_content, 'text/javascript'),
}

response = requests.post(url, headers=headers, files=files)
```

## 获取 API Key

1. 登录 [Cloudflare Dashboard](https://dash.cloudflare.com/)
2. 进入 **My Profile** > **API Tokens**
3. 在 **API Keys** 部分查看 **Global API Key**

## 优势

✅ **更简单**：无需编译 Go 程序  
✅ **更轻量**：只需 Python + requests  
✅ **更直接**：直接 HTTP API 调用  
✅ **更透明**：可以清楚看到每个 API 请求  
✅ **更易调试**：标准 HTTP 工具都可以使用  

## 测试

运行测试脚本：
```bash
python test_api.py
```

运行部署示例：
```bash
python example_deploy.py
```

## 文件清单

### 修改的文件
- `backend/main.py` - 完全重写
- `backend/requirements.txt` - 添加 requests
- `frontend/src/pages/Accounts.js` - 改为登录界面
- `frontend/src/pages/Pages.js` - 新增完整功能
- `Dockerfile` - 移除 Go 构建
- `docker-entrypoint.sh` - 简化启动脚本
- `.gitignore` - 添加 Python venv

### 新增的文件
- `README.md` - 项目文档
- `test_api.py` - API 测试脚本
- `example_deploy.py` - 部署示例
- `CHANGES.md` - 本文件

## 参考

- [Cloudflare API 文档](https://developers.cloudflare.com/api/)
- [Pages API 参考](https://developers.cloudflare.com/api/operations/pages-project-list-projects)
- CloudflreDocs 仓库（SDK 文档）
