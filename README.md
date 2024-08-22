[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=FelipeAJdev_dev-cloud-challenge&metric=alert_status)](https://sonarcloud.io/dashboard?id=FelipeAJdev_dev-cloud-challenge)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=FelipeAJdev_dev-cloud-challenge&metric=bugs)](https://sonarcloud.io/dashboard?id=FelipeAJdev_dev-cloud-challenge)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=FelipeAJdev_dev-cloud-challenge&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=FelipeAJdev_dev-cloud-challenge)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=FelipeAJdev_dev-cloud-challenge&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=FelipeAJdev_dev-cloud-challenge)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=FelipeAJdev_dev-cloud-challenge&metric=sqale_index)](https://sonarcloud.io/dashboard?id=FelipeAJdev_dev-cloud-challenge)

# Dev Cloud Challenge

This project is a RESTful API developed in Go as part of a challenge for a Web Development and Cloud course. The main goal of this project is to implement a CRUD (Create, Read, Update, Delete) application for managing student data within a school, utilizing PostgreSQL for data persistence.

## Features

- **CRUD Operations**: Manage student data with create, read, update, and delete operations.
- **RESTful API**: Follows REST principles, making it easy to integrate with other systems.
- **Database Integration**: Uses PostgreSQL for storing and managing data.
- **Database Migrations**: Automatically applies database migrations to set up the schema.
- **Containerization**: The project includes Docker configurations for easy deployment and testing.
- **Database Management**: Includes pgAdmin for managing the PostgreSQL database through a web interface.

## Technologies Used

- **Go**: The programming language used to develop the application.
- **PostgreSQL**: Database used for data persistence.
- **Docker**: Used for containerization and easy setup of the development environment.
- **GitHub Actions**: Set up for CI/CD to automate the testing and deployment processes.

## How to Run the Project

### Prerequisites

- **Go 1.19 or higher**
- **Docker** (if using the Docker setup)
- **PostgreSQL** (if running the database locally)

### Docker Compose Setup

The project includes a `docker-compose.yml` file that sets up the following services:

- **PostgreSQL**: A PostgreSQL database server.
- **pgAdmin**: A web-based database management tool for PostgreSQL.

#### Steps to Run

1. **Clone the repository**:
   ```bash
   git clone https://github.com/FelipeAJdev/dev-cloud-challenge.git
   cd dev-cloud-challenge
   ```

2. **Set up the environment**:
   - Ensure the environment variables for the database are configured in a `.env` file:
     ```
     WSRS_DATABASE_PORT=5432
     WSRS_DATABASE_USER=postgres
     WSRS_DATABASE_PASSWORD=yourpassword
     WSRS_DATABASE_NAME=wsrs
     ```

3. **Run the services using Docker Compose**:
   ```bash
   docker-compose up
   ```

   - This will start both the PostgreSQL database and the pgAdmin interface. pgAdmin will be accessible at `http://localhost:8081` with the default credentials provided in the `docker-compose.yml` file.

4. **Apply Database Migrations**:
   - Ensure that the migrations are automatically applied when the application starts. If needed, you can run the migrations manually using the tool or method defined in your Go application.

5. **Run the application**:
   ```bash
   go run main.go
   ```

6. **Access the API**:
   - The API will be available at `http://localhost:8080`.

## Endpoints

### Students

- **GET /students**: Retrieve all students.
- **POST /students**: Create a new student record.
- **PUT /students/{id}**: Update an existing student record.
- **DELETE /students/{id}**: Delete a student record.

## Documentation

The API documentation, including the Swagger UI, is available at the following link:

[Swagger Documentation](https://dev-cloud-challenge-b3f5485f2dcf.herokuapp.com/swagger/index.html)

## Contribution

Feel free to fork the repository, submit issues, and open pull requests. Contributions are welcome!<!--
## License

This project is open-source and available under the [MIT License](LICENSE).
-->
