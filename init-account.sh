#!/bin/bash
# 快速初始化Cloudflare账号脚本

set -e

echo "🚀 Cloudflare Account Initialization"
echo "===================================="
echo

# 提供的凭证
EMAIL="tqa88tawlq@downnaturer.me"
TOKEN="4e2dd4818267ebd2ab8d1aa2e7f9bf4151b70"

# 检查cfm是否存在
CFM_PATH="./cloudflare-manager/cfm"
if [ ! -f "$CFM_PATH" ]; then
    echo "❌ Error: cfm binary not found at $CFM_PATH"
    echo "Please build first: cd cloudflare-manager && go build -o cfm"
    exit 1
fi

echo "📝 Adding default account..."
$CFM_PATH account add default --token "$TOKEN" --email "$EMAIL"

echo
echo "✅ Account initialized successfully!"
echo

echo "📊 Account information:"
$CFM_PATH account info

echo
echo "🎯 Next steps:"
echo "  1. Test CLI: $CFM_PATH zone list"
echo "  2. Start backend: cd backend && python main.py"
echo "  3. Build frontend: cd frontend && npm install && npm start"
echo "  4. Or use Docker: docker build -t cfm . && docker run -p 7860:7860 cfm"
echo

echo "🌐 Available services:"
echo "  - Zones: $CFM_PATH zone --help"
echo "  - DNS: $CFM_PATH dns --help"
echo "  - Workers: $CFM_PATH worker --help"
echo "  - KV: $CFM_PATH kv --help"
echo "  - R2: $CFM_PATH r2 --help"
echo "  - Pages: $CFM_PATH pages --help"
