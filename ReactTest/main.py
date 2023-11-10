from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from starlette.responses import FileResponse
from starlette.staticfiles import StaticFiles
import uvicorn

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=['*'],
    allow_headers=['*'],
)

app.mount('/', StaticFiles(directory='./test-app/build'), name='UI')

@app.get('/')
def root():
    return FileResponse('./test-app/build/index.html')

@app.exception_handler(404)
async def err(request, exc):
    return FileResponse('./test-app/build/index.html')


if __name__ == '__main__':
    uvicorn.run(
        'main:app',
        host='localhost',
        port=8000,
        reload=True,
    )