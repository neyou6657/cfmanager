# 部署指南 - Hugging Face Spaces

## 📦 项目结构

```
cloudflare-manager/
├── Dockerfile                    # Docker构建文件
├── docker-entrypoint.sh         # 容器启动脚本
├── README_HUGGINGFACE.md        # Hugging Face Space README
├── .dockerignore                # Docker忽略文件
├── cloudflare-manager/          # Go CLI后端
│   ├── main.go
│   ├── commands/                # 命令实现
│   ├── client/                  # Cloudflare客户端
│   ├── config/                  # 配置管理
│   └── utils/                   # 工具函数
├── backend/                     # Python FastAPI后端
│   ├── main.py                  # API服务器
│   └── requirements.txt         # Python依赖
└── frontend/                    # React前端
    ├── package.json
    ├── public/
    └── src/
        ├── App.js               # 主应用
        ├── pages/               # 页面组件
        └── index.js
```

## 🚀 部署到 Hugging Face Spaces

### 方法 1: 通过 Web 界面

1. **创建新 Space**
   - 访问 https://huggingface.co/new-space
   - Space name: `cloudflare-manager`
   - License: `MIT`
   - Select SDK: `Docker`
   - Space hardware: `CPU basic` (免费)

2. **上传文件**
   - 点击 "Files" 标签
   - 上传以下文件到根目录：
     - `Dockerfile`
     - `docker-entrypoint.sh`
     - `README_HUGGINGFACE.md` (重命名为 `README.md`)
     - `.dockerignore`
   - 上传 `cloudflare-manager/` 目录
   - 上传 `backend/` 目录
   - 上传 `frontend/` 目录

3. **设置环境变量（可选）**
   - 点击 "Settings" 标签
   - 在 "Variables and secrets" 部分添加：
     ```
     CLOUDFLARE_EMAIL=tqa88tawlq@downnaturer.me
     CLOUDFLARE_TOKEN=4e2dd4818267ebd2ab8d1aa2e7f9bf4151b70
     ```

4. **等待构建完成**
   - Space 会自动构建 Docker 镜像
   - 构建时间约 5-10 分钟
   - 构建完成后自动启动

### 方法 2: 使用 Git

1. **克隆 Space 仓库**
   ```bash
   git clone https://huggingface.co/spaces/YOUR_USERNAME/cloudflare-manager
   cd cloudflare-manager
   ```

2. **复制项目文件**
   ```bash
   # 从本项目复制所有需要的文件
   cp /path/to/project/Dockerfile .
   cp /path/to/project/docker-entrypoint.sh .
   cp /path/to/project/README_HUGGINGFACE.md README.md
   cp /path/to/project/.dockerignore .
   cp -r /path/to/project/cloudflare-manager .
   cp -r /path/to/project/backend .
   cp -r /path/to/project/frontend .
   ```

3. **提交并推送**
   ```bash
   git add .
   git commit -m "Initial deployment"
   git push
   ```

4. **Space 会自动构建和部署**

## 🔧 本地测试

在部署到 Hugging Face 之前，建议先本地测试：

```bash
# 构建 Docker 镜像
docker build -t cloudflare-manager .

# 运行容器
docker run -p 7860:7860 \
  -e CLOUDFLARE_EMAIL="tqa88tawlq@downnaturer.me" \
  -e CLOUDFLARE_TOKEN="4e2dd4818267ebd2ab8d1aa2e7f9bf4151b70" \
  cloudflare-manager

# 访问应用
open http://localhost:7860
```

## 📊 Dockerfile 解析

