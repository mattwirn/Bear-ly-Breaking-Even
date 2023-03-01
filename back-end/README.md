To start the backend, cd into the back-end folder and run the command in the terminal:

> go run .

It should say connected to database and then any request made should pop up in the terminal.
To see what response you should get when calling the api, for signup it should say
> Signed up

If you see that with no errors it worked and created a new user in the database.
The back end is running on localhost:8080 and you should see the frontend when you visit it.
If not, something is broken.

If you want to check out more about the signup function, its located in *handlers/posts.go*
