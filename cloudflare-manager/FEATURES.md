# 功能特性详解

## 🎯 核心功能

### 1. 多账号管理

轻松管理多个Cloudflare账号，支持快速切换。

**支持的操作:**
- ✅ 添加多个账号
- ✅ 列出所有账号
- ✅ 切换当前账号
- ✅ 删除账号
- ✅ 查看账号信息

**使用示例:**
```bash
# 添加账号
cfm account add personal --token TOKEN1 --email personal@email.com
cfm account add company --token TOKEN2 --email work@company.com

# 列出账号（带当前账号标记）
cfm account list
# 输出:
# CURRENT  NAME      EMAIL                ACCOUNT_ID
# *        personal  personal@email.com   abc123
#          company   work@company.com     xyz789

# 切换账号
cfm account switch company
# ✓ Switched to account 'company'

# 查看当前账号详情
cfm account info
```

**配置存储:**
- 位置: `~/.cloudflare-manager.yaml`
- 权限: 自动设置为600（仅owner可读写）
- 加密: Token以明文存储，请保护好配置文件

---

### 2. Zone（域名）管理

完整的域名生命周期管理。

**支持的操作:**
- ✅ 创建新域名
- ✅ 列出所有域名
- ✅ 查看域名详情
- ✅ 删除域名
- ✅ 清除缓存

**使用示例:**
```bash
# 创建域名（自动扫描DNS记录）
cfm zone create example.com
# 输出:
# ✓ Zone 'example.com' created successfully
#   Zone ID: 1234567890abcdef
#   Status:  pending
# 
# Nameservers:
#   - alice.ns.cloudflare.com
#   - bob.ns.cloudflare.com
# 
# Update your domain's nameservers to the ones listed above.

# 列出所有域名
cfm zone list
# NAME          ID              STATUS  NAME_SERVERS
# example.com   1234567890ab    active  alice.ns.cloudflare.com ...
# test.com      9876543210fe    active  alice.ns.cloudflare.com ...

# 查看域名详情
cfm zone info example.com
# Zone Information:
#   Name:              example.com
#   ID:                1234567890abcdef
#   Status:            active
#   Plan:              Free Website
#   Development Mode:  ✗
#   Created On:        2024-01-15T10:30:00Z
#   Modified On:       2024-01-15T10:30:00Z
# 
# Nameservers:
#   - alice.ns.cloudflare.com
#   - bob.ns.cloudflare.com

# 清除所有缓存
cfm zone purge example.com --everything

# 清除指定文件缓存
cfm zone purge example.com --files https://example.com/style.css,https://example.com/script.js
```

**Zone状态说明:**
- `pending`: 等待Nameserver更新
- `active`: 已激活，正常运行
- `moved`: 已移动到其他账号
- `deleted`: 已删除

---

### 3. DNS记录管理

全功能DNS记录操作，支持所有记录类型。

**支持的记录类型:**
- A, AAAA, CNAME, MX, TXT, NS, SRV, CAA, PTR, LOC, HTTPS, SVCB, CERT, DNSKEY, DS, NAPTR, SMIMEA, SSHFP, TLSA, URI

**支持的操作:**
- ✅ 列出DNS记录（支持类型过滤）
- ✅ 创建DNS记录
- ✅ 更新DNS记录
- ✅ 删除DNS记录
- ✅ 导出DNS记录（BIND格式）

**使用示例:**
```bash
# 列出所有DNS记录
cfm dns list example.com

# 按类型过滤
cfm dns list example.com --type A

# 创建A记录（启用Cloudflare代理）
cfm dns create example.com A www 192.0.2.1 --proxied
# 输出:
# ✓ DNS record created successfully
#   ID:      abc123
#   Type:    A
#   Name:    www.example.com
#   Content: 192.0.2.1

# 创建CNAME记录
cfm dns create example.com CNAME blog www.example.com

# 创建MX记录
cfm dns create example.com MX @ mail.example.com --priority 10

# 创建TXT记录（SPF）
cfm dns create example.com TXT @ "v=spf1 include:_spf.google.com ~all"

# 设置自定义TTL
cfm dns create example.com A api 192.0.2.2 --ttl 3600

# 更新DNS记录
cfm dns update example.com abc123 192.0.2.3

# 删除DNS记录
cfm dns delete example.com abc123

# 导出DNS配置
cfm dns export example.com > backup.txt
```

**DNS记录选项:**
- `--proxied`: 启用Cloudflare代理（橙色云朵）
- `--ttl`: 设置TTL（1=自动，60-86400秒）
- `--priority`: 设置优先级（MX/SRV记录）

**常见DNS配置模板:**

**Web服务器:**
```bash
cfm dns create example.com A @ 192.0.2.1 --proxied
cfm dns create example.com CNAME www @ --proxied
```

