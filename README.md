<p align="center">
  <img src="https://github.com/k5sha/goBook/blob/master/media/logo_.png" alt="Logo" width="256"/>
</p>

<div align="center">
  
  ![Go Badge](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
  ![Docker Badge](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
  ![REST API Badge](https://img.shields.io/badge/REST_API-34B7B1?style=for-the-badge&logo=swagger&logoColor=white)
  ![GORM Badge](https://img.shields.io/badge/GORM-00A9E0?style=for-the-badge&logo=gorm&logoColor=white)
  ![Gin Badge](https://img.shields.io/badge/Gin-8B8B00?style=for-the-badge&logo=gin&logoColor=white)
</div >

# GoBook
This project demonstrates how to create a simple but scalable API using Go (Golang), Gin, and GORM. 
It includes endpoints for managing workbooks with functions such as adding, searching, updating, and deleting workbook records. 
The project also integrates PostgreSQL for data storage and uses the Aconfig package for configuration management. 
This guide provides a step-by-step approach to installing and configuring the API via Docker.

### How it works:
- **Gin** is used for handling HTTP requests, routing, and middleware, making the API fast and efficient. It simplifies tasks like validation and error handling.
- **GORM** is used to interact with the PostgreSQL database, allowing seamless CRUD operations with Go structs instead of raw SQL queries, making database management easier and more efficient.
<p align="center">
  <img src="https://github.com/k5sha/goBook/blob/master/media/how.jpg" alt="How work" width="726"/>
</p>

### Startup
- via Docker
```bash
docker compose -f docker-compose.dev.yml up -d
```
- via bash 
```bash
go mod download
go run cmd/main.go
```

### Configuration
The application uses a configuration file, **config.yaml**, to define environment-specific settings, such as the server port and database connection details.

Here is an example of the config.local.yaml file:
> [!NOTE]
> If you are using docker to run the database on localhost, you should use `db` instead of `localhost` in the `database_dsn` clause of the config.yaml file
```yaml
port: ':3000'
database_dsn: 'postgres://postgres:password@localhost:5432/books_db?sslmode=disable'
```
#### Explanation:
- `port: ':3000'`
This sets the port on which the API server will run. By default, the server will listen on port 3000.

- `database_dsn: 'postgres://postgres:password@localhost:5432/books_db?sslmode=disable'`
  
This is the Data Source Name (DSN) for connecting to the PostgreSQL database. It contains the necessary credentials and database information:

- `postgres`: The username for the database.
- `password`: The password for the database.
- `localhost`: The hostname of the database server (in this case, it’s running locally).
- `5432`: The port on which the PostgreSQL database is listening (default is 5432).
- `books_db`: The name of the database to connect to.
- `sslmode=disable`: This disables SSL encryption for the connection, which is commonly used in local development environments.
  
By editing this file, you can easily change the server's port or adjust the database connection settings without modifying the source code.

### Nice to have features (backlog)
- [ ]  Add swagger
- [ ]  Add a separate author model to the database
- [ ]  Add tests
- [ ]  Add comments
- [ ]  Add to config the 'release' mode

### Author:
**Yurii (k5sha) Yevtushenko**
