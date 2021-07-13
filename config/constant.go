package config

// define the config constant here
const (
	// log level
	DefaultLogLevel = "info"
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
	KubernetesMod_Auto       = 1 // auto set kubernetes model, try out side first, and try in side, last set disable
	KubernetesMod_InCluster  = 2 // the program run as a pod in the kubernetes cluster
	KubernetesMod_OutCluster = 3 // the program run out of kubernetes

	KubernetesConfig_InitFail = "INIT-FAIL" // if kube-config init fail, set Kubernetes.Config to it.
	KubernetesConfig_InSide   = "IN-SIDE"   // if kubernetes.Mod is KubernetesMod_InCluster, set to it.
	KubernetesConfig_Disable  = "DISABLE"   // if kubernetes.Mod is KubernetesMod_Disable, set to it.
	KubernetesConfig_Default  = ""          // the Kubernetes.Config default value

	// System ====================== System Constant

	// inner invoke response type
	InnerInvokeResponseType_Value   = "value"
	InnerInvokeResponseType_Copy    = "copy"
	InnerInvokeResponseType_Package = "package" // is default type
)
