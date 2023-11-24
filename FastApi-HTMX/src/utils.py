from fastapi.security import OAuth2PasswordBearer
from passlib.context import CryptContext


# oauth_scheme is for setting up authenticated routes
# passwd_context is for hashing and verifying against hashes
# secret_key if for verifying JWT tokens
# token_timeout is the time for tokens to be invalid


OAUTH_SCHEME = OAuth2PasswordBearer(tokenUrl='/login')
PASSWD_CONTEXT = CryptContext(
    schemes=['bcrypt'], 
    deprecated='auto'
)
SECRET_KEY = '1e389a7c4086367ff5751c1518133b14cc352a6154818c11307fed31f9b117f9'
ALGORITHM = 'HS256'
TOKEN_TIMEOUT = 30
