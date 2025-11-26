# 安装指南

## 系统要求

- **操作系统**: Linux, macOS, Windows
- **Go版本**: 1.21 或更高（仅编译时需要）
- **网络**: 需要访问Cloudflare API

## 安装方式

### 方式 1: 使用预编译二进制（推荐）

如果项目已包含编译好的二进制文件：

```bash
# Linux/macOS
chmod +x cfm
sudo mv cfm /usr/local/bin/

# 或者添加到PATH
mkdir -p ~/bin
mv cfm ~/bin/
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### 方式 2: 从源码编译

#### 安装Go

**Linux/macOS:**
```bash
# 使用包管理器（推荐）
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# macOS (Homebrew)
brew install go

# 或者手动安装
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
```

**Windows:**
1. 下载安装包: https://go.dev/dl/
2. 运行安装程序
3. 重启命令提示符

#### 编译项目

```bash
# 1. 克隆或进入项目目录
cd cloudflare-manager

# 2. 下载依赖
go mod tidy

# 3. 编译
go build -o cfm

# 4. 验证编译
./cfm --version

# 5. 安装到系统（可选）
# Linux/macOS
sudo mv cfm /usr/local/bin/

# Windows
# 将cfm.exe移动到PATH中的目录
```

### 方式 3: 使用Go安装

```bash
# 直接从源码安装（需要先配置好Git和Go）
go install github.com/your-username/cloudflare-manager@latest

# 程序会被安装到 $GOPATH/bin/cloudflare-manager
# 确保 $GOPATH/bin 在你的 PATH 中
```

## 配置

### 1. 获取API Token

1. 访问 [Cloudflare Dashboard](https://dash.cloudflare.com/profile/api-tokens)
2. 点击 **Create Token**
3. 选择权限模板或自定义：

**推荐权限设置:**
```
Account:
  - Account Settings: Read
  - Account Firewall Access Rules: Read

Zone:
  - Zone: Read
  - Zone: Edit
  - DNS: Edit
  - Workers Routes: Edit

User:
  - User Details: Read
```

4. 创建Token并复制

### 2. 添加账号

```bash
# 基本用法
cfm account add myaccount --token YOUR_API_TOKEN

# 带邮箱（可选）
cfm account add myaccount --token YOUR_API_TOKEN --email your@email.com

# 验证
cfm account list
cfm account info
```

### 3. 配置文件

配置会自动保存到 `~/.cloudflare-manager.yaml`:

```yaml
current_account: myaccount
accounts:
  - name: myaccount
    api_token: your_api_token_here
    account_id: auto_detected
    email: your@email.com
```

**⚠️ 安全提示:**
- 配置文件权限会自动设置为 600
- 不要将配置文件提交到Git
- 定期更新API Token
- 为不同用途创建不同的Token

## 验证安装

```bash
# 1. 检查版本
cfm --version
# 输出: cfm version 1.0.0

# 2. 查看帮助
cfm --help

# 3. 测试账号
cfm account list

# 4. 测试API连接
cfm account info
# 应该显示你的账号信息

# 5. 测试基本功能
cfm zone list
```

## 常见安装问题

### 问题 1: "command not found: cfm"

**原因**: 程序不在PATH中

**解决:**
```bash
# 方案1: 移动到PATH目录
sudo mv cfm /usr/local/bin/

# 方案2: 添加当前目录到PATH
export PATH=$PATH:$(pwd)

# 方案3: 使用完整路径
/path/to/cfm --version
```

### 问题 2: "permission denied"

**原因**: 文件没有执行权限

**解决:**
```bash
chmod +x cfm
./cfm --version
```

### 问题 3: 编译失败 "go: command not found"

**原因**: Go未安装或不在PATH中

**解决:**
```bash
# 检查Go是否安装
which go

# 如果没有，按照上面的步骤安装Go
```

### 问题 4: "failed to load config"

**原因**: 配置文件权限或格式问题

**解决:**
```bash
# 检查配置文件
ls -la ~/.cloudflare-manager.yaml

# 修复权限
chmod 600 ~/.cloudflare-manager.yaml

# 重新创建配置
rm ~/.cloudflare-manager.yaml
cfm account add myaccount --token YOUR_TOKEN
```

### 问题 5: "failed to get account ID"

**原因**: Token权限不足或无效

**解决:**
1. 检查Token是否有 "Account > Account Settings > Read" 权限
2. 在Cloudflare Dashboard验证Token是否有效
3. 重新创建Token并更新配置

## 卸载

### 删除程序

```bash
# 如果安装在 /usr/local/bin
sudo rm /usr/local/bin/cfm

# 如果在其他位置
which cfm  # 查找位置
rm $(which cfm)
```

### 删除配置

```bash
# 删除配置文件
rm ~/.cloudflare-manager.yaml

# 删除示例脚本生成的文件（如果有）
rm -f dns-*.txt zones.txt
```

## 升级

### 从源码升级

```bash
cd cloudflare-manager
git pull
go build -o cfm
sudo mv cfm /usr/local/bin/
```

### 使用Go升级

```bash
go install github.com/your-username/cloudflare-manager@latest
```

## Docker使用（可选）

如果不想安装Go，可以使用Docker：

```bash
# 构建镜像
docker build -t cfm .

# 使用别名
alias cfm='docker run --rm -v $HOME/.cloudflare-manager.yaml:/root/.cloudflare-manager.yaml cfm'

# 使用
cfm --version
cfm zone list
```

## 下一步

安装完成后：

1. 阅读 [QUICKSTART.md](QUICKSTART.md) 快速开始
2. 查看 [README_CN.md](README_CN.md) 了解所有功能
3. 运行 `./demo.sh` 查看使用演示
4. 查看 [FEATURES.md](FEATURES.md) 了解详细功能

## 获取帮助

- 📖 文档: 查看本仓库的Markdown文件
- 🐛 问题: 在GitHub Issues提交
- 💬 讨论: 在GitHub Discussions参与
- 📧 联系: your@email.com

---

**祝你使用愉快！** 🎉
