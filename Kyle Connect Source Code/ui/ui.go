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
    <div class="error-container">
        <div class="button-wrapper">
        <h3>The friends you selected have been added to your selected server.</h3>
        <button class="back-button">
                <a href="http://localhost:8086/addFriendToServer/addFriendToServer.go">Back</a>
            </button>
        </div>
    </div>

    <style>
    h3 {
        text-align: center;
    }

    .back-button {
        padding: 10px 20px;
        background-color: #007bff   ;
        color: white;
        font-size: 16px;
        border-radius: 6px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    a {
        color: black;
        text-decoration: none;
    }

    .button-wrapper {
        text-align: center;
        margin-top: 20px;
    }

    body {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .error-container {
        text-align: center;
        background-color: #fff;
        padding: 40px;
        border: 1px solid #dee2e6;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .error-message {
        color: #dc3545;
        font-size: 22px;
        margin-bottom: 20px;
    }
    </style>
`

var FriendsDeletedFromServer = `
	
    <div class="error-container">
        <div class="button-wrapper">
        <h3> The friends you selected have been removed from your selected server.</h3>
        <button class="back-button">
                <a href="http://localhost:8087/deleteFriendFromServer/deleteFriendFromServer.go">Back</a>
            </button>
        </div>
    </div>

    <style>
    h3 {
        text-align: center;
    }

    .back-button {
        padding: 10px 20px;
        background-color: #007bff   ;
        color: white;
        font-size: 16px;
        border-radius: 6px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    a {
        color: black;
        text-decoration: none;
    }

    .button-wrapper {
        text-align: center;
        margin-top: 20px;
    }

    body {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .error-container {
        text-align: center;
        background-color: #fff;
        padding: 40px;
        border: 1px solid #dee2e6;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .error-message {
        color: #dc3545;
        font-size: 22px;
        margin-bottom: 20px;
    }
    </style>
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

    <style>
        h3 {
            padding-left: 15px; /* Moves the link to the right */
        }
    </style>
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

var CannotAddThisFriendToThisServer = `
   <style>
    body {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .error-container {
        text-align: center;
        background-color: #fff;
        padding: 40px;
        border: 1px solid #dee2e6;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .error-message {
        color: #dc3545;
        font-size: 22px;
        margin-bottom: 20px;
    }

    .back-button {
        padding: 10px 20px;
        background-color: #007bff   ;
        color: white;
        font-size: 16px;
        border-radius: 6px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    a {
        color: black;
        text-decoration: none;
    }
    </style>

    <div class="error-container">
        <h3>You cannot add someone who is already in the server.</h3> 
        <button class="back-button"><a href="http://localhost:8086/addFriendToServer/addFriendToServer.go">Back</a></button>
    </div>
`

var DeleteFriendFromServer = `
    <h3>You can remove someone because they are already in the server.</h3>
`

var CannotDeleteSomeoneWhoIsNotInTheServer = `
    <style>
    body {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .error-container {
        text-align: center;
        background-color: #fff;
        padding: 40px;
        border: 1px solid #dee2e6;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .error-message {
        color: #dc3545;
        font-size: 22px;
        margin-bottom: 20px;
    }

    .back-button {
        padding: 10px 20px;
        background-color: #007bff   ;
        color: white;
        font-size: 16px;
        border-radius: 6px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    a {
        color: black;
        text-decoration: none;
    }
    </style>

    <div class="error-container">
        <h3>Oops! You cannot remove this person because they are not in the server.</h3>
        <button class="back-button"><a href="http://localhost:8087/deleteFriendFromServer/deleteFriendFromServer.go">Back</a></button>
    </div>
`

var CannotSendFriendRequest = `
    <style>
    body {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .error-container {
        text-align: center;
        background-color: #fff;
        padding: 40px;
        border: 1px solid #dee2e6;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .error-message {
        color: #dc3545;
        font-size: 22px;
        margin-bottom: 20px;
    }

    .back-button {
        padding: 10px 20px;
        background-color: #007bff   ;
        color: white;
        font-size: 16px;
        border-radius: 6px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    a {
        color: black;
        text-decoration: none;
    }
    </style>

    <div class="error-container">
        <h3>You're already friends, so no need to send another request!</h3>
        <button class="back-button"><a href="http://localhost:8082/actions/addFriend.go">Back</a></button>
    </div>
`

var NoUsernameWithThatName = `
    <style>
    body {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .error-container {
        text-align: center;
        background-color: #fff;
        padding: 40px;
        border: 1px solid #dee2e6;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .error-message {
        color: #dc3545;
        font-size: 22px;
        margin-bottom: 20px;
    }

    .back-button {
        padding: 10px 20px;
        background-color: #007bff   ;
        color: white;
        font-size: 16px;
        border-radius: 6px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    a {
        color: black;
        text-decoration: none;
    }
    </style>
    
    <div class="error-container">
        <h3>Hmm, we couldn't find a username with that name in Kyle Connect.</h3>
        <button class="back-button"><a href="http://localhost:8082/actions/addFriend.go">Try again</a></button>
    </div>
`

var CannotUseThisServerName = `
    <style>
    body {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .error-container {
        text-align: center;
        background-color: #fff;
        padding: 40px;
        border: 1px solid #dee2e6;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .error-message {
        color: #dc3545;
        font-size: 22px;
        margin-bottom: 20px;
    }

    .back-button {
        padding: 10px 20px;
        background-color: #007bff   ;
        color: white;
        font-size: 16px;
        border-radius: 6px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    a {
        color: black;
        text-decoration: none;
    }
    </style>
    
    <div class="error-container">
        <h3>It looks like you've already created a server with this name. Try choosing a different name!</h3>
        <button class="back-button"><a href="http://localhost:8084/createServer/createServer.go">Try again</a></button>
    </div>
`

var CannotUseThisChannelName = `
    <style>
    body {
        background-color: #f8f9fa;
        font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }

    .error-container {
        text-align: center;
        background-color: #fff;
        padding: 40px;
        border: 1px solid #dee2e6;
        border-radius: 12px;
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
    }

    .error-message {
        color: #dc3545;
        font-size: 22px;
        margin-bottom: 20px;
    }

    .back-button {
        padding: 10px 20px;
        background-color: #007bff   ;
        color: white;
        font-size: 16px;
        border-radius: 6px;
        border: none;
        cursor: pointer;
        transition: background-color 0.3s ease;
    }

    .back-button:hover {
        background-color: #0056b3;
    }

    a {
        color: black;
        text-decoration: none;
    }
    </style>
    
    <div class="error-container">
        <h3>It looks like you've already created a channel with this name. Try choosing a different name!</h3>
        <button class="back-button"><a href="http://localhost:8085/createChannel/createChannel.go">Try again</a></button>
    </div>
`

var CannotAddNobodyToAServer = `
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
		<h3>Oops! You need to select someone to add to the server.</h3>
		<a href="http://localhost:8086/addFriendToServer/addFriendToServer.go"><button class="btn">Try Again</button></a>
	</div>
`

var CannotDeleteNobodyFromAServer = `
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
		<h3>Oops! You need to select someone to delete from the server.</h3>
		<a href="http://localhost:8087/deleteFriendFromServer/deleteFriendFromServer.go"><button class="btn">Try Again</button></a>
	</div>
`

var ThereIsNoChannelsInThisServer = `
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
        <h3>No channels have been created in this server yet.</h3>
		<a href="http://localhost:8085/createChannel/createChannel.go"><button class="btn">Create a Channel</button></a>
	</div>
`

var YouCannotSendAFriendRequestToYourself = `
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
        <h3>You cannot send a friend request to yourself!</h3>
		<a href="http://localhost:8082/actions/addFriend.go"><button class="btn">Back</button></a>
	</div>
`