Go REST API – Now with Gorilla Mux & Validation
This is an upgraded CRUD (Create, Read, Update, Delete) API built with Go.
Originally written using only the net/http package, it now includes:

🦍 Gorilla Mux for cleaner routing and subrouters

✅ Request validation for better input handling

📌 Features
✅ Built with Go’s net/http + github.com/gorilla/mux

✅ RESTful API for basic product/task management

✅ In-memory data storage (slice or map)

✅ Clean, modular route handling

✅ Validation for request bodies (e.g. empty fields, data types)

✅ Ready for middleware (auth, logging, etc.)

🚀 Getting Started
bash
Copy
Edit
go run main.go
Make sure Go is installed (version 1.20+ recommended).

🔄 What's New in This Version?
Refactored to use Gorilla Mux instead of plain net/http

Added subrouters for method-specific grouping

Added validation logic to improve data integrity

Codebase structured for easier scalability

🔗 GitHub
github.com/abrshDev/go-microservice
