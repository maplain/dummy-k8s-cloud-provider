package guestcluster

import (
	vcfg "gitlab.eng.vmware.com/core-build/cloud-provider-guest-cluster/pkg/common/config"
	cm "gitlab.eng.vmware.com/core-build/cloud-provider-guest-cluster/pkg/common/connectionmanager"
	k8s "gitlab.eng.vmware.com/core-build/cloud-provider-guest-cluster/pkg/common/kubernetes"
)

// Provider is an implementation of cloud provider Interface for guest cluster.
type Provider struct {
	cfg               *vcfg.Config
	connectionManager *cm.ConnectionManager
	informMgr         *k8s.InformerManager
}
