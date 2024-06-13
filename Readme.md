# Go REST API Example

## Overview

This is a simple example of a REST API written in Go. It is a simple API that gives an opportunity to create, get, update and delete comments.

## Getting Started

To get started with the API, you will need to have Go go-task installed. You can install it by running the following command:

```bash
go get -u github.com/go-task/task/v3
```

Then, you can run the API by running the following command:

```bash
task run
```

You will also need to install docker and docker-compose. You can install them by running the following command:

```bash
brew install docker docker-compose
```

You will be able to access the API at http://localhost:8080/api/v1/comments.

## Endpoints

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