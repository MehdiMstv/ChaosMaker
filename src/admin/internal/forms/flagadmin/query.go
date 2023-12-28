package flagadmin

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/gin-gonic/gin"
)

func GetFlagsByService(c *gin.Context, conn db.Connection) {
	serviceName := c.Request.FormValue("service_name")

	rows, err := conn.Query("SELECT name, value FROM flags WHERE service_name = ?", serviceName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	flags := gin.H{}

	// Iterate through the results and parse values
	for _, v := range rows {
		flags[v["name"].(string)] = map[string]string{"value": v["value"].(string)}
	}

	c.JSON(200, flags)
}
