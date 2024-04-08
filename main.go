package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asynched/idempotent-transaction-api/controllers"
	"github.com/asynched/idempotent-transaction-api/database"
	"github.com/asynched/idempotent-transaction-api/domain/repositories"
	"github.com/labstack/echo/v4"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)
	log.SetPrefix(fmt.Sprintf("[%d] [main] ", os.Getpid()))
}

func main() {
	app := echo.New()
	app.HideBanner = true
	app.HidePort = true

	sqlite := database.GetSQLite()
	defer sqlite.Close()

	redis := database.GetRedis()
	defer redis.Close()

	accountRepository := repositories.NewAccountRepository(sqlite)
	accountController := controllers.NewAccountController(accountRepository)

	transactionRepository := repositories.NewTransactionRepository(sqlite)
	transactionController := controllers.NewTransactionController(accountRepository, transactionRepository, redis)

	app.POST("/v1/accounts", accountController.Create)
	app.GET("/v1/accounts/:id/transactions", transactionController.ListAll)

	app.POST("/v1/transactions", transactionController.Create)

	log.Println("Server started on: http://localhost:8080")
	if err := app.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
