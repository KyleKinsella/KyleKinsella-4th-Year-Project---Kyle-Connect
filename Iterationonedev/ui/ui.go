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
            <li>Add friend</li>
            <li>Create a server</li>
            <li>Add friend to server</li>
        </ul>
    </div>


    <div class="in">
        <input type="text" placeholder="this is going to be used for sending messages to a friend">
    </div>


    <br><br><br>
    <hr class="dashed">


    <div class="ser">
        <h3 class="servers">Servers</h3>
        <!--this is going to pull all of your servers that you have made from the database or from your servers list, and populate the ul, li below-->
        <ul>
            <li>this is going to have all of your servers that you have made, <br>below is an example what it will look like:<br><br></li>

            <li>Drinks</li>
            <li>Sports</li>
            <li>Movies</li>
        </ul>
    </div>
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