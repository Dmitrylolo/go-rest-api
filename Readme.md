# Go REST API Example

## Overview

This is a simple example of a REST API written in Go. It is a simple API that gives an opportunity to create, get, update and delete comments.

## Getting Started

To get started with the API, you will need to have Go, Docker, Docker Compose and Taskfile installed on your machine. Once you have these tools installed, you can follow these steps to get the API up and running:

1. Clone the repository:

```bash
git clone https://github.com/go-rest-api-example/go-rest-api-example.git
```

2. Change into the project directory:

```bash
cd go-rest-api-example
```

3. Build the Docker and run server: 
```bash
task run
```
4. You will be able to access the API at http://localhost:8080/api/v1/comments.

## Endpoints

### Create a Comment

To create a comment, you can use the following endpoint:

```
POST /api/v1/comments
```

The request body should contain the comment data.

### Get All Comments

To get all comments, you can use the following endpoint:

```
GET /api/v1/comments
```

### Get a Comment

To get a comment, you can use the following endpoint:

```
GET /api/v1/comments/:id
```

### Update a Comment

To update a comment, you can use the following endpoint:

```
PUT /api/v1/comments/:id
```

The request body should contain the updated comment data.

### Delete a Comment

To delete a comment, you can use the following endpoint:

```
DELETE /api/v1/comments/:id
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details