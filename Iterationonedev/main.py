from flask import Flask, render_template, request
import DBcm
import mysql.connector

import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart
from cryptography.fernet import Fernet

config = {
    "user": "root",
    "password": "",
    "database": "users",
    "host": "localhost"
}


connection = mysql.connector.connect(
    user = "root",
    password = "",
    database = "users",
    host = "localhost"
)


def encryptPassword(password):
    s = password
    key = Fernet.generate_key()
    fernet = Fernet(key)
    encrypt = fernet.encrypt(s.encode())

    return encrypt



def send(recipientEmail, subject, body): 
    message = MIMEMultipart()
    message['from'] = sendedEmail
    message['to'] = recipientEmail
    message["subject"] = subject
    message.attach(MIMEText(body, 'plain'))

    try:
        # Set up the SMTP server
        with smtplib.SMTP(smtp_server, smtp_port) as server:
            server.starttls()  # Upgrade the connection to secure
            server.login(sendedEmail, sender_password)  # Login to your email account
            server.send_message(message)  # Send the email
        print("Email sent successfully")
    except Exception as e:
        print(f"Failed to send email: {e}")


# Email credentials
sendedEmail = 'kylekinsella10@gmail.com'
sender_password = 'xocq rwwb zqzd ipga'  # Use an app password or environment variable for security

# SMTP server configuration
smtp_server = 'smtp.gmail.com'
smtp_port = 587  # Use 587 for TLS


INSERT = """
    INSERT INTO communicators
    (username, email, password)
    VALUES
    (%s, %s, %s)
 """

app = Flask(__name__)



@app.route('/')
def home():
    return render_template('index.html')


@app.route('/home')
def makeAccount():
    return render_template('makeAccount.html')



@app.route('/login')
def login():
    return render_template('login.html')



@app.route('/submit', methods=['POST'])
def submitForm():
    name = request.form.get('name')
    email = request.form.get('email')
    password = request.form.get('password')
    password = encryptPassword(password)

    data = (name, email, password)

    with DBcm.UseDatabase(config) as db:
        db.execute(INSERT, data)


    recipient_email = email
    subject = "Welcome"
    body = "Hi " + name + ",\n" + "Welcome to Kyle Connect."
    send(recipient_email, subject, body)

    return f"Your account has been created {name}. Please check your email {name}."


if __name__ == '__main__':
    app.run(debug=True)