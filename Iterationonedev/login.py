import sqlite3
import mysql.connector
from flask import Flask, render_template, request
from cryptography.fernet import Fernet

app = Flask(__name__)


def encryptPassword(password):
    s = password
    key = Fernet.generate_key()
    fernet = Fernet(key)
    encrypt = fernet.encrypt(s.encode())

    return encrypt


def loginUser(email, password):
    email = request.form.get('email')
    password = request.form.get('password')

    connection = mysql.connector.connect(
    user = "root",
    password = "",
    database = "users",
    host = "localhost"
    )

    cursor = connection.cursor()
    query = "SELECT email, password FROM communicators WHERE email=%s AND password=%s"
    cursor.execute(query, (email, password))
    
    result = cursor.fetchall()
    for row in result:
        print(row)

    cursor.close()
    connection.close()


@app.route('/', methods=['GET', 'POST'])
def loginScreen():
    return render_template('login.html')



@app.route('/submit', methods=['POST'])
def submitForm():
    email = request.form.get('email')
    password = request.form.get('password')
    password = encryptPassword(password)
    loginUser(email, password)
    return render_template('ui.html')
    # return f"Your have logged into your account. Make sure to use this email to login in: {email}."



if __name__ == '__main__':
    app.run(debug=True)