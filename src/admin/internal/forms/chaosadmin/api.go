package chaosadmin

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/gin-gonic/gin"
	"io"
)

func SetChaosStatus(c *gin.Context, conn db.Connection) {
	chaosID := c.Request.FormValue("id")
	row, err := conn.Query("SELECT status FROM chaos WHERE id = ?", chaosID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	chaosStatus := row[0]["status"].(string)
	switch chaosStatus {
	case "Unknown":
		chaosStatus = "Getting Requests"
	case "Getting Requests":
		chaosStatus = "Invoking Requests"
	case "Invoking Requests":
		chaosStatus = "Done"
		chaosResult, _ := getResult(c)
		fmt.Println(chaosResult)
		_, err = conn.Query("UPDATE chaos SET status=?, result=? WHERE id=?", chaosStatus, chaosResult, chaosID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{})
	default:
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = conn.Query("UPDATE chaos SET status=? WHERE id=?", chaosStatus, chaosID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}

func getResult(c *gin.Context) (string, error) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
