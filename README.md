# my-js-server/my-js-server/README.md

# My JS Server with SQLite Database

This project is a simple JavaScript server built with Node.js and Express, utilizing an SQLite database. It includes an HTML form for user input and demonstrates basic CRUD operations.

## Project Structure

```
my-js-server
├── src
│   ├── index.js          # Entry point of the application
│   ├── database.js       # Database connection and operations
│   └── public
│       ├── index.html    # HTML form for user input
│       └── styles.css     # Styles for the HTML form
├── Dockerfile             # Instructions to build the Docker image
├── docker-compose.yml     # Defines services for the application
├── package.json           # npm configuration file
└── README.md              # Project documentation
```

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd my-js-server
   ```

2. Build and run the application using Docker Compose:

   ```
   docker-compose up --build
   ```

### Usage

- Access the application at `http://localhost:3000`.
- Fill out the form and submit to interact with the SQLite database.

### Technologies Used

- Node.js
- Express
- SQLite
- Docker

### License

This project is licensed under the MIT License.