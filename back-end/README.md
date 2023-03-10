# Handler Function Documentation
## Signup Handler
Signup is a POST function used to create a user. It takes in username and password strings in json, hashes the password, and stores the user data in a database. A cookie is also made for automatic login, and the newly created user is sent back to the front end.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's username and password.
If the function returns "Username already taken", the username input already exists. All accounts must be unique. If the account was successfully created, the function will return "Signed up"

##

The json received should look like this for both Signup and Login:
```
{
   "username": "example",
   "password": "example",
}
```
## Login Handler
Login is a POST function to enter the website with an existing user. It checks to see if the input username exists, and if the input password matches the associated password. If both are valid, it creates a cookie for a login session.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's username and password.
If the function returns "Invalid username or password", the username doesn't exist or the password was incorrect. If the username exists and the input password matches the stored password, the function will return "Logged in".

## Dashboard Handler
Dashboard is a GET function that provides the front end with a json of all the user's fields. It identifies the user via their session cookie.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's session cookie.
If the session cookie does not exist, it returns "Server side error", otherwise it returns a json with all the user's fields.
