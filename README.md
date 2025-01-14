# File Storage and Sharing Application Backend

## Overview

This is a backend-only application for file storage and sharing, built using Go. It allows users to securely upload, store, and retrieve files.

## Features

- **File Upload and Download**: Basic file storage and retrieval.
- **File Sharing**: Generate shareable links for files.
- **Authentication**: JWT-based authentication.
- **Scalable Storage**: Supports local or S3-compatible storage backends.

## Technology Stack

- **Backend**: Go, GorillaMux framework
- **Storage**: Local filesystem or MinIO
- **Database**: PostgreSQL

## Folder Structure

```
file-storage-app/
├── cmd/
│   └── main.go                # Entry point for the application
├── config/
│   ├── config.go              # Configuration loader
│   └── app.yaml               # Application configuration file
├── internal/
│   ├── auth/
│   │   ├── auth.go            # JWT-based authentication logic
│   │   └── middleware.go      # Authentication middleware
│   ├── file/
│   │   ├── upload.go          # File upload handling
│   │   ├── download.go        # File download handling
│   │   └── service.go         # Core file service logic
│   └── utils/
│       ├── logger.go          # Logging utility
│       └── helpers.go         # Common helper functions
├── pkg/
│   ├── database/
│   │   ├── postgres.go        # PostgreSQL connection and setup
│   │   └── migrations/
│   │       └── init.sql       # Initial database schema
│   └── storage/
│       ├── storage.go         # Storage interface
│       ├── local.go           # Local filesystem implementation
│       └── s3.go              # S3-compatible storage implementation
├── test/
│   ├── auth_test.go           # Unit tests for authentication
│   ├── file_test.go           # Unit tests for file handling
│   └── utils_test.go          # Unit tests for utilities
├── docker-compose.yml         # Docker Compose setup
├── Dockerfile                 # Dockerfile for the Go app
├── Makefile                   # Makefile for build and automation
├── go.mod                     # Go module file
└── LICENSE                    # License file
```

## Setup and Installation

### Prerequisites

- Go (version 1.20 or higher)
- Docker and Docker Compose
- PostgreSQL

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/Annany2002/backend-in-go.git
   cd file-storage-app
   ```

2. Configure the application:

   - Update `config/app.yaml` with your environment settings.

3. Build and run the application:

   ```bash
   make build
   make run
   ```

4. API is available at `http://localhost:8080`.

### Running with Docker

1. Build and start the containers:

   ```bash
   docker-compose up --build
   ```

2. API is available at `http://localhost:8080`.

## Testing

Run unit tests using:

```bash
make test
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contact

For any questions or feedback, please contact:

- **Name**: Annany Vsishwakarma
- **Email**: shazam6061@gmail.com
- **GitHub**: [Annany2002](https://github.com/Annany2002)
