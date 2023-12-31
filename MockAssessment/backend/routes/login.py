from fastapi import Depends, APIRouter, Form, HTTPException, status, Request
from fastapi.responses import HTMLResponse
from fastapi.security import OAuth2PasswordRequestForm
from fastapi.templating import Jinja2Templates

from typing import Annotated

from jose import JWTError, jwt
import uuid
from datetime import datetime, timedelta

from models import Token, SignupForm
from utils import get_db, TOKEN_TIMEOUT, ALGORITHM, SECRET_KEY, pwd_context


router = APIRouter()
templates = Jinja2Templates('Templates')


@router.get('/', response_class=HTMLResponse)
async def login_page(request: Request):
    context = {'request': request, 'usr_page': 'home'}
    return templates.TemplateResponse('loginPage.html', context)


@router.post('/', response_model=Token)
async def login(form_data: Annotated[OAuth2PasswordRequestForm, Depends()]):
    auth_exception = HTTPException(
        status_code=status.HTTP_404_NOT_FOUND,
        detail='Email or Password is Invalid'
    )

    # auth login details
    user = get_db().cursor().execute(
        'SELECT * FROM users WHERE email = ?',
        (form_data.username,) 
    ).fetchone()

    if user is None or user == [] or user == {}:
        raise auth_exception
    if not pwd_context.verify(form_data.password, user['password']):
        raise auth_exception

    # generate token
    token = 'Bearer ' + str(jwt.encode(
        claims={
            'data': user['uuid'],
            'exp': datetime.utcnow() + timedelta(TOKEN_TIMEOUT)
        },
        key=SECRET_KEY,
        algorithm=ALGORITHM
    ))
    response = HTMLResponse('<h1>logged in</h1>')
    response.set_cookie(
        key='session',
        value=token
    )
    return response






@router.get('/new', response_class=HTMLResponse)
async def signup(request: Request):
    context = {'request': request}
    return templates.TemplateResponse('signupPage.html', context)


@router.post('/new')
async def signup(signup_data: SignupForm):
    if signup_data.password != signup_data.conf_password:
        return HTTPException(
            status_code=status.HTTP_409_CONFLICT,
            detail='Passwords do not match'
        )
    if signup_data.account_type != 'Teacher' or signup_data.account_type != 'Student':
        return HTTPException(
            status_code=status.HTTP_409_CONFLICT,
            detail='Account type invalid'
        )
    
    try:
        con = get_db()
        con.cursor().execute(
            'INSERT INTO users VALUES (?, ?, ?, ?, ?)',
            (
                str(uuid.uuid4()),
                signup_data.email,
                pwd_context.hash(secret=signup_data.password),
                signup_data.account_type,
                signup_data.username,
            )
        )
        con.commit()
        user = con.cursor().execute(
            'SELECT uuid FROM users WHERE email = ?',
            (signup_data.email,)
        ).fetchone()
    except Exception:
        return HTTPException(
            status_code=status.HTTP_409_CONFLICT,
            detail='Failed to add new account'
        )

    return Token(
        access_token=jwt.encode(
            claims={
                'data': user['uuid'],
                'exp': datetime.utcnow() + timedelta(TOKEN_TIMEOUT)
            },
            key=SECRET_KEY,
            algorithm=ALGORITHM
        ),
        token_type='Bearer'
    )