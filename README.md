# URL Shortener Service

A highly scalable and decoupled URL Shortener service built with **Golang**. This project provides an API to shorten long URLs and handle redirection from short URLs back to the original URLs. The system is designed with scalability and decoupling in mind, using components like **API Gateway**, **Redis** for caching, and **MongoDB** for persistent storage.

## Features

- Shorten long URLs via an API endpoint
- Retrieve and redirect to original URLs from short links
- Cache frequently accessed URLs using Redis
- Fully decoupled architecture for scalability
- Ready for containerization and horizontal scaling

## Architecture

The system is designed with a fully decoupled architecture:

- **API Gateway**: Handles incoming client requests and routes them to the appropriate services.
- **URL Shortener Service**: Responsible for generating short URLs, storing the mappings, and handling redirection logic.
- **Cache Layer (Redis)**: Caches frequently accessed URL mappings to reduce database load.
- **Database Layer (MongoDB)**: Stores the mappings of original and shortened URLs for persistence.
- **Monitoring & Logging**: Tracks metrics and logs events like URL shortening and redirection.

### Diagram
To be add in future

## Project Structure

```bash
.
├── cmd                 # Main application entry points
├── pkg                 # Core packages (shortening logic, database, cache)
├── internal            # Internal packages (URL validation, business logic)
├── api                 # API layer with REST endpoints
├── config              # Configuration files for database, cache, etc.
├── docs                # Documentation files (PlantUML diagrams, API docs)
└── README.md           # This file
