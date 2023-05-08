# User Service REST API
This repository contains a RESTful API for user service, which provides endpoints for user authentication, user management, and favorite cities management.

## Authentication Endpoints
### Registration
To register a new user, send a POST request to `/api/auth/registration` with the following request body:
```json
{
  "email": "email",
  "password": "password"
}
```
### Login
To authenticate a user, send a POST request to `/api/auth/login` with the following request body:
```json
{
  "email": "email",
  "password": "password"
}
```
On successful authentication, the server returns a JSON object containing a JWT token, which can be used to authenticate subsequent requests.
## User Management Endpoints
### Delete User
To delete a user, send a DELETE request to `/api/user/:id` with the following request body:
```json
{
  "email": "email",
  "password": "password"
}
```
This endpoint requires a valid JWT token to be included in the request header.
### Get All Users
To get all users, send a GET request to `/api/user/all`.

## Favorite Cities Endpoints
### Get Favorite Cities
To get a user's favorite cities, send a GET request to `/api/user/favorite-cities/:id`.

### Add Favorite City
To add a new favorite city for a user, send a POST request to `/api/user/favorite-cities` with the following request body:
```json
{
  "user_id": "id",
  "city_name": "city_name"
}
```
This endpoint requires a valid JWT token to be included in the request header.
### Delete Favorite City
To delete a favorite city for a user, send a DELETE request to `/api/user/favorite-cities` with the following request body:
```json
{
  "user_id": "id",
  "city_name": "city_name"
}
```
This endpoint requires a valid JWT token to be included in the request header.