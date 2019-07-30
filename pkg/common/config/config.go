/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/

package config

import (
	"fmt"
	"io"

	gcfg "gopkg.in/gcfg.v1"
)

func ReadConfig(config io.Reader) (Config, error) {
	if config == nil {
		return Config{}, fmt.Errorf("no cloud provider config file given")
	}

	cfg := Config{}

	err := gcfg.FatalOnly(gcfg.ReadInto(&cfg, config))
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
