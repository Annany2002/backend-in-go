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
│   └── main.go            # Application entry point
├── config/
│   ├── config.go          # Configuration loader
│   └── app.yaml           # Application configuration
├── internal/
│   ├── auth/              # Authentication logic
│   ├── file/              # File handling logic
│   └── utils/             # Utility functions
├── pkg/
│   ├── database/          # Database connection
│   └── storage/           # Storage backend abstraction
├── test/                  # Unit tests
├── docker-compose.yml     # Docker setup for development
├── Dockerfile             # Dockerfile for the application
├── Makefile               # Build and run automation
├── go.mod                 # Go module file
└── README.md              # Project documentation
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
