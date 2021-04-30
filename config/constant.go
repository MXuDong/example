package config

// define the config constant here
const (
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
