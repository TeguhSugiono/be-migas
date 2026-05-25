package controllers

import (
	"BackendEsp32/helpers"
	"BackendEsp32/models"
	"BackendEsp32/services"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func getLastAlarmStatus(db *sql.DB) string {

	var lastStatus string

	query := `
	SELECT alarm_status
	FROM gas_log
	ORDER BY id DESC
	LIMIT 1
	`

	err := db.QueryRow(query).Scan(&lastStatus)

	if err != nil {
		return ""
	}

	return lastStatus
}

func StoreGas(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	var input models.GasInput

	if err := c.ShouldBindJSON(&input); err != nil {

		helpers.JsonResponse(
			c,
			400,
			false,
			"Invalid request",
			nil,
		)

		return
	}

	// VALIDASI DEVICE
	isValid := validateDevice(
		db,
		input.DeviceID,
		input.DeviceToken,
	)

	if !isValid {

		c.JSON(401, gin.H{
			"success": false,
			"message": "Device tidak valid",
		})

		return
	}

	waSent := "NO"

	lastStatus := getLastAlarmStatus(db)

	// AMAN -> BAHAYA
	if lastStatus != "BAHAYA" && input.AlarmStatus == "BAHAYA" {

		err := services.SendWA(
			"⚠️ GAS BOCOR TERDETEKSI!",
		)

		if err == nil {
			waSent = "YES"
		}
	}

	// BAHAYA -> AMAN
	if lastStatus == "BAHAYA" && input.AlarmStatus == "AMAN" {

		err := services.SendWA(
			"✅ Kondisi gas kembali normal",
		)

		if err == nil {
			waSent = "YES"
		}
	}

	query := `
	INSERT INTO gas_log
	(device_id,status_gas, alarm_status, wifi_status, wa_sent)
	VALUES (?, ?, ?, ?,?)
	`

	_, err := db.Exec(
		query,
		input.DeviceID,
		input.StatusGas,
		input.AlarmStatus,
		input.WifiStatus,
		waSent,
	)

	if err != nil {

		helpers.JsonResponse(
			c,
			500,
			false,
			err.Error(),
			nil,
		)

		return
	}

	// updateLastSeen := `
	// 	UPDATE devices
	// 	SET last_seen = NOW()
	// 	WHERE device_id = ?
	// 	`

	updateLastSeen := `
		UPDATE devices
		SET
			last_seen = NOW(),
			last_seen_unix = ?
		WHERE device_id = ?
		`

	db.Exec(
		updateLastSeen,
		time.Now().Unix(),
		input.DeviceID,
	)

	helpers.JsonResponse(
		c,
		200,
		true,
		"Data berhasil disimpan",
		nil,
	)
}

func validateDevice(
	db *sql.DB,
	deviceID string,
	deviceToken string,
) bool {

	var total int

	query := `
	SELECT COUNT(*)
	FROM devices
	WHERE device_id = ?
	AND device_token = ?
	AND is_active = 1
	`

	err := db.QueryRow(
		query,
		deviceID,
		deviceToken,
	).Scan(&total)

	if err != nil {
		return false
	}

	return total > 0
}