**邮件服务器:**
```bash
cfm dns create example.com MX @ mail.example.com --priority 10
cfm dns create example.com A mail 192.0.2.10
cfm dns create example.com TXT @ "v=spf1 ip4:192.0.2.10 ~all"
cfm dns create example.com TXT _dmarc "v=DMARC1; p=quarantine; rua=mailto:dmarc@example.com"
```

**子域名:**
```bash
cfm dns create example.com A api 192.0.2.20
cfm dns create example.com A admin 192.0.2.30
cfm dns create example.com CNAME dev www.example.com
```

---

### 4. Worker部署与路由

一键部署Cloudflare Workers并配置路由规则。

**支持的操作:**
- ✅ 部署Worker脚本
- ✅ 删除Worker
- ✅ 列出Worker路由
- ✅ 创建路由规则
- ✅ 删除路由规则

**使用示例:**

**1. 部署简单的Worker:**
```bash
# 创建Worker脚本
cat > hello.js << 'EOF'
export default {
  async fetch(request) {
    return new Response('Hello, World!', {
      headers: { 'content-type': 'text/plain' }
    });
  }
};
EOF

# 部署Worker
cfm worker deploy hello-world hello.js
# ✓ Worker 'hello-world' deployed successfully
```

**2. 部署API Worker:**
```bash
cat > api.js << 'EOF'
export default {
  async fetch(request) {
    const url = new URL(request.url);
    
    const routes = {
      '/api/status': { status: 'ok', version: '1.0' },
      '/api/time': { time: new Date().toISOString() },
    };
    
    const data = routes[url.pathname] || { error: 'Not found' };
    
    return new Response(JSON.stringify(data), {
      headers: { 'content-type': 'application/json' }
    });
  }
};
EOF

cfm worker deploy api-worker api.js
```

**3. 配置路由规则:**
```bash
# 将整个域名路由到Worker
cfm worker route create example.com "example.com/*" hello-world

# 将子域名路由到Worker
cfm worker route create example.com "api.example.com/*" api-worker

# 将特定路径路由到Worker
cfm worker route create example.com "example.com/api/*" api-worker
cfm worker route create example.com "example.com/admin/*" admin-worker

# 列出所有路由
cfm worker route list example.com
# PATTERN                 WORKER       ID
# example.com/*           hello-world  route123
# api.example.com/*       api-worker   route456
# example.com/api/*       api-worker   route789

# 删除路由
cfm worker route delete example.com route123
```

**Worker路由模式:**
- `example.com/*` - 匹配所有路径
- `example.com/api/*` - 匹配/api开头的所有路径
- `*.example.com/*` - 匹配所有子域名
- `example.com/static/*.jpg` - 匹配特定文件类型

**Worker开发最佳实践:**
1. 本地测试Worker脚本后再部署
2. 使用环境变量管理配置
3. 为不同环境使用不同的Worker
4. 定期备份Worker脚本
5. 使用版本控制管理Worker代码

---

### 5. Pages项目管理

管理Cloudflare Pages项目和部署。

**支持的操作:**
- ✅ 列出所有Pages项目
- ✅ 查看项目详情
- ✅ 删除项目
- ✅ 列出部署历史
- ✅ 查看部署详情

**使用示例:**
```bash
# 列出所有Pages项目
cfm pages list
# NAME       SUBDOMAIN    DOMAINS  CREATED_ON
# my-blog    my-blog      2        2024-01-15
# docs-site  docs-site    1        2024-01-20

# 查看项目详情
cfm pages info my-blog
# Pages Project Information:
#   Name:         my-blog
#   Subdomain:    my-blog.pages.dev
#   Created On:   2024-01-15 10:30:00
# 
# Domains:
#   - blog.example.com
#   - www.blog.example.com

# 查看部署历史
cfm pages deployment list my-blog
# ID            ENVIRONMENT  STATUS   CREATED_ON
# abc123def456  production   success  2024-01-15 10:30:00
# xyz789uvw012  preview      success  2024-01-14 15:20:00

# 查看部署详情
cfm pages deployment info my-blog abc123def456
# Deployment Information:
#   ID:          abc123def456
#   Environment: production
#   Status:      success
#   URL:         https://abc123def456.my-blog.pages.dev
#   Created On:  2024-01-15 10:30:00

# 删除项目
cfm pages delete old-project
```

---

### 6. 缓存管理

快速清除Zone缓存。

**支持的操作:**
- ✅ 清除所有缓存
- ✅ 清除指定文件缓存

**使用示例:**
```bash
# 清除所有缓存
cfm zone purge example.com --everything
# ✓ Cache purged successfully

# 清除指定文件
cfm zone purge example.com --files https://example.com/style.css,https://example.com/app.js
# ✓ Cache purged successfully
```

