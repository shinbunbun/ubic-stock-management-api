package controllers

import (
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"
)

func (c *Controller) FindUserByID(id string) (int, string, error) {
	user, err := c.Interactor.FindUserByID(id)
	if err != nil {
		return internalErrorMessage("Failed to found user")
	}
	return jsonDump(user)
}

func (c *Controller) FindUserByEmail(id string) (int, string, error) {
	user, err := c.Interactor.FindUserByEmail(id)
	if err != nil {
		return internalErrorMessage("Failed to find user")
	}
	return jsonDump(user)
}

func (c *Controller) DeleteUserByID(id string) (int, string, error) {
	err := c.Interactor.DeleteUserByID(id)
	if err != nil {
		return internalErrorMessage("Failed to delete user")
	}
	return message("Successful delete user!")
}

func (c *Controller) CreateUser(email string, name string, password string) (int, string, error) {
	id, err := c.Interactor.CreateUser(email, name, password)
	if err != nil {
		return internalErrorMessage("Failed to create user")
	}
	return jsonDump(&domain.User{
		ID:    id,
		Name:  name,
		Email: email,
	})
}
