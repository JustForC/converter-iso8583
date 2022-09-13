package models

import (
	"github.com/moov-io/iso8583/field"
)

type EchoTestRequestISO struct {
	MTI                              *field.Numeric `index:"0"`
	PrimaryBitMap                    *field.Bitmap  `index:"2"`
	SecondaryBitMP                   *field.Bitmap  `index:"1"`
	TransmisionDateAndTime           *field.Numeric `index:"7"`
	SystemTraceAuditNumber           *field.Numeric `index:"11"`
	SettlementDate                   *field.Numeric `index:"15"`
	NetworkManagementInformationCode *field.Numeric `index:"70"`
}

type EchoTestRequestJSON struct {
	PrimaryBitMap          string `json:"primary_bit_map"`
	SecondaryBitMP         string `json:"secondary_bit_map"`
	SystemTraceAuditNumber string `json:"system_trace_audit_number"`
	SettlementDate         string `json:"settlement_date"`
}
