# TODOLIST 
Simple Rest API for todolist

## A Little bit about us
### 🛠️ Built with :
- [MongoDB Atlas](https://www.mongodb.com)
- [Gorilla Mux](https://github.com/gorilla/mux)
<br />

## 🏁 Getting started
1. Clone the repo.
2. Move to ``todolist`` directory
```bash
cd todolist/
```
3. Run go
```bash
go run main.go
```
4. API Running in localhost:8000
```
⚫ get all data
  GET localhost:8000/todolist

⚫ get data by id
  GET localhost:8000/todolist/{id}

⚫ create list
  POST localhost:8000/todolist
  Content-Type: applicatiopn/json
  {
    "name"    : "My List"
    "status"  : "Completed"
  }

⚫ update list
  PUT localhost:8000/todolist/{id}
  Content-Type: applicatiopn/json
  {
    "name"    : "My List"
    "status"  : "Completed"
  }

⚫ delete list
  DELETE localhost:8000/todolist/{id}
```
