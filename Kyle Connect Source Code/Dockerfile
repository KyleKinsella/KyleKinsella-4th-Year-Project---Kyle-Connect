FROM golang:1.23

# Change working directory.
WORKDIR /usr/src/app 

# Copy files and directories:
# COPY makeAccount ./
# COPY makeAccount/main ./
COPY . .
# COPY createChannel ./
COPY utils ./ 
COPY login ./
# COPY sendMessageToChannelInServer ./
# COPY actions ./
# COPY sendMessage ./
# COPY createServer ./
# COPY addFriendToServer ./
# COPY deleteFriendFromServer ./
# COPY welcomePage ./

# Describe which ports your application is listening on.
EXPOSE 8081/tcp

# Specify default commands.
CMD ["go", "run", "login.go"] 