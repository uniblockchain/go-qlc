/*
 * Copyright (c) 2019 QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package commands

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/qlcchain/go-qlc/rpc"
	"github.com/spf13/cobra"
)

func walletRemove() {
	var accountP string
	if interactive {
		account := Flag{
			Name:  "account",
			Must:  true,
			Usage: "account for wallet",
			Value: "",
		}
		c := &ishell.Cmd{
			Name: "walletremove",
			Help: "remove a wallet",
			Func: func(c *ishell.Context) {
				args := []Flag{account}
				if HelpText(c, args) {
					return
				}
				if err := CheckArgs(c, args); err != nil {
					Warn(err)
					return
				}
				accountP = StringVar(c.Args, account)

				err := removeWallet(accountP)
				if err != nil {
					Warn(err)
					return
				}
			},
		}
		shell.AddCmd(c)
	} else {
		var wrCmd = &cobra.Command{
			Use:   "walletremove",
			Short: "remove wallet",
			Run: func(cmd *cobra.Command, args []string) {
				err := removeWallet(accountP)
				if err != nil {
					cmd.Println(err)
					return
				}
			},
		}
		wrCmd.Flags().StringVarP(&accountP, "account", "a", "", "wallet address")
		rootCmd.AddCommand(wrCmd)
	}
}

func removeWallet(accountP string) error {
	client, err := rpc.Dial(endpointP)
	if err != nil {
		return err
	}
	defer client.Close()
	err = client.Call(nil, "wallet_remove", accountP)
	if err != nil {
		return err
	}
	s := fmt.Sprintf("remove wallet %s success", accountP)
	if interactive {
		Info(s)
	} else {
		fmt.Println(s)
	}
	return nil
}
