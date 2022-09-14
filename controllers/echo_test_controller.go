package controllers

import (
	"converter-iso8583/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
		Name: "ISO 8583 v1987 Hex",
		Fields: map[int]field.Field{
			0: field.NewString(&field.Spec{
				Length:      4,
				Description: "Message Type Indicator",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			1: field.NewBitmap(&field.Spec{
				Length:      16,
				Description: "Secondary Bit Map",
				Enc:         encoding.BytesToASCIIHex,
				Pref:        prefix.Hex.Fixed,
			}),
			7: field.NewString(&field.Spec{
				Length:      10,
				Description: "Transmission Date & Time",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			11: field.NewString(&field.Spec{
				Length:      6,
				Description: "Systems Trace Audit Number (STAN)",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			70: field.NewString(&field.Spec{
				Length:      3,
				Description: "Network Management Information Code",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
		},
	}

	message := iso8583.NewMessage(spec)

	err = message.Field(0, "0800")
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": err.Error(),
		})
	}
	// Hasilnya selalu sama meskipun inputnya berbeda
	err = message.Field(1, "asd")
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": err.Error(),
		})
	}
	err = message.Field(11, "123456")
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": err.Error(),
		})
	}
	message.Field(70, "301")

	rawMessage, err := message.Pack()
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": fmt.Sprintf("% x", rawMessage),
	})
}
