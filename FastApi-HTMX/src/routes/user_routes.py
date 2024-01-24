from fastapi import APIRouter, Cookie, HTTPException, status
from fastapi.requests import Request
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from typing import Annotated
from jose import jwt
from utils import SECRET_KEY, ALGORITHM, is_loggedin
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
        async with aiosqlite.connect('data/DataBase.sqlite3') as db:
            user = await db.execute(
                'SELECT * FROM users WHERE user_id = ?',
                (user_id,),
            )
            user = await user.fetchone()
            if user == None or user == ():
                raise Exception

            return templates.TemplateResponse(
                name='pages/userDetailsPage.html',
                status_code=status.HTTP_200_OK,
                context={
                    'request': request, 
                    'logged_in': await is_loggedin(access_token),
                    'user_data': list(user)
                },
                headers={
                    'HX-Redirect': '/user/', 
                }
            )


    except Exception as e:
        print(e)
        return templates.TemplateResponse(
            name='pages/userDetailsPage.html',
            status_code = status.HTTP_401_UNAUTHORIZED,
            context={
                'request': request, 
                'logged_in': await is_loggedin(access_token),
                'user_data': None 
            },
            headers={
                'HX-Redirect': '/user/', 
            }
        )

        