## Front End

The front end has testing for components, as well as end-to-end testing of the login, signup, and home pages. A new dashboard page is in process now that the front and back ends are integrated. 

The front end has two component tests: 
1. A component test for the logo that checks if the logo renders and is visible
2. A test fort the Page Header to check that it renders and displays

The  four e2e tests are as follows:
1. Testing that the Log in button redirects from the Home page to the Login page
2. Testing that the Sign up button redirects from the Home page to the Signup page
3. Testing that the Logo/Home button redirects from the Login page to the Home page
4. Testing that the Logo/Home button redirects from the Signup page to the Home page


## Back End
In the back end, to integrate the front end and back end we used a proxy server. To make it easier to integrate this tool, we switched our server manager from GIN to gorillaMux.
After converting our server/database set up and our handlers to gorillaMux, we created unit tests to validate our Signup and Login handler functions.

There were 5 unit tests in total.
1. The first one involved signing up with a unique username and checking that the account was created.
2. The second test involved signing up with the same username to verify that the account would not be made since it already existed.
3. The third test would check that logging in with the appropriate username and password, the same ones used in the Signup test, would result in giving access to the user.
4. The fourth test checked that logging in with an incorrect username would result in an "Invalid username or password" error.
5. The fifth test checked that logging in with an incorrect password would result in an "Invalid username or password" error.
