# Go JWT example

Go application that implements JWT based authentication and authorization. 

This JWT example using RSA256 algorithm claims, so need private and public key. 
You can use [Cryptotools](https://cryptotools.net/rsagen) for generate your own private and public key.
Change contents of `jwt-private.key` and `jwt-public.key` inside `files/cert` with your own key.

First, download the library:

```sh
go get 
```

To run this application, build and run the Go binary:

```sh
go build
./go-jwt
```

Now, using any HTTP client (like [Postman](https://www.getpostman.com/apps)) make a sign-in request with the appropriate credentials:

```
POST http://localhost:2345/login

{"email":"admin@mail.com","password":"admin"}
```

After receive the refresh token, place token to Authorization header at HTTP Client.
You can now try hitting the home route from the same client to get the home message:

```
GET http://localhost:2345/home
```

You can use

