from fastapi import FastAPI, Request, Depends
from fastapi.responses import HTMLResponse
from fastapi.middleware.cors import CORSMiddleware
from fastapi.templating import Jinja2Templates
import uvicorn

from typing import Annotated

from routes.login_routes import router as login_router
from utils import OAUTH_SCHEME


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
    login_router,
    prefix='/api/login'
)




# routes for html main pages

@app.get('/', response_class=HTMLResponse)
async def index_page(request: Request):
    return templates.TemplateResponse(
        'pages/indexPage.html',
        {'request': request},
        headers={'HX-Redirect': '/'} #work-around for refreshing the page instead of using an <a> tag
    )

@app.get('/login', response_class=HTMLResponse)
async def login_page(request: Request):
    return templates.TemplateResponse(
        name='pages/loginPage.html',
        context={'request': request},
        headers={'HX-Redirect': '/login'}
    )
@app.get('/authtest')
async def test(token: Annotated[str, Depends(OAUTH_SCHEME)]):
    return HTMLResponse(
        '<h1>IT WORKED</h1>',
    )




if __name__ == '__main__':
    uvicorn.run(
        'main:app',
        host='localhost',
        port=8000,
        reload=True,
    )
