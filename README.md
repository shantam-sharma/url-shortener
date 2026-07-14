# 🚀 Go URL Shortener

Go • React • PostgreSQL • Docker • TypeScript • Render • Vercel

A production-ready URL shortening service built with **Go**, **PostgreSQL**, **React**, and **Docker**. Users can generate unique short links, create custom aliases, and instantly redirect to the original URL through a clean REST API.

**🔗 Live Demo:** https://url-shortner-pink-phi.vercel.app

**⚙️ Backend API:** https://url-shortener-api-udkx.onrender.com

The application follows a layered architecture (Handler → Service → Repository) with environment-based configuration, PostgreSQL persistence, and Dockerized local development.

---

## 📸 Preview

### Home

![Homepage](assets/homepage.png)

### Shortened URL

![Result](assets/result.png)

---

## ✨ Features

- 🔗 Shorten long URLs instantly
- ✏️ Create custom aliases
- 🎲 Automatic unique short code generation
- 🚀 Fast HTTP redirects
- 📊 Track click count
- 🗄️ PostgreSQL database
- 🐳 Docker support for local development
- 🌐 Production deployment with Render, Vercel & Neon
- ⚡ RESTful API
- 🎨 Responsive React frontend built with Tailwind CSS

---

# 🛠 Tech Stack

## Backend

- Go
- pgx
- Standard net/http package
- Repository Pattern
- Service Layer Architecture

## Frontend

- React
- TypeScript
- Vite
- Tailwind CSS

## Infrastructure

- Docker
- Docker Compose
- PostgreSQL

## Deployment

- Render (Backend)
- Vercel (Frontend)
- Neon (PostgreSQL)

---

# 🏗️ Project Architecture

```
                   React Frontend
                          │
                          │ HTTP
                          ▼
                 Go REST API (net/http)
                          │
                 Service Layer
                          │
                Repository Layer
                          │
                    PostgreSQL
```

The backend follows a layered architecture to keep business logic separate from HTTP handlers and database operations.

```
Handler
   │
Service
   │
Repository
   │
Database
```

---

# 📂 Project Structure

```
url-shortener
│
├── cmd/
│   └── server/
│
├── internal/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── repository/
│   └── service/
│
├── migrations/
│
├── frontend/
│   ├── src/
│   ├── public/
│   └── package.json
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

# 🚀 Running Locally

## Clone the repository

```bash
git clone https://github.com/shantam-sharma/url-shortener.git

cd url-shortener
```

---

## Backend

Copy the example environment file and update values if needed.

```bash
cp .env.example .env
```

Start the backend and PostgreSQL.

```bash
docker compose up --build
```

The API will be available at:

```
http://localhost:8080
```

Note: Commit example env files (`.env.example`, `.env.docker.example`) — do not commit real `.env` files.

---

## Frontend

Move into the frontend directory and install dependencies.

```bash
cd frontend
npm install
```

Create the frontend env file from the example if needed.

```bash
cp .env.example .env
```

Run the development server.

```bash
npm run dev
```

Frontend will run on

```
http://localhost:5173
```

---

# 🌐 Production Deployment

| Service | Platform |
|----------|----------|
| Frontend | Vercel |
| Backend | Render |
| Database | Neon PostgreSQL |

---

# 📡 API

Base URL (Production): https://url-shortener-api-udkx.onrender.com

## Create Short URL

**POST**

```
/api/v1/urls
```

### Request

```json
{
    "url": "https://google.com",
    "alias": "google"
}
```

### Successful Response

```json
{
    "short_url": "https://url-shortener-api-udkx.onrender.com/google",
    "original_url": "https://google.com",
    "short_code": "google"
}
```

---

## Redirect

```
GET /{shortCode}
```

Example

```
GET /google
```

Automatically redirects to

```
https://google.com
```

---

# ⚙️ Design Decisions

- Layered architecture (Handler → Service → Repository)
- PostgreSQL for persistent storage
- Environment-based configuration for local and production deployments
- CORS middleware for secure frontend communication
- Docker support for reproducible development environments
- Clean separation between frontend and backend services

---

# 🔮 Future Improvements

- JWT Authentication
- URL Expiration
- Rate Limiting
- Redis Caching
- Analytics Dashboard
- QR Code Generation
- Custom Domains
- Admin Dashboard
- Unit & Integration Tests

---

# 👨‍💻 Author

**Shantam Sharma**

GitHub: https://github.com/shantam-sharma

LinkedIn: https://www.linkedin.com/in/shantam-sharma (update with your LinkedIn URL)

---

# 📄 License

This project is licensed under the MIT License.
