# Ascii-Art-Web-Dockerize
## Table of Contents


- [Description](#description)
- [Dockerization](#dockerization)
- [Authors](#authors)
- [References](#references)

## Description

The `ASCII Art Web Generator` is a web application that converts user input text into ASCII art. Users can enter text, select a font style, and view the generated ASCII art directly on the web page. The application is built using Go and serves the web pages and handles requests efficiently.
Dockerization

To simplify deployment and ensure a consistent environment across different machines, the application has been Dockerized. Docker is a platform that allows developers to package applications into containers—standardized executable components that combine application source code with the operating system (OS) libraries and dependencies required to run that code in any environment.

### By Dockerizing the ASCII Art Web Generator, we achieve the following benefits:

- Portability: The application runs consistently across different environments (development, testing, production) without compatibility issues.
- Isolation: The application and its dependencies are isolated from other applications, preventing conflicts.
- Scalability: Docker makes it easier to scale the application horizontally by running multiple containers.
- Simplified Deployment: Deploying the application becomes straightforward, as all dependencies are included within the Docker image.

### How Docker Works

- Dockerfile: A script that contains instructions on how to build a Docker image. It specifies the base image, copies the application files, installs dependencies, and defines the command to run the application.
- Docker Image: A snapshot of the application and its environment, created based on the Dockerfile. The image can be used to run containers.
- Docker Container: A running instance of a Docker image. Containers are lightweight and can be easily started, stopped, and moved across different environments.

## How to Run Locally
1. Clone the repository:
    ```go
    git clone https://learn.zone01kisumu.ke/git/masman/ascii-art-web-dockerize.git
    cd ascii-art-web-dockerize
    ```
2. Install Go and set up your Go environment if you haven't already. You can download and install Go from the official Go website.

3. Run the application:

    ```go
    go run main.go
    ```
4. Open your web browser and navigate to http://localhost:8080/.

### Implementation Details: Algorithm

The ASCII Art Generator utilizes a series of Go packages to manage different aspects of the application. Here's an overview of the implementation:

1. **Main Server Setup:** The main function sets up an HTTP server that listens on port 8080. It defines several route handlers: 
`/` - for the homepage
`/submit` - for handling form submissions

2. **Frontend HTML:** The HTML file serves as the user interface for the application. It includes a form where users can input text and select a font file. The result area displays the generated ASCII art.

3. **ASCII Art Generation:** The core logic for converting text into ASCII art is handled in the functions package. The Ascii function reads the font data, processes the input text, and generates the corresponding ASCII art.

4 **Backend Handlers:** The backend handlers manage the routing and processing of user requests. The HandleRequest function handles form submissions, generates ASCII art, and returns the result to be displayed on the web page.

With these details, you should be able to understand and run the ASCII Art Web Generator application effectively. Feel free to explore the code and modify it according to your needs.

## Dockerization
**How to Build and Run with Docker**

1. Build the Docker image:
    ```sh
    docker build -t ascii-art-web .
    ```
2. Run the Docker container:

    ```bash
    docker run -d -p 8080:8080 --name web-app ascii-art-web
    ```
3. Open your web browser and navigate to http://localhost:8080/.

## Authors

- [Malika](https://learn.zone01kisumu.ke/git/masman)
- [Cliff](https://learn.zone01kisumu.ke/git/clomollo)

## References

- ASCII MANUAL and HTML

## License
This project is licensed under the MIT protection act. [LICENSE](./LICENSE)