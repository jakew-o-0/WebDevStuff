from pydantic import BaseModel, Field, EmailStr
from fastapi import Form
from typing import Annotated


class Token(BaseModel):
    access_token: str = Field(...)
    token_type: str = Field(...)

class SignupForm(BaseModel):
    username: str = Field(max_length=20, min_length=4)
    email: EmailStr = Field(...)
    password: str = Field(max_length=20, min_length=4)
    conf_password: str = Field(max_length=20, min_length=4)
    account_type: str = Field(...)

class NewTask(BaseModel):
    name: str = Field(min_length=5, max_length=30)
    class_name: str = Field(min_length=5, max_length=30)
    due_date: str = Field(min_length=5, max_length=30)