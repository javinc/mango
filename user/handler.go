package user

import (
	"net/http"
	"strings"

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
	id := c.Param("id")
	module.SetContext(c)

	switch c.Request.Method {
	case http.MethodGet:
		// list
		if id == "" {
			filter := Object{
				Name:  c.Query("filter.name"),
				Email: c.Query("filter.email"),
			}

			option := Option{
				Slice:  c.Query("slice"),
				Order:  c.Query("order"),
				Filter: filter,
			}

			d, err := service.Find(option)
			if err != nil {
				module.Error("GET_RESOURCE_"+strings.ToUpper(resourceName), err.Error())
			}

			module.Output(d)

			return
		}

		// detail
		d, err := service.Get(id)
		if err != nil {
			module.Error("GET_RESOURCE_"+strings.ToUpper(resourceName), err.Error())

			return
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
			module.Error("POST_RESOURCE_"+strings.ToUpper(resourceName), err.Error())
		}

		module.Output(d)

		return
	case http.MethodPatch:
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
			module.Error("PUT_RESOURCE_"+strings.ToUpper(resourceName), err.Error())

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
			module.Error("DELETE_RESOURCE_"+strings.ToUpper(resourceName), err.Error())

			return
		}

		module.Output(d)

		return
	}

	module.Error("METHOD_NOT_ALLOWED",
		c.Request.Method+" method not allowed in this endpoint")
}