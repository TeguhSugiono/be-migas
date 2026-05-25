package controllers

import (
	"BackendEsp32/models"
	"BackendEsp32/utils"
	"database/sql"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func CreateDevice(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	userID := c.GetInt("user_id")

	var input models.CreateDeviceInput

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	deviceID := "GAS-" + utils.RandomString(6)

	deviceToken := utils.RandomString(20)
	pairCode := utils.RandomString(6)

	query := `
	INSERT INTO devices (
		device_id,
		device_token,
		device_name,
		pair_code,
		seller_id
	)
	VALUES (?, ?, ?, ?,?)
	`

	_, err := db.Exec(
		query,
		deviceID,
		deviceToken,
		input.DeviceName,
		pairCode,
		userID,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Device berhasil dibuat",
		"data": gin.H{
			"device_id":    deviceID,
			"device_token": deviceToken,
			"pair_code":    pairCode,
		},
	})
}

func PairDevice(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	customerID := c.GetInt("user_id")

	var input models.PairDeviceInput

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	var total int

	queryCheck := `
	SELECT COUNT(*)
	FROM devices
	WHERE device_id = ?
	AND pair_code = ?
	AND customer_id IS NULL
	`

	err := db.QueryRow(
		queryCheck,
		input.DeviceID,
		input.PairCode,
	).Scan(&total)

	if err != nil || total == 0 {

		c.JSON(400, gin.H{
			"success": false,
			"message": "Device tidak valid",
		})

		return
	}

	queryUpdate := `
	UPDATE devices
	SET customer_id = ?,
	    paired_at = NOW()
	WHERE device_id = ?
	`

	_, err = db.Exec(
		queryUpdate,
		customerID,
		input.DeviceID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Device berhasil dipairing",
	})
}
func MyDevices(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	customerID := c.GetInt("user_id")

	query := `
		SELECT
			id,
			device_id,
			device_name,
			is_active,
			last_seen,
			last_seen_unix,
			created_at
		FROM devices
		WHERE customer_id = ?
		ORDER BY id DESC
		`

	rows, err := db.Query(
		query,
		customerID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	defer rows.Close()

	var devices []gin.H

	for rows.Next() {

		var id int
		var deviceID string
		var deviceName string
		var isActive bool
		var createdAt string
		var lastSeen sql.NullTime
		var lastSeenUnix sql.NullInt64

		rows.Scan(
			&id,
			&deviceID,
			&deviceName,
			&isActive,
			&lastSeen,
			&lastSeenUnix,
			&createdAt,
		)

		onlineStatus := false

		if lastSeenUnix.Valid {

			nowUnix :=
				time.Now().Unix()

			diff :=
				nowUnix - lastSeenUnix.Int64

			if diff <= 15 {

				onlineStatus = true
			}
		}

		devices = append(devices, gin.H{
			"id":          id,
			"device_id":   deviceID,
			"device_name": deviceName,
			"is_active":   isActive,
			"created_at":  createdAt,
			"online":      onlineStatus,
			"last_seen":   lastSeen.Time,
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    devices,
	})
}

func LatestDeviceStatus(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	deviceID := c.Param("device_id")

	query := `
	SELECT
		device_id,
		status_gas,
		alarm_status,
		wifi_status,
		created_at
	FROM gas_log
	WHERE device_id = ?
	ORDER BY id DESC
	LIMIT 1
	`

	var result gin.H

	var statusGas bool
	var alarmStatus string
	var wifiStatus bool
	var createdAt string
	var device string

	err := db.QueryRow(
		query,
		deviceID,
	).Scan(
		&device,
		&statusGas,
		&alarmStatus,
		&wifiStatus,
		&createdAt,
	)

	if err != nil {

		c.JSON(404, gin.H{
			"success": false,
			"message": "Data tidak ditemukan",
		})

		return
	}

	result = gin.H{
		"device_id":    device,
		"status_gas":   statusGas,
		"alarm_status": alarmStatus,
		"wifi_status":  wifiStatus,
		"created_at":   createdAt,
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    result,
	})
}

func DeviceHistory(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	deviceID := c.Param("device_id")

	query := `
	SELECT
		status_gas,
		alarm_status,
		wifi_status,
		created_at
	FROM gas_log
	WHERE device_id = ?
	ORDER BY id DESC
	LIMIT 50
	`

	rows, err := db.Query(
		query,
		deviceID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})

		return
	}

	defer rows.Close()

	var histories []gin.H

	for rows.Next() {

		var statusGas bool
		var alarmStatus string
		var wifiStatus bool
		var createdAt string

		rows.Scan(
			&statusGas,
			&alarmStatus,
			&wifiStatus,
			&createdAt,
		)

		histories = append(histories, gin.H{
			"status_gas":   statusGas,
			"alarm_status": alarmStatus,
			"wifi_status":  wifiStatus,
			"created_at":   createdAt,
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    histories,
	})
}
