# cms_automotoresmw
Web application called cms_automotoresmw, which lets car dealer shops administrate car fleet, clients, sales, and more.

| **Method** | **URL Pattern**   | **Handler**       | **Action**                                    |
|------------|-------------------|-------------------|-----------------------------------------------|
| GET        | /                 |  home             | Display the home page                         |
| GET        | /car/view/:id     | carView           | Display a specific car                        |
| GET        | /car/create       | carCreate         | Displa a HTML form for crating a new car      |
| POST       | /car/create       | carCreatePost     | Create a new car                              |
| GET        | /user/signup      | userSignup        | Display a HTML form for signing up a new user |
| POST       | /user/signupPost  | userSignupPost    | Create a new user                             |
| GET        | /user/login       | userLogin         | Display a HTML form for logging in a user     |
| POST       | /user/login       | userLoginPost     | Authenticate and login the user               |
| POST       | /user/logout      | userLogoutPost    | Logout the user                               |
| GET        | /static/*filepath | http.FileServer   | Serve a specific static file                  |

# Test the endpoint throw the terminal.
* $ curl -i -X GET  http://localhost:4000/
* $ curl -i -X GET  http://localhost:4000/car/view/123.
* $ curl -i -X GET http://localhost:4000/car/create
* $ curl -iL -X POST http://localhost:4000/car/create
* $ curl -i -X GET http://localhost:4000/static/
