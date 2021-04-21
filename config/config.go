package config

// define the config here
const (
	// I18n value
	// ============================== server mods: run, debug
	ServerMod_Debug = iota // debug mod
	ServerMod_Run          // run mod
	// ============================== docker mods: image, disable
	DockerMod_Disable = iota // run common
	DockerMod_Image          // run as image
	// ============================== kubernetes mods: in-cluster, out-cluster, disable
	KubernetesMod_Disable    = iota // disable kubernetes
	KubernetesMod_InCluster         // the program run as a pod in the kubernetes cluster
	KubernetesMod_OutCluster        // the program run out of kubernetes
)

// default values
const (
	DefaultServerPort    = ":3000"               // default gin server port is 3000
	DefaultServerMod     = ServerMod_Debug       // default server mod is ServerMod_Debug
	DefaultDockerMod     = DockerMod_Disable     // default docker mod is DockerMod_Disable
	DefaultKubernetesMod = KubernetesMod_Disable // default kubernetes mod is KubernetesMod_Disable
)

type config struct {
	ServerConfig     ServerConfig     `json:"server_config"`     // the server config
	KubernetesConfig KubernetesConfig `json:"kubernetes_config"` // the kubernetes config
	DockerConfig     DockerConfig     `json:"docker_config"`     // the docker config
}

// ServerConfig define the program behavior, suck like port, mod
type ServerConfig struct {
	Port string `json:"run_port"` // the server run port
	Mod  int    `json:"run_mod"`  // the server run mod
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
	},
	DockerConfig: DockerConfig{
		Mod: DefaultDockerMod,
	},
	KubernetesConfig: KubernetesConfig{
		Mod: DefaultKubernetesMod,
	},
}
