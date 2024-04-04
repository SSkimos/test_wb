package employees

import (
	"github.com/curtrika/reader/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	var employee models.Employee

	if result := h.DB.First(&employee, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &employee)
}
