#!/bin/bash
set -e

echo "🚀 Starting Cloudflare Manager..."

# Initialize default account if credentials are provided
if [ -n "$CLOUDFLARE_TOKEN" ]; then
    echo "📝 Initializing default account..."
    /app/cloudflare-manager/cfm account add default --token "$CLOUDFLARE_TOKEN" --email "$CLOUDFLARE_EMAIL" || true
fi

echo "🌐 Starting FastAPI backend on port 7860..."
cd /app/backend

# Start FastAPI with frontend serving
python -c "
from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from fastapi.responses import FileResponse
import uvicorn
import os
import sys

# Import the main app
sys.path.insert(0, '/app/backend')
from main import app as api_app

# Mount frontend
app = FastAPI()

# Mount API routes
app.mount('/api', api_app)

# Serve frontend static files
if os.path.exists('/app/frontend/build'):
    app.mount('/static', StaticFiles(directory='/app/frontend/build/static'), name='static')
    
    @app.get('/{full_path:path}')
    async def serve_frontend(full_path: str):
        file_path = f'/app/frontend/build/{full_path}'
        if os.path.exists(file_path) and os.path.isfile(file_path):
            return FileResponse(file_path)
        return FileResponse('/app/frontend/build/index.html')

if __name__ == '__main__':
    print('✅ Cloudflare Manager is running!')
    print('📊 Dashboard: http://0.0.0.0:7860')
    print('📡 API Docs: http://0.0.0.0:7860/api/docs')
    uvicorn.run(app, host='0.0.0.0', port=7860)
" &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?
