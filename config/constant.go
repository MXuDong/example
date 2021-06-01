package config

// define the config constant here
const (
	// ============================== controller mods: run, debug
	ServerMod_Debug = 0 // debug mod
	ServerMod_Run   = 1 // run mod
	// tcp values
	DefaultTcpPort            = 8000
	DefaultTcpIpAddress       = "0.0.0.0"
	DefaultTcpMaxHandlerCount = 10
	// ============================== docker mods: image, disable
	DockerMod_Disable = 0 // run common
	DockerMod_Image   = 1 // run as image
	// ============================== kubernetes mods: in-cluster, out-cluster, disable
	KubernetesMod_Disable    = 0 // disable kubernetes
	KubernetesMod_InCluster  = 1 // the program run as a pod in the kubernetes cluster
	KubernetesMod_OutCluster = 2 // the program run out of kubernetes

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
