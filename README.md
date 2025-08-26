# My IP
================

My IP is a web application that displays the user's IP address and its reversed version. The application is built using Go as the backend language and uses a PostgreSQL database to store IP addresses.

## Table of Contents
---------------

- [My IP](#my-ip)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
  - [Project Structure](#project-structure)
  - [Backend](#backend)
  - [Frontend](#frontend)
  - [Database](#database)
  - [API Endpoints](#api-endpoints)
  - [Build and Deployment](#build-and-deployment)
  - [Testing](#testing)

## Getting Started
---------------

To get started with the project, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/your-repo-name.git`
2. Change into the project directory: `cd your-repo-name`
3. Install the required dependencies: `go get -u ./...`
4. Start the application: `go run main.go`

## Project Structure
-----------------

The project is structured as follows:

* `client`: contains the frontend code, including the HTML, CSS, and JavaScript files.
* `server`: contains the backend code, including the Go files and the database schema.
* `internal`: contains the internal packages, including the database and storage packages.
* `migrations`: contains the database migration files.

## Backend
---------

The backend is built using Go and uses the Gin framework to handle HTTP requests. The backend is responsible for:

* Handling API requests
* Interacting with the database
* Reversing IP addresses

## Frontend
---------

The frontend is built using HTML, CSS, and JavaScript. The frontend is responsible for:

* Displaying the user's IP address and its reversed version
* Handling user input
* Making API requests to the backend

## Database
---------

The database is built using PostgreSQL and is used to store IP addresses. The database schema is defined in the `migrations` directory.

## API Endpoints
-------------

The application has the following API endpoints:

* `GET /myip`: returns the user's IP address and its reversed version
* `GET /health`: returns the health status of the application

## Build and Deployment
----------------------

The application is built and deployed using Docker. The `Dockerfile` is located in the `server` directory and is used to build the Docker image. The `docker-compose.yml` file is used to define the Docker containers.

## Testing
---------

The application has unit tests and integration tests. The unit tests are located in the `internal` directory and the integration tests are located in the `server` directory.

I hope this helps! Let me know if you have any questions or need further clarification.