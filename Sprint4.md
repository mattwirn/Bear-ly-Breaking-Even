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

In sprint 4, we wrote 4 cypress tests which tested the login error handling and the dashboard functionality.

- The first test for the login error message tests the component with the text, html and css for the message. It tested whether the text renders on the page in plain html.
- The second login error test checks whether the entire component—the html, css, and text—renders on the screen when necessary. It types information into the text fields and clicks the login button twice: once with correct information and once with incorrect info. 
- The third test we wrote checks the component for the tables in the dashboard page. It first renders the plain html onto the screen. Then it checks if adding and deleting rows works.
- The final test checks whether the tables render onto the screen properly. It first visits the dashboard page, and then clicks the “show expenses” button and checks if the tables render.

## Back End
For the new features, we tested the helper getter methods that the Dashboard handler uses to get data from the tables.

 - "TestDashboardUsername" tests the "getUsername" method to see if it returns the correct username.
 - "TestDashboardIncome" tests the "getIncome" method to see if it returns the correct income.
 - "TestDashboardTotalSpent" tests the "getTotalSpent" method to see if it indeed returns the correct amount for the user's total spent.
 - "TestDashboardHU" tests the "getHomeUts" method to see if it indeed returns the correct total amount spent for the Home & Utilities category.
 - "TestDashboardT" tests the "getTravel" method to see if it indeed returns the correct total amount spent for the Travel category.
 - "TestDashboardIncomeF" tests the "getFood" method to see if it indeed returns the correct total amount spent for the Food category.
 - "TestDashboardIncomeE" tests the "getEnt" method to see if it indeed returns the correct total amount spent for the Entertainment category.
 - "TestDashboardIncomeH" test the "getHealth" method to see if it indeed returns the correct total amount spent for the Health category.
