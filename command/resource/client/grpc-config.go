// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package client

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// GRPCAddrEnvName defines an environment variable name which sets the gRPC
	// address for consul connect envoy. Note this isn't actually used by the api
	// client in this package but is defined here for consistency with all the
	// other ENV names we use.
	GRPCAddrEnvName = "CONSUL_GRPC_ADDR"

	// GRPCTLSEnvName defines an environment variable name which sets the gRPC
	// communication mode. Default is false in plaintext mode.
	GRPCTLSEnvName = "CONSUL_GRPC_TLS"

	// GRPCCAFileEnvName defines an environment variable name which sets the
	// CA file to use for talking to Consul gRPC over TLS.
	GRPCCAFileEnvName = "CONSUL_GRPC_CACERT"

	// GRPCCAPathEnvName defines an environment variable name which sets the
	// path to a directory of CA certs to use for talking to Consul gRPC over TLS.
	GRPCCAPathEnvName = "CONSUL_GRPC_CAPATH"
)

type GRPCConfig struct {
	// Address is the optional address of the Consul server in format of host:port.
	// It doesn't include schema
	Address string

	// GRPCTLS is the optional boolean flag to determine the communication protocol
	GRPCTLS bool

	// CAFile is the optional path to the CA certificate used for Consul
	// communication, defaults to the system bundle if not specified.
	CAFile string

	// CAPath is the optional path to a directory of CA certificates to use for
	// Consul communication, defaults to the system bundle if not specified.
	CAPath string
}

func GetDefaultGRPCConfig() *GRPCConfig {
	return &GRPCConfig{
		Address: "127.0.0.1:8502",
	}
}

func LoadGRPCConfig(defaultConfig *GRPCConfig) (*GRPCConfig, error) {
	if defaultConfig == nil {
		defaultConfig = GetDefaultGRPCConfig()
	}

	overwrittenConfig, err := loadEnvToDefaultConfig(defaultConfig)
	if err != nil {
		return nil, err
	}

	return overwrittenConfig, nil
}

func loadEnvToDefaultConfig(config *GRPCConfig) (*GRPCConfig, error) {
	if addr := os.Getenv(GRPCAddrEnvName); addr != "" {
		if strings.HasPrefix(strings.ToLower(addr), "https://") {
			config.GRPCTLS = true
		}
		config.Address = removeSchemaFromGRPCAddress(addr)
	}

	if tlsMode := os.Getenv(GRPCTLSEnvName); tlsMode != "" {
		doTLS, err := strconv.ParseBool(tlsMode)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", GRPCTLSEnvName, err)
		}
		if doTLS {
			config.GRPCTLS = true
		}
	}

	if caFile := os.Getenv(GRPCCAFileEnvName); caFile != "" {
		config.CAFile = caFile
	}

	if caPath := os.Getenv(GRPCCAPathEnvName); caPath != "" {
		config.CAPath = caPath
	}

	return config, nil
}
