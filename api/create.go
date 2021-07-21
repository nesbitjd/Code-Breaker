package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create creates a database entry
func Create(c *gin.Context) {

	db := database.Setup(c)

	// Migrate the schema
	db.AutoMigrate(&types.Product{})

	product := &types.Product{}
	err := c.Bind(product)
	if err != nil {
		retErr := fmt.Errorf("Unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	fmt.Printf("create: %+v\n", db.Create(product))

	resp := fmt.Sprintf("created entry %+v", product.Code)
	c.JSON(http.StatusCreated, resp)
}
