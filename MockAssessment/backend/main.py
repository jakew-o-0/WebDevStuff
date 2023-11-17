from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware 
import uvicorn

from routes.login import router as login_router


app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=['*'],
    allow_headers=['*'],
)

app.include_router(
    login_router,
)

def test():
    from utils import get_db
    get_db().cursor().execute(
        'CREATE TABLE tasks (task_id TEXT, owner_id TEXT, title: TEXT, due_date: TEXT, class_name: TEXT, )',

    )
if __name__ == '__main__':
    uvicorn.run(
        'main:app',
        host='localhost',
        port=8000,
        reload=True,
    )