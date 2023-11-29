from fastapi import APIRouter, Cookie
from fastapi.requests import Request
from fastapi.templating import Jinja2Templates
from typing import Annotated
from jose import jwt
from utils import SECRET_KEY, ALGORITHM
import aiosqlite


router = APIRouter()
templates = Jinja2Templates('templates')



@router.get('/')
async def get_user_details(request: Request, access_token: Annotated[str | None, Cookie()] = None):
    try:
        # validate access_token
        token = jwt.decode(access_token, SECRET_KEY, ALGORITHM)
        user_id = token.get('data')
        if user_id == None:
            raise Exception
        
        # fetch user data from db
        with aiosqlite.connect('data/DataBase.sqlite3') as db:
            user = await db.execute(
                'SELECT * FROM users WHERE user_id = ?',
                (user_id,)
            )
            user = await user.fetchone()

        if user == None or user == ():
            raise Exception

    except Exception:
        pass

    # create response
    return templates.TemplateResponse(
        'pages/userDetailsPage.html',
        {'request': request, 'user_data': user}
    )