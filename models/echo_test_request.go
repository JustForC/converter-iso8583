package models

import "github.com/ideazxy/iso8583"

type EchoTestRequestISO struct {
	SecondaryBitMap                  *iso8583.Alphanumeric `field:"1" length:"16" encode:"ascii"`
	TransmissionDateAndTime          *iso8583.Numeric      `field:"7" length:"10" encode:"ascii"`
	SystemTraceAuditNumber           *iso8583.Numeric      `field:"11" length:"6" encode:"ascii"`
	AdditionalData                   *iso8583.Lllvar       `field:"48" length:"20" encode:"ascii"`
	NetworkManagementInformationCode *iso8583.Numeric      `field:"70" length:"3" encode:"ascii"`
}

type EchoTestRequestJSON struct {
	SecondaryBitMap        string `json:"secondary_bit_map"`
	SystemTraceAuditNumber string `json:"system_trace_audit_number"`
	SettlementDate         string `json:"settlement_date"`
	AdditionalData         string `json:"additional_data"`
}
