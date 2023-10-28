from fastapi import Depends, FastAPI, Form, HTTPException, status
from fastapi.encoders import jsonable_encoder
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm

from pydantic import BaseModel, Field
from typing import Annotated

from passlib.context import CryptContext
from jose import JWTError, jwt
import uuid
from datetime import datetime, timedelta

from motor.motor_asyncio import AsyncIOMotorClient
import uvicorn



app = FastAPI()
oauth2_scheme = OAuth2PasswordBearer(tokenUrl='token')

client = AsyncIOMotorClient('mongodb://localhost:27017')
db = client['AuthTestDB']

pwd_context = CryptContext(schemes=['bcrypt'], deprecated='auto')
SECRET_KEY = '5847030be3818c9311f436ee7b4f4d300e3010580b0c98933fd5baf617b6c869'
ALGORITHM = 'HS256'
TOKEN_TIMEOUT = 330




class Token(BaseModel):
    access_token: str
    token_type: str

class User(BaseModel):
    id: str = Field(default_factory=uuid.uuid4, alias='_id')
    username: str = Field(...)
    hashed_password: str = Field(...)

    class config:
        populate_by_name = True



@app.post('/token', response_model=Token)
async def access_token(form_data: Annotated[OAuth2PasswordRequestForm, Depends()]):
    # authenticate user
    validation_exception =  HTTPException(
        status_code=status.HTTP_404_NOT_FOUND,
        detail="Incorrect Username or Password"
    )

    if ((user := await db['Users'].find_one({'username': form_data.username})) is None):
        raise validation_exception
    user = User(**user)

    if not pwd_context.verify(form_data.password, user.hashed_password):
        raise validation_exception

    # create token: exp, uuid
    token_expires = datetime.utcnow() + timedelta(TOKEN_TIMEOUT)
    token_data = {
        'data': user.id,
        'exp': token_expires
    }
    jwt_token = jwt.encode(
        claims=token_data,
        key=SECRET_KEY,
        algorithm=ALGORITHM,
    )
    
    # return token
    return {"access_token": jwt_token, "token_type": "Bearer"}


@app.post('/users/new')
async def create_user(username: Annotated[str, Form()], password: Annotated[str, Form()]):
    user = jsonable_encoder(
        User(
            username=username,
            hashed_password=pwd_context.hash(password)
        )
    )
    user_inserted = await db['Users'].insert_one(user)
    if user_inserted is None:
        raise HTTPException(
            status_code = status.HTTP_500_INTERNAL_SERVER_ERROR,
            detail="db failed to create user"
        )
    return HTTPException(
        status_code=status.HTTP_201_CREATED,
        detail="user created"
    )


@app.get('/users/')
async def get_user_passwd(token: Annotated[str, Depends(oauth2_scheme)]):
    auth_err = HTTPException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail='Could not authenticate'
    )

    try:
        response = jwt.decode(token, SECRET_KEY, ALGORITHM)
        user_id = response.get('data')
        if user_id is None:
            raise auth_err
    except(JWTError):
        raise auth_err

    
    user = await db['Users'].find_one({'_id': user_id})
    return user




if __name__ == "__main__":
    uvicorn.run(
        'main:app',
        host='localhost',
        port=8000,
        reload=True
    )
