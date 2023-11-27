from typing import Annotated
from fastapi import Cookie, FastAPI, Request
from fastapi.exceptions import RequestValidationError
from fastapi.responses import HTMLResponse
from fastapi.middleware.cors import CORSMiddleware
from fastapi.templating import Jinja2Templates
import uvicorn

from routes.login_routes import router as auth_router



app = FastAPI()
templates = Jinja2Templates('templates')

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=['*'],
    allow_headers=['*'],
)

app.include_router(
    auth_router,
    prefix='/auth'
)




async def is_loggedin(access_token):
    logged_in = False
    if access_token != None or access_token == '':
       logged_in = True
    return logged_in


# routes for html main pages

@app.get('/', response_class=HTMLResponse)
async def index_page(request: Request, access_token: Annotated[str | None, Cookie()] = None):
    return templates.TemplateResponse(
        name='pages/indexPage.html',
        context={
            'request': request,
            'logged_in': await is_loggedin(access_token)
        },
        headers={'HX-Redirect': '/'}
    )

@app.get('/login', response_class=HTMLResponse)
async def login_page(request: Request, access_token: Annotated[str | None, Cookie()] = None):
    return templates.TemplateResponse(
        name='pages/loginPage.html',
        context={
            'request': request,
            'logged_in': await is_loggedin(access_token)
        },
        headers={'HX-Redirect': '/login'}
    )

@app.get('/signup', response_class=HTMLResponse)
async def signup_page(request: Request, access_token: Annotated[str | None, Cookie()] = None):
    return templates.TemplateResponse(
        name='pages/signupPage.html',
        context={
            'request': request,
            'logged_in': await is_loggedin(access_token)
        },
        headers={'HX-Redirect': '/signup'}
    )

@app.get('/user/home', response_class=HTMLResponse)
async def user_home_page(request: Request, access_token: Annotated[str | None, Cookie()] = None):
    return templates.TemplateResponse(
        name='pages/userHomePage.html',
        context={
            'request': request,
            'logged_in': await is_loggedin(access_token)
        },
        headers={'HX-Redirect': '/user/home'}
    )




if __name__ == '__main__':
    uvicorn.run(
        'main:app',
        host='localhost',
        port=8000,
        reload=True,
    )
