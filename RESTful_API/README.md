Develop RESTful API with Go and Gin
https://go.dev/doc/tutorial/web-service-gin 

Writing a RESTful web service API with Go 
and Gin Web Framework

Gin simplifies coding tasks associated with building web apps and web services.

Use Gin to route requests, retrieve request details and marshal JSON for responses.

Buld RESTful
Representational State Transfer is a software architectural style for the WWw.

It means that a server will respond with the representation of a resource (most often a HTML document).

Design a RESTful API server with two endpoints. 

API will provide access to a store selling vintage records.
Endpoints for a client to get and add albums for users.

When building an API, you usually start by designing the endpoints.

APIs users have more success if the endpoints are easy to understand.

/albums
GET
POST

/albums/:id
GET

data will be stored in memory for simplicity

APIs normally interact with a database

storing in memory means the set of albums will be lost each time you stop the server and recreated when you start it

add the github.com/gin-gonic/gin module as a dependency for your module. Use a dot argument to mean “get dependencies for code in the current directory.”

go get . 

go run .
directory containing main.go. Use a dot argument to mean “run code in the current directory.”

use curl to make a request to your running web service
curl http://localhost:8080/albums



                                                                                                        



