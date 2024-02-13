package flagadmin

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func GetFlagsByService(c *gin.Context, conn db.Connection) {
	serviceName := c.Request.FormValue("service_name")
	isStaging := c.Request.FormValue("is_staging")

	rows, err := conn.Query("SELECT name, value, staging_value, type FROM flags WHERE service_name = ?", serviceName)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	flags := gin.H{}
	valueKey := "value"
	if isStaging == "true" {
		valueKey = "staging_value"
	}

	// Iterate through the results and parse values
	for _, v := range rows {
		fmt.Println(v)
		stringValue := v[valueKey].(string)
		stringName := v["name"].(string)
		fmt.Println(v)
		switch v["type"].(int64) {
		case 0:
			flags[stringName] = stringValue
		case 1:
			if strings.ToLower(stringValue) == "true" {
				flags[stringName] = true
				continue
			}
			if strings.ToLower(stringValue) == "false" {
				flags[stringName] = false
				continue
			}
			c.JSON(500, gin.H{"Error": fmt.Sprintf("Invalid boolean flag for flag name: %s, value: %s", stringName, stringValue)})
			return
		case 2:
			flags[stringName], err = strconv.Atoi(stringValue)
			if err != nil {
				c.JSON(500, gin.H{"Error": fmt.Sprintf("Invalid integer flag for flag name: %s, value: %s", stringName, stringValue)})
				return
			}
		default:
			c.JSON(400, gin.H{"Error": "Bad request!"})
		}

	}

	c.JSON(200, flags)
}
