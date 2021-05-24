package model

/**
This file use to add invoke model for inner func.

Like error return and so on.
*/

type TracePostParam struct {
	// TargetRespCode to limit response code. All of http.code is support.
	// But for any code, only set to response code, but not effect the action of code.
	// Like if set response with 204, for http protocol, it should return with nothing.
	// But it also write response value.
	// This param support more than one response
	// For example, if want to return 200 for 30% and 204 for 70%, it can be write: 30%200,70%204
	// The present only for int, if all value's sum not equals 100, it will re-cal.
	// For more, please see `pkg.util.random.AnalyzePercentage`
	TargetRespCode string `json:"target_resp_code"`

	// The BeforeSleep will block process the request next trace. If it is last item of trace,
	// the next step is AfterSleep.
	// The BeforeSleep also support present value, like:
	BeforeSleep string `json:"before_sleep"`

	// The AfterSleep will block process when next trace is response.
	AfterSleep string `json:"after_sleep"`

	// The ResponseType define the response. Default is package
	// The type:
	// 1. InnerInvokeResponseType_Value = value, use ResponseExpendField value as response value, can set empty value.
	// 2. InnerInvokeResponseType_Copy = copy, copy next request's response as this response
	// 3. InnerInvokeResponseType_Package = package, package will packing the response of next request to ResponseResp.Object
	// Default type is response_type
	ResponseType string `json:"response_type"`

	// ResponseExpendField for expand some info, can be empty
	ResponseExpendField string `json:"response_expend_field"`

	// If next url is nil, it's mean trace is end
	NextUrl string `json:"next_url"`
}

type TraceResponseResp struct {
	ProcessStartTime int64          `json:"process_start_time"`
	ProcessStopTime  int64          `json:"process_stop_time"`
	RequestSettings  TracePostParam `json:"request_settings"`
	Object           interface{}    `json:"object"` // any object, but usually to save next response
	Value            interface{}    `json:"value"`  // the response want to return value
}