```dockerfile
# 阶段1: 构建前端
FROM node:18-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

# 阶段2: 编译 Go 程序
FROM golang:1.21-alpine AS go-builder
WORKDIR /app
COPY cloudflare-manager/ ./cloudflare-manager/
WORKDIR /app/cloudflare-manager
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o cfm

# 阶段3: 最终镜像
FROM python:3.11-slim
WORKDIR /app

# 安装系统依赖
RUN apt-get update && apt-get install -y curl

# 复制 Go 二进制文件
COPY --from=go-builder /app/cloudflare-manager/cfm /app/cloudflare-manager/cfm

# 安装 Python 后端
COPY backend/ ./backend/
WORKDIR /app/backend
RUN pip install --no-cache-dir -r requirements.txt

# 复制前端构建
COPY --from=frontend-builder /app/frontend/build /app/frontend/build

# 暴露端口
EXPOSE 7860

# 启动脚本
COPY docker-entrypoint.sh /app/
RUN chmod +x /app/docker-entrypoint.sh
CMD ["/app/docker-entrypoint.sh"]
```

## 🎯 端口和路由

- **主应用**: http://localhost:7860/
- **API文档**: http://localhost:7860/api/docs
- **健康检查**: http://localhost:7860/health

## 🔐 凭证管理

### 自动初始化（推荐）

通过环境变量自动初始化默认账号：

```dockerfile
ENV CLOUDFLARE_EMAIL="tqa88tawlq@downnaturer.me"
ENV CLOUDFLARE_TOKEN="4e2dd4818267ebd2ab8d1aa2e7f9bf4151b70"
```

### 手动添加账号

部署后，在 Web UI 中：
1. 进入 "Accounts" 页面
2. 点击 "Add Account"
3. 输入账号信息

## 🐛 故障排查

### 构建失败

检查日志：
```bash
# Hugging Face Spaces 会显示构建日志
# 常见问题：
# 1. 文件路径错误
# 2. 依赖安装失败
# 3. 端口冲突
```

### 运行时错误

```bash
# 查看容器日志
docker logs <container-id>

# 进入容器调试
docker exec -it <container-id> /bin/bash

# 测试 Go CLI
/app/cloudflare-manager/cfm --version

# 测试 FastAPI
curl http://localhost:7860/health
```

### 前端无法加载

检查：
1. 前端构建是否成功
2. 静态文件路径是否正确
3. API代理配置是否正确

## 📈 性能优化

### 多阶段构建优化

- ✅ 前端使用 Alpine 镜像
- ✅ Go 静态编译，无需运行时
- ✅ 最终镜像基于 slim 版本
- ✅ 清理不必要的文件

### 镜像大小

预期大小：~500MB
- Python base: ~200MB
- 前端静态文件: ~5MB
- Go 二进制: ~15MB
- 依赖和系统文件: ~280MB

## 🔒 安全建议

1. **不要提交敏感信息到 Git**
   ```bash
   # .gitignore 已配置
   .cloudflare-manager.yaml
   *.env
   ```

2. **使用 Secrets 管理敏感配置**
   - 在 Hugging Face Spaces Settings 中配置
   - 不要硬编码在 Dockerfile 中

3. **限制 API Token 权限**
   - 只授予需要的最小权限
   - 定期轮换 Token

4. **添加认证（生产环境）**
   ```python
   # 在 FastAPI 中添加
   from fastapi.security import HTTPBasic
   ```

## 📚 相关文档

- [Hugging Face Spaces 文档](https://huggingface.co/docs/hub/spaces)
- [Docker 多阶段构建](https://docs.docker.com/build/building/multi-stage/)
- [FastAPI 部署](https://fastapi.tiangolo.com/deployment/)
- [React 生产构建](https://create-react-app.dev/docs/production-build/)

## ✅ 部署检查清单

在部署前确认：

- [ ] 所有文件已上传
- [ ] Dockerfile 语法正确
- [ ] 环境变量已配置
- [ ] 端口设置为 7860
- [ ] README.md 已更新
- [ ] 本地测试通过
- [ ] Git 提交完成

## 🎉 部署成功后

访问你的 Space：
```
https://huggingface.co/spaces/YOUR_USERNAME/cloudflare-manager
```

享受管理 Cloudflare 的便利！🚀

---

**问题反馈**: 如有问题，请在 GitHub Issues 或 Hugging Face Discussions 中提出。
