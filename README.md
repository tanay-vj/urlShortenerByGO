# 🔗 urlShortenerByGO

A fast, simple, and reliable URL shortening service built with **Go** using clean **MVC architecture**.

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)
![Architecture](https://img.shields.io/badge/Architecture-MVC-green?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)

## 🚀 Features

- ✅ **Lightning Fast** - MD5-based 8-character hash generation
- ✅ **Clean Architecture** - Professional MVC pattern implementation
- ✅ **RESTful API** - JSON request/response format
- ✅ **In-Memory Storage** - Ultra-fast data access
- ✅ **Comprehensive Documentation** - Built-in API guide
- ✅ **Error Handling** - Proper HTTP status codes

## 🏗️ Architecture Overview

This project follows clean MVC architecture with the following structure:

- **📁 models/** → Data structures & domain objects
- **🎮 controllers/** → HTTP handlers & request/response logic  
- **⚙️ services/** → Business logic & validation
- **💾 repositories/** → Data access layer & storage
- **🔧 utils/** → Helper functions & utilities
- **🚀 main.go** → Entry point & dependency injection
- **📦 go.mod** → Go module definition

### Data Flow:

## 📋 API Endpoints

| Method | Endpoint | Description | Request | Response |
|--------|----------|-------------|---------|----------|
| `GET` | `/` | API documentation homepage | - | HTML page |
| `POST` | `/shorten` | Create shortened URL | `{"url": "..."}` | `{"short_url": "abc123"}` |
| `GET` | `/redirect/{hash}` | Redirect to original URL | - | HTTP 302 redirect |


## 🔧 Installation & Usage

### Prerequisites
```bash
# Ensure Go is installed (1.19+ recommended)
go version

# 1. Clone the repository
git clone https://github.com/tanay-vj/urlShortenerByGO.git
cd urlShortenerByGO

# 2. Run the application
go run main.go

# 3. Access the service
# Homepage: http://localhost:8080/ (View API documentation)
# API: Use Postman or cURL for /shorten endpoint
# Redirect: Browser can directly use /redirect/{hash}
```

## 💡 Complete Example Workflow

### Step 1: View API Documentation
Visit the homepage in your browser: `http://localhost:8080/`

### Step 2: Choose a URL to Shorten
Example: `https://www.linkedin.com/in/tanayvijay/`

### Step 3: Create Short URL (Postman)
- **Method:** POST
- **URL:** `http://localhost:8080/shorten`
- **Headers:** `Content-Type: application/json`
- **Body:**
```json
{
  "url": "https://www.linkedin.com/in/tanayvijay/"
}
```
- **Expected Response:**
 ```json
{
  "short_url": "f4a5b6c7"
}
```
### Step 4: Use the short URL
- **Visit in browser :** `http://localhost:8080/redirect/f4a5b6c7`
- **✅ Result :** Automatically redirects to LinkedIn profile!
