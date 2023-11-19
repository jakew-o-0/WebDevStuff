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
        userid =  res.get('data')

        if userid is None:
            raise JWTError
        if not validate_user(userid):
            raise JWTError
        return userid

    except JWTError:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail='Invalid Token'
        )

def validate_user(userid):
    user = get_db().cursor().execute(
        'SELECT * FROM users WHERE uuid = ?',
        (userid,)

    ).fetchone()
    if user is None or user == {} or user == []:
        return False
    return True


@router.post('/')
async def new_task(task_data: NewTask, token: Annotated[str, Depends(oauth_scheme)]):
    user_id = auth_token(token)

    try:
        con = get_db()

        # create task
        task_id = str(uuid4())
        con.cursor().execute(
            'INSERT INTO tasks VALUES (?, ?, ?, )',
            (
                task_id,
                user_id,
                task_data.name,
                task_data.due_date,
                task_data.class_name,
                task_data.content,
            )
        )

        # create empty groups
        group_ids = []
        for i in range(0, task_data.num_of_groups):
            group_ids.append(str(uuid4()))
            con.cursor().execute(
                'INSERT INTO groups VALUES (?, ?, ?, ?)',
                (
                    group_ids[-1],
                    i,
                    0.0,
                    0.0,
                )
            )

        # create many to many relation between the task and the groups
        for i in group_ids:
            con.cursor().execute(
                'INSERT INTO tasks_groups VALUES (?, ?)',
                (
                    task_id,
                    i,
                )
            )
    except Exception:
        return HTTPException(
            status_code=status.HTTP_409_CONFLICT,
            detail='DB failure with creating task'
        )

    return HTTPException(
        status_code=status.HTTP_201_CREATED,
        detail='Task created'
    )

@router.get('/')
async def get_tasks(token: Annotated[str, Depends(oauth_scheme)]):
    #get all tasks related to the user qerying
    user_id = auth_token(token)
    user = get_db().cursor().execute(
        'SELECT * FROM users WHERE uuid = ?',
        (user_id,)
    ).fetchone()

    if user['user_type'].upper() == 'TEACHER':
        try:
            # get all tasks related to the user
            tasks = get_db().cursor().execute(
                'SELECT * FROM tasks WHERE user_id = ?',
                (user_id,)
            ).fetchall()

            # get all group_id's related to the task
            groups = []
            for i in tasks:
                groups_ids = get_db().cursor().execute(
                    'SELECT * FROM tasks_groups WHERE task_id = ?',
                    (i['task_id'],)
                ).fetchall()
                for i in groups_ids:
                    groups.append(i)

                
            # get all group data and add that to the correct task_data
            for i in groups:
                group_data = get_db().cursor().execute(
                    'SELECT * FROM groups WHERE group_id = ?',
                    (i['group_id'],)
                ).fetchone()

                for j in tasks:
                    if j['task_id'] == i['task_id']:
                        j.append(group_data)

            return tasks
        except Exception:
            pass
    else:
        pass
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
async def submit_work():
    #student accounts can submit work to their group
    pass
