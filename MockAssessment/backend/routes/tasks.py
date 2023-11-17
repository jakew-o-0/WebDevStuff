from fastapi import APIRouter, Depends, HTTPException, status
from typing import Annotated
from jose import jwt, JWTError
from uuid import uuid4

from utils import get_db, oauth_scheme, SECRET_KEY, ALGORITHM
from models import NewTask


router = APIRouter()


def auth_token(token):
    try:
        res = jwt.decode(token, SECRET_KEY, ALGORITHM)
        if userid := res.get('data') is None:
            raise JWTError
        else:
            return userid
    except JWTError:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail='Invalid Token'
        )


@router.post('/')
async def new_task(task_data: NewTask, token: Annotated[str, Depends]):
    #new row in tasks, create group table
    user_id = auth_token(token)
    con = get_db()
    con.cursor().execute(
        'INSERT INTO tasks VALUES (?, ?, ?, )',
        (
            str(uuid4()),
            user_id,
            
        )
    )
    pass

@router.get('/')
async def get_tasks():
    #get all tasks related to the user qerying
    pass

@router.put('/')
async def update_task():
    #teacher accounts can update their own tasks
    pass

@router.delete('/')
async def delete_tasks():
    # teacher accounts can delete their own tasks
    pass


@router.post('/addStudent')
async def add_student():
    # add relation between student and task id's
    pass

@router.post('/submitWork')
async def add_student():
    #student accounts can submit work to their group
    pass