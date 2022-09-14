package models

type EchoTestRequestJSON struct {
	SecondaryBitMap                  string `json:"secondary_bit_map"`
	SystemTraceAuditNumber           string `json:"system_trace_audit_number"`
	SettlementDate                   string `json:"settlement_date"`
	AdditionalData                   string `json:"additional_data"`
	NetworkManagementInformationCode string `json:"network_management_information_code"`
}
