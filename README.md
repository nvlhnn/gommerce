# Gommerce

Gommerce is a backend RESTful API application that serves as the backbone for an e-commerce platform. It provides various endpoints to manage products, shopping carts, user cutomer, and order processing. Built using Go programming language with the Gin framework, Gommerce utilizes PostgreSQL for database storage, Redis for caching, and includes rate limiting functionality to manage API requests effectively.

## Features
- Gommerce provides essential e-commerce functionalities including product search, user authentication, shopping cart management, and order processing.
- Utilize caching to enhance browsing speed and performance
- Implement rate limiting to ensure fair usage and prevent abuse 

## Instalation

To install and run the Gommerce application locally, follow these steps:

1.  Clone the repository
 ```bash
git clone https://github.com/nvlhnn/gommerce.git
```

2. Navigate to the project directory
 ```bash
cd gommerce
```

3. Build the Docker image
 ```bash
docker-compose build
```

4. Start the application
 ```bash
docker-compose up
```

## Configuration
Gommerce application can be configured using environment variables. You can customize database settings, memory store configurations, and more by modifying the .env file in the project root directory.

```bash
DB_PORT=5432
DB_USER=nvlhnn
DB_PASSWORD=here
DB_NAME=gommerce

JWT_SECRET=jwtsecret
GO111MODULE=on
IS_SEEDING=true

REDIS_ADDR=host.docker.internal
REDIS_PASS=
REDIS_DB=0
REDIS_PORT=6379
```

## Postman Documentation
Explore the API endpoints and interact with the Gommerce application using the [Postman Documentation](https://elements.getpostman.com/redirect?entityId=30413366-2613f7c1-bec3-4623-b685-be4d29fd2995&entityType=collection) select **gommerce-local** as the enviroment variable.

