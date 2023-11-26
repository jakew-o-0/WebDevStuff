from typing import Annotated
from fastapi import APIRouter, Form, Request, Depends
from fastapi.security import OAuth2PasswordRequestForm
from fastapi.responses import HTMLResponse
from utils import SECRET_KEY, ALGORITHM, PASSWD_CONTEXT, TOKEN_TIMEOUT, OAUTH_SCHEME
from jose import JWTError, jwt
from datetime import datetime, timedelta
import aiosqlite


router = APIRouter()




@router.post('/', response_class=HTMLResponse)
async def login(form_data: Annotated[OAuth2PasswordRequestForm, Depends()]):
    #validate login credentials
    user_id = ''
    try:
        async with aiosqlite.connect('data/DataBase.sqlite3') as db:
            user = await db.execute('SELECT * FROM users WHERE email = ?', (form_data.username,))
            user = await user.fetchone()
            print(user)

            if user is None or user == () or user == []:
                raise Exception
            if not PASSWD_CONTEXT.verify(form_data.password, user[2]):
                raise Exception

            user_id = user[0]
            del user
    except:
        return HTMLResponse(
            '<h2>Invalid Email or Password<h2/>'
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
    token = 'Bearer ' + str(token)

    response = HTMLResponse('<h2>Loggedin</h2>')
    response.set_cookie(
        key='access_token',
        value=token
    )
    return response





@router.get('/token', response_class=HTMLResponse)
async def refresh_token(token: Annotated[str, Depends(OAUTH_SCHEME)]):
    try:
        decoded_token = jwt.decode(token, SECRET_KEY, ALGORITHM)
        user_id = decoded_token.get('data')

        if user_id is None:
            raise JWTError
    except JWTError:
        return HTMLResponse('').set_cookie(
            key='access_token',
            value=''
        )

    new_token = jwt.encode(
        claims={
            'data': user_id,
            'exp': datetime.utcnow() + timedelta(minutes=TOKEN_TIMEOUT)
        },
        key=SECRET_KEY,
        algorithm=ALGORITHM
    )

    return HTMLResponse('').set_cookie(
        key='access_token',
        value=new_token
    )






@router.post('/new', response_class=HTMLResponse)
async def signup(request: Request,
                 first_name: Annotated[str, Form()],
                 last_name: Annotated[str, Form()],
                 password: Annotated[str, Form()],
                 conf_password: Annotated[str, Form()],
                 ):
    #validate information
    #insert into db
    #response
    pass
