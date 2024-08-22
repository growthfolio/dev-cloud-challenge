
# Dev Cloud Challenge 

This project is a RESTful API developed in Go as part of a challenge for a Web Development and Cloud course. The main goal of this project is to implement a CRUD (Create, Read, Update, Delete) application for managing student data within a school, utilizing PostgreSQL for data persistence.

## Features

- **CRUD Operations**: Manage student data with create, read, update, and delete operations.
- **RESTful API**: Follows REST principles, making it easy to integrate with other systems.
- **Database Integration**: Uses PostgreSQL for storing and managing data.
- **Containerization**: The project includes Docker configurations for easy deployment and testing.

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

### Steps to Run

1. **Clone the repository**:
   ```bash
   git clone https://github.com/FelipeAJdev/dev-cloud-challenge.git
   cd dev-cloud-challenge
   ```

2. **Set up the environment**:
   - If using Docker, you can simply run:
     ```bash
     docker-compose up
     ```
   - If running locally, make sure to configure your PostgreSQL credentials in the environment variables.

3. **Run the application**:
   ```bash
   go run main.go
   ```

4. **Access the API**:
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

Feel free to fork the repository, submit issues, and open pull requests. Contributions are welcome!
