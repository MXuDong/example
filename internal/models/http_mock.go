package models

// ResponseCodePresent 返回状态码及相关比例，按照比例进行返回，值会自动计算绝对值
type ResponseCodePresent struct {
	Code    int `json:"code"`
	Present int `json:"present"`
}

// HttpMock 表示当前请求中的请求参数
type HttpMock struct {

	// Method 表示期望本次请求的请求方法
	Method string `json:"method"`

	// ResponseCodePresent 表示本次请求返回的状态码，如果存在多个，则按照权重进行随机选择
	// 如果为空值则所有空数据加和平均计算，负权重会计算其绝对值
	ResponseCodePresent []ResponseCodePresent `json:"codes"`

	// PreDuration 本次请求进行前的操作延时，延时该时间后执行后续操作。
	// 单位（秒）
	PreDuration int `json:"pre_duration"`

	// PostDuration 本次请求进行后的操作延时，操作完成后延时改时间。
	// 单位（秒）
	PostDuration int `json:"post_duration"`

	// Type 表示本次操作的类型，忽略大小写:
	// (empty, others) default: 啥也不做，结果为空
	// - log: 打印日志，并将日志内容记录到结果中
	// - time: 将此时时间记录到结果中
	// - continue: 执行调用下一个请求（仅在链调用时可用，其他情况会转换为 default），并将调用结果记录到结果中
	Type int `json:"type"`

	// ResponseType 表示本次返回值的类型，忽略大小写
	// (empty, others) default : 返回本次操作结果数据
	// - value: 返回 ExtValue 中的数据，不做任何处理
	// - None: 不返回任何数据
	// - Packing: 对结果进行标准化包装并返回, 标准化包装格式： PackingRes
	ResponseType int `json:"response_type"`

	// ExtValue 扩展数据
	ExtValue string `json:"ext_value"`
}

// HttpMockChain
type HttpMockChain struct {
	Mocks []HttpMock `json:"mocks"`
}

// PackingRes
type PackingRes struct {
	StartTime    int64    `json:"start_time"`
	EndTime      int64    `json:"end_time"`
	Type         string   `json:"type"`
	Res          string   `json:"res"`
	RequestParam HttpMock `json:"request_param"`
	Code         string   `json:"code"`
}
