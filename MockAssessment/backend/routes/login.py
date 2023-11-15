from fastapi import Depends, APIRouter, Form, HTTPException, status
from fastapi.security import OAuth2PasswordRequestForm

from pydantic import BaseModel, Field
from typing import Annotated

from passlib.context import CryptContext
from jose import JWTError, jwt
import uuid
from datetime import datetime, timedelta

from models import Token
from utils import get_db


router = APIRouter()

@router.post('/', response_model=Token)
async def login(form_data: Annotated[OAuth2PasswordRequestForm, Depends()]):
    auth_exception = HTTPException(
        status_code=status.HTTP_404_NOT_FOUND,
        detail='Email or Password is Invalid'
    )

    # auth login details
    data = (
        form_data.username,
    )
    query = '''
        SELECT *
        FROM users
        WHERE email = ?
    '''
    cur = get_db().cursor()
    user = cur.execute(query, data).fetchone()
    print(user['email'])

    # generate token

    return Token(token_type='', access_token='')