package config

import "github.io/MXuDong/example/pkg/constant"

// default values
const (
	DefaultServerPort    = ":3000"               // default gin controller port is 3000
	DefaultServerMod     = ServerMod_Debug       // default controller mod is ServerMod_Debug
	DefaultDockerMod     = DockerMod_Disable     // default docker mod is DockerMod_Disable
	DefaultKubernetesMod = KubernetesMod_Disable // default kubernetes mod is KubernetesMod_Disable
)

// config define all info for project
type config struct {
	ServerConfig     ServerConfig     `json:"server_config"`     // the controller config
	KubernetesConfig KubernetesConfig `json:"kubernetes_config"` // the kubernetes config
	DockerConfig     DockerConfig     `json:"docker_config"`     // the docker config
}

// ServerConfig define the program behavior, suck like port, mod
type ServerConfig struct {
	Port string `json:"run_port"` // the controller run port
	Mod  int    `json:"run_mod"`  // the controller run mod

	EnableTcpServer    bool   `json:"enable_tcp_server"`     // if true, enable the tcp server
	TcpServerPort      int    `json:"tcp_server_port"`       // if EnableTcpServer is true, tcp will try listen this port
	TcpNetWork         string `json:"tcp_net_work"`          // tcp's net work, support tcp, tcp4, tcp6, unix and unixpacket
	TcpAddress         string `json:"tcp_address"`           // tcp's address
	TcpMaxHandlerCount int    `json:"tcp_max_handler_count"` // max handler of tcp

	EnableUdpServer    bool   `json:"enable_udp_server"` // if true, enable the udp server
	UdpServerPort      int    `json:"udp_server_port"`
	UdpNetWork         string `json:"udp_net_work"`
	UdpAddress         string `json:"udp_address"`
	UdpMaxHandlerCount int    `json:"udp_max_handler_count"`
	UdpReadBuffer      int    `json:"udp_read_buffer"`
}

// KubernetesConfig define the program run with kubernetes, use KubernetesMod to switch kubernetes mod with feature of kubernetes support
type KubernetesConfig struct {
	Mod int `json:"kubernetes_mod"` // the kubernetes mod, the values: KubernetesMod_Disable, KubernetesMod_InCluster, KubernetesMod_OutCluster
}

// DockerConfig define the program run with docker, use DockerMod to switch docker mod with feature of docker support
type DockerConfig struct {
	Mod int `json:"docker_mod"` // the docker mod, the value: DockerMod_Disable, DockerMod_Image
}

// ============================== the config instance
var Config = config{
	ServerConfig: ServerConfig{
		Mod:  DefaultServerMod,
		Port: DefaultServerPort,

		// tcp server
		EnableTcpServer:    true,
		TcpNetWork:         constant.TcpProtocol,
		TcpServerPort:      DefaultTcpPort,
		TcpAddress:         DefaultTcpIpAddress,
		TcpMaxHandlerCount: DefaultTcpMaxHandlerCount,

		// udp server
		EnableUdpServer:    true,
		UdpServerPort:      DefaultUdpPort,
		UdpNetWork:         constant.UdpProtocol,
		UdpAddress:         DefaultUdpIpAddress,
		UdpMaxHandlerCount: DefaultUdpMaxHandlerCount,
		UdpReadBuffer:      DefaultUdpReaderBufferSize,
	},
	DockerConfig: DockerConfig{
		Mod: DefaultDockerMod,
	},
	KubernetesConfig: KubernetesConfig{
		Mod: DefaultKubernetesMod,
	},
}
