/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubernetes

import (
	"time"

	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
	listerv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/sample-controller/pkg/signals"
)

func noResyncPeriodFunc() time.Duration {
	return 0
}

// NewInformer creates a newk8s client based on a service account
func NewInformer(client *clientset.Interface) *InformerManager {
	return &InformerManager{
		client:          client,
		stopCh:          signals.SetupSignalHandler(),
		informerFactory: informers.NewSharedInformerFactory(*client, noResyncPeriodFunc()),
	}
}

// GetSecretListener creates a lister to use
func (im *InformerManager) GetSecretListener() listerv1.SecretLister {
	if im.secretInformer == nil {
		im.secretInformer = im.informerFactory.Core().V1().Secrets()
	}

	return im.secretInformer.Lister()
}

// AddNodeListener hooks up add, update, delete callbacks
func (im *InformerManager) AddNodeListener(add, remove func(obj interface{}), update func(oldObj, newObj interface{})) {
	if im.nodeInformer == nil {
		im.nodeInformer = im.informerFactory.Core().V1().Nodes().Informer()
	}

	im.nodeInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    add,
		UpdateFunc: update,
		DeleteFunc: remove,
	})
}

// Listen starts the Informers
func (im *InformerManager) Listen() {
	go im.informerFactory.Start(im.stopCh)
}
