package guestcluster

import (
	"io"

	"k8s.io/klog"

	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/kubernetes/pkg/controller"

	vcfg "gitlab.eng.vmware.com/core-build/cloud-provider-guest-cluster/pkg/common/config"
	cm "gitlab.eng.vmware.com/core-build/cloud-provider-guest-cluster/pkg/common/connectionmanager"
	k8s "gitlab.eng.vmware.com/core-build/cloud-provider-guest-cluster/pkg/common/kubernetes"
)

const (
	// ProviderName is the name of the cloud provider registered with
	// Kubernetes.
	ProviderName string = "vsphere-guest-cluster"
)

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		cfg, err := vcfg.ReadConfig(config)
		if err != nil {
			return nil, err
		}
		return newProvider(cfg)
	})
}

// Creates new Controller node interface and returns
func newProvider(cfg *vcfg.Config) (*Provider, error) {
	cp := Provider{
		cfg: cfg,
	}
	return cp, nil
}

// Initialize initializes the cloud provider.
func (cp *Provider) Initialize(clientBuilder controller.ControllerClientBuilder) {
	client, err := clientBuilder.Client(cp.cfg.Global.ServiceAccount)
	if err == nil {
		klog.V(1).Info("Kubernetes Client Init Succeeded")

		cp.informMgr = k8s.NewInformer(client)

		connMgr := cm.NewConnectionManager(cp.cfg, cp.informMgr.GetSecretListener())
		cp.connectionManager = connMgr

		cp.informMgr.Listen()

		if !cp.cfg.Global.APIDisable {
			klog.V(1).Info("Starting the API Server")
			cp.server.Start()
		} else {
			klog.V(1).Info("API Server is disabled")
		}
	} else {
		klog.Errorf("Kubernetes Client Init Failed: %v", err)
	}
}

// LoadBalancer returns a balancer interface. Also returns true if the
// interface is supported, false otherwise.
func (cp *Provider) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	klog.V(1).Info("The vSphere cloud provider does not support load balancers")
	return nil, false
}

// Instances returns an instances interface. Also returns true if the
// interface is supported, false otherwise.
func (cp *Provider) Instances() (cloudprovider.Instances, bool) {
	klog.V(1).Info("Enabling Instances interface on vSphere cloud provider")
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface
// is supported, false otherwise.
func (cp *Provider) Zones() (cloudprovider.Zones, bool) {
	klog.V(1).Info("Enabling Zones interface on vSphere cloud provider")
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface
// is supported, false otherwise.
func (cp *Provider) Clusters() (cloudprovider.Clusters, bool) {
	klog.V(1).Info("The vSphere cloud provider does not support clusters")
	return nil, false
}

// Routes returns a routes interface along with whether the interface
// is supported.
func (cp *Provider) Routes() (cloudprovider.Routes, bool) {
	klog.V(1).Info("The vSphere cloud provider does not support routes")
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (cp *Provider) ProviderName() string {
	return ProviderName
}

// HasClusterID returns true if a ClusterID is required and set/
func (cp *Provider) HasClusterID() bool {
	return true
}
