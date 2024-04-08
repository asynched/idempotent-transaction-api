package controllers

import (
	"errors"

	"github.com/asynched/idempotent-transaction-api/domain/entities"
	"github.com/asynched/idempotent-transaction-api/domain/repositories"
	"github.com/labstack/echo/v4"
)

type AccountController struct {
	accountRepository *repositories.AccountRepository
}

func NewAccountController(accountRepository *repositories.AccountRepository) *AccountController {
	return &AccountController{accountRepository}
}

type CreateAccountDto struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func (d *CreateAccountDto) Validate() error {
	if d.Name == "" {
		return errors.New("name is required")
	}

	if d.Cpf == "" {
		return errors.New("cpf is required")
	}

	if len(d.Cpf) != 11 {
		return errors.New("cpf must have 11 characters")
	}

	for _, char := range d.Cpf {
		if char < '0' || char > '9' {
			return errors.New("cpf must have only numbers")
		}
	}

	return nil
}

func (ctrl *AccountController) Create(c echo.Context) error {
	data := CreateAccountDto{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(400, echo.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := data.Validate(); err != nil {
		return c.JSON(400, echo.Map{
			"message": "Provided data is invalid",
			"error":   err.Error(),
		})
	}

	account, err := ctrl.accountRepository.Save(entities.Account{
		Name: data.Name,
		Cpf:  data.Cpf,
	})

	if err != nil {
		return c.JSON(500, echo.Map{
			"message": "Failed to create account",
			"error":   err.Error(),
		})
	}

	return c.JSON(201, account)
}
