# Movie CRUD API

A simple **Movie CRUD API** built using **Go (Gin)** and **GORM** for managing movies with **user authentication**. This project allows users to **create**, **read**, **update**, and **delete** movies. It also provides JWT-based authentication to access protected routes.

## Table of Contents

- [Overview](#overview)
- [Technologies Used](#technologies-used)
- [Setup and Installation](#setup-and-installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Database Schema](#database-schema)
- [Contributing](#contributing)
- [License](#license)

## Overview

This is a simple **Movie CRUD API** that allows users to manage movies by performing CRUD operations. It uses **Go (Gin)** framework for handling HTTP requests and **GORM** ORM for PostgreSQL database interaction. The project implements **user authentication** through JWT tokens.

### Features:
- **JWT-based Authentication** for secure access
- **Create, Read, Update, Delete** (CRUD) functionality for movies
- Simple **role-based access control** (admin access for movie management)

## Technologies Used

- **Go (Golang)**: A statically typed language for building efficient and scalable applications.
- **Gin**: A web framework for Go, offering a high-performance HTTP web server.
- **GORM**: An ORM (Object Relational Mapping) library for Go, used for interacting with the database.
- **PostgreSQL**: A powerful, open-source relational database system for storing movie and user data.
- **JWT**: JSON Web Token for secure user authentication.
- **GitHub**: To host the repository.

## Setup and Installation

### Prerequisites

Ensure you have the following installed:
- Go (1.18+)
- PostgreSQL
- Git

### Clone the repository

Start by cloning the repository:

```bash
git clone https://github.com/SAUIR04/movie-crud-api.git
cd movie-crud-api

Install dependencies
Install the required Go modules using:

bash
Копировать
go mod tidy
Setup PostgreSQL Database
Make sure your PostgreSQL database is running.

Create a new database called moviecrudapi (or modify the connection string in the code if necessary).

bash
Копировать
psql -U postgres
CREATE DATABASE moviecrudapi;
Update the database connection string in the main.go file if needed.

Run the API
Once everything is set up, run the API with:

bash
Копировать
go run ./cmd/main.go
The server should be running at http://localhost:8080.

Usage
Authentication
To access the protected endpoints, you need to authenticate and obtain a JWT token.

Login to receive a token:

bash
Копировать
POST /login
Request Body:

json
Копировать
{
  "username": "admin",
  "password": "admin123"
}
Response:

json
Копировать
{
  "token": "your-jwt-token-here"
}
Use the token in subsequent requests by including it in the Authorization header as Bearer <token>.

Endpoints
Public Endpoints
GET /api/movies – Get a list of all movies.

GET /api/movies/{id} – Get details of a specific movie by ID.

Protected Endpoints (Admin Only)
These endpoints require authentication using a JWT token.

POST /api/movies – Create a new movie.

PUT /api/movies/{id} – Update an existing movie.

DELETE /api/movies/{id} – Delete a movie by ID.

Example Request:
bash
Копировать
curl -X POST http://localhost:8080/api/movies \
-H "Authorization: Bearer <token>" \
-d '{
  "title": "Inception",
  "description": "A thief who steals corporate secrets...",
  "image_url": "http://example.com/inception.jpg",
  "genre": "Sci-Fi",
  "release_date": "2010-07-16"
}'
Database Schema
Here is the structure of the PostgreSQL database:

users table:
sql
Копировать
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  email TEXT
);
movies table:
sql
Копировать
CREATE TABLE movies (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  image_url TEXT,
  genre TEXT,
  release_date DATE
);
Contributing
If you'd like to contribute to this project, feel free to fork the repository, create a new branch, and submit a pull request. Here's how you can contribute:

Fork the repository.

Create a new branch for your feature or bugfix.

Make your changes.

Test your changes thoroughly.

Create a pull request with a description of what you've done.

