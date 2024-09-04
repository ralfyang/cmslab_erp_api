
# ERP Feature Enhancement and Expansion API System

## Project Overview
This project is an API system designed to improve and expand the functionality of ERP systems. The API was developed to integrate and extend various ERP features.

## Key Features
- **ERP Integration**: Enhances and expands functions through integration with existing ERP systems.
- **Developed with Golang**: Implemented using Go to maintain lightweight and high performance.
- **MS SQL Query Support**: Provides queries optimized for MS SQL databases.
- **Scalable Interface**: Designed with a scalable interface to allow easy integration with various ERP and CMS systems.
- **CRUD API Support**: Provides a RESTful API for Creating, Reading, Updating, and Deleting employee information.

## Development Environment
- **Language**: Golang
- **Database**: MS SQL
- **Libraries**:
  - Gorilla Mux (HTTP router for routing)
  - Go MSSQL driver (for MS SQL database connection)
  - Godotenv (for loading environment variables)

## API Usage

### 1. Get All Employees (GET)
- **URL**: `/api/employees`
- **Method**: `GET`
- **Response**: Returns a JSON array of all employees.
  
  ```bash
  curl http://localhost:8090/api/employees
  ```

  ```json
  [
      {
          "id": 1,
          "name": "John Doe",
          "job_position": "Developer",
          "salary": 50000.00
      }
  ]
  ```

### 2. Get Employee by ID (GET)
- **URL**: `/api/employee/{id}`
- **Method**: `GET`
- **Path Parameters**: `id` - The ID of the employee to retrieve
- **Response**: Returns employee details in JSON format.
  
  ```bash
  curl http://localhost:8090/api/employee/1
  ```

  ```json
  {
      "id": 1,
      "name": "John Doe",
      "job_position": "Developer",
      "salary": 50000.00
  }
  ```

### 3. Create a New Employee (POST)
- **URL**: `/api/employee`
- **Method**: `POST`
- **Request Body**: 
  - `name`: Employee's name
  - `job_position`: Employee's position
  - `salary`: Employee's salary
  
  ```bash
  curl -X POST http://localhost:8090/api/employee   -H "Content-Type: application/json"   -d '{"name": "Jane Doe", "job_position": "Manager", "salary": 60000.00}'
  ```

  **Response**:
  ```json
  {"message": "Employee created"}
  ```

### 4. Update an Employee (PUT)
- **URL**: `/api/employee/{id}`
- **Method**: `PUT`
- **Path Parameters**: `id` - The ID of the employee to update
- **Request Body**: 
  - `name`: Employee's name
  - `job_position`: Employee's position
  - `salary`: Employee's salary

  ```bash
  curl -X PUT http://localhost:8090/api/employee/1   -H "Content-Type: application/json"   -d '{"name": "Jane Doe", "job_position": "Director", "salary": 70000.00}'
  ```

  **Response**:
  ```json
  {"message": "Employee updated"}
  ```

### 5. Delete an Employee (DELETE)
- **URL**: `/api/employee/{id}`
- **Method**: `DELETE`
- **Path Parameters**: `id` - The ID of the employee to delete
  
  ```bash
  curl -X DELETE http://localhost:8090/api/employee/1
  ```

  **Response**:
  ```json
  {"message": "Employee deleted"}
  ```

## Installation and Running

1. Clone the API system to your local environment:
   ```bash
   git clone https://github.com/your-repo/erp-api-system.git
   ```

2. Prepare the configuration file. Copy the `.env_TEMP` file to `.env` and set the environment variables as follows:

   ```bash
   cp .env_TEMP .env
   ```

   Set the following variables in the `.env` file:

   ```env
   DB_HOST=       # Database host address
   DB_PORT=       # Database port
   DB_USER=       # Database username
   DB_PASSWORD=   # Database password
   DB_NAME=       # Database name

   SERVER_PORT=8090  # The port on which the API server will run
   ```

3. Install the necessary Go packages and run the API server:
   ```bash
   go mod tidy   # Install the required dependencies
   go run main.go
   ```

