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
	ProgramLogConfig *ProgramLogConfig `json:"program_log_config"` // the log config
	ServerConfig     *ServerConfig     `json:"server_config"`      // the controller config
	KubernetesConfig *KubernetesConfig `json:"kubernetes_config"`  // the kubernetes config
	DockerConfig     *DockerConfig     `json:"docker_config"`      // the docker config
}

// the program log config, the log can't change behavior of the server log
type ProgramLogConfig struct {
	LogLevel string `json:"log_level"` // the log level
}

// ServerConfig define the program behavior, suck like port, mod
type ServerConfig struct {
	Port string `json:"run_port"` // the controller run port
	Mod  int    `json:"run_mod"`  // the controller run mod

	EnableTcpServer bool   `json:"enable_tcp_server"` // if true, enable the tcp server
	TcpServerPort   int    `json:"tcp_server_port"`   // if EnableTcpServer is true, tcp will try listen this port
	TcpNetWork      string `json:"tcp_net_work"`      // tcp's net work, support tcp, tcp4, tcp6, unix and unixpacket
	TcpAddress      string `json:"tcp_address"`       // tcp's address
	MaxHandlerCount int    `json:"max_handler_count"` // max handler of tcp

	EnableUdpServer bool `json:"enable_udp_server"` // if true, enable the udp server

}

// KubernetesConfig define the program run with kubernetes, use KubernetesMod to switch kubernetes mod with feature of kubernetes support
type KubernetesConfig struct {
	// the kubernetes mod, the values: KubernetesMod_Disable, KubernetesMod_InCluster, KubernetesMod_OutCluster
	Mod int `json:"kubernetes_mod"`

	// if the out-side of kubernetes cluster, it is the config path to kube-client config.
	// if empty, set to home/.kube/config
	// if Mod is disable(0), set to string 'DISABLE', if in-side, set to IN-SIDE(it is auto).
	// if init fail, set to 'INIT
	Config string `json:"config"`
}

// Enable will return true when enable the kubernetes feature
func (k *KubernetesConfig) Enable() bool {
	return k.Mod != KubernetesMod_Disable
}

// DockerConfig define the program run with docker, use DockerMod to switch docker mod with feature of docker support
type DockerConfig struct {
	Mod int `json:"docker_mod"` // the docker mod, the value: DockerMod_Disable, DockerMod_Image
}

// ============================== the config instance
var Config = config{
	ProgramLogConfig: &ProgramLogConfig{
		LogLevel: DefaultLogLevel,
	},
	ServerConfig: &ServerConfig{
		Mod:  DefaultServerMod,
		Port: DefaultServerPort,

		// tcp server
		EnableTcpServer: true,
		TcpNetWork:      constant.TcpProtocol,
		TcpServerPort:   DefaultTcpPort,
		TcpAddress:      DefaultTcpIpAddress,
		MaxHandlerCount: DefaultTcpMaxHandlerCount,

		// udp server
		EnableUdpServer: false,
	},
	DockerConfig: &DockerConfig{
		Mod: DefaultDockerMod,
	},
	KubernetesConfig: &KubernetesConfig{
		Mod:    KubernetesMod_Auto,
		Config: KubernetesConfig_Default,
	},
}
