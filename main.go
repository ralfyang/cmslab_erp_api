package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "unicode/utf8"

    _ "github.com/denisenkom/go-mssqldb"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "golang.org/x/text/encoding/korean"
    "golang.org/x/text/transform"
    "io/ioutil"
)

var db *sql.DB

// ConvertEUC_KRtoUTF8: EUC-KR에서 UTF-8로 변환
func ConvertEUC_KRtoUTF8(input string) (string, error) {
    reader := transform.NewReader(strings.NewReader(input), korean.EUCKR.NewDecoder())
    utf8Text, err := ioutil.ReadAll(reader)
    if err != nil {
        return "", err
    }
    return string(utf8Text), nil
}

// UTF-8 인코딩을 확인하고 필요한 경우만 변환하는 함수
func ConvertEUC_KRtoUTF8IfNecessary(input string) (string, error) {
    // 입력이 이미 UTF-8 형식인 경우 변환하지 않음
    if utf8.ValidString(input) {
        return input, nil
    }
    // 변환이 필요한 경우 EUC-KR에서 UTF-8로 변환
    return ConvertEUC_KRtoUTF8(input)
}

// GetEmployees: 모든 직원 목록 가져오기
func GetEmployees(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    rows, err := db.Query("SELECT id, name, job_position, salary FROM employees")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    employees := []map[string]interface{}{}
    for rows.Next() {
        var id int
        var name, jobPosition string
        var salary float64
        err = rows.Scan(&id, &name, &jobPosition, &salary)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        employee := map[string]interface{}{
            "id":           id,
            "name":         name,
            "job_position": jobPosition,
            "salary":       salary,
        }
        employees = append(employees, employee)
    }

    json.NewEncoder(w).Encode(employees)
}

// 법인카드 결제이력 조회 API
func GetCards(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    // 쿼리 파라미터 가져오기
    cCode := r.URL.Query().Get("c_code")
    if cCode == "" {
        cCode = "7000" // 기본 값
    }

    rows, err := db.Query(`
        SELECT
            C_CODE,       -- 코드 (소유자 관련 정보)
            CLIENT_NOTE,  -- 적요
            TRADE_PLACE,  -- 가맹점
            DOCU_STAT,    -- 전표처리
            MCC_CODE_NAME -- 코스트센터
        FROM NEOE.CARD_TEMP
        WHERE C_CODE = @p1
    `, cCode)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    cards := []map[string]interface{}{}
    for rows.Next() {
        var cCode, clientNote, tradePlace, docuStat, mccCodeName sql.NullString
        err = rows.Scan(&cCode, &clientNote, &tradePlace, &docuStat, &mccCodeName)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // 인코딩 문제 해결: 각 항목을 UTF-8로 변환
        tradePlaceUtf8, err := ConvertEUC_KRtoUTF8IfNecessary(tradePlace.String)
        if err != nil {
            tradePlaceUtf8 = tradePlace.String // 변환 실패 시 원래 문자열 사용
        }

        clientNoteUtf8, err := ConvertEUC_KRtoUTF8IfNecessary(clientNote.String)
        if err != nil {
            clientNoteUtf8 = clientNote.String
        }

        mccCodeNameUtf8, err := ConvertEUC_KRtoUTF8IfNecessary(mccCodeName.String)
        if err != nil {
            mccCodeNameUtf8 = mccCodeName.String
        }

        card := map[string]interface{}{
            "c_code":       cCode.String,
            "client_note":  clientNoteUtf8,
            "trade_place":  tradePlaceUtf8,
            "docu_stat":    docuStat.String,
            "mcc_code_name": mccCodeNameUtf8,
        }
        cards = append(cards, card)
    }

    json.NewEncoder(w).Encode(cards)
}

// main 함수
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

    // API 엔드포인트 설정
    r.HandleFunc("/api/employees", GetEmployees).Methods("GET")
    r.HandleFunc("/api/cards", GetCards).Methods("GET")

    // 서버 시작
    log.Printf("Server running on port %s", serverPort)
    log.Fatal(http.ListenAndServe(":"+serverPort, r))
}

