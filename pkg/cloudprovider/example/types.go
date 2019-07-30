package example

import (
	vcfg "github.com/maplain/dummy-k8s-cloud-provider/pkg/common/config"
	k8s "github.com/maplain/dummy-k8s-cloud-provider/pkg/common/kubernetes"
)

// Provider is an implementation of cloud provider Interface
type Provider struct {
	cfg       *vcfg.Config
	informMgr *k8s.InformerManager
}
