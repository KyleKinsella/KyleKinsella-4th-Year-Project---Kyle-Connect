package ui

var UI = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Home Page</title>
</head>
<body>
	<h1> Welcome to Kyle Connect, {{.Email}}</h1>

	<div class="act">
        <!-- <hr id="a1" class="dashed"> -->
        <h3 class="actions">Actions</h3>
        <ul>
            <li>Add friend (Send Friend Request)</li>
            <li>Send message to a friend</li>
            <li>Create a server</li>
            <li>Send message to a server</li>
            <li>Add friends to server</li>
            <li>Delete friends from server</li>
        </ul>
    </div>

    <div class="in">
        <input type="text" placeholder="this is going to be used for sending messages to a friend">
    </div>

    <br><br><br>
    <hr class="dashed">
</body>
</html>
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

var AdminOfServer = `
	<h3>You created the above servers to add and delete friends.</h3>
    <a href="addFriendToServer/addFriendToServer.go">Add Friend To Server</a> <br>
    <a href="deleteFriendFromServer/deleteFriendFromServer.go">Delete Friend From Server</a> 
`