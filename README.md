# Pond Management API

## Introduction

Welcome to the Pond Management API, a robust and efficient RESTful web service designed to simplify pond and fish population management. This API is built with Golang, GIN, and JWT for security, ensuring a seamless experience for developers and end-users.

## Key Endpoints

### User Authentication

- `POST /login`: Securely log in and obtain an access token.
- `POST /register`: Register as a new user and gain access to the full suite of features.

### Pond Management

- `GET /pond`: Retrieve a list of all ponds associated with the logged-in user.
- `POST /pond`: Create a new pond with ease.
- `GET /pond/:id_pond`: Access detailed information about a specific pond.
- `PUT /pond/:id_pond`: Update and fine-tune pond details.
- `DELETE /pond/:id_pond`: Remove ponds that are no longer needed.

### Fish Population Tracking

- `GET /pond/:id_pond/fish`: Explore the fish populations residing in a specific pond.
- `POST /pond/:id_pond/fish`: Add new fish to the selected pond.
- `GET /pond/:id_pond/fish/:id_fish`: Dive deep into the specifics of an individual fish.
- `PUT /pond/:id_pond/fish/:id_fish`: Update the information of a particular fish.
- `DELETE /pond/:id_pond/fish/:id_fish`: Remove fish from the pond as required.

### Stact
- Golang
- Gin
- JWT