4. Once the server starts successfully, the API will be accessible at `http://localhost:8090`.

## Database Initialization

To create the necessary tables in the `ralftest` database, run the following SQL query:

```sql
USE ralftest;

CREATE TABLE employees (
    id INT PRIMARY KEY IDENTITY(1,1),
    name NVARCHAR(100),
    job_position NVARCHAR(100),
    salary DECIMAL(10, 2)
);
```

---

# ERP 기능 개선 및 확장을 위한 연동형 API 시스템

## 프로젝트 개요
이 프로젝트는 ERP 시스템의 기능을 개선하고 확장하기 위해 설계된 API 시스템입니다. 이 API는 다양한 ERP 기능과 통합 및 연동될 수 있도록 개발되었습니다.

## 주요 기능
- **ERP 기능 연동**: 기존 ERP 시스템과의 연동을 통해 기능을 개선 및 확장.
- **Golang 기반 개발**: Go 언어로 API 시스템을 구현하여 경량성과 고성능을 유지.
- **MS SQL 쿼리 지원**: MS SQL 데이터베이스에 적합한 쿼리 제공.
- **확장성 있는 인터페이스**: 다양한 ERP 및 CMS 시스템과의 확장성 높은 통합 인터페이스 제공.
- **CRUD API 제공**: 직원 정보의 생성(Create), 조회(Read), 업데이트(Update), 삭제(Delete) 작업을 수행할 수 있는 RESTful API 제공.

## 개발 환경
- **언어**: Golang
- **데이터베이스**: MS SQL
- **라이브러리**: 
  - Gorilla Mux (라우팅을 위한 HTTP 라우터)
  - Go MSSQL 드라이버 (MS SQL 데이터베이스 연결)
  - Godotenv (환경 변수 로드를 위한 패키지)

## API 사용 방법

### 1. 전체 직원 목록 조회 (GET)
- **URL**: `/api/employees`
- **Method**: `GET`
- **Response**: 모든 직원 정보를 JSON 배열로 반환합니다.
  
  ```bash
  curl http://localhost:8090/api/employees
  ```

### 2. 특정 직원 조회 (GET)
- **URL**: `/api/employee/{id}`
- **Method**: `GET`
- **Path Parameters**: `id` - 조회할 직원의 ID

  ```bash
  curl http://localhost:8090/api/employee/1
  ```

### 3. 새로운 직원 생성 (POST)
- **URL**: `/api/employee`
- **Method**: `POST`

  ```bash
  curl -X POST http://localhost:8090/api/employee   -H "Content-Type: application/json"   -d '{"name": "Jane Doe", "job_position": "Manager", "salary": 60000.00}'
  ```

### 4. 직원 정보 업데이트 (PUT)
- **URL**: `/api/employee/{id}`
- **Method**: `PUT`

  ```bash
  curl -X PUT http://localhost:8090/api/employee/1   -H "Content-Type: application/json"   -d '{"name": "Jane Doe", "job_position": "Director", "salary": 70000.00}'
  ```

### 5. 직원 정보 삭제 (DELETE)
- **URL**: `/api/employee/{id}`
- **Method**: `DELETE`

  ```bash
  curl -X DELETE http://localhost:8090/api/employee/1
  ```

## 설치 및 실행

1. API 시스템을 클론하여 로컬 환경에 설치합니다:
   ```bash
   git clone https://github.com/your-repo/erp-api-system.git
   ```

2. 설정 파일을 준비합니다:
   ```bash
   cp .env_TEMP .env
   ```

3. Go 패키지 설치 후 서버 실행:
   ```bash
   go mod tidy
   go run main.go
   ```

## 데이터베이스 초기화
`ralftest` 데이터베이스에 필요한 테이블을 생성하려면 아래 쿼리를 실행하세요:

```sql
USE ralftest;

CREATE TABLE employees (
    id INT PRIMARY KEY IDENTITY(1,1),
    name NVARCHAR(100),
    job_position NVARCHAR(100),
    salary DECIMAL(10, 2)
);
```

---

Developed by **CMS Lab**

