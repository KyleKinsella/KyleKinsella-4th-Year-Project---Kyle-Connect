FROM golang:1.23

WORKDIR /usr/src/app

COPY makeAccount ./
COPY makeAccount/main ./
COPY . .

EXPOSE 8080

CMD ["go", "run", "signup.go"] 