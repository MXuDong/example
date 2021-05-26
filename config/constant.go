package config

// define the config constant here
const (
	// ============================== controller mods: run, debug
	ServerMod_Debug = iota // debug mod
	ServerMod_Run          // run mod
	// tcp values
	DefaultTcpPort            = 8000
	DefaultTcpIpAddress       = "127.0.0.1"
	DefaultTcpMaxHandlerCount = 10
	// ============================== docker mods: image, disable
	DockerMod_Disable = iota // run common
	DockerMod_Image          // run as image
	// ============================== kubernetes mods: in-cluster, out-cluster, disable
	KubernetesMod_Disable    = iota // disable kubernetes
	KubernetesMod_InCluster         // the program run as a pod in the kubernetes cluster
	KubernetesMod_OutCluster        // the program run out of kubernetes

	// System ====================== System Constant

	// inner invoke response type
	InnerInvokeResponseType_Value   = "value"
	InnerInvokeResponseType_Copy    = "copy"
	InnerInvokeResponseType_Package = "package" // is default type
)

const (
	_ = ServerMod_Debug
	_ = ServerMod_Run
	_ = DockerMod_Disable
	_ = DockerMod_Image
	_ = KubernetesMod_Disable
	_ = KubernetesMod_InCluster
	_ = KubernetesMod_OutCluster
	_ = InnerInvokeResponseType_Value
	_ = InnerInvokeResponseType_Copy
	_ = InnerInvokeResponseType_Package
)
