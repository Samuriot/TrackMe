# Go Backend (Built with Fiber Framework)
This will be the backend responsible for handling all requests from the frontend applications. The primary purpose of this Backend is to serve an API server for any frontend application (if authenticated properly) to use for fetching data

## General Datapath for Requests Through Backend
```
HTTP -> Routes -> Middleware -> Handlers -> Service -> Repository -> Database
```

## Routes
- This directory serves the paths for HTTP requests to navigate (e.g, GET /api/users)

## Models
- This directory holds all the data models for our application, pulling the data from a DB

## Middleware
- This directory will hold all things middleware. All HTTP requests will pass through the middleware to attach items like JWT, context, etc.

## Handlers
- This directory will hold the HTTP controller logic for every request. They will be responsible for error coding, sending 200 OK, etc.

## Services
- This directory will handle the business logic of the backend. They will handle data validation, business logic rules, etc.

## Repository
- This directory will handle querying the Mongo Collection relative to that specific request.