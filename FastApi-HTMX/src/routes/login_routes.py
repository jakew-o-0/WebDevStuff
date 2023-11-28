from typing import Annotated
from fastapi import APIRouter, Form, Request, Depends, Cookie
from fastapi.security import OAuth2PasswordRequestForm
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from utils import SECRET_KEY, ALGORITHM, PASSWD_CONTEXT, TOKEN_TIMEOUT
from jose import jwt
from datetime import datetime, timedelta
from uuid import uuid4
import aiosqlite
import re


router = APIRouter()
templates = Jinja2Templates('templates')




@router.post('/login', response_class=HTMLResponse)
async def login(request: Request, form_data: Annotated[OAuth2PasswordRequestForm, Depends()]):
    #validate login credentials
    user_id = ''
    try:
        async with aiosqlite.connect('data/DataBase.sqlite3') as db:
            user = await db.execute('SELECT * FROM users WHERE email = ?', (form_data.username,))
            user = await user.fetchone()

            if user is None or user == () or user == []:
                raise Exception
            if not PASSWD_CONTEXT.verify(form_data.password, user[2]):
                raise Exception

            user_id = user[0]
            del user
    except Exception:
        return templates.TemplateResponse(
            'components/loginError.html',
            { 'request': request }
        )

    #create access token
    token = jwt.encode(
        claims={
            'data': user_id,
            'exp': datetime.utcnow() + timedelta(minutes=TOKEN_TIMEOUT)
        },
        key=SECRET_KEY,
        algorithm=ALGORITHM
    )

    response = HTMLResponse(
        headers={'HX-Redirect': '/user/home/'}
    )
    response.set_cookie(
        key='access_token',
        value=token
    )
    return response




@router.get('/logout', response_class=HTMLResponse)
async def logout():
    response = HTMLResponse(
        headers={'HX-Redirect': '/'}
    )
    response.delete_cookie(
        key='access_token'
    )
    return response




@router.get('/refresh-token', response_class=HTMLResponse)
async def refresh_token(access_token: Annotated[str | None, Cookie()] = None):
    # validate token
    try:
        if access_token is None:
            raise Exception

        decoded_token = jwt.decode(access_token, SECRET_KEY, ALGORITHM)
        user_id = decoded_token.get('data')

        if user_id is None:
            raise Exception
    except Exception:
        response = HTMLResponse()
        response.delete_cookie(
            key='access_token',
        )
        return response

    new_token = jwt.encode(
        claims={
            'data': user_id,
            'exp': datetime.utcnow() + timedelta(minutes=TOKEN_TIMEOUT)
        },
        key=SECRET_KEY,
        algorithm=ALGORITHM
    )

    response = HTMLResponse()
    response.set_cookie(
        key='access_token',
        value=new_token
    )
    return response






@router.post('/signup', response_class=HTMLResponse)
async def signup(request: Request,
                 first_name: Annotated[str, Form()],
                 last_name: Annotated[str, Form()],
                 email: Annotated[str, Form()],
                 password: Annotated[str, Form()],
                 conf_password: Annotated[str, Form()],
                 ):
    #validate information
    errs = []
    if password != conf_password:
        errs.append('passwords_match_err')

    if not re.match(r'^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[!@#$%^&*,.?+_=]).{8,}$', password):
        errs.append('password_format_err')

    if not re.match(r'^([a-zA-Z])\w+$', first_name):
        errs.append('fname_format_err')

    if not re.match(r'^([a-zA-Z])\w+$', last_name):
        errs.append('lname_format_err')

    if not re.match(r'^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$', email):
        errs.append('email_format_err')

    async with aiosqlite.connect('data/DataBase.sqlite3') as db:
        user = await db.execute('SELECT * FROM users WHERE email = ?', (email,))
        user = await user.fetchone()
        if not(user == [] or user == () or user == None): 
            errs.append('email_exists_err')

    if errs != []:
        return templates.TemplateResponse(
            'components/signupError.html',
            { 'request': request, 'errs': errs }
        )

    #insert into db
    user_id = str(uuid4())
    async with aiosqlite.connect('data/DataBase.sqlite3') as db:
        await db.execute(
            'INSERT INTO users VALUES (?,?,?,?,?)',
            (
                user_id,
                email,
                PASSWD_CONTEXT.hash(password),
                first_name,
                last_name,
            )
        )
        await db.commit()

    # create token response
    token = jwt.encode(
        claims={
            'data': user_id,
            'exp': datetime.utcnow() + timedelta(minutes=TOKEN_TIMEOUT)

        },
        algorithm=ALGORITHM,
        key=SECRET_KEY
    )
    response = HTMLResponse(
        headers={'HX-Redirect': '/user/home/'}
    )
    response.set_cookie(
        key='access_token',
        value=token
    )
    return response
