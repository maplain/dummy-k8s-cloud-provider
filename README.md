# An Example Cloud Provider

## Cloud controller manager

This repository provides tools and scripts for building and testing an example `Kubernetes cloud-controller-manager`.

## Building the cloud provider

This section outlines how to build the cloud provider.

### Building locally

Build locally with the following command:

```shell
$ git clone https://github.com/maplain/dummy-k8s-cloud-provider && \
  go build -o cloud-controller-manager cmd/cloud-controller-manager/main.go
```

The project uses [Go modules](https://github.com/golang/go/wiki/Modules) and:
* Requires Go 1.11+
* Should not be cloned into the `$GOPATH`

### Create a k8s cluster
Refer to either [minikube](https://github.com/kubernetes/minikube) or [kind](https://github.com/kubernetes-sigs/kind) to create a local k8s cluster

### Try it out!
Assume you have a valid kubeconfig at ~/.kube/config.

```shell
export KUBECONFIG=~/.kube/config
➜  git:(master) ✗ ./cloud-controller-manager --authentication-kubeconfig=$KUBECONFIG  --authorization-kubeconfig=$KUBECONFIG --kubeconfig=$KUBECONFIG --cloud-config=examples/cloud-config
I0730 16:24:38.251072   35148 serving.go:319] Generated self-signed cert in-memory
I0730 16:24:38.492400   35148 controllermanager.go:117] Version: v0.0.0-master+$Format:%h$
I0730 16:24:38.493066   35148 secure_serving.go:116] Serving securely on [::]:10258
I0730 16:24:38.493120   35148 leaderelection.go:235] attempting to acquire leader lease  kube-system/cloud-controller-manager...
I0730 16:24:54.762640   35148 leaderelection.go:245] successfully acquired lease kube-system/cloud-controller-manager
I0730 16:24:54.762714   35148 event.go:258] Event(v1.ObjectReference{Kind:"Endpoints", Namespace:"kube-system", Name:"cloud-controller-manager", UID:"3fca798d-3702-4ca2-9d60-892031705b83", APIVersion:"v1", ResourceVersion:"348911", FieldPath:""}): type: 'Normal' reason: 'LeaderElection' fangyuanl-a01.vmware.com_6da701e8-2a14-43ec-816b-c3b2593dd4cf became leader
I0730 16:24:54.774763   35148 node_controller.go:71] Sending events to api server.
I0730 16:24:54.774851   35148 controllermanager.go:244] Started "cloud-node"
E0730 16:24:54.774895   35148 node_controller.go:113] failed to get instances from cloud provider
I0730 16:24:54.779789   35148 node_lifecycle_controller.go:77] Sending events to api server
W0730 16:24:54.779857   35148 core.go:60] failed to start cloud node lifecycle controller: cloud provider does not support instances
W0730 16:24:54.779874   35148 controllermanager.go:241] Skipping "cloud-node-lifecycle"
I0730 16:24:54.784825   35148 controllermanager.go:244] Started "service"
I0730 16:24:54.784846   35148 core.go:91] Will not configure cloud provider routes for allocate-node-cidrs: false, configure-cloud-routes: true.
I0730 16:24:54.784857   35148 service_controller.go:198] Starting service controller
W0730 16:24:54.784863   35148 controllermanager.go:241] Skipping "route"
I0730 16:24:54.784905   35148 controller_utils.go:1029] Waiting for caches to sync for service controller
I0730 16:24:54.889495   35148 controller_utils.go:1036] Caches are synced for service controller
``` 

open another shell:
```shell
kubectl create -f examples/loadbalancer.yaml
```

You should be able to see:
```shell
I0730 16:25:28.176294   35148 loadbalancer.go:34] Hold on! Let me create a dummy Load Balancer ironman for you :)
I0730 16:25:28.176315   35148 event.go:258] Event(v1.ObjectReference{Kind:"Service", Namespace:"default", Name:"ironman", UID:"e98739b8-73d3-499b-80b6-5ac3c395087d", APIVersion:"v1", ResourceVersion:"348972", FieldPath:""}): type: 'Normal' reason: 'EnsuringLoadBalancer' Ensuring load balancer
I0730 16:25:28.176358   35148 event.go:258] Event(v1.ObjectReference{Kind:"Service", Namespace:"default", Name:"ironman", UID:"e98739b8-73d3-499b-80b6-5ac3c395087d", APIVersion:"v1", ResourceVersion:"348972", FieldPath:""}): type: 'Warning' reason: 'UnAvailableLoadBalancer' There are no available nodes for LoadBalancer
I0730 16:25:28.176368   35148 event.go:258] Event(v1.ObjectReference{Kind:"Service", Namespace:"default", Name:"ironman", UID:"e98739b8-73d3-499b-80b6-5ac3c395087d", APIVersion:"v1", ResourceVersion:"348972", FieldPath:""}): type: 'Normal' reason: 'EnsuredLoadBalancer' Ensured load balancer
```

Then run:
```shell
kubectl delete -f examples/loadbalancer.yaml
```
to clean it up.

### Documents
For more details, refer to this [Google Doc](https://docs.google.com/document/d/1Tz9llq1XY4lV4qwswbXDDYw8EB8vxhZcY5YBJBy9U0U/edit#)
