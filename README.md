# User Microservice

This is a simple user management microservice built using Go Fiber, a web framework for Go. It provides basic CRUD (Create, Read, Update, Delete) operations for managing user data.

## Features

- **Create User**: Add a new user with a unique ID, name, and email address.
- **Read User**: Retrieve details of a specific user by their ID.
- **Update User**: Update the name and email address of an existing user.
- **Delete User**: Remove a user from the system based on their ID.
- **Get All Users**: Fetch a list of all users currently stored in the system.

## Installation

1. Make sure you have Go installed. If not, you can download it from the [official Go website](https://golang.org/dl/).

2. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/aselasperera/User-Microservice
   ```

3. Navigate to the project directory:
   ```bash
   cd user-microservice
   ```

4. Install dependencies:
   ```bash
   go mod tidy
   ```

5. Run the application:
   ```bash
   go run main.go
   ```

The microservice will start running on `http://localhost:3000`.

## Usage

You can interact with the microservice using various HTTP methods (GET, POST, PUT, DELETE) to manage user data. Here are some example API endpoints:

- **GET /users**: Get a list of all users.
- **GET /users/{id}**: Get details of a specific user by ID.
- **POST /users**: Create a new user.
- **PUT /users/{id}**: Update an existing user.
- **DELETE /users/{id}**: Delete a user.

You can use tools like `curl`, Postman, or any HTTP client to test these endpoints.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
