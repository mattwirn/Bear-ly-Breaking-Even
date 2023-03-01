## Link to Video Presentation
[Click me](https://drive.google.com/file/d/1LeO5o12e_EzsODLL3Xp02SgG4LOaInyV/view?usp=sharing)

## Front End

The front end has testing for components, as well as end-to-end testing of the login, signup, and home pages. A new dashboard page is in process now that the front and back ends are integrated. 

The front end has two component tests: 
1. A component test for the logo that checks if the logo renders and is visible
2. A component test for the Page Header to check that it renders and displays

The  four e2e tests are as follows:
1. Testing that the Log in button redirects from the Home page to the Login page
2. Testing that the Sign up button redirects from the Home page to the Signup page
3. Testing that the Logo/Home button redirects from the Login page to the Home page
4. Testing that the Logo/Home button redirects from the Signup page to the Home page


## Back End
In the back end, to integrate the front end and back end we used a reverse proxy. To make it easier to integrate this tool, we switched our router from GIN to gorillaMux and had to rewrite almost our entire back end. We then added more features, improving our current handlers and middleware and adding another handler, Dashboard, which we will focus on expanding upon and adding new features to in the next sprint. 
After converting our server set up and handlers to gorillaMux, we created table-driven unit tests to validate our Signup and Login handler functions. 

There were 5 unit tests in total.

#### Signup
- Check that signing up with a unique username would result in "Signed Up"
- Check that signing up with the same username would result in an "Username already exists" error
#### Login
- Check that logging in with the appropriate username and password, the same ones used in the Signup test, would result in "Logged In"
- Check that logging in with an incorrect username would result in an "Invalid username or password" error
- Check that logging in with an incorrect password would result in an "Invalid username or password" error


## API Documentation

### Signup Handler
Signup is a POST function used to create a user. It takes in username and password strings in json, hashes the password, and stores the user data in a database. A cookie is also made for automatic login, and the newly created user is sent back to the front end.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's username and password.
If the function returns "Username already taken", the username input already exists. All accounts must be unique. If the account was successfully created, the function will return "Signed up"

### Login Handler
Login is a POST function to enter the website with an existing user. It checks to see if the input username exists, and if the input password matches the associated password. If both are valid, it creates a cookie for a login session.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's username and password.
If the function returns "Invalid username or password", the username doesn't exist or the password was incorrect. If the username exists and the input password matches the stored password, the function will return "Logged in".

##

The sent JSON paramter should look like this for both Signup and Login:
```Go
{
   "username": "example",
   "password": "example",
}
```

## 

### Dashboard Handler
Dashboard is a GET function that provides the front end with a json of all the user's fields. It identifies the user via their session cookie.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's session cookie.
If the session cookie does not exist, it returns "Server side error", otherwise it returns a json with all the user's fields.