**何时清除缓存:**
- 网站更新后内容未刷新
- CSS/JS文件更新但浏览器显示旧版本
- 图片更新但仍显示旧图片
- 需要立即反映内容变更

---

## 🎨 高级功能

### 批量操作

虽然工具本身不直接支持批量操作，但可以通过shell脚本实现：

```bash
# 批量创建DNS记录
for subdomain in api admin cdn media; do
  cfm dns create example.com A $subdomain 192.0.2.100
done

# 批量添加域名
for domain in site1.com site2.com site3.com; do
  cfm zone create $domain
done

# 批量部署Worker到多个路由
for route in "/api/*" "/admin/*" "/dashboard/*"; do
  cfm worker route create example.com "example.com$route" api-worker
done
```

### 自动化脚本

```bash
#!/bin/bash
# 自动配置新域名

DOMAIN=$1
IP=$2

echo "配置域名: $DOMAIN"
echo "IP地址: $IP"

# 1. 创建域名
cfm zone create $DOMAIN

# 2. 等待DNS传播
sleep 10

# 3. 配置基本DNS记录
cfm dns create $DOMAIN A @ $IP --proxied
cfm dns create $DOMAIN CNAME www @ --proxied

# 4. 配置邮件记录
cfm dns create $DOMAIN MX @ mail.$DOMAIN --priority 10
cfm dns create $DOMAIN A mail $IP

# 5. 配置SPF记录
cfm dns create $DOMAIN TXT @ "v=spf1 ip4:$IP ~all"

echo "✓ 域名配置完成！"
```

### 环境变量支持

虽然当前版本不直接支持环境变量，但可以通过wrapper脚本实现：

```bash
#!/bin/bash
# cfm-wrapper.sh

if [ -n "$CLOUDFLARE_ACCOUNT" ]; then
  cfm account switch $CLOUDFLARE_ACCOUNT
fi

cfm "$@"
```

使用:
```bash
export CLOUDFLARE_ACCOUNT=production
./cfm-wrapper.sh zone list
```

---

## 📊 对比其他工具

| 功能 | CFM | Cloudflare CLI | Wrangler | Terraform |
|------|-----|----------------|----------|-----------|
| 多账号管理 | ✅ | ❌ | ❌ | ⚠️ |
| DNS管理 | ✅ | ✅ | ❌ | ✅ |
| Worker部署 | ✅ | ❌ | ✅ | ✅ |
| Worker路由 | ✅ | ❌ | ✅ | ✅ |
| Pages管理 | ✅ | ❌ | ✅ | ✅ |
| 简单易用 | ✅ | ⚠️ | ⚠️ | ❌ |
| 配置文件 | YAML | JSON | TOML | HCL |

**CFM的优势:**
- ✨ 专为多账号场景设计
- 🚀 简单直观的命令行界面
- 🎯 覆盖常用功能，无需学习复杂配置
- 📦 单一二进制文件，无需依赖
- 🔄 快速账号切换

---

## 🔮 未来计划

- [ ] KV命名空间管理
- [ ] R2存储桶管理
- [ ] D1数据库管理
- [ ] 证书管理
- [ ] 防火墙规则管理
- [ ] 分析和日志查询
- [ ] Worker Durable Objects支持
- [ ] 批量操作优化
- [ ] 交互式模式
- [ ] 配置模板系统
- [ ] API限速优化
- [ ] 更多输出格式（JSON, CSV）

---

## 💡 使用技巧

### 1. 命令别名

在 `~/.bashrc` 或 `~/.zshrc` 中添加:
```bash
alias cfm-prod='cfm account switch production && cfm'
alias cfm-dev='cfm account switch development && cfm'
alias cfm-dns='cfm dns'
alias cfm-worker='cfm worker'
```

### 2. 快速切换环境

```bash
# 开发环境操作
cfm account switch dev && cfm zone list

# 生产环境操作
cfm account switch prod && cfm zone list
```

### 3. 命令链式操作

```bash
# 创建域名并立即配置DNS
cfm zone create example.com && \
cfm dns create example.com A @ 192.0.2.1 --proxied && \
cfm dns create example.com CNAME www @ --proxied
```

### 4. 输出重定向

```bash
# 保存DNS配置
cfm dns list example.com > dns-$(date +%Y%m%d).txt

# 导出所有域名列表
cfm zone list > zones.txt
```

### 5. 错误处理

```bash
# 检查命令是否成功
if cfm zone create example.com; then
  echo "域名创建成功"
  cfm dns create example.com A @ 192.0.2.1
else
  echo "域名创建失败"
  exit 1
fi
```
