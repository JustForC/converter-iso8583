package controllers

import (
	"converter-iso8583/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/prefix"
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

	spec := &iso8583.MessageSpec{
		Fields: map[int]field.Field{
			0: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "Message Type",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.L,
			}),
			1: field.NewBitmap(&field.Spec{
				Length:      16,
				Description: "Secondary Bit Map",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.L,
			}),
			2: field.NewBitmap(&field.Spec{
				Length:      16,
				Description: "Primary Bit Map",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.L,
			}),
			7: field.NewNumeric(&field.Spec{
				Length:      10,
				Description: "Transmission Date and Time",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.L,
			}),
			11: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "System Trace Audit Number",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.L,
			}),
			15: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "Settlement Date",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.L,
			}),
			70: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "Network Management Information Code",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.L,
			}),
		},
	}

	// return c.JSON(http.StatusOK, input)

	message := iso8583.NewMessage(spec)

	message.MTI("0800")
	err = message.Field(1, input.SecondaryBitMP)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	err = message.Field(2, input.PrimaryBitMap)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	err = message.Field(7, time.Now().UTC().Format("060102150405"))
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	err = message.Field(11, input.SystemTraceAuditNumber)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	err = message.Field(15, input.SettlementDate)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	err = message.Field(70, "201")
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}

	requestMessage, err := message.Pack()
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, requestMessage)
}
