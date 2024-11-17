package ui 

var UI = `
<!DOCTYPE html>
<html>
<head>
    <title>Kyle Connect - Home Page</title>
</head>
<body>

	<h1> Welcome to Kyle Connect</h1>
	<p> {{.Email}}</p>
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