## Link to Video Presentation


## Front End

In this sprint, the dashboard page was worked on and new functionality was added:
- Welcome message with username
- View current yearly income amount
- Edit current income value via input field and submit button
- Maintain current values if field is left empty
- Toggle between edit and view options 
- View current monthly expenses in 6 general categories: Home and Utilities, Transportation, Food, Entertainment, and Other expenses.
- Edit each individual category of expenses and maintain current values if empty fields are submitted
- Toggle between view and edit modes for current monthly expenses
- View estimated total yearly expenses from monthly input values
- Show capital after expenses and monthly excess income based on input income and expense values
Towards the end of the sprint, we found out there was some disconnect between the front and back end in how expenses were implemented. These issues will be resolved in Sprint 4 and allow for personalized data saved on a user object instead of requiring placeholder values and user input on every refresh.

Additionally, we improved the signup page and added error-checking funtionality:
- Return error message if input Username is already taken
- Return error message if password and confirm password inputs do not match
- Redirect from the signup page to dashboard upon successful signup

The front end has new component tests: 
1. 

The  new e2e tests are as follows:
1. 


## Back End


### Dashboard Handler
Dashboard is a GET function that provides the front end with a json of all the user's fields. It identifies the user via their session cookie.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's session cookie.
If the session cookie does not exist, it returns "Server side error", otherwise it returns a json with all the user's fields.
