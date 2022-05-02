# aspire-mini-api

This repo contains an attempt to create a mini version of the API at Aspire App.
Although some punches have been pulled, significant effort was made to manage the data models and the business logic required.

The API routes are located in the handlers directory and the models directory contains the data and the relationships.
There is no database connection and it keeps a slice of users as JSON in memory.

## Build instructions

Clone this repository and execute the below command to start the server

```go run main.go```

This will start the server at `localhost:8080` by default.