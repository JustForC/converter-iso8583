package controllers

import (
	"converter-iso8583/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ideazxy/iso8583"
	"github.com/labstack/echo/v4"
)

func EchoTestControllers(c echo.Context) error {
	input := models.EchoTestRequestJSON{}

	err := json.NewDecoder(c.Request().Body).Decode(&input)
	if err != nil {
		log.Printf("Error read request body with err: %s", err)
		return c.JSON(500, map[string]interface{}{
			"message": "Error read request body",
			"page":    nil,
			"data":    nil,
			"error":   nil,
		})
	}
	defer c.Request().Body.Close()

	data := &models.EchoTestRequestISO{
		SecondaryBitMap:                  iso8583.NewAlphanumeric(input.SecondaryBitMap),
		TransmissionDateAndTime:          iso8583.NewNumeric(time.Now().UTC().Format("0601150405")),
		SystemTraceAuditNumber:           iso8583.NewNumeric(input.SystemTraceAuditNumber),
		AdditionalData:                   iso8583.NewLllvar([]byte(input.AdditionalData)),
		NetworkManagementInformationCode: iso8583.NewNumeric("001"),
	}
	msg := iso8583.NewMessage("0800", data)
	msg.MtiEncode = iso8583.ASCII
	b, err := msg.Bytes()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": fmt.Sprintf("%x", b),
	})
}
