package ui

var UI = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Home Page</title>
</head>
<body>
	<h1> Welcome to Kyle Connect, {{.Email}}</h1> <br>

    <hr class="dashed">

    <h3 class="actions">Actions</h3>

	<style>
	h1 {
		text-align: center;
		margin-bottom: 10px;
        font-size: 28px;
	}
	</style>
</body>
</html>
`

var Line = `
    <br><br><br>
    <hr class="dashed">
`

var UIERROR = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Home Page</title>
</head>
<body>

	<h1> Incorrect Email or Password, try again.</h1>

</body>
</html>
`


var FriendRequestAccepted = `
	<p> You have accepted the friend request. Your friends list will update now.</p>
`

var FriendRequestDeclined = `
	<p> You have declined the friend request.</p>
`

var FriendsAddedToServer = `
	<p> The friends you selected have been added to your selected server.</p>
`

var FriendsDeletedFromServer = `
	<p> The friends you selected have been removed from your selected server.</p>
`

var Adding = `
	<h1> Select the friends you wish to add to your server.</h1>
`

var Removing = `
	<h1> Select the friends you wish to remove from your server.</h1>
`

// var AdminOfServer = `
// 	<h3>You created the above servers to add and delete friends.</h3>
//     <a href="addFriendToServer/addFriendToServer.go">Add Friend To Server</a> <br>
//     <a href="deleteFriendFromServer/deleteFriendFromServer.go">Delete Friend From Server</a> 
// `

var TEST = `
    <h3>You created the above servers to add and delete friends.</h3>
`

var ThisUserNameIsAlreadyTaken = `
	<style>
        body {
            font-family: 'Arial', sans-serif;
            background: linear-gradient(to right, #141e30, #243b55);
            text-align: center;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }

        .container {
            background: rgba(255, 255, 255, 0.1);
            padding: 30px;
            border-radius: 12px;
            box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
            width: 100%;
            max-width: 400px;
        }

        h1 {
            margin-bottom: 10px;
            font-size: 28px;
        }

        h3 {
            margin-bottom: 20px;
            font-size: 16px;
            opacity: 0.8;
			color: black;
        }

        .input-container {
            align-items: center;
            background: white;
            padding: 12px;
            border-radius: 8px;
            margin-bottom: 15px;
            transition: 0.3s ease-in-out;
            border: 2px solid transparent;
        }

        .input-container:hover {
            border-color: #007bff;
        }

        .icon {
            margin-right: 10px;
            color: #007bff;
            font-size: 18px;
        }

        input {
            border: none;
            outline: none;
            width: 100%;
            font-size: 16px;
            padding: 5px;
        }

        input:focus {
            border-bottom: 2px solid #007bff;
        }

        .btn {
            background-color: #ff4b4b;
            color: white;
            padding: 12px;
            border: none;
            border-radius: 8px;
            cursor: pointer;
            width: 100%;
            font-size: 18px;
            font-weight: bold;
            transition: 0.3s;
        }

        .btn:hover {
            background-color: #d43f3f;
        }
	</style>

	<div class="input-container">
		<h3>Oops! That username is already taken. Please try a different one.</h3>
		<a href="http://localhost:8080/"><button class="btn">Try Again</button></a>
	</div>
`

var YourAccountHasBeenMade = `
	<style>
        body {
            font-family: 'Arial', sans-serif;
            background: linear-gradient(to right, #141e30, #243b55);
            text-align: center;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
        }

        .container {
            background: rgba(255, 255, 255, 0.1);
            padding: 30px;
            border-radius: 12px;
            box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
            width: 100%;
            max-width: 400px;
        }

        h1 {
            margin-bottom: 10px;
            font-size: 28px;
        }

        h3 {
            margin-bottom: 20px;
            font-size: 16px;
            opacity: 0.8;
			color: black;
        }

		.input-container {            
            align-items: center;
            background: white;
            padding: 12px;
            border-radius: 8px;
            margin-bottom: 15px;
            transition: 0.3s ease-in-out;
            border: 2px solid transparent;
        }

        .input-container:hover {
            border-color: #007bff;
        }

        .icon {
            margin-right: 10px;
            color: #007bff;
            font-size: 18px;
        }

        input {
            border: none;
            outline: none;
            width: 100%;
            font-size: 16px;
            padding: 5px;
        }

        input:focus {
            border-bottom: 2px solid #007bff;
        }

		.btn {
            background-color: #007bff;
            color: #243b55;
            padding: 12px;
            border: none;
            border-radius: 8px;
            cursor: pointer;
            width: 100%;
            font-size: 18px;
            font-weight: bold;
            transition: 0.3s;
        }

        .btn:hover {
            background-color: #0056b3;
        }
	</style>

	<div class="input-container">
		<h3>Your account has been made {{.Username}}.</h3> 
		<a href="http://localhost:8081/login/login.go"><button class="btn">Login</button></a>
	</div>
`