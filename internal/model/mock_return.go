package model

import "time"

type MockReturn struct {
	InvokeTime      time.Time   `json:"invoke_time"`      // the mock request start time
	EndTime         time.Time   `json:"end_time"`         // the mock response end time
	ValueByte       []byte      `json:"value_byte"`       // the value (convert to the bytes)
	ValueSize       uint64      `json:"value_size"`       // the size of value (convert to the bytes)
	RequestProtocol string      `json:"request_protocol"` // mock protocol
	InvokingSetting interface{} `json:"invoking_setting"` // invoke setting
}
