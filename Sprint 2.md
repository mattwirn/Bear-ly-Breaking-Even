## Front End

## Back End
In the back end, to integrate the front end and back end we used a proxy server. To make it easier to integrate this tool, we switched our server manager from GIN to gorillaMux.
After converting our server/database set up and our handlers to gorillaMux, we created unit tests to validate our Signup and Login handler functions.

There were 5 unit tests in total.
1. The first one involved signing up with a unique username and checking that the account was created.
2. The second test involved signing up with the same username to verify that the account would not be made since it already existed.
3. The third test would check that logging in with the appropriate username and password, the same ones used in the Signup test, would result in giving access to the user.
4. The fourth test checked that logging in with an incorrect username would result in an "Invalid username or password" error.
5. The fifth test checked that logging in with an incorrect password would result in an "Invalid username or password" error.
