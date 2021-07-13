package config

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

// InitKubernetesConfig will set kubernetes from config
func InitKubernetesConfig() {
	try_init() // try init
}

// try_init will init kubernetes config, if fail, set mod to disable, and config to FAIL or DISABLE
func try_init() {
	c := Ctl.Config

	switch c.KubernetesConfig.Mod {
	case KubernetesMod_Disable:
		// disable the kubernetes feature
		c.KubernetesConfig.Config = KubernetesConfig_Disable
		Ctl.Logrus.Infoln("Kubernetes feature is disable")
		return
	case KubernetesMod_InCluster:
		cs := initInSideConfig()
		if cs != nil {
			c.KubernetesConfig.Config = KubernetesConfig_InSide
			Ctl.ClientSet = cs
			Ctl.Logrus.Infoln("Init kube config by in-side mode")
		}else{
			Ctl.Logrus.Warn("Kube-Client-Set nil, set mod to disable")
			// set mod to disable
			Ctl.Config.KubernetesConfig.Mod = KubernetesMod_Disable
		}
	case KubernetesMod_OutCluster:
		if c.KubernetesConfig.Config == KubernetesConfig_Default {
			if home := homedir.HomeDir(); home != "" {
				c.KubernetesConfig.Config = filepath.Join(home, ".kube", "config")
			}
		}
		cs := initOutSideConfig(c.KubernetesConfig.Config)
		if cs != nil {
			Ctl.ClientSet = cs
			Ctl.Logrus.Infoln("Init kube config by out-side mode")
		}else{
			Ctl.Logrus.Warn("Kube-Client-Set nil, set mod to disable")
			// set mod to disable
			Ctl.Config.KubernetesConfig.Mod = KubernetesMod_Disable
		}
	//case KubernetesMod_Auto:
	default:
		// other value set to AUTO
		Ctl.Logrus.Warn("Init mode is AUTO")
		configPath := c.KubernetesConfig.Config // save path
		c.KubernetesConfig.Mod = KubernetesMod_InCluster
		try_init()
		if c.KubernetesConfig.Config == KubernetesConfig_InitFail {
			// in-side fail
			Ctl.Logrus.Warn("Try init in-side fail")
		} else {
			return
		}
		c.KubernetesConfig.Config = configPath // reset path
		c.KubernetesConfig.Mod = KubernetesMod_OutCluster
		try_init()
		if c.KubernetesConfig.Config == KubernetesConfig_InitFail {
			// in-side fail
			Ctl.Logrus.Warn("Try init out-side fail")
		} else {
			return
		}
		c.KubernetesConfig.Mod = KubernetesMod_Disable
		try_init()
		return
	}
}

func initOutSideConfig(kubeconfig string) *kubernetes.Clientset {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		Ctl.Logrus.Errorf("Init kube config fail, mod is out-side of kubenretes, err : %v", err)
		Ctl.Config.KubernetesConfig.Config = KubernetesConfig_InitFail
		return nil
	}
	cs, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		Ctl.Logrus.Errorf("Init kube config fail, mod is out-side of kubenretes, err : %v", err)
		Ctl.Config.KubernetesConfig.Config = KubernetesConfig_InitFail
		return nil
	}
	return cs
}

func initInSideConfig() *kubernetes.Clientset {
	kubeConfig, err := rest.InClusterConfig()
	if err != nil {
		Ctl.Logrus.Errorf("Init kube config fail, mod is in-side of kubenretes, err : %v", err)
		Ctl.Config.KubernetesConfig.Config = KubernetesConfig_InitFail
		return nil
	}

	cs, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		Ctl.Logrus.Errorf("Init kube config fail, mod is out-side of kubenretes, err : %v", err)
		Ctl.Config.KubernetesConfig.Config = KubernetesConfig_InitFail
		return nil
	}
	return cs
}
