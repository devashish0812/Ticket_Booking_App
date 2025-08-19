# Authentication & Authorization Flow

## 1. User Signup
- Frontend calls **`POST /signup`** → User Service  
- User Service:
  - Hashes the password
  - Stores user in **MongoDB**
- Responds with success (no token yet)

---

## 2. User Login
- Frontend calls **`POST /login`** with email + password → User Service  
- User Service:
  - Validates credentials from DB
  - Generates **Access Token (JWT)** + **Refresh Token**
  - Sends:
    - **Access Token** → in response body (or header)
    - **Refresh Token** → as **HTTP-only Cookie**
- Frontend stores Access Token in memory (not in localStorage for security).

---

## 3. Accessing Other Services
- Frontend includes **Access Token** in `Authorization: Bearer <token>` header.  
- Target Service (e.g., Ticket Service):
  - Verifies JWT using shared secret/public key
  - If valid → proceed
  - If expired → return `401 Unauthorized`

---

## 4. Refreshing Token
- If Access Token expired:
  - Frontend calls **`POST /refresh`** → User Service
  - Refresh Token (from HTTP-only cookie) is sent automatically with request
  - User Service validates refresh token
  - If valid:
    - Issues new Access Token
    - (Optionally rotates Refresh Token)
  - Sends new Access Token back to frontend

---

## 5. Logout
- Frontend calls **`POST /logout`** → User Service
- User Service:
  - Clears refresh token cookie
  - (Optionally invalidates token in DB/Redis)
