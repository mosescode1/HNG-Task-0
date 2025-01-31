# HNG-Task-0

This is a simple web application built with the [Fiber](https://gofiber.io/) web framework for Go. The application serves a single endpoint that returns a JSON response with the current date and time, an email address, and a GitHub URL.

## Requirements

- Go 1.23.5 or later

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/mosescode1/HNG-Task-0.git
    cd HNG-Task-0
    ```

2. Install the dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:

    ```sh
    go run main.go
    ```

2. Open your browser and navigate to `http://localhost:3000`. You should see a JSON response similar to the following:

    ```json
    {
        "email": "iammoses19@gmail.com",
        "current_datetime": "2023-10-05T14:48:00Z",
        "github_url": "https://github.com/mosescode1/HNG-Task-0.git"
    }
    ```

## Project Structure

- : The main application file that sets up the Fiber web server and defines the endpoint.

## Dependencies

This project uses the following dependencies:

- [Fiber](https://github.com/gofiber/fiber): An Express-inspired web framework for Go.
- [time](https://pkg.go.dev/time): A standard Go package for working with dates and times.