package serviceadmin

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/gin-gonic/gin"
)

func GetChaosStagingAddress(c *gin.Context, conn db.Connection) {
	serviceName := c.Request.FormValue("name")
	row, err := conn.Query("SELECT staging_address FROM services WHERE name = ?", serviceName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(row) != 1 {
		c.JSON(400, gin.H{"error": "No Staging Address Found or Multiple Found"})
		return
	}

	c.JSON(200, gin.H{"staging_address": row[0]["staging_address"]})
}
