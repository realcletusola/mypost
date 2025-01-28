# Go Post Sample Application

This is a **Post Sample Application** built with Go, designed to demonstrate a standard and structured project architecture, PostgreSQL integration, database migrations, and concurrent CRUD functionality using **goroutines** and **mutexes** to ensure thread-safe operations.

---

## Features

- **PostgreSQL Integration**: A robust database connection to store and manage posts and categories.
- **Database Migrations**: Easy-to-use migration scripts to set up and update the database schema.
- **Goroutines & Mutex**: Efficient concurrent data handling, with mutexes to prevent race conditions in a multi-threaded environment.
- **Standard Go Project Architecture**: A modular and maintainable structure for scaling and production readiness.
- **RESTful API with Chi Router**: Lightweight and fast routing for handling HTTP requests.

---


---

## Endpoints

1. **Create Post**  
   - **Method**: `POST`  
   - **Endpoint**: `/posts`  
   - **Payload**:  
     ```json
     {
       "title": "Post Title",
       "content": "Post Content",
       "category": {
         "name": "Category Name"
       }
     }
     ```
   - **Response**: `201 Created`

2. **Get All Posts**  
   - **Method**: `GET`  
   - **Endpoint**: `/posts`  
   - **Response**: List of all posts.

3. **Get Post by ID**  
   - **Method**: `GET`  
   - **Endpoint**: `/posts/{id}`  
   - **Response**: Post details.

---

## Getting Started

### Prerequisites

1. **Go**: Install [Go](https://golang.org/doc/install).  
2. **PostgreSQL**: Install PostgreSQL and set up a database for the application.  
3. **Migrate Database**: Run the database migrations to set up the schema.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/realcletusola/mypost.git
   cd mypost
   ```

2. Install dependencies:
	```bash
	go mod tidy 
	```

3. Run Migrations 
	```bash
	go run internal/db/database.go
	```

4. Start server 
	```bash
	go run cmd/mypost/main.go
	```
