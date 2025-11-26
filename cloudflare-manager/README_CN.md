# Cloudflare 多账号管理器 (CFM)

<div align="center">

🚀 **功能强大的Cloudflare多账号CLI管理工具**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

[English](README.md) | 简体中文

</div>

## ✨ 功能特性

- 🔐 **多账号管理** - 轻松管理多个Cloudflare账号，一键切换
- 🌐 **Zone管理** - 创建、查询、删除域名配置
- 📝 **DNS记录** - 完整的DNS CRUD操作，支持所有记录类型
- ⚡ **Worker部署** - 一键部署Cloudflare Workers脚本
- 🔗 **Worker路由** - 将域名路径路由到Worker
- 📄 **Pages管理** - 管理Cloudflare Pages项目和部署
- 🗑️ **缓存管理** - 快速清除Zone缓存
- 🎯 **简单易用** - 清晰的命令行界面，符合直觉

## 📦 安装

### 方式1: 从源码编译

```bash
git clone https://github.com/your-username/cloudflare-manager.git
cd cloudflare-manager
go build -o cfm
sudo mv cfm /usr/local/bin/
```

### 方式2: 使用Go安装

```bash
go install github.com/cloudflare-manager@latest
```

### 方式3: 下载预编译二进制

访问 [Releases页面](https://github.com/your-username/cloudflare-manager/releases) 下载对应平台的二进制文件。

## 🚀 快速开始

### 1. 获取API Token

1. 访问 [Cloudflare Dashboard](https://dash.cloudflare.com/profile/api-tokens)
2. 点击 **Create Token**
3. 选择模板或自定义权限：
   - ✅ Account > Account Settings > Read
   - ✅ Zone > Zone > Edit
   - ✅ Zone > DNS > Edit
   - ✅ Account > Workers Scripts > Edit
   - ✅ Account > Pages > Edit
4. 创建并复制Token

### 2. 添加账号

```bash
cfm account add myaccount --token YOUR_API_TOKEN --email your@email.com
```

### 3. 开始使用

```bash
# 查看账号
cfm account list

# 列出域名
cfm zone list

# 创建DNS记录
cfm dns create example.com A www 192.0.2.1 --proxied
```

## 📖 使用指南

### 账号管理

```bash
# 添加账号
cfm account add production --token YOUR_TOKEN --email admin@example.com

# 列出所有账号
cfm account list

# 切换账号
cfm account switch production

# 删除账号
cfm account remove old-account

# 查看当前账号信息
cfm account info
```

### Zone（域名）管理

```bash
# 列出所有域名
cfm zone list

# 创建新域名
cfm zone create example.com

# 查看域名详情
cfm zone info example.com

# 删除域名
cfm zone delete example.com

# 清除缓存
cfm zone purge example.com --everything
```

### DNS记录管理

```bash
# 列出DNS记录
cfm dns list example.com

# 按类型过滤
cfm dns list example.com --type A

# 创建A记录
cfm dns create example.com A www 192.0.2.1 --proxied

# 创建CNAME记录
cfm dns create example.com CNAME blog www.example.com

# 创建MX记录
cfm dns create example.com MX @ mail.example.com --priority 10

# 更新DNS记录
cfm dns update example.com RECORD_ID 192.0.2.2

# 删除DNS记录
cfm dns delete example.com RECORD_ID

# 导出DNS配置
cfm dns export example.com > dns-backup.txt
```

### Worker管理

```bash
# 部署Worker
cfm worker deploy my-worker ./worker.js

# 删除Worker
cfm worker delete my-worker

# 列出所有路由
cfm worker route list example.com

# 创建路由（将域名路由到Worker）
cfm worker route create example.com "example.com/*" my-worker
cfm worker route create example.com "api.example.com/*" api-worker

# 删除路由
cfm worker route delete example.com ROUTE_ID
```

### Pages项目管理

```bash
# 列出所有Pages项目
cfm pages list

# 查看项目信息
cfm pages info my-blog

# 查看部署历史
cfm pages deployment list my-blog

# 查看部署详情
cfm pages deployment info my-blog DEPLOYMENT_ID

# 删除项目
cfm pages delete my-blog
```

## 🎯 实战案例

### 案例1: 完整配置新域名

```bash
# 1. 切换到生产账号
cfm account switch production

# 2. 添加域名到Cloudflare
cfm zone create myapp.com

# 3. 配置DNS记录
cfm dns create myapp.com A @ 192.0.2.1 --proxied
cfm dns create myapp.com CNAME www @ --proxied
cfm dns create myapp.com MX @ mail.myapp.com --priority 10
cfm dns create myapp.com TXT @ "v=spf1 include:_spf.google.com ~all"

# 4. 验证配置
cfm dns list myapp.com
```

### 案例2: 部署API Worker

```bash
# 1. 创建Worker脚本 (api-worker.js)
cat > api-worker.js << 'EOF'
export default {
  async fetch(request) {
    return new Response(JSON.stringify({
      status: 'ok',
      message: 'API is running'
    }), {
      headers: { 'content-type': 'application/json' }
    });
  }
};
EOF

# 2. 部署Worker
cfm worker deploy api-worker ./api-worker.js

# 3. 配置路由
cfm worker route create myapp.com "api.myapp.com/*" api-worker

# 4. 测试
curl https://api.myapp.com/
```

### 案例3: 多环境管理

```bash
# 配置开发环境
cfm account add dev --token DEV_TOKEN
cfm account switch dev
cfm zone create dev.myapp.com
cfm dns create dev.myapp.com A @ 192.0.2.10

# 配置生产环境
cfm account add prod --token PROD_TOKEN
cfm account switch prod
cfm zone create myapp.com
cfm dns create myapp.com A @ 192.0.2.20

# 快速切换
cfm account list
cfm account switch dev
cfm account switch prod
```

## 🏗️ Worker示例

项目包含一个功能完整的Worker示例 `example-worker.js`:

```javascript
export default {
  async fetch(request, env, ctx) {
    const url = new URL(request.url);
    
    switch (url.pathname) {
      case '/':
        return new Response('Hello from Cloudflare Worker!');
      case '/api/hello':
        return new Response(JSON.stringify({
          message: 'Hello, API!',
          timestamp: new Date().toISOString()
        }), {
          headers: { 'content-type': 'application/json' }
        });
      default:
        return new Response('404 Not Found', { status: 404 });
    }
  }
};
```

部署：

```bash
cfm worker deploy hello ./example-worker.js
cfm worker route create example.com "example.com/*" hello
```

## ⚙️ 配置文件

配置文件保存在 `~/.cloudflare-manager.yaml`:

```yaml
current_account: production
accounts:
  - name: production
    api_token: ********************************
    account_id: abc123def456
    email: admin@example.com
  - name: development
    api_token: ********************************
    account_id: xyz789uvw012
    email: dev@example.com
```

## 🔐 安全建议

1. **使用最小权限原则**: 只授予必要的权限
2. **定期轮换Token**: 建议每90天更换一次API Token
3. **保护配置文件**: 配置文件权限已自动设置为600
4. **不要提交Token**: 将`.cloudflare-manager.yaml`加入`.gitignore`
5. **分离环境**: 为开发/生产环境使用不同的账号和Token

## 📊 命令速查表

| 功能 | 命令 |
|------|------|
| 添加账号 | `cfm account add <name> --token <token>` |
| 切换账号 | `cfm account switch <name>` |
| 列出域名 | `cfm zone list` |
| 创建域名 | `cfm zone create <domain>` |
| 创建DNS | `cfm dns create <domain> <type> <name> <content>` |
| 列出DNS | `cfm dns list <domain>` |
| 部署Worker | `cfm worker deploy <name> <file>` |
| 创建路由 | `cfm worker route create <domain> <pattern> <worker>` |
| 列出Pages | `cfm pages list` |

## 🛠️ 技术栈

- **语言**: Go 1.21+
- **SDK**: [cloudflare-go](https://github.com/cloudflare/cloudflare-go)
- **CLI框架**: [Cobra](https://github.com/spf13/cobra)
- **配置**: YAML

## 🤝 贡献

欢迎贡献！请随时提交Issue或Pull Request。

1. Fork项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 📝 常见问题

<details>
<summary><b>Q: 如何获取API Token?</b></summary>

访问 https://dash.cloudflare.com/profile/api-tokens 创建Token。推荐使用"Edit zone DNS"模板或自定义权限。
</details>

<details>
<summary><b>Q: 支持哪些DNS记录类型?</b></summary>

支持所有Cloudflare DNS记录类型：A、AAAA、CNAME、MX、TXT、NS、SRV、CAA、PTR等。
</details>

<details>
<summary><b>Q: Worker路由规则语法?</b></summary>

- `example.com/*` - 匹配所有路径
- `example.com/api/*` - 匹配/api开头的路径
- `*.example.com/*` - 匹配所有子域名
</details>

<details>
<summary><b>Q: 如何备份DNS配置?</b></summary>

```bash
cfm dns export example.com > dns-backup.txt
```
</details>

<details>
<summary><b>Q: 支持批量操作吗?</b></summary>

当前版本主要支持单条记录操作。批量操作可以通过shell脚本实现：

```bash
for domain in domain1.com domain2.com domain3.com; do
  cfm zone create $domain
done
```
</details>

## 📜 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 🔗 相关链接

- [Cloudflare API文档](https://developers.cloudflare.com/api/)
- [Cloudflare Workers文档](https://developers.cloudflare.com/workers/)
- [Cloudflare Pages文档](https://developers.cloudflare.com/pages/)
- [DNS最佳实践](https://developers.cloudflare.com/dns/manage-dns-records/reference/dns-record-types/)

## ⭐ Star History

如果这个项目对你有帮助，请给个Star！⭐

---

<div align="center">
Made with ❤️ for the Cloudflare community
</div>
