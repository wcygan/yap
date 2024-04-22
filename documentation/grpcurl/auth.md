# gRPC references for the auth endpoints

## Register

```
grpcurl -proto proto/auth/v1/auth.proto -d '{"username": "your_username", "password": "your_password"}' -plaintext localhost:50050  auth.AuthService/Register
```

## Login

```
grpcurl -proto proto/auth/v1/auth.proto -d '{"username": "your_username", "password": "your_password"}' -plaintext localhost:50050 auth.AuthService/Login
```

## Validate

```
grpcurl -proto proto/auth/v1/auth.proto -d '{"access_token": "your_access_token"}' -plaintext localhost:50050 auth.AuthService/Validate
```

## Refresh

```
grpcurl -proto proto/auth/v1/auth.proto -d '{"refresh_token": "your_refresh_token"}' -plaintext localhost:50050 auth.AuthService/Refresh
```