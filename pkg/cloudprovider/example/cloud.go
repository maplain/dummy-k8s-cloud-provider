package example

import (
	"io"

	"k8s.io/klog"

	cloudprovider "k8s.io/cloud-provider"

	vcfg "github.com/maplain/dummy-k8s-cloud-provider/pkg/common/config"
	k8s "github.com/maplain/dummy-k8s-cloud-provider/pkg/common/kubernetes"
)

const (
	// ProviderName is the name of the cloud provider registered with
	// Kubernetes.
	ProviderName string = "example-cloud-provider"
)

func init() {
	cloudprovider.RegisterCloudProvider(ProviderName, func(config io.Reader) (cloudprovider.Interface, error) {
		cfg, err := vcfg.ReadConfig(config)
		if err != nil {
			return nil, err
		}
		return newProvider(&cfg)
	})
}

// Creates new Controller node interface and returns
func newProvider(cfg *vcfg.Config) (*Provider, error) {
	cp := &Provider{
		cfg: cfg,
	}
	return cp, nil
}

// Initialize initializes the cloud provider.
func (cp *Provider) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	client, err := clientBuilder.Client(cp.cfg.Global.ServiceAccount)
	if err == nil {
		klog.V(1).Info("Kubernetes Client Init Succeeded")

		cp.informMgr = k8s.NewInformer(&client)
		cp.informMgr.Listen()

	} else {
		klog.Errorf("Kubernetes Client Init Failed: %v", err)
	}
}

// LoadBalancer returns a balancer interface. Also returns true if the
// interface is supported, false otherwise.
func (cp *Provider) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	klog.V(1).Info("The cloud provider does not support load balancers")
	return &loadBalancer{}, true
}

// Instances returns an instances interface. Also returns true if the
// interface is supported, false otherwise.
func (cp *Provider) Instances() (cloudprovider.Instances, bool) {
	klog.V(1).Info("Enabling Instances interface on cloud provider")
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface
// is supported, false otherwise.
func (cp *Provider) Zones() (cloudprovider.Zones, bool) {
	klog.V(1).Info("Enabling Zones interface on cloud provider")
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface
// is supported, false otherwise.
func (cp *Provider) Clusters() (cloudprovider.Clusters, bool) {
	klog.V(1).Info("The cloud provider does not support clusters")
	return nil, false
}

// Routes returns a routes interface along with whether the interface
// is supported.
func (cp *Provider) Routes() (cloudprovider.Routes, bool) {
	klog.V(1).Info("The cloud provider does not support routes")
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
