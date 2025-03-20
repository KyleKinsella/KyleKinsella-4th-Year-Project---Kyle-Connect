package ui

var UI = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Home Page</title>
</head>
<body>
	<h1> Welcome to Kyle Connect, {{.Email}}</h1>

    <h3 class="actions">Actions</h3>
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