package controllers

import (
	"bytes"
	"encoding/json"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/domain"
)

func (c *Controller) FindUserByID(id string) (int, string, error) {
	user, err := c.Interactor.FindUserByID(id)
	if err != nil {
		return internalErrorMessage("Failed to found user")
	}
	return jsonDump(user)
}

func (c *Controller) FindUserByEmail(id string) (string, error) {
	user, err := c.Interactor.FindUserByEmail(id)
	if err != nil {
		return "", err
	}

	byte, err := json.Marshal(user)
	var buf bytes.Buffer

	json.HTMLEscape(&buf, byte)

	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *Controller) DeleteUserByID(id string) (string, error) {
	err := c.Interactor.DeleteUserByID(id)
	if err != nil {
		return "", err
	}

	byte, err := json.Marshal(map[string]interface{}{
		"message": "Successful Delete!",
	})
	var buf bytes.Buffer

	json.HTMLEscape(&buf, byte)

	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *Controller) CreateUser(email string, name string, password string) (string, error) {
	id, err := c.Interactor.CreateUser(email, name, password)
	if err != nil {
		return "", err
	}

	byte, err := json.Marshal(&domain.User{
		ID:    id,
		Name:  name,
		Email: email,
	})

	var buf bytes.Buffer

	json.HTMLEscape(&buf, byte)

	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
