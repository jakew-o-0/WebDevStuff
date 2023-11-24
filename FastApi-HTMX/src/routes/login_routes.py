from typing import Annotated
from fastapi import APIRouter, Request, Depends
from fastapi.security import OAuth2PasswordRequestForm
from src.utils import SECRET_KEY, ALGORITHM, PASSWD_CONTEXT, TOKEN_TIMEOUT
import aiosqlite


router = APIRouter()


@router.post('/')
async def login(reqest: Request, form_data: Annotated[OAuth2PasswordRequestForm, Depends()]):
    #validate login credentials
    async with aiosqlite.connect('data/DataBase.sqlite3') as db:
        await db.execute('SELECT * FROM users WHERE email == ?', (form_data.username,))


    #create response cookie

    #return cookie as hears to response
    pass
