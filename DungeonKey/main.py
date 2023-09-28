from flask import Flask, render_template, url_for, request, flash
from pymongo import MongoClient

import dataModels as dm 
from forms import LoginForm, SignupForm



app = Flask(__name__)
app.config['SECRET_KEY'] = "supper secret token"
client = MongoClient()
db = client.DungeonKeyDB
UsrCluster = db.Users


CurrentSession = ''






@app.route('/', methods=['GET', 'POST'])
def MainMenu():
    return render_template('home.html')


@app.route('/login', methods=['GET', 'POST'])
def Login():
    form = LoginForm()

    if form.validate_on_submit():
        user = UsrCluster.find_one( {'email': form.email.data,
                                     'password': form.password} )
        if user is None:
            flash("There is no user")
        flash("logged in")


    return render_template('login.html', 
                            form=form)


@app.route('/login/new', methods=['GET', 'POST'])
def Signup():
    form = SignupForm()

    if request.method == "POST" and form.validate():
        user = UsrCluster.find_one( {'email': form.email.data} )
        if user is None:
            UsrCluster.insert_one(dm.UserModel(form.name.data,
                                               form.email.data,
                                               form.password.data))
        form.name.data = ''
        form.email.data = ''
        form.password.data = ''
        flash("successfull signup")

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