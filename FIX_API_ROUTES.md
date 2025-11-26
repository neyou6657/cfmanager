# API 路由 404 错误修复说明

## 问题描述

用户报告 `POST /api/auth` 返回 404 Not Found 错误，尽管后端代码中确实定义了该路由。

## 根本原因

在 `docker-entrypoint.sh` 文件中，FastAPI 应用被挂载到 `/api` 路径下：

```python
app.mount('/api', api_app)
```

这意味着所有在 `backend/main.py` 中定义的路由都会自动添加 `/api` 前缀。

之前的路由定义：
```python
@app.post("/api/auth")  # 实际路径变成: /api/api/auth
```

当应用被挂载到 `/api` 后，路由路径会变成 `/api/api/auth`，导致访问 `/api/auth` 时返回 404。

## 解决方案

从 `backend/main.py` 中的所有路由定义中移除 `/api` 前缀。

修改后：
```python
@app.post("/auth")  # 挂载后实际路径: /api/auth ✓
```

## 修改的路由

所有 API 路由的前缀都已从 `/api/xxx` 改为 `/xxx`：

- `/api/auth` → `/auth`
- `/api/accounts` → `/accounts`
- `/api/zones` → `/zones`
- `/api/workers` → `/workers`
- `/api/pages` → `/pages`
- `/api/kv/namespaces` → `/kv/namespaces`
- `/api/r2/buckets` → `/r2/buckets`
- 以及所有相关的子路由

## 测试结果

修复后，所有路由均可正常访问：
- ✅ POST /api/auth - 401 (路由找到，认证失败是预期行为)
- ✅ GET /api/accounts - 401 (路由找到)
- ✅ GET /api/zones - 401 (路由找到)
- ✅ GET /api/workers - 401 (路由找到)
- ✅ GET /api/pages - 401 (路由找到)
- ✅ GET /api/kv/namespaces - 401 (路由找到)
- ✅ GET /api/r2/buckets - 401 (路由找到)

注意：返回 401 状态码是因为未提供有效的认证信息，这是正常行为，说明路由已正确配置。

## FastAPI 最佳实践

在使用 `app.mount()` 时，应该：
1. 在子应用中定义路由时不包含挂载前缀
2. 在挂载时统一指定前缀

这样可以使子应用更具可移植性和可重用性。
