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
<<<<<<< HEAD
async def get_user_details(request: Request, access_token: Annotated[str | None, Cookie()] = None):

    try:
        # validate access_token
=======
async def get_user_details(request: Request, access_token: Annotated[str | None, Cookie()]):
    # validate access_token
    try:
        if access_token == None:
            return Exception

>>>>>>> 394ae330b1b87929071042faf72587c2ad87513f
        token = jwt.decode(access_token, SECRET_KEY, ALGORITHM)
        user_id = token.get('data')

        if user_id == None:
            raise Exception
<<<<<<< HEAD
=======

    except Exception:
        return HTMLResponse(
            status_code=status.HTTP_401_UNAUTHORIZED
        )
>>>>>>> 394ae330b1b87929071042faf72587c2ad87513f

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
<<<<<<< HEAD

        
=======
        user = await user.fetchone()
        if user == None or user == ():
            return HTMLResponse(
                status_code=status.HTTP_401_UNAUTHORIZED
            )

        return templates.TemplateResponse(
            name='pages/userDetailsPage.html',
            status_code=status.HTTP_200_OK,
            context={
                'request': request, 
                'user_data': user,
                'logged_in': await is_loggedin(access_token)
            },
            headers={
                'HX-Redirect': '/user/', 
            }
        )


>>>>>>> 394ae330b1b87929071042faf72587c2ad87513f
