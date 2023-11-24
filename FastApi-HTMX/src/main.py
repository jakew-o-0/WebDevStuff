from fastapi import FastAPI
import uvicorn
import aiosqlite

app = FastAPI()

if __name__ == '__main__':
    uvicorn.run(
        'app:main',
        host='localhost',
        port=8000,
        reload=True,
    )
