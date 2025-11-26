// Cloudflare Worker示例脚本
// 这是一个功能完整的Worker，展示了常见的使用场景

export default {
  async fetch(request, env, ctx) {
    const url = new URL(request.url);
    
    // 路由处理
    switch (url.pathname) {
      case '/':
        return handleRoot(request);
      case '/api/hello':
        return handleAPI(request);
      case '/json':
        return handleJSON(request);
      case '/redirect':
        return Response.redirect('https://cloudflare.com', 301);
      case '/proxy':
        return handleProxy(request);
      default:
        return new Response('404 Not Found', { status: 404 });
    }
  },
};

// 处理根路径
function handleRoot(request) {
  const html = `
<!DOCTYPE html>
<html>
<head>
  <title>Cloudflare Worker Demo</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      max-width: 800px;
      margin: 50px auto;
      padding: 20px;
      line-height: 1.6;
    }
    h1 { color: #f38020; }
    code { 
      background: #f4f4f4; 
      padding: 2px 5px; 
      border-radius: 3px; 
    }
    .endpoint {
      background: #f9f9f9;
      padding: 15px;
      margin: 10px 0;
      border-left: 4px solid #f38020;
    }
  </style>
</head>
<body>
  <h1>🚀 Cloudflare Worker Demo</h1>
  <p>这是一个运行在Cloudflare边缘网络上的Worker示例。</p>
  
  <h2>可用的端点：</h2>
  
  <div class="endpoint">
    <h3>GET /</h3>
    <p>显示此页面</p>
  </div>
  
  <div class="endpoint">
    <h3>GET /api/hello</h3>
    <p>返回JSON格式的问候消息</p>
  </div>
  
  <div class="endpoint">
    <h3>GET /json</h3>
    <p>返回请求信息的JSON</p>
  </div>
  
  <div class="endpoint">
    <h3>GET /redirect</h3>
    <p>重定向到Cloudflare官网</p>
  </div>
  
  <div class="endpoint">
    <h3>GET /proxy</h3>
    <p>反向代理示例</p>
  </div>
  
  <hr>
  <p>
    <strong>请求信息：</strong><br>
    IP: ${request.headers.get('cf-connecting-ip')}<br>
    Country: ${request.cf?.country || 'Unknown'}<br>
    User-Agent: ${request.headers.get('user-agent')}
  </p>
</body>
</html>
  `;
  
  return new Response(html, {
    headers: { 
      'content-type': 'text/html;charset=UTF-8',
      'cache-control': 'public, max-age=3600'
    },
  });
}

// 处理API请求
function handleAPI(request) {
  const data = {
    message: 'Hello from Cloudflare Worker!',
    timestamp: new Date().toISOString(),
    ip: request.headers.get('cf-connecting-ip'),
    country: request.cf?.country,
    city: request.cf?.city,
    colo: request.cf?.colo,
  };
  
  return new Response(JSON.stringify(data, null, 2), {
    headers: { 
      'content-type': 'application/json',
      'access-control-allow-origin': '*',
    },
  });
}

// 返回请求详细信息
function handleJSON(request) {
  const requestInfo = {
    url: request.url,
    method: request.method,
    headers: Object.fromEntries(request.headers),
    cf: request.cf,
  };
  
  return new Response(JSON.stringify(requestInfo, null, 2), {
    headers: { 
      'content-type': 'application/json',
    },
  });
}

// 反向代理示例
async function handleProxy(request) {
  // 代理到另一个服务
  const targetUrl = 'https://api.github.com/zen';
  
  const response = await fetch(targetUrl, {
    headers: {
      'user-agent': 'Cloudflare-Worker-Proxy',
    },
  });
  
  // 修改响应头
  const newResponse = new Response(response.body, response);
  newResponse.headers.set('x-proxied-by', 'cloudflare-worker');
  
  return newResponse;
}
