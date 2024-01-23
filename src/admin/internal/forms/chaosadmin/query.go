package chaosadmin

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/gin-gonic/gin"
)

func SetChaosStatus(c *gin.Context, conn db.Connection) {
	chaosID := c.Request.FormValue("chaos_id")
	fmt.Print(chaosID)
	row, err := conn.Query("SELECT status FROM chaoses WHERE id = ?", chaosID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	chaosStatus := row[0]["status"].(string)
	switch chaosStatus {
	case "starting":
		chaosStatus = "getting requests"
	case "getting requests":
		chaosStatus = "invoking requests"
	case "invoking requests":
		chaosStatus = "gathering results"
	case "gathering results":
		chaosStatus = "finished"
	default:
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	_, err = conn.Query("UPDATE chaoses SET status=? WHERE id=?", chaosStatus, chaosID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{})
}
