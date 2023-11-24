from fastapi.security import OAuth2PasswordBearer
from passlib.context import CryptContext
import sqlite3

oauth_scheme = OAuth2PasswordBearer(tokenUrl='/login')
pwd_context = CryptContext(schemes=['bcrypt'], deprecated='auto')
SECRET_KEY = '5847030be3818c9311f436ee7b4f4d300e3010580b0c98933fd5baf617b6c869'
ALGORITHM = 'HS256'
TOKEN_TIMEOUT = 330

def dict_factory(cur, row):
    d = {}
    for idx,column in enumerate(cur.description):
        d[column[0]] = row[idx]
    return d

def get_db():
    con = sqlite3.connect('./data/GibjohnTutoringDB')
    con.row_factory = dict_factory
    return con

    