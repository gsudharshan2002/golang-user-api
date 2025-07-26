# User API - Go + MongoDB

A simple REST API for user management built with **Go (Golang)** and **MongoDB Atlas** for interview .

---

## Name 

**Sudharshan**


---

##  Tech Stack

* **Backend**: Go (net/http)
* **Database**: MongoDB Atlas (Cloud MongoDB)
* **Testing**: Built-in `testing` package + MongoDB `mtest`

---

##  Project Setup

### 1. Install Go (1.24+)




###  2. Install MongoDB Atlas

* Go to [https://www.mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)
* Create a free cluster
* Create a database & collection (e.g., `users`)
* Whitelist your IP and get your connection string (starts with `mongodb+srv://`)

> Example:

```bash
mongodb+srv://<username>:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority
```

### ✅ 3. Clone the project

```bash
git clone https://github.com/yourusername/user-api.git
cd user-api
```

### 4. Set environment variables

Update your MongoDB connection string in `config/db.go`:

```go
clientOptions := options.Client().ApplyURI("your-mongodb-connection-string")


replace ur string 
```

---

##  Running Unit Tests

```bash
go test ./controller -v
```



##  Running the Server

```bash
go run main.go
```

Visit: [http://localhost:8080]

we can also change port

You can test your API using **Postman** 

---

## Sample API Call (POST User)

```



Content-Type: application/json

{
  "name": "Sudharshan",
  "dob": "2002-01-01",
  "address": "Chennai",
  "description": "Engineer"
}
```
url for all HTTP request

POST ---->/users
GET ----> /users
put ----> /users/id(dynamic)
DELETE ----> /users/id(dynamic)


---





---



