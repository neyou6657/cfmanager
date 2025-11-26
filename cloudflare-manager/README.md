# Cloudflare Multi-Account Manager

一个功能强大的Cloudflare多账号管理CLI工具，支持域名托管、DNS管理、Worker部署和路由等功能。

## 功能特性

- ✅ **多账号管理** - 支持添加、切换多个Cloudflare账号
- ✅ **Zone/域名管理** - 创建、列表、删除、查询域名信息
- ✅ **DNS记录管理** - 完整的DNS记录CRUD操作
- ✅ **Worker部署** - 部署、管理Cloudflare Workers
- ✅ **Worker路由** - 将域名路由到Worker
- ✅ **Pages项目管理** - 管理Cloudflare Pages项目和部署
- ✅ **缓存管理** - 清除Zone缓存
- ✅ **域名绑定** - 为Pages项目绑定自定义域名

## 安装

### 从源码编译

```bash
git clone https://github.com/your-username/cloudflare-manager.git
cd cloudflare-manager
go build -o cfm
sudo mv cfm /usr/local/bin/
```

### 使用Go安装

```bash
go install github.com/cloudflare-manager@latest
```

## 快速开始

### 1. 添加Cloudflare账号

```bash
cfm account add myaccount --token YOUR_API_TOKEN --email your@email.com
```

API Token可以在Cloudflare Dashboard创建：
https://dash.cloudflare.com/profile/api-tokens

推荐的权限：
- Account > Account Settings > Read
- Zone > Zone > Edit
- Zone > DNS > Edit
- Account > Workers Scripts > Edit
- Account > Pages > Edit

### 2. 查看账号列表

```bash
cfm account list
```

### 3. 切换账号

```bash
cfm account switch myaccount
```

## 使用指南

### 账号管理

```bash
# 添加账号
cfm account add <name> --token <api-token> [--email <email>]

# 列出所有账号
cfm account list

# 切换账号
cfm account switch <name>

# 删除账号
cfm account remove <name>

# 查看当前账号信息
cfm account info
```

### Zone/域名管理

```bash
# 列出所有域名
cfm zone list

# 创建新域名
cfm zone create example.com [--jump-start]

# 查看域名信息
cfm zone info example.com

# 删除域名
cfm zone delete example.com

# 清除缓存
cfm zone purge example.com [--everything] [--files file1,file2]
```

### DNS记录管理

```bash
# 列出DNS记录
cfm dns list example.com [--type A]

# 创建DNS记录
cfm dns create example.com A www 1.2.3.4 [--ttl 3600] [--proxied]

# 更新DNS记录
cfm dns update example.com <record-id> 5.6.7.8 [--ttl 3600]

# 删除DNS记录
cfm dns delete example.com <record-id>

# 导出DNS记录（BIND格式）
cfm dns export example.com

# 导入DNS记录
cfm dns import example.com /path/to/bind-file.txt
```

### Worker管理

```bash
# 列出所有Workers
cfm worker list

# 部署Worker
cfm worker deploy my-worker /path/to/worker.js

# 查看Worker脚本
cfm worker get my-worker

# 删除Worker
cfm worker delete my-worker

# 管理Workers subdomain
cfm worker subdomain get
cfm worker subdomain set mycompany
```

### Worker路由

```bash
# 列出路由规则
cfm worker route list example.com

# 创建路由规则（将域名路由到Worker）
cfm worker route create example.com "example.com/*" my-worker
cfm worker route create example.com "api.example.com/*" api-worker

# 删除路由规则
cfm worker route delete example.com <route-id>
```

### Pages项目管理

```bash
# 列出所有Pages项目
cfm pages list

# 查看项目信息
cfm pages info my-project

# 删除项目
cfm pages delete my-project

# 列出部署
cfm pages deployment list my-project

# 查看部署信息
cfm pages deployment info my-project <deployment-id>

# 管理自定义域名
cfm pages domain list my-project
cfm pages domain add my-project www.example.com
cfm pages domain delete my-project www.example.com
```

## 完整工作流示例

### 示例1: 托管域名并部署Worker

```bash
# 1. 添加Cloudflare账号
cfm account add production --token <your-token>

# 2. 创建域名
cfm zone create example.com

# 3. 添加DNS记录
cfm dns create example.com A @ 192.0.2.1 --proxied
cfm dns create example.com CNAME www @ --proxied

# 4. 部署Worker
cfm worker deploy hello-worker ./worker.js

# 5. 将域名路由到Worker
cfm worker route create example.com "example.com/api/*" hello-worker

# 6. 验证配置
cfm dns list example.com
cfm worker route list example.com
```

### 示例2: 管理多个账号

```bash
# 添加多个账号
cfm account add personal --token <token1>
cfm account add company --token <token2>

# 查看所有账号
cfm account list

# 在个人账号上操作
cfm account switch personal
cfm zone list

# 切换到公司账号
cfm account switch company
cfm worker list
```

### 示例3: Pages项目绑定域名

```bash
# 1. 为Pages项目添加自定义域名
cfm pages domain add my-blog blog.example.com

# 2. 创建DNS记录指向Pages
cfm dns create example.com CNAME blog <project>.pages.dev --proxied

# 3. 验证域名状态
cfm pages domain list my-blog
```

## Worker脚本示例

创建一个简单的Worker脚本 `worker.js`:

```javascript
export default {
  async fetch(request, env, ctx) {
    return new Response('Hello from Cloudflare Worker!', {
      headers: { 'content-type': 'text/plain' },
    });
  },
};
```

部署：

```bash
cfm worker deploy hello ./worker.js
cfm worker route create example.com "example.com/*" hello
```

## 配置文件

配置文件位于 `~/.cloudflare-manager.yaml`:

```yaml
current_account: myaccount
accounts:
  - name: myaccount
    api_token: your_api_token_here
    account_id: your_account_id
    email: your@email.com
  - name: anotheraccount
    api_token: another_token
    account_id: another_account_id
    email: another@email.com
```

## 常见问题

### 如何获取API Token?

1. 访问 https://dash.cloudflare.com/profile/api-tokens
2. 点击 "Create Token"
3. 选择 "Edit zone DNS" 模板或自定义权限
4. 复制生成的Token

### Worker路由规则说明

路由规则支持通配符：
- `example.com/*` - 匹配所有路径
- `example.com/api/*` - 只匹配 /api 路径
- `*.example.com/*` - 匹配所有子域名

### DNS记录类型

支持的记录类型：A, AAAA, CNAME, MX, TXT, SRV, CAA, NS, PTR等

## 技术栈

- Go 1.21+
- cloudflare-go SDK
- cobra (CLI框架)
- yaml (配置管理)

## 贡献

欢迎提交Issue和Pull Request！

## 许可证

MIT License

## 相关链接

- [Cloudflare API文档](https://developers.cloudflare.com/api/)
- [Cloudflare Workers文档](https://developers.cloudflare.com/workers/)
- [Cloudflare Pages文档](https://developers.cloudflare.com/pages/)
