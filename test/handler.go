package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/module"
)

// Login test
type Login struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

var service Service

// Handler test
func Handler(c *gin.Context) {
	module.SetContext(c)

	id := c.Param("id")

	switch c.Request.Method {
	case http.MethodGet:
		// detail
		if id != "" {
			d, err := service.Get(id)
			if err != nil {
				module.Error("GET_RESOURCE_TEST", err.Error())

				return
			}

			module.Output(d)

			return
		}

		// list
		d, err := service.Find()
		if err != nil {
			module.Error("GET_RESOURCE_TEST", err.Error())
		}

		module.Output(d)

		return
	case http.MethodPost:
		var payload Object
		err := c.BindJSON(&payload)
		if err != nil {
			module.Panic("REQUIRED_FIELDS", "field is required")

			return
		}

		d, err := service.Create(payload)
		if err != nil {
			module.Error("POST_RESOURCE_TEST", err.Error())
		}

		module.Output(d)

		return
	case http.MethodPut:
		if id == "" {
			module.Error("RESOURCE_ID_REQUIRED", "resource id is missing")

			return
		}

		var payload Object
		err := c.BindJSON(&payload)
		if err != nil {
			module.Panic("REQUIRED_FIELDS", "field is required")

			return
		}

		d, err := service.Update(payload, id)
		if err != nil {
			module.Error("PUT_RESOURCE_TEST", err.Error())

			return
		}

		module.Output(d)

		return
	case http.MethodDelete:
		if id == "" {
			module.Error("RESOURCE_ID_REQUIRED", "resource id is missing")

			return
		}

		d, err := service.Remove(id)
		if err != nil {
			module.Error("DELETE_RESOURCE_TEST", err.Error())

			return
		}

		module.Output(d)

		return
	}

	module.Error("METHOD_NOT_ALLOWED",
		c.Request.Method+" method not allowed in this endpoint")
}