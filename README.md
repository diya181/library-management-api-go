Library Management API
Project Description:-
This project is a backend Library Management System developed using Go (Golang) and PostgreSQL. It provides REST APIs to manage books, perform checkout and return operations, calculate fines for late returns, and handle reservations when books are unavailable.
The system demonstrates database operations, transaction handling, and structured backend development using Gin and GORM.

*Features:-
The application supports the following features:
Add new books to the library
View all available books
Checkout books
Return books
Automatically calculate fine for late returns (₹10 per day)
Reserve books when all copies are checked out
Fine Calculation Logic
If a book is returned after the due date, a fine is calculated based on the number of late days.

Fine = 10 × number of late days

The due date is automatically set to 7 days from the checkout date.

*Technologies Used:-
Go (Golang)
Gin Web Framework
GORM (Object Relational Mapping)
PostgreSQL

*Project Structure:-
The project is organized into separate folders:
database – contains database connection logic
models – contains data models such as Book, User, Checkout, and Reservation
handlers – contains API handler functions
main.go – application entry point

*How to Run:-
Install Go
Install PostgreSQL
Create a database named library
Update the database credentials in database/db.go
Run the application using:

go run main.go
The server will start on http://localhost:8080

Notes:-
This project is designed for academic purposes and does not include authentication. It focuses on backend logic, database interaction, and API development.
