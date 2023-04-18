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

The order of the provided json is as follows:
- ***Username***
- ***Income***
- ***Total Spent***
- Expenses of ***Home and Utilities***
    - Total of Expenses
    - Individual Expenses
- Expenses of ***Travel***
    - Total of Expenses
    - Individual Expenses
- Expenses of ***Food***
    - Total of Expenses
    - Individual Expenses
- Expenses of ***Entertainment***
    - Total of Expenses
    - Individual Expenses
- Expenses of ***Health***
    - Total of Expenses
    - Individual Expenses

The first three entries can be guaranteed to be the **Username**, **Income**, and **Total Spent**. The rest of the expenses can be distinguished by the *ExpenseType* field. Here is an example of what one expense category may look like:

```JSON
{
   "ExpenseType": "Food",
   "Total": 122
},
{
   "ExpenseType": "Food",
   "ExpenseName": "Grocery Store",
   "Amount": 109
},
{
   "ExpenseType": "Food",
   "ExpenseName": "Fast Food",
   "Amount": 13
}
```


## Update Income Handler
### Path - "dashboard/income/update"
This is a POST function that edits the income of the user. It takes in a json that contains the *Username*, and *Amount*. Here is an example input:
```JSON
{
    "Username": "test",
    "Amount": 6900
}
```
If the update is successful, the returned message will be "*Income Updated*." If the Username is incorrect, an error will return stating "*Failed to find user, username does not exist*."


## Input Format of ExpenseTypes
When inputting the ExpenseType field to the json to be sent to the API, there are 5 possible types of expenses to choose from. They must be written in these specific ways in order for the input to be valid and function correctly:
```
Home & Utilities --> "HomeUtils"
Travel           --> "Travel"
Food             --> "Food"
Entertainment    --> "Ent"
Health           --> "Health"
```


## Add Expense Handler
### Path - "dashboard/expense/add"
This is a POST function that adds a new expense to the database. It takes in a json that contains the *Username*, *ExpenseType*, *ExpenseName*, *Amount*. Here is an example input:
```JSON
{
   "Username": "ex",
   "ExpenseType": "Food",
   "ExpenseName": "Grocery Store",
   "Amount": 109
}
```
If the addition is successful, the returned message will be "*Expense Added*." If the Expense Type is incorrect, an error will return stating "*Failed to create expense, expense type not found*." If the Username is incorrect, an error will return stating "*Failed to find user, username does not exist*."


## Update Expense Handler
### Path - "dashboard/expense/update"
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


## Delete Expense Handler
### Path - "dashboard/expense/delete"
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