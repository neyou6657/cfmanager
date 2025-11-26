#!/bin/bash
# Cloudflare Manager Demo Script
# 这个脚本演示如何使用CFM工具

set -e

CFM="./cfm"
DEMO_DOMAIN="example.com"
WORKER_NAME="demo-worker"

echo "======================================"
echo "Cloudflare Manager (CFM) 演示脚本"
echo "======================================"
echo

# 检查cfm是否存在
if [ ! -f "$CFM" ]; then
    echo "错误: 找不到cfm可执行文件"
    echo "请先运行: go build -o cfm"
    exit 1
fi

echo "📋 步骤 1: 查看版本"
echo "$ $CFM --version"
$CFM --version
echo

echo "📋 步骤 2: 查看帮助"
echo "$ $CFM --help"
$CFM --help | head -20
echo "... (更多内容请运行 '$CFM --help')"
echo

echo "📋 步骤 3: 账号管理"
echo "提示: 使用你的Cloudflare API Token"
echo "$ $CFM account add myaccount --token YOUR_TOKEN"
echo
echo "列出账号:"
echo "$ $CFM account list"
echo
echo "切换账号:"
echo "$ $CFM account switch myaccount"
echo

echo "📋 步骤 4: Zone管理"
echo "创建域名:"
echo "$ $CFM zone create $DEMO_DOMAIN"
echo
echo "列出所有域名:"
echo "$ $CFM zone list"
echo
echo "查看域名详情:"
echo "$ $CFM zone info $DEMO_DOMAIN"
echo

echo "📋 步骤 5: DNS记录管理"
echo "创建A记录:"
echo "$ $CFM dns create $DEMO_DOMAIN A www 192.0.2.1 --proxied"
echo
echo "创建CNAME记录:"
echo "$ $CFM dns create $DEMO_DOMAIN CNAME blog www.$DEMO_DOMAIN"
echo
echo "列出DNS记录:"
echo "$ $CFM dns list $DEMO_DOMAIN"
echo
echo "按类型过滤:"
echo "$ $CFM dns list $DEMO_DOMAIN --type A"
echo

echo "📋 步骤 6: Worker部署"
echo "部署Worker:"
echo "$ $CFM worker deploy $WORKER_NAME example-worker.js"
echo
echo "创建路由:"
echo "$ $CFM worker route create $DEMO_DOMAIN \"$DEMO_DOMAIN/*\" $WORKER_NAME"
echo
echo "列出路由:"
echo "$ $CFM worker route list $DEMO_DOMAIN"
echo

echo "📋 步骤 7: Pages管理"
echo "列出Pages项目:"
echo "$ $CFM pages list"
echo
echo "查看项目详情:"
echo "$ $CFM pages info my-project"
echo
echo "查看部署历史:"
echo "$ $CFM pages deployment list my-project"
echo

echo "📋 步骤 8: 缓存管理"
echo "清除所有缓存:"
echo "$ $CFM zone purge $DEMO_DOMAIN --everything"
echo
echo "清除指定文件:"
echo "$ $CFM zone purge $DEMO_DOMAIN --files https://$DEMO_DOMAIN/style.css"
echo

echo "======================================"
echo "✨ 演示完成！"
echo "======================================"
echo
echo "🚀 快速开始:"
echo "  1. 获取API Token: https://dash.cloudflare.com/profile/api-tokens"
echo "  2. 添加账号: $CFM account add myaccount --token YOUR_TOKEN"
echo "  3. 查看文档: cat README_CN.md"
echo "  4. 快速指南: cat QUICKSTART.md"
echo
echo "💡 提示:"
echo "  - 运行 '$CFM command --help' 获取详细帮助"
echo "  - 配置文件位于: ~/.cloudflare-manager.yaml"
echo "  - 示例Worker: example-worker.js"
echo
echo "📚 完整文档:"
echo "  - README_CN.md     - 中文完整文档"
echo "  - QUICKSTART.md    - 快速开始指南"
echo "  - FEATURES.md      - 功能详解"
echo "  - PROJECT_SUMMARY.md - 项目总结"
echo
