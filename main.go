package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/denisenkom/go-mssqldb"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

var db *sql.DB

func main() {
    // .env 파일 읽기
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    serverPort := os.Getenv("SERVER_PORT")

    // 데이터베이스 연결 문자열 구성 (보안을 위해 출력하지 않음)
    connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    // 데이터베이스 연결 테스트
    db, err = sql.Open("sqlserver", connString)
    if err != nil {
        log.Fatalf("Error creating connection pool: %s", err.Error())
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %s", err.Error())
    } else {
        log.Println("Successfully connected to database!")
    }

    // API 라우팅 설정
    r := mux.NewRouter()

    // 기본 API 엔드포인트 설정 (예시)
    r.HandleFunc("/api/employees", GetEmployees).Methods("GET")
    r.HandleFunc("/api/employee/{id}", GetEmployeeByID).Methods("GET")
    r.HandleFunc("/api/employee", CreateEmployee).Methods("POST")
    r.HandleFunc("/api/employee/{id}", UpdateEmployee).Methods("PUT")
    r.HandleFunc("/api/employee/{id}", DeleteEmployee).Methods("DELETE")

    // 서버 시작
    log.Printf("Server running on port %s", serverPort)
    log.Fatal(http.ListenAndServe(":"+serverPort, r))
}

// API 핸들러 함수들

// GetEmployees: 모든 직원 목록 가져오기
func GetEmployees(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    rows, err := db.Query("SELECT id, name, position FROM employees")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    employees := []map[string]interface{}{}
    for rows.Next() {
        var id int
        var name, position string
        err = rows.Scan(&id, &name, &position)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        employee := map[string]interface{}{
            "id":       id,
            "name":     name,
            "position": position,
        }
        employees = append(employees, employee)
    }

    json.NewEncoder(w).Encode(employees)
}

// GetEmployeeByID: ID로 직원 정보 가져오기
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    id := vars["id"]

    row := db.QueryRow("SELECT id, name, position FROM employees WHERE id = @p1", id)
    var employee struct {
        ID       int    `json:"id"`
        Name     string `json:"name"`
        Position string `json:"position"`
    }

    err := row.Scan(&employee.ID, &employee.Name, &employee.Position)
    if err == sql.ErrNoRows {
        http.Error(w, "Employee not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(employee)
}

// CreateEmployee: 직원 정보 생성
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var employee struct {
        Name     string `json:"name"`
        Position string `json:"position"`
    }

    err := json.NewDecoder(r.Body).Decode(&employee)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


    _, err = db.Exec("INSERT INTO employees (name, position) VALUES (@p1, @p2)", employee.Name, employee.Position)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Employee created"})
}

// UpdateEmployee: 직원 정보 업데이트
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    id := vars["id"]

    var employee struct {
        Name     string `json:"name"`
        Position string `json:"position"`
    }

    err := json.NewDecoder(r.Body).Decode(&employee)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = db.Exec("UPDATE employees SET name = @p1, position = @p2 WHERE id = @p3", employee.Name, employee.Position, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Employee updated"})
}

// DeleteEmployee: 직원 정보 삭제
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := db.Exec("DELETE FROM employees WHERE id = @p1", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Employee deleted"})
}



