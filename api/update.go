package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Update updates a database entry
func Update(c *gin.Context) {

	db := database.Setup(c)

	// Migrate the schema
	db.AutoMigrate(&types.Product{})

	code := c.Param("code")
	product := &types.Product{}
	err := c.Bind(product)
	if err != nil {
		retErr := fmt.Errorf("Unable to parse json body: %w", err)
		c.Error(retErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, retErr.Error())
		return
	}

	fmt.Printf("update: %+v\n", db.Model(&product).Where("Code = ?", code).Updates(types.Product{Code: product.Code, Price: product.Price}))

	resp := fmt.Sprintf("updated entry %+v", product.Code)
	c.JSON(http.StatusCreated, resp)
}
