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

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"syreclabs.com/go/faker"
)

var (
	ErrorSupplierIsNotFound     = errors.New("supplier isn't found")
	ErrorSuppliersIsMoreThanOne = errors.New("cli value of supplier is more than one")
	ErrorSupplierFileIsNotFound = errors.New("file is not found by path")
)

// supplierCmd represents the supplier command
var supplierCmd = &cobra.Command{
	Use:   "supplier",
	Short: "Check exist supplier of channel api .",
	Long:  `Check exist supplier of channel api .`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err error
		)

		if len(args) > 1 {
			lgr.Logger().WithField("Supplier", "error").WithFields(logrus.Fields{"error": ErrorSuppliersIsMoreThanOne,}).Error(err)
			return ErrorSuppliersIsMoreThanOne
		}

		v := args[0]
		if err = CheckIsUUIDTypeOfSupplierFlagValue(v); err == nil {
			if err = CheckSupplierExist(v); err != nil {
				return err
			}
		}

		if err != nil {
			_, err = CheckIsPathTypeOfSupplierFlagValue(v)
			if err == nil {
				supplier := new(swclient.RequestPostSupplierEnvelope)

				data, err := ioutil.ReadFile(v)
				if err != nil {
					lgr.Logger().WithField("Supplier", "error").WithFields(logrus.Fields{"error": ErrorSupplierFileIsNotFound,}).Error(err)
				}

				if err = json.Unmarshal([]byte(data), &supplier); err != nil {
					return err
				}
				supplier.Meta.ReqId = uuid.New().String()
				supplier.Supplier.Id = uuid.New().String()
				supplier.Supplier.Code = faker.App().Author() + faker.App().Name()

				lgr.Logger().WithFields(logrus.Fields{"file data": string(data),}).Debug(" data from supplier json file")
				lgr.Logger().WithFields(logrus.Fields{"Supplier": supplier,}).Debug(supplier)

				ResponsePostSupplierEnvelope, resp, err := channelsClient.Client.SuppliersApi.CreateSupplier(ctx, *supplier)

				lgr.Logger().WithFields(logrus.Fields{"ResponsePostSupplierEnvelope": ResponsePostSupplierEnvelope, "create supplier response resp statusCode": resp.StatusCode, "err": err}).Debug("ResponsePostSupplierEnvelope")

				if err != nil {
					lgr.Logger().WithFields(logrus.Fields{"ResponsePostSupplierEnvelope": ResponsePostSupplierEnvelope, "response resp statusCode": resp.StatusCode, "create supplier body": resp.Body, "err": err}).Error("Channel api create supplier error")
					return err
				}
			}
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(supplierCmd)
}

func CheckSupplierExist(supplierID string) error {
	getSupplier, resp, err = channelsClient.Client.SuppliersApi.GetSupplier(ctx, supplierID)

	if getSupplier.Supplier == nil {
		return ErrorSupplierIsNotFound
	}

	lgr.Logger().WithField("Supplier", "done").WithFields(logrus.Fields{"Supplier response": getSupplier, "response status": resp.StatusCode}).Info("Supplier creation was done")

	channelsClient.SetSupplierID(getSupplier.Supplier.Id)

	return nil
}

// CheckIsUUIDTypeOfSupplierFlagValue is used to check value is uuid type
func CheckIsUUIDTypeOfSupplierFlagValue(value string) error {
	_, err := uuid.Parse(value)

	return err
}

// CheckIsUUIDTypeOfSupplierFlagValue is used to check value is uuid type
func CheckIsPathTypeOfSupplierFlagValue(value string) (*os.FileInfo, error) {
	var (
		err      error
		fileInfo os.FileInfo
	)
	if fileInfo, err = os.Stat(value); os.IsNotExist(err) {}

	return &fileInfo, err
}
