package utils

import (
	"context"
	"errors"
	"net/http"
	"os"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/NickTaporuk/channels_booking_clients/channels"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	ErrorSupplierIsNotFound = errors.New("supplier isn't found")
)

// CheckIsUUIDTypeOfSupplierFlagValue is used to check value is uuid type
func CheckIsUUIDTypeOfSupplierFlagValue(value string) error {
	_, err := uuid.Parse(value)

	return err
}

// CheckIsUUIDTypeOfSupplierFlagValue is used to check value is uuid type
func CheckIsPathTypeOfSupplierFlagValue(value string) error {
	var (
		err error
	)

	_, err = os.Stat(value);
	return err
}

// CheckSupplierExist is used to check exist supplier by uuid
func CheckSupplierExist(supplierID string, channelsClient *channels.ChannelsClient, ctx context.Context, lgr *logger.LocalLogger) error {
	var (
		supplier swclient.ResponseGetSupplierEnvelope
		resp     *http.Response
		err      error
	)
	supplier, resp, err = channelsClient.Client.SuppliersApi.GetSupplier(ctx, supplierID)
	if err != nil {
		return err
	}

	if supplier.Supplier == nil {
		return ErrorSupplierIsNotFound
	}

	lgr.Logger().WithField("Supplier", "done").WithFields(logrus.Fields{"Supplier response": supplier, "response status": resp.StatusCode}).Info("Supplier creation was done")

	channelsClient.SetSupplierID(supplier.Supplier.Id)

	return nil
}
