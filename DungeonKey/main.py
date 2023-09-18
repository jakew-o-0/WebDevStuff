from flask import Flask, render_template, url_for, request, flash
from flask_sqlalchemy import SQLAlchemy
from datetime import datetime

from forms import LoginForm, SignupForm



app = Flask(__name__)
app.config['SECRET_KEY'] = "supper secret token"
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///users.db'
db = SQLAlchemy(app)

CurrentSession = ''


class UserModel(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(200), nullable=False)
    email=db.Column(db.String(120), nullable=False, unique=True)
    password = db.Column(db.String(200), nullable=False)
    dateAdded = db.Column(db.DateTime, default=datetime.utcnow)

    def __repr__(self):
        return '<Name %r>' % self.name





@app.route('/', methods=['GET', 'POST'])
def MainMenu():
    return render_template('home.html')


@app.route('/login', methods=['GET', 'POST'])
def Login():
    email = None
    password = None
    form = LoginForm()

    if form.validate_on_submit():
        pass

    return render_template('login.html', 
                            email=email, password=password, form=form)


@app.route('/login/new', methods=['GET', 'POST'])
def Signup():
    form = SignupForm()

    if request.method == "POST" and form.validate():
        user = UserModel.query.filter_by(email=form.email.data)
        if user is None:
            user = UserModel(name=form.name.data, email=form.email.data, password=form.password.data)

        CurrentSession = form.email.data
        form.name.data = ''
        form.email.data = ''
        form.password.data = ''

        flash("successfull signup")
        
        user = UserModel.query.order_by(UserModel.dateAdded)
        print(user)

    return render_template('signup.html', 
                            form=form)



@app.errorhandler(404)
def Error404(e):
    return render_template('error.html', etype=400, e=e)

@app.errorhandler(500)
def Error404(e):
    return render_template('error.html', etype=500, e=e)


if __name__ == "__main__":
    app.run(debug=True)