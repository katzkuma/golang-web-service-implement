package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 宣告全域變數 db
var db *sql.DB

// 宣告 Response 格式
type Resp struct {
	code    int    `json:'Code'`
	message string `json:'Message'`
	result  IsOK   `json:'Result'`
}

type IsOK struct {
	IsOK bool `json:"IsOK"`
}

//
func main() {
	var err error
	// 連接 golang-web 資料庫
	db, err = sql.Open("mysql", "root:passpass@tcp(127.0.0.1:3306)/golang-web")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// 設置連接池數量
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	// 實體化 Gin 框架
	apiServer := gin.Default()

	// 設定路由路徑
	apiServer.POST("v1/user/create", create)
	apiServer.POST("v1/user/delete", delete)
	apiServer.POST("v1/user/pwd/change", pwdChange)
	apiServer.GET("v1/user/login", login)

	// 執行伺服器，監聽 8080 port
	apiServer.Run(":8080")
}

// 1. 新增會員
func create(context *gin.Context) {
	// 宣告 Response 變數
	var httpCode int
	var code int
	var message string
	var isOK bool

	// 取得 API 參數
	Account := context.PostForm("Account")
	Password := context.PostForm("Password")

	// 插入資料，若帳號已存在則插入失敗
	rs, err := db.Exec(`
		INSERT INTO Account(Account,Password) 
		SELECT ?, ? FROM DUAL 
		WHERE NOT EXISTS
		(SELECT Account FROM Account WHERE Account = ?)
	`, Account, Password, Account)

	if err != nil {
		log.Panicln("user insert error", err.Error())
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}

	// 設定 Response 參數
	if id > 0 {
		httpCode = http.StatusOK
		code = 0
		message = ""
		isOK = true
	} else {
		httpCode = http.StatusOK
		code = 2
		message = ""
		isOK = false
	}

	// 呼叫 Response 函式
	contextResp(context, httpCode, code, message, isOK)
}

// 2. 刪除會員
func delete(context *gin.Context) {
	// 宣告 Response 變數
	var httpCode int
	var code int
	var message string
	var isOK bool

	// 取得 API 參數
	Account := context.PostForm("Account")

	// 刪除資料，回傳所刪除資料行數
	rs, err := db.Exec(`
		DELETE FROM Account WHERE Account = ?
	`, Account)

	if err != nil {
		log.Panicln("user insert error", err.Error())
	}

	rows, err := rs.RowsAffected()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}

	// 設定 Response 參數
	if rows > 0 {
		httpCode = http.StatusOK
		code = 0
		message = ""
		isOK = true
	} else {
		httpCode = http.StatusOK
		code = 2
		message = ""
		isOK = false
	}

	// 呼叫 Response 函式
	contextResp(context, httpCode, code, message, isOK)
}

// 3. 修改會員密碼
func pwdChange(context *gin.Context) {
	// 宣告 Response 變數
	var httpCode int
	var code int
	var message string
	var isOK bool

	// 取得 API 參數
	Account := context.PostForm("Account")
	Password := context.PostForm("Password")

	// 更新資料，回傳所更新資料行數
	rs, err := db.Exec(`
		UPDATE Account SET Password = ? WHERE Account = ?
	`, Password, Account)

	if err != nil {
		log.Panicln("user insert error", err.Error())
	}

	rows, err := rs.RowsAffected()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}

	// 設定 Response 參數
	if rows > 0 {
		httpCode = http.StatusOK
		code = 0
		message = ""
		isOK = true
	} else {
		httpCode = http.StatusOK
		code = 2
		message = ""
		isOK = false
	}

	// 呼叫 Response 函式
	contextResp(context, httpCode, code, message, isOK)
}

// 4. 驗證帳號密碼
func login(context *gin.Context) {
	// 宣告 Response 變數
	var httpCode int
	var code int
	var message string
	var isOK bool

	// 宣告資料庫回傳的變數
	var respAccount string
	var respPassword string

	// 取得 API 參數
	Account := context.Query("Account")
	Password := context.Query("Password")

	// 驗證帳號密碼是否存在
	rows := db.QueryRow(`
		SELECT Account, Password FROM Account WHERE Account = ? AND Password = ?
	`, Account, Password)

	err := rows.Scan(&respAccount, &respPassword)

	// 設定 Response 參數
	if err != nil {
		httpCode = http.StatusBadRequest
		code = 2
		message = "Login Failed"
		isOK = false
	} else {
		httpCode = http.StatusOK
		code = 0
		message = ""
		isOK = true
	}

	// 呼叫 Response 函式
	contextResp(context, httpCode, code, message, isOK)
}

func contextResp(context *gin.Context, httpCode int, code int, message string, isOK bool) {
	// Http Response 函式
	context.JSON(httpCode, gin.H{
		"Code":    code,
		"Message": message,
		"Result": IsOK{
			IsOK: isOK,
		},
	})
}
