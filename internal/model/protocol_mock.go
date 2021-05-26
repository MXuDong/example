package model

type TcpMockParam struct {
	ConnectTime     int    `json:"connect_time"`      // Tcp Connect time, second
	RemoteIpAddress string `json:"remote_ip_address"` // tcp server address
	RemoteIpPort    string `json:"remote_ip_port"`    // tcp server port
	SendByteCount   int    `json:"send_byte_count"`   // Tcp send byte pre second
	// tcp protocol, default is tcp, support tcp4, tcp, tcp6, unix, unixpacket
	Protocol string `json:"protocol"`
}
