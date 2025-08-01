// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/srl-labs/containerlab/cmd/common"
	"github.com/srl-labs/containerlab/core"
	"github.com/srl-labs/containerlab/runtime"
	"github.com/srl-labs/containerlab/utils"
)

var cntName string

// upgradeCmd represents the version command.
var disableTxOffloadCmd = &cobra.Command{
	Use:   "disable-tx-offload",
	Short: "disables tx checksum offload on eth0 interface of a container",

	PreRunE: common.CheckAndGetRootPrivs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		opts := []core.ClabOption{
			core.WithTimeout(common.Timeout),
			core.WithRuntime(common.Runtime,
				&runtime.RuntimeConfig{
					Debug:            common.Debug,
					Timeout:          common.Timeout,
					GracefulShutdown: common.Graceful,
				},
			),
			core.WithDebug(common.Debug),
		}
		c, err := core.NewContainerLab(opts...)
		if err != nil {
			return err
		}

		node, err := c.GetNode(cntName)
		if err != nil {
			return err
		}

		err = node.ExecFunction(ctx, utils.NSEthtoolTXOff(cntName, "eth0"))
		if err != nil {
			return err
		}

		log.Infof("Tx checksum offload disabled for eth0 interface of %s container", cntName)
		return nil
	},
}

func init() {
	toolsCmd.AddCommand(disableTxOffloadCmd)
	disableTxOffloadCmd.Flags().StringVarP(&cntName, "container", "c", "", "container name to disable offload in")
	_ = disableTxOffloadCmd.MarkFlagRequired("container")
}
