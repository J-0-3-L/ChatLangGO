# 🚀 ChatLanGO API – Go + Gin + Gorm

API básica estilo red social (similar a Twitter), desarrollada en **Go**, usando:

- Gin Web Framework  
- Gorm ORM  
- SQLite  
- JWT para autenticación  
- bcrypt para protección de contraseñas  
- Swagger

Incluye autenticación, usuarios y CRUD completo de posts.

## 📁 Estructura del Proyecto

```plaintext
cmd/
  └── app/
      └── main.go
internal/
  └── auth/
  └── config/
  └── docs/
  └── models/
  └── posts/
  └── tools/
  └── scan_route.go
.gitignore
Chatdb.db
README.md
go.mod
go.sum
```

---

## 🛠 Tecnologías

- **Go (Golang)**
- **Gin**
- **Gorm**
- **SQLite**
- **JWT**
- **bcrypt**
- **Swagger**

## ⚙️ Instalación

Clona el proyecto:

```bash
git clone https://github.com/J-0-3-L/ChatLangGO
cd tu-proyecto
```

```bash  
go mod tidy
```

## ▶️ Ejecutar el Servidor
```bash
go run cmd/app/main.go
```
```bash
http://localhost:8080
```
