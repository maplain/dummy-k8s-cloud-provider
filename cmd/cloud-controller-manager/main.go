/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/

// The external controller manager is responsible for running controller loops that
// are cloud provider dependent. It uses the API to listen to new events on resources.

package main

import (
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/maplain/dummy-k8s-cloud-provider/pkg/cloudprovider/example"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/cloud-controller-manager/app"
)

func init() {
	// Those flags are not used, but it is referenced in vendors, hence it's still registered and hidden from users.
	_ = goflag.String("cloud-provider-gce-lb-src-cidrs", "", "not used")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewCloudControllerManagerCommand()
	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// utilflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.
	// utilflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	// setup for azure
	var versionFlag *pflag.Value
	command.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "cloud-provider" {
			flag.Value.Set(example.ProviderName)
			flag.DefValue = example.ProviderName
			return
		}

		// Set unwanted flags as hidden.
		if flag.Name == "cloud-provider-gce-lb-src-cidrs" {
			flag.Hidden = true
		}
	})

	pflag.CommandLine.VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "version" {
			versionFlag = &flag.Value
		}
	})

	command.Use = example.ProviderName
	innerRun := command.Run
	command.Run = func(cmd *cobra.Command, args []string) {
		if versionFlag != nil && (*versionFlag).String() != "false" {
			os.Exit(0)
		}
		innerRun(cmd, args)
	}

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
