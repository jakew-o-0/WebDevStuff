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
    con = get_db()
    con.cursor().execute(
        'CREATE TABLE tasks (task_id TEXT, owner_id TEXT, groups_id TEXT, name TEXT, due_date TEXT, class_name TEXT, content TEXT)'
    )
    con.cursor().execute(
        'CREATE TABLE users_tasks (user_id TEXT, task_id TEXT)'
    )
    con.cursor().execute(
        'CREATE TABLE tasks_groups (task_id TEXT, group_id TEXT)'
    )
    con.commit()

    uvicorn.run(
        'main:app',
        host='localhost',
        port=8000,
        reload=True,
    )