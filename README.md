# gin-middleware API

A simple API created using gin framework for authentication and logging requests using middleware.

Server port: `8080`
JWT Token used: `HMAC SHA-256 algorithm `

## 1. Middleware:

Two middleware functions are applied to all routes:
1. LoggerMiddleware: Logs information about incoming requests, including the HTTP method, URL path, client IP address, and request processing time.
2. AuthMiddleware: Validates JWT tokens for routes under the `/sam-api` group. It checks for the presence of the Authorization header containing a valid JWT token.

## 2. Login Endpoint (/sam-login):

Handles `POST` request for user authentication and returns a JWT token in the response body.

## 3. Secured Endpoint (/sam-api/secured):

Only accessible if the `GET` request includes a `valid JWT token` in the _Authorization header_ and returns a message indicating a secure route in the response body.

## 4. Authentication:

If the request made to the **secured endpoint** has a valid token in the _Authorization header_ is valid, then the middleware allows the the request to proceed to the `Secured Handler`.

# Steps to run the API:

1. Run `go mod init gin-middleware`
2. Run `go mod tidy`
3. Finally run `go run main.go`

## 1. Login endpoint on Postman:

1. Use POST request `localhost:8080/sam-login`
2. In body, select `x-www-form-urlencoded`
3. Add the key value pair as follows, 
```
username: sam
password: gin-gonic
```

4. We will get a unique `<token>` in the response body as follows,
```
{
    "token": <token>
}
```

## 2. Secured Endpoint:

1. Use GET request `localhost:8080/sam-api/secured`
2. In Headers, add key value pair as follows where we paste the `<token>` from the login endpoint's response,
```
Authorization: <token>
```
3. The desired output is as follows,
```
{
    "message": "You have entered a secured route"
}
```