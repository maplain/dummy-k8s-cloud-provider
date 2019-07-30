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

package config

// Config is used to read and store information from the cloud configuration file
type Config struct {
	Global struct {
		// Name of the secret were vCenter credentials are present.
		SecretName string `gcfg:"secret-name"`
		// Secret Namespace where secret will be present that has vCenter credentials.
		SecretNamespace string `gcfg:"secret-namespace"`
		// The kubernetes service account used to launch the cloud controller manager.
		// Default: cloud-controller-manager
		ServiceAccount string `gcfg:"service-account"`
		// Secret directory in the event that:
		// 1) we don't want to use the k8s API to listen for changes to secrets
		// 2) we are not in a k8s env, namely DC/OS, since CSI is CO agnostic
		// Default: /etc/cloud/credentials
		SecretsDirectory string `gcfg:"secrets-directory"`
		// Disable the CCM API
		// Default: true
		APIDisable bool `gcfg:"api-disable"`
		// Configurable CCM API port
		// Default: 43001
		APIBinding string `gcfg:"api-binding"`
	}
}
