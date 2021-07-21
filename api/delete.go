package api

import (
	"Projects/code_breaker/database"
	"Projects/code_breaker/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete deletes database entry
func Delete(c *gin.Context) {

	db := database.Setup(c)

	// Migrate the schema
	db.AutoMigrate(&types.Product{})

	code := c.Param("code")

	fmt.Printf("delete: %+v\n", db.Where("Code = ?", code).Delete(&types.Product{}))

	resp := fmt.Sprintf("deleted entry %+v\n", code)
	c.JSON(http.StatusOK, resp)
}
