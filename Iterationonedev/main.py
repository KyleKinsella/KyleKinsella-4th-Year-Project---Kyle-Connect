from flask import Flask, render_template, request
import DBcm
import mysql.connector

config = {
    "user": "root",
    "password": "",
    "database": "users",
    "host": "localhost"
}


INSERT = """
    INSERT INTO communicators
    (username, email, password)
    VALUES
    (%s, %s, %s)
 """

app = Flask(__name__)

@app.route('/')
def home():
    return render_template('makeAccount.html')


@app.route('/submit', methods=['POST'])
def submitForm():
    print(request.form)  # Debugging step
    
    name = request.form.get('name')
    email = request.form.get('email')
    password = request.form.get('password')

    data = (name, email, password)

    print(f"name: {name}, email: {email}, password: {password}")

    with DBcm.UseDatabase(config) as db:
        db.execute(INSERT, data)


    return f"form submitted! name: {name}, email: {email}, password: {password}"

if __name__ == '__main__':
    app.run(debug=True, port=5000)