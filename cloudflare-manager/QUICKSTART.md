# 快速开始指南

## 1. 编译安装

```bash
# 编译
cd cloudflare-manager
go build -o cfm

# 安装到系统（可选）
sudo mv cfm /usr/local/bin/

# 或者使用本地路径
alias cfm="./cfm"
```

## 2. 添加你的Cloudflare账号

```bash
# 使用你的API Token添加账号
cfm account add myaccount --token 4e2dd4818267ebd2ab8d1aa2e7f9bf4151b70
```

**获取API Token:**
1. 访问 https://dash.cloudflare.com/profile/api-tokens
2. 点击 "Create Token"
3. 选择权限（推荐：Edit zone DNS + Workers Scripts）
4. 创建并复制Token

## 3. 验证账号

```bash
# 查看所有账号
cfm account list

# 查看当前账号信息
cfm account info
```

## 4. 管理域名

```bash
# 创建新域名
cfm zone create example.com

# 列出所有域名
cfm zone list

# 查看域名详情
cfm zone info example.com
```

## 5. DNS记录管理

```bash
# 创建A记录
cfm dns create example.com A www 192.0.2.1 --proxied

# 创建CNAME记录
cfm dns create example.com CNAME blog www.example.com

# 列出所有DNS记录
cfm dns list example.com

# 删除DNS记录
cfm dns delete example.com <record-id>
```

## 6. 部署Worker

```bash
# 部署Worker脚本
cfm worker deploy hello-world example-worker.js

# 创建路由规则（将域名路由到Worker）
cfm worker route create example.com "example.com/*" hello-world

# 查看所有路由
cfm worker route list example.com
```

## 7. Pages项目管理

```bash
# 列出所有Pages项目
cfm pages list

# 查看项目信息
cfm pages info my-project

# 查看部署历史
cfm pages deployment list my-project
```

## 完整示例：从零到部署Worker

```bash
# 1. 添加账号
cfm account add production --token YOUR_TOKEN

# 2. 创建域名
cfm zone create myapp.com

# 3. 添加DNS记录
cfm dns create myapp.com A @ 192.0.2.1 --proxied
cfm dns create myapp.com CNAME www @ --proxied

# 4. 部署Worker
cfm worker deploy api-worker example-worker.js

# 5. 将API路径路由到Worker
cfm worker route create myapp.com "myapp.com/api/*" api-worker

# 6. 验证配置
cfm dns list myapp.com
cfm worker route list myapp.com

# 7. 访问测试
curl https://myapp.com/api/hello
```

## 多账号切换

```bash
# 添加多个账号
cfm account add personal --token TOKEN1
cfm account add company --token TOKEN2

# 查看所有账号
cfm account list

# 切换账号
cfm account switch company

# 现在所有操作都在company账号上执行
cfm zone list
```

## 常用命令速查

```bash
# 账号管理
cfm account add <name> --token <token>
cfm account list
cfm account switch <name>

# Zone管理
cfm zone list
cfm zone create <domain>
cfm zone info <domain>
cfm zone purge <domain> --everything

# DNS管理
cfm dns list <domain>
cfm dns create <domain> <type> <name> <content>
cfm dns delete <domain> <record-id>

# Worker管理
cfm worker deploy <name> <file>
cfm worker route create <domain> <pattern> <worker-name>
cfm worker route list <domain>
cfm worker route delete <domain> <route-id>

# Pages管理
cfm pages list
cfm pages info <project>
cfm pages deployment list <project>
```

## 配置文件

配置文件自动保存在 `~/.cloudflare-manager.yaml`:

```yaml
current_account: myaccount
accounts:
  - name: myaccount
    api_token: your_token_here
    account_id: auto_detected_id
    email: your@email.com
```

## 获取帮助

```bash
# 主帮助
cfm --help

# 命令帮助
cfm account --help
cfm zone --help
cfm dns --help
cfm worker --help
cfm pages --help

# 子命令帮助
cfm dns create --help
cfm worker route create --help
```

## 故障排查

### 问题1: "No accounts configured"
```bash
# 解决：添加账号
cfm account add myaccount --token YOUR_TOKEN
```

### 问题2: "Failed to get account ID"
```bash
# 检查Token权限
# Token需要至少有 "Account > Account Settings > Read" 权限
```

### 问题3: "Zone not found"
```bash
# 使用完整域名或Zone ID
cfm zone list  # 查看所有zone
cfm dns list example.com  # 使用完整域名
```

## 最佳实践

1. **使用有限权限的Token**: 为每个用途创建专用Token
2. **定期备份DNS**: 使用 `cfm dns export` 导出配置
3. **测试Worker**: 先在测试域名上验证Worker功能
4. **监控账号**: 定期运行 `cfm account info` 检查配置
5. **文档化路由**: 记录你的Worker路由规则

## 安全提示

⚠️ **重要**: 
- 永远不要分享你的API Token
- 配置文件包含敏感信息，权限设置为600
- 使用版本控制时，将 `.cloudflare-manager.yaml` 加入 `.gitignore`
- 定期轮换API Token
- 为不同环境使用不同账号

## 下一步

- 查看完整 [README.md](README.md) 了解所有功能
- 阅读 [Cloudflare Workers文档](https://developers.cloudflare.com/workers/)
- 学习 [DNS最佳实践](https://developers.cloudflare.com/dns/)
- 探索 [Pages部署](https://developers.cloudflare.com/pages/)
