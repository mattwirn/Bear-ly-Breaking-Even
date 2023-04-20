## Link to Video Presentation
[Click me]

## Front End

In this sprint, the dashboard page was further worked on, with many aspects changed and new functionality added:
- Welcome message changed
- Page format changed to be centered
- Added interactive tabs for each individual expense type
- Created dynamic tables for each tab and expense group 
- Rows of individual expenses can now be added and deleted 
- Information remains on tables after component reload
- Total amount from "Expense Amount" row dynamically added and displayed for each category
- Removed Login and Signup redirect options from Dashboard page

New Cypress tests have been written to test the new features:
-

## Back End
For the new features, we tested the helper getter methods that the Dashboard handler uses to get data from the tables.

"TestDashboardUsername" tests the "getUsername" method to see if it returns the correct username.
"TestDashboardIncome" tests the "getIncome" method to see if it returns the correct income.
"TestDashboardTotalSpent" tests the "getTotalSpent" method to see if it indeed returns the correct amount for the user's total spent.
"TestDashboardHU" tests the "getHomeUts" method to see if it indeed returns the correct total amount spent for the Home & Utilities category.
"TestDashboardT" tests the "getTravel" method to see if it indeed returns the correct total amount spent for the Travel category.
"TestDashboardIncomeF" tests the "getFood" method to see if it indeed returns the correct total amount spent for the Food category.
"TestDashboardIncomeE" tests the "getEnt" method to see if it indeed returns the correct total amount spent for the Entertainment category.
"TestDashboardIncomeH" test the "getHealth" method to see if it indeed returns the correct total amount spent for the Health category.
