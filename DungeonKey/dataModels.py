from datetime import datetime

def UserModel(Name, Email, Password):
    return {
        'dateAdded': datetime.utcnow,
        'name': Name,
        'email': Email,
        'Password': Password,
    }