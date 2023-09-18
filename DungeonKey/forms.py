from flask_wtf import FlaskForm
from wtforms import StringField, EmailField, PasswordField, SubmitField
from wtforms.validators import DataRequired, Email, EqualTo, ValidationError


class SignupForm(FlaskForm):
    name = StringField(label='Name: ',
                            validators=[DataRequired()])
    email = EmailField(label='Email: ',
                            validators=[DataRequired()])
    password = PasswordField(label='Password: ', 
                                 validators=[DataRequired()])
    submit = SubmitField(label='Unlock')


class LoginForm(FlaskForm):
    email = EmailField(label='Email: ',
                            validators=[DataRequired()])
    password = PasswordField(label='Password: ', 
                                 validators=[DataRequired()])
    submit = SubmitField(label='Unlock')