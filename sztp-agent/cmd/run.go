/*
SPDX-License-Identifier: Apache-2.0
Copyright (C) 2022-2023 Intel Corporation
Copyright (c) 2022 Dell Inc, or its subsidiaries.
Copyright (C) 2022 Red Hat.
*/
// Package cmd implements the CLI commands
package cmd

import (
	"github.com/opiproject/sztp/sztp-agent/pkg/secureagent"
	"github.com/spf13/cobra"
)

// NewRunCommand returns the run command
func NewRunCommand() *cobra.Command {
	var (
		bootstrapURL             string
		serialNumber             string
		devicePassword           string
		devicePrivateKey         string
		deviceEndEntityCert      string
		bootstrapTrustAnchorCert string
	)

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Exec the run command",
		RunE: func(cmd *cobra.Command, args []string) error {
			a := secureagent.NewAgent(bootstrapURL, serialNumber, devicePassword, devicePrivateKey, deviceEndEntityCert, bootstrapTrustAnchorCert)
			return a.RunCommand()
		},
	}

	flags := cmd.Flags()
	// TODO this options should be retrieved automatically instead of requests in the agent
	// Opened discussion to define the procedure: https://github.com/opiproject/sztp/issues/2
	flags.StringVar(&bootstrapURL, "bootstrap-url", "", "Bootstrap server URL")
	flags.StringVar(&serialNumber, "serial-number", "", "Device's serial number")
	flags.StringVar(&devicePassword, "device-password", "", "Device's password")
	flags.StringVar(&devicePrivateKey, "device-private-key", "", "Device's private key")
	flags.StringVar(&deviceEndEntityCert, "device-end-entity-cert", "", "Device's End Entity cert")
	flags.StringVar(&bootstrapTrustAnchorCert, "bootstrap-trust-anchor-cert", "", "Bootstrap server trust anchor Cert")

	return cmd
}
