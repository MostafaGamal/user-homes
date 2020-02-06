package controllers

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"net/http"
	"olarm/app"
	"olarm/app/dtos"
	"olarm/app/infrastructure"
	"olarm/app/validations"
)

type Users struct {
	*revel.Controller
	db *gorm.DB
}

func (c *Users) Inject(con *app.Context) {
	c.db = con.DB
	return
}

func (c Users) Add() revel.Result {
	response := make(map[string]interface{})
	userData := &dtos.UserDTO{}

	if err := c.Params.BindJSON(&userData); err != nil {
		response["response"] = errors.New("invalid user request body").Error()
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(response)
	}

	userData.Validate(c.Validation)
	if c.Validation.HasErrors() {
		response["response"] = c.Validation.ErrorMap()
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(response)
	}

	newUser := &infrastructure.User{
		Name:     userData.Name,
		Password: userData.Password,
		Email:    userData.Email,
		Address:  userData.Address,
		Age:      userData.Age,
	}

	if err := c.db.Create(&newUser).Error; err != nil {
		response["response"] = errors.New("error adding new user").Error()
		c.Response.Status = http.StatusInternalServerError
		return c.RenderJSON(response)
	}

	return c.RenderText("User Added Successfully")
}

func (c Users) AddHome() revel.Result {
	response := make(map[string]interface{})
	homeData := &dtos.HomeDTO{}

	if err := c.Params.BindJSON(&homeData); err != nil {
		response["response"] = errors.New("invalid home request body").Error()
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(response)
	}

	homeData.Validate(c.Validation)
	if c.Validation.HasErrors() {
		response["response"] = c.Validation.ErrorMap()
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(response)
	}

	sensorsData, err := json.Marshal(homeData.Sensors)
	if err != nil {
		response["response"] = errors.New("error marshal home sensors data").Error()
		c.Response.Status = http.StatusInternalServerError
		return c.RenderJSON(response)
	}

	newHome := &infrastructure.Home{
		ID:      homeData.ID,
		Serial:  homeData.Serial,
		Sensors: string(sensorsData),
		UserID:  homeData.UserID,
	}

	if err := c.db.Create(&newHome).Error; err != nil {
		response["response"] = errors.New("error adding new home").Error()
		c.Response.Status = http.StatusInternalServerError
		return c.RenderJSON(response)
	}

	return c.RenderText("Home Added Successfully")
}

func (c Users) View(userID string) revel.Result {
	response := make(map[string]interface{})

	validations.ValidateID(c.Validation, userID)
	if c.Validation.HasErrors() {
		response["response"] = c.Validation.ErrorMap()
		c.Response.Status = http.StatusBadRequest
		return c.RenderJSON(response)
	}

	userData := &infrastructure.User{}
	if err := c.db.Where("id = ?", userID).First(&userData).Error; err != nil {
		response["response"] = errors.New("error retrieve user data").Error()
		c.Response.Status = http.StatusInternalServerError
		return c.RenderJSON(response)
	}

	home := &infrastructure.Home{}
	var homes []dtos.HomeDTO

	rows, err := c.db.Model(&infrastructure.Home{}).Where("user_id = ?", userID).Rows()
	if err != nil {
		response["response"] = errors.New("error retrieve user's home data").Error()
		c.Response.Status = http.StatusInternalServerError
		return c.RenderJSON(response)
	}

	for rows.Next() {
		if err = c.db.ScanRows(rows, &home); err != nil {
			response["response"] = errors.New("error in retrieving home data").Error()
			c.Response.Status = http.StatusInternalServerError
			return c.RenderJSON(response)
		}
		homes = append(homes, *dtos.ToHomeDTO(home))
	}

	userFullData := &dtos.UserDataDTO{
		User:  dtos.UserDTO{
			Name:     userData.Name,
			Password: userData.Password,
			Email:    userData.Email,
			Address:  userData.Address,
			Age:      userData.Age,
		},
		Homes: homes,
	}

	return c.RenderJSON(userFullData)
}
