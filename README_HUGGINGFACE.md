---
title: Cloudflare Manager
emoji: ☁️
colorFrom: blue
colorTo: orange
sdk: docker
pinned: false
license: mit
app_port: 7860
---

# Cloudflare Multi-Account Manager 🚀

A comprehensive web-based management platform for Cloudflare services with multi-account support.

## Features

### Core Services
- 🔐 **Multi-Account Management** - Manage multiple Cloudflare accounts
- 🌐 **Zone Management** - Create, configure, and delete domains
- 📝 **DNS Records** - Full CRUD operations for all DNS record types
- ⚡ **Workers** - Deploy and manage Cloudflare Workers
- 📄 **Pages** - Manage Cloudflare Pages projects
- 💾 **KV Storage** - Workers KV namespace and key-value management
- 🗄️ **R2 Storage** - Object storage bucket management

### Technology Stack
- **Frontend**: React + Ant Design
- **Backend**: FastAPI (Python)
- **CLI**: Go
- **Deployment**: Docker

## Quick Start

### Using on Hugging Face Spaces

1. Click the "Deploy" button above
2. Wait for the Docker container to build and start
3. Access the web interface at the provided URL
4. Add your Cloudflare API token in the Accounts section

### Local Development

```bash
# Clone the repository
git clone <your-repo>
cd cloudflare-manager

# Build and run with Docker
docker build -t cloudflare-manager .
docker run -p 7860:7860 cloudflare-manager
```

Open http://localhost:7860 in your browser.

## Configuration

### Environment Variables

- `CLOUDFLARE_EMAIL`: Your Cloudflare account email
- `CLOUDFLARE_TOKEN`: Your Cloudflare API token

### Getting API Token

1. Go to https://dash.cloudflare.com/profile/api-tokens
2. Click "Create Token"
3. Use "Edit zone DNS" template or create custom permissions
4. Copy the token and add it in the Accounts page

### Recommended Permissions

```
Account:
  - Account Settings: Read
Zone:
  - Zone: Read, Edit
  - DNS: Edit
  - Workers Routes: Edit
  - Workers Scripts: Edit
```

## API Documentation

FastAPI provides interactive API documentation:
- Swagger UI: http://localhost:7860/api/docs
- ReDoc: http://localhost:7860/api/redoc

## Features Overview

### 1. Accounts
- Add multiple Cloudflare accounts
- Switch between accounts instantly
- Manage API tokens securely

### 2. Zones
- Create new zones
- View zone statistics
- Purge cache
- Delete zones

### 3. DNS Records
- Create A, AAAA, CNAME, MX, TXT, and more
- Enable Cloudflare proxy (orange cloud)
- Bulk operations support
- Export/Import DNS records

### 4. Workers
- Deploy JavaScript workers
- Manage worker routes
- View worker logs
- Configure worker settings

### 5. Pages
- List all Pages projects
- View deployment history
- Manage custom domains
- Rollback deployments

### 6. KV Storage
- Create KV namespaces
- Store key-value pairs
- Bulk operations
- Search by prefix

### 7. R2 Storage
- Create R2 buckets
- Configure locations
- Manage bucket settings
- Monitor usage

## Security Notes

⚠️ **Important Security Considerations:**

1. **API Tokens**: Keep your API tokens secure and never commit them to git
2. **Permissions**: Use tokens with minimal required permissions
3. **Access Control**: This is a single-user application. Do not expose publicly without authentication
4. **HTTPS**: Always use HTTPS in production
5. **Token Rotation**: Regularly rotate your API tokens

## Architecture

```
┌─────────────────┐
│  React Frontend │  (Port 7860)
│   (Ant Design)  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  FastAPI        │  (Python Backend)
│  REST API       │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Go CLI Tool    │  (Cloudflare SDK)
│  (cfm binary)   │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Cloudflare API │
└─────────────────┘
```

## Troubleshooting

### Cannot connect to Cloudflare API
- Check your API token is valid
- Verify token permissions
- Check network connectivity

### Frontend not loading
- Ensure port 7860 is accessible
- Check Docker logs: `docker logs <container-id>`
- Verify frontend build completed successfully

### Worker deployment fails
- Ensure worker script is valid JavaScript
- Check account has Workers enabled
- Verify API token has Workers permissions

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details

## Support

For issues and questions:
- GitHub Issues: [Your Repo]
- Documentation: See README files in subdirectories

## Acknowledgments

- [Cloudflare](https://www.cloudflare.com/) for the amazing platform
- [FastAPI](https://fastapi.tiangolo.com/) for the web framework
- [React](https://reactjs.org/) for the frontend library
- [Ant Design](https://ant.design/) for the UI components
- [Go](https://golang.org/) for the CLI tool

---

Made with ❤️ for the Cloudflare community
