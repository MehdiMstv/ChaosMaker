package chaosadmin

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/gin-gonic/gin"
)

func SetChaosStatus(c *gin.Context, conn db.Connection) {
	chaosID := c.Request.FormValue("chaos_id")
	fmt.Print(chaosID)
	row, err := conn.Query("SELECT status FROM chaos WHERE id = ?", chaosID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	chaosStatus := row[0]["status"].(string)
	switch chaosStatus {
	case "Unknown":
		chaosStatus = "Starting"
	case "Starting":
		chaosStatus = "Getting Requests"
	case "Getting Requests":
		chaosStatus = "Invoking Requests"
	case "Invoking Requests":
		chaosStatus = "Done"
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
