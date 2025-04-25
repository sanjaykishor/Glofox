# Glofox

Glofox is a SaaS platform for boutique fitness studios, gyms, and wellness centers that allows owners to manage their classes and bookings efficiently.

## Features

### Classes Management
- Create and manage fitness classes
- Set class capacity, start date, and end date
- Retrieve class details and listings

### Bookings Management
- Create bookings for members
- Book specific classes or general appointments
- View bookings by date or ID

## Getting Started

### Prerequisites
- Go 1.23 or higher
- Git

### Installation

1. Clone the repository
```bash
git clone https://github.com/sanjaykishor/Glofox.git
cd Glofox
```

2. Install dependencies
```bash
go mod download
```

### Running the Application

#### Using Go
```bash
# Run the application
go run ./cmd/glofox/main.go
```

#### Using Make
```bash
# Build the application
make build

# Run the application
make run
```

#### Using Docker
```bash
# Build the Docker image
docker build -t glofox .

# Run the container
docker run -p 8080:8080 glofox
```

## API Documentation

The API server runs on port 8080 by default.

### Base URL
```
http://localhost:8080/api/v1
```

### Classes API

#### Create a Class
- **URL**: `/classes`
- **Method**: `POST`
- **Request Body**:
```json
{
    "name": "Yoga Class",
    "start_date": "2025-04-25",
    "end_date": "2025-04-26",
    "capacity": 15
}
```
- **Success Response** (201 Created):
```json
{
    "success": true,
    "message": "Class created successfully",
    "data": {
        "id": "c0e3bcde-1d22-4c7b-a788-15c8f815b35d",
        "name": "Yoga Class",
        "start_date": "2025-04-25T00:00:00Z",
        "end_date": "2025-04-26T00:00:00Z",
        "capacity": 15
    }
}
```
- **Error Response** (400 Bad Request):
```json
{
    "success": false,
    "error": "invalid start date format, use YYYY-MM-DD"
}
```

#### Get All Classes
- **URL**: `/classes`
- **Method**: `GET`
- **Success Response** (200 OK):
```json
{
    "success": true,
    "data": [
        {
            "id": "c0e3bcde-1d22-4c7b-a788-15c8f815b35d",
            "name": "Yoga Class",
            "start_date": "2025-04-25T00:00:00Z",
            "end_date": "2025-04-26T00:00:00Z",
            "capacity": 15
        },
        {
            "id": "f8a9d724-5c31-4b9d-8c0e-7d45d0aab123",
            "name": "HIIT Training",
            "start_date": "2025-04-26T00:00:00Z",
            "end_date": "2025-04-27T00:00:00Z", 
            "capacity": 10
        }
    ]
}
```

#### Get Class by ID
- **URL**: `/classes/:id`
- **Method**: `GET`
- **Success Response** (200 OK):
```json
{
    "success": true,
    "data": {
        "id": "c0e3bcde-1d22-4c7b-a788-15c8f815b35d",
        "name": "Yoga Class",
        "start_date": "2025-04-25T00:00:00Z",
        "end_date": "2025-04-26T00:00:00Z",
        "capacity": 15
    }
}
```
- **Error Response** (404 Not Found):
```json
{
    "success": false,
    "error": "class not found"
}
```

### Bookings API

#### Create a Booking
- **URL**: `/bookings`
- **Method**: `POST`
- **Request Body**:
```json
{
    "name": "USER A",
    "date": "2025-04-25",
    "class_id": "c0e3bcde-1d22-4c7b-a788-15c8f815b35d"
}
```
- **Success Response** (201 Created):
```json
{
    "success": true,
    "message": "Booking created successfully",
    "data": {
        "id": "d7a8e931-2b41-5f6c-9d0e-8f12a3b45c67",
        "member_name": "USER A",
        "class_id": "c0e3bcde-1d22-4c7b-a788-15c8f815b35d",
        "date": "2025-04-25T00:00:00Z",
        "created_at": "2025-04-24T14:30:45Z"
    }
}
```
- **Error Response** (400 Bad Request):
```json
{
    "success": false,
    "error": "invalid date format, use YYYY-MM-DD"
}
```

#### Get All Bookings
- **URL**: `/bookings`
- **Method**: `GET`
- **Success Response** (200 OK):
```json
{
    "success": true,
    "data": [
        {
            "id": "d7a8e931-2b41-5f6c-9d0e-8f12a3b45c67",
            "member_name": "USER A",
            "class_id": "c0e3bcde-1d22-4c7b-a788-15c8f815b35d",
            "date": "2025-04-25T00:00:00Z",
            "created_at": "2025-04-24T14:30:45Z"
        },
        {
            "id": "e8b9f042-3c52-6d7e-0e1f-9g23b4c56d78",
            "member_name": "Jane Smith",
            "class_id": "f8a9d724-5c31-4b9d-8c0e-7d45d0aab123",
            "date": "2025-04-25T00:00:00Z",
            "created_at": "2025-04-24T15:45:12Z"
        }
    ]
}
```

#### Get Booking by ID
- **URL**: `/bookings/:id`
- **Method**: `GET`
- **Success Response** (200 OK):
```json
{
    "success": true,
    "data": {
        "id": "d7a8e931-2b41-5f6c-9d0e-8f12a3b45c67",
        "member_name": "USER A",
        "class_id": "c0e3bcde-1d22-4c7b-a788-15c8f815b35d",
        "date": "2025-04-25T00:00:00Z",
        "created_at": "2025-04-24T14:30:45Z"
    }
}
```
- **Error Response** (404 Not Found):
```json
{
    "success": false,
    "error": "booking not found"
}
```

#### Get Bookings by Date
- **URL**: `/bookings/date/:date`
- **Method**: `GET`
- **URL Example**: `/bookings/date/2025-04-15`
- **Success Response** (200 OK):
```json
{
    "success": true,
    "data": [
        {
            "id": "d7a8e931-2b41-5f6c-9d0e-8f12a3b45c67",
            "member_name": "USER A",
            "class_id": "c0e3bcde-1d22-4c7b-a788-15c8f815b35d",
            "date": "2025-04-25T00:00:00Z",
            "created_at": "2025-04-24T14:30:45Z"
        },
        {
            "id": "f9c0e143-4d65-7g86-1h2i-3j45k6l78m90",
            "member_name": "USER B",
            "class_id": "a1b2c3d4-5e6f-7g8h-9i0j-1k2l3m4n5o6p",
            "date": "2025-04-25T00:00:00Z",
            "created_at": "2025-04-24T09:15:30Z"
        }
    ]
}
```
- **Error Response** (400 Bad Request):
```json
{
    "success": false,
    "error": "invalid date format, use YYYY-MM-DD"
}
```

## Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...
```

## Project Structure

```
Glofox/
├── cmd/
│   └── glofox/           # Application entry point
├── internal/
│   ├── handler/          # HTTP handlers
│   ├── middleware/       # HTTP middleware
│   ├── repository/       # Data access layer
│   ├── router/           # HTTP router setup
│   ├── validation/       # Validation logic
│   └── service/          # Business logic
├── bin/                  # Compiled binaries
├── Dockerfile            # Docker build configuration
├── Makefile              # Build automation
├── go.mod                # Go module definition
└── README.md             # Documentation
```

## License

This project is licensed under the [MIT License](LICENSE).
