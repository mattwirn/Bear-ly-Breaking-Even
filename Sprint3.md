## Link to Video Presentation
[Click me](https://drive.google.com/file/d/1-2zIlwyRai0SwN_aaDNTX0nISqlCmkCL/view?usp=sharing)

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

Additionally, error handling has been implemented in the signup page for both usernames that are taken and passwords that don’t match:
- Return error message if input Username is already taken/exists in the database
- Return error message if password and confirm password inputs do not match
- Redirect from the signup page to dashboard upon successful signup

Three new Cypress tests have been written to test the new features:
- Testing the e2e aspect of the error handling
- Testing the error-message-rendering component on the signup page.
- Test if the function that enables the user to edit their income and different expense categories works as intended.


## Back End

In this sprint, we improved the Dashboard handler and added more features to it that are integral to our application. The dashboard handler is a GET function that currently returns all info about a user in a JSON. Most of that information is new to this sprint which we have added functionality for. The first one is an Income section on the dashboard. Upon the signup of a new user, a default value of 0 is defined. An Update Income handler is provided whenever the user wants to change that value. The most important feature we have added this sprint is expenses. We currently have 5 different types of expenses which we plan to expand in the next sprint. The categories are Home & Utilities, Transportation, Food, Education, and Health. Each individual expense has a name and an amount. We provided adders, updaters, and deleters for each type of expense to give the user the most control over their expenses. More detailed information about the expenses functionality can be found in the API Documentation.

For our unit tests, we added 11 new tests to maximize coverage for our new handlers.

- For the InputIncome handler, the only fields it needs from the front-end are the username and the amount. Since the amount can be any number, the only possible error is if the front end inputs a username that does not exist. Thus, only 2 tests are done for this handler - one with a valid username and one with an invalid username. The valid one will return "Income Updated" and the invalid one will return "Failed to find user, username does not exist"

- For the AddExpense handler, the ExpenseName and Amount fields are arbitrary. Thus an error might only occur if the username passed by the front end is invalid or if the ExpenseType is invalid. If the username does not exist, the handler returns "Failed to create expense"; if the ExpenseType is not formatted as is necessary, it will return "Failed to create expense, expense type not found". Otherwise, if successful, it returns "Expense Added".

- The UpdateExpense handler has the most input fields. It again has the username and ExpenseType, which if incorrect returns similarly to previously. However it also takes in the OldExpenseName and OldAmount, to identify which expense entry must be updated. If either of these are not valid for the given user, it also returns "Failed to update expense". The NewExpenseName and NewAmount are arbitrary.

- The DeleteExpense is similar to the AddExpense handler, except the ExpenseName and Amount fields are not arbitrary. If either are invalid, the handler will return "Failed to delete expense". It will also return this if the username is invalid, and if the ExpenseType is invalid, it will specify that it "Failed to delete expense, expense type not found".

## API Documentation
### Signup Handler
Signup is a POST function used to create a user. It takes in username and password strings in json, hashes the password, and stores the user data in a database. A cookie is also made for automatic login, and the newly created user is sent back to the front end.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's username and password.
If the function returns "Username already taken", the username input already exists. All accounts must be unique. If the account was successfully created, the function will return "Signed up"

### Login Handler
Login is a POST function to enter the website with an existing user. It checks to see if the input username exists, and if the input password matches the associated password. If both are valid, it creates a cookie for a login session.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's username and password.
If the function returns "Invalid username or password", the username doesn't exist or the password was incorrect. If the username exists and the input password matches the stored password, the function will return "Logged in".

The json received should look like this for both Signup and Login:
```JSON
{
   "username": "example",
   "password": "example",
}
```


### Dashboard Handler
Dashboard is a GET function that provides the front end with a json of all the user's fields. It identifies the user via their session cookie.
The function itself has 2 parameters, a write json body, "**w**", and a read json body, "**r**". "**w**" is the body the function will write to and send back to the front end, and "**r**" is the body that the function receives with the user's session cookie.
If the session cookie does not exist, it returns "Server side error", otherwise it returns a json with all the user's fields.

The order of the provided json is as follows:
- ***Username***
- ***Income***
- Expenses of ***Home and Utilities***
- Expenses of ***Transportation***
- Expenses of ***Food***
- Expenses of ***Education***
- Expenses of ***Health***

The first two entries can be guaranteed to be the **Username** and **Income**, with the rest being expenses that can be distinguished by the *ExpenseType* field. Here is an example of what one expense may look like:

```JSON
{
   "ExpenseType": "Food",
   "ExpenseName": "Grocery Store",
   "Amount": 109
}
```


### Update Income Handler
#### Path - "dashboard/income/update"
This is a POST function that edits the income of the user. It takes in a json that contains the *Username*, and *Amount*. Here is an example input:
```JSON
{
    "Username": "test",
    "Amount": 6900
}
```
If the update is successful, the returned message will be "*Income Updated*." If the Username is incorrect, an error will return stating "*Failed to find user, username does not exist*."


### Input Format of ExpenseTypes
When inputting the ExpenseType field to the json to be sent to the API, there are 5 possible types of expenses to choose from. They must be written in these specific ways in order for the input to be valid and function correctly:
```
Home & Utilities --> "HomeUtils"
Transportation   --> "Trans"
Food             --> "Food"
Education        --> "Edu"
Health           --> "Health"
```


### Add Expense Handler
#### Path - "dashboard/expense/add"
This is a POST function that adds a new expense to the database. It takes in a json that contains the *Username*, *ExpenseType*, *ExpenseName*, *Amount*. Here is an example input:
```JSON
{
   "Username": "ex",
   "ExpenseType": "Food",
   "ExpenseName": "Grocery Store",
   "Amount": 109
}
```
If the delete is successful, the returned message will be "*Expense Added*." If the Expense Type is incorrect, an error will return stating "*Failed to create expense, expense type not found*." If the Username is incorrect, an error will return stating "*Failed to create expense*."


### Update Expense Handler
#### Path - "dashboard/expense/update"
This is a PUT function that updates an exisiting expense with new value(s). It takes in a json that contains the *Username*, *ExpenseType*, *OldExpenseName*, *OldAmount*, *NewExpenseName*, *NewAmount*. Here is an example input:
```JSON
{
   "Username": "ex",
   "ExpenseType": "Food",
   "OldExpenseName": "Grocery Store",
   "OldAmount": 109,
   "NewExpenseName": "Fast Food",
   "NewAmount": 12
}
```
If the update is successful, the returned message will be "*Expense Updated*." If the Expense Type is incorrect, an error will return stating "*Failed to update expense, expense type not found*." If any of the fields Username, OldExpenseName, or OldAmount is incorrect, an error will return stating "*Failed to find expense*."


### Delete Expense Handler
#### Path - "dashboard/expense/delete"
This is a DELETE function that deletes an existing expense from the database. It takes in a json that contains the *Username*, *ExpenseType*, *ExpenseName*, *Amount*. Here is an example input:
```JSON
{
   "Username": "ex",
   "ExpenseType": "Food",
   "ExpenseName": "Fast Food",
   "Amount": 12
}
```
If the delete is successful, the returned message will be "*Expense Deleted*." If the Expense Type is incorrect, an error will return stating "*Failed to delete expense, expense type not found*." If any of the fields Username, ExpenseName, or Amount is incorrect, an error will return stating "*Failed to find expense*."
