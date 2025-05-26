
### 📁 `internal/middleware/README.md`

````md
# Middleware Package

This package defines middleware functions for securing routes and managing HTTP headers. Middleware functions are built for use with the Gin framework.

---

## 📄 Files

- `auth.go`: Handles authentication and authorization via static and dynamic tokens.
- `header.go`: Sets security-related and CORS HTTP headers for incoming requests.

---

## 🔐 Authentication Middleware (`auth.go`)

### ✅ StaticProtectedRoute

```go
func StaticProtectedRoute() gin.HandlerFunc
````

* Verifies token, role, and admin ID from the request.
* Validates token from DB using the static token strategy.
* Uses `security.VerifyStaticToken()` for verification.
* Adds `adminID` and `role` to the request context.

> Route Example: `/check-static`

### 🔄 DynamicProtectedRoute

```go
func DynamicProtectedRoute() gin.HandlerFunc
```

* Same validation flow as the static variant.
* Uses the dynamic token stored in DB and `security.VerifyDynamicToken()`.

> Route Example: `/check-dynamic`

---

### 🧩 Internal Logic

#### `getRequestInfo(*gin.Context) (*authMeta, error)`

* Parses and validates:

  * Authorization token (from headers)
  * Role (from URL param)
  * Admin ID (from URL param)
* Validates the admin from the DB and ensures role consistency.

#### `authMeta.checkStaticWithDB() error`

* Fetches the static token for the admin from DB.
* Compares it to the provided token.

#### `authMeta.checkDynamicWithDB() error`

* Same as above but for dynamic tokens.

> Uses services from `admin` and `credentials` core packages.

---

## 🛡 Header Middleware (`header.go`)

### `SetHeaders`

```go
func SetHeaders() gin.HandlerFunc
```

Sets important HTTP headers for security and CORS support:

| Header                         | Description                                 |
| ------------------------------ | ------------------------------------------- |
| `Access-Control-Allow-Origin`  | Allows requests from any origin (`*`)       |
| `Content-Security-Policy`      | Disables all sources (`default-src 'none'`) |
| `X-Content-Type-Options`       | Prevents MIME sniffing                      |
| `Access-Control-Allow-Methods` | Specifies allowed HTTP methods              |
| `Access-Control-Allow-Headers` | Lists permitted headers                     |

If the request method is `OPTIONS`, the request is terminated with `204 No Content`.

---

## 🧠 Best Practices

* Group routes by protection level (`static`, `dynamic`) using corresponding middleware.
* Apply `SetHeaders` globally for all routes to enforce security.

---

## 🗂 Related Components

* `pkg/security`: Token generation and validation.
* `internal/core/admin`: Admin entity and DB interaction.
* `internal/core/credentials`: Token retrieval logic.
* `internal/response`: Standardized API responses.


