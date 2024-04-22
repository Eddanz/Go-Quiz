# Go Quiz

A simple quiz application implemented in Go, featuring a server, a command-line interface (CLI), and API handlers.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Server](#server)
  - [CLI](#cli)
- [API and Models](#api-and-models)
  - [HTTP Handlers](#http-handlers)
  - [Data Models](#data-models)
- [License](#license)

## Installation

### Prerequisites
- Go installed

### Steps
1. Clone the repository: `git clone https://github.com/your-username/quiz.git`
2. Navigate to the project directory: `cd quiz`

## Usage

### Server

The server component of the application exposes endpoints to retrieve questions and submit answers.

#### How to Run
1. Navigate to the `/cmd/server` directory.
2. Build the server binary: `go build -o quiz-server.exe`
3. Run the server: `./quiz-server.exe`

#### Endpoints

- `/questions`: GET endpoint to retrieve questions.
- `/answers`: POST endpoint to submit answers.

### CLI

The CLI component allows users to interact with the quiz via the command line.

#### How to Run
1. Navigate to the `/cmd/cli` directory.
2. Build the CLI binary: `go build -o quiz-cli.exe`
3. Run the CLI: `./quiz-cli.exe`

## API and Models

### HTTP Handlers

API handlers for serving questions and receiving answers.

#### How to Run
1. Navigate to the `/api` directory.
2. Build the handlers binary: `go build -o handlers.exe`
3. Run the handlers: `./handlers.exe`

### Data Models

Data models for representing questions and quizzes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
