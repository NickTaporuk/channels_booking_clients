/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

/*import (
	"errors"
	"io/ioutil"

	"github.com/NickTaporuk/channels_booking_clients/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	SupplierCommandName = "supplier"
)
var (
	ErrorSupplierFileIsNotFound = errors.New("file is not found by path")
)

// supplierCmd represents the supplier command
var supplierCmd = &cobra.Command{
	Use:   SupplierCommandName,
	Short: "Check exist supplier of channel api .",
	Long:  `Check exist supplier of channel api .`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err error
		)

		v := args[0]
		if err = utils.CheckIsUUIDTypeOfSupplierFlagValue(v); err == nil {
			if err = utils.CheckSupplierExist(v, channelsClient, ctx, lgr); err != nil {
				return err
			}
		}

		if err != nil {
			err = utils.CheckIsPathTypeOfSupplierFlagValue(v)
			if err == nil {

				data, err := ioutil.ReadFile(v)
				if err != nil {
					lgr.Logger().WithField("Supplier", "error").WithFields(logrus.Fields{"error": ErrorSupplierFileIsNotFound,}).Error(err)
					return err
				}

				err = channelsClient.CreateSupplier(&data)
				if err != nil {
					return err
				}
			}
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(supplierCmd)
	rootCmd.MarkFlagRequired("supplier")

}*/
