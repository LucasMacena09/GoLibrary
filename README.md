# GoLibrary API

This is my first API built with **Go** using the **Gin** framework. The project is a simple **Bookstore API** that allows users to:
- Retrieve all books
- Retrieve a book by its ID
- Add new books
- Checkout (borrow) books
- Return books

## Why a Slice Instead of a Database?
Since this is my first API, I chose to use a **slice** (`[]book`) instead of a database for simplicity. This approach allows me to focus on learning API structuring, routing, and request handling without dealing with database setup and migrations.

In a production environment, this would be replaced by a proper database such as **PostgreSQL, MySQL, or MongoDB**.

## Project Structure
```
GoLibrary/
├── main.go
├── models/
│   ├── book.go
├── services/
│   ├── book_service.go
├── handlers/
│   ├── book_handler.go
├── routes/
│   ├── routes.go
```

## Running the API
### Prerequisites:
- Go installed (version 1.23.4 or later)
- Gin framework

### Steps:
```sh
git clone https://github.com/LucasMacena09/GoLibrary.git
cd GoLibrary
go run main.go
```

## Endpoints
| Method | Endpoint         | Description            |
|--------|-----------------|------------------------|
| GET    | `/books`        | Get all books         |
| GET    | `/books/:id`    | Get a book by ID      |
| POST   | `/books`        | Add a new book        |
| PATCH  | `/checkout?id=` | Checkout a book      |
| PATCH  | `/return?id=`   | Return a book        |

## Future Improvements
- Implement a **database** to persist book data.
- Add **user authentication**.
- Implement **unit tests** for better reliability.

This project is just the beginning of my journey into Go backend development! 🚀

