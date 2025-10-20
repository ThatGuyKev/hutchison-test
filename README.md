# Hutchison Code Test - Dogs API

This is a full-stack web application 
- Frontend: React
- Backend: Go(Golang)
- Database: SQLite

## Prerequisites
Make sure you have the following installed
- [Go](https://go.dev/dl/)
- [Node.js](https://nodejs.org/)
- [yarn](https://yarnpkg.com/)



## Backend Setup (Go)
```bash
cd backend
go mod tidy
go run ./cmd
```

The server will run 
```bash
Server is listening on port 5050
```
API ENDPOINTS
| Method | Endpoint       | Description         |
| ------ | -------------- | ------------------- |
| POST   | `/api/dogs`    | create a new dog    |
| GET    | `/api/dogs`    | list all dogs       |
| GET    | `/api/dogs/id` | get a dog by id     |
| PUT    | `/api/dogs/id` | update existing dog |
| DELET  | `/api/dogs`    | delte existing dog  |



## Frontend Setup (React)

```bash 
cd frontend
yarn
```


Create .env file with BACKEND_BASE_URL

```bash
echo "BACKEND_BASE_URL=localhost:5050/api" >> .env
```

```bash
yarn dev
```