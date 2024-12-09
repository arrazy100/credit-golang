## How to Run

    docker-compose -f docker-compose.dev.yaml up --build

Above command will run these:
- Database (Postgres)
- Golang Service
- Apply Migrations
- Run Seeds

Swagger http://localhost:8080/swagger/index.html

## Implemented OWASP

1. A01:2021-Broken Access Control<br/>
Users can only access resources they are authorized to.

2. A07:2021-Identification and Authentication Failures<br/>
Authentication helps prevents unauthorized access.

3. A03:2021-Injection<br/>
<t/>- SQL Injection is prevented by using GORM for database queries.<br/>
<t/>- Sanitizing user input before executing commands by using bluemonday.

4. A04:2021-Insecure Design<br/>
    The service is designed to prevent common security vulnerabilities such as SQL Injection and Cross-Site Scripting (XSS).

5. A05:2021-Security Misconfiguration<br/>
	The service properly configured CORS to ensure that only authorized domains can make requests to the server.

## Database Diagram

https://dbdiagram.io/d/Credit-Golang-6752c295e9daa85acae0ebbc
