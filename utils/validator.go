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
	ErrorProductIsNotFound  = errors.New("product isn't found")
	ErrorRateIsNotFound     = errors.New("rate isn't found")
)

// CheckIsUUIDTypeOfSupplierFlagValue is used to check value is uuid type
func CheckIsUUIDTypeOfFlagValue(value string) error {
	_, err := uuid.Parse(value)

	return err
}

// CheckIsPathTypeOfFlagValue is used to check value is uuid type
func CheckIsPathTypeOfFlagValue(value string) error {
	var (
		err error
	)

	_, err = os.Stat(value);
	return err
}

// CheckSupplierExist is used to check exist supplier by uuid
func CheckSupplierExist(supplierID string, channelsClient *channels.ChannelsClient, ctx *context.Context, lgr *logger.LocalLogger) error {
	var (
		supplier swclient.ResponseGetSupplierEnvelope
		resp     *http.Response
		err      error
	)
	supplier, resp, err = channelsClient.Client.SuppliersApi.GetSupplier(*ctx, supplierID)
	if err != nil {
		return err
	}

	if supplier.Supplier == nil {
		return ErrorSupplierIsNotFound
	}

	lgr.Logger().WithField("Supplier", "found").WithFields(logrus.Fields{"Supplier response": supplier, "response status": resp.StatusCode}).Info("Supplier was found")

	channelsClient.SetSupplierID(supplier.Supplier.Id)

	return nil
}

// CheckProductExist is used to check exist supplier by uuid
func CheckProductExist(supplierID, productID string, channelsClient *channels.ChannelsClient, ctx *context.Context, lgr *logger.LocalLogger) error {
	var (
		product swclient.ResponseGetProductEnvelope
		resp    *http.Response
		err     error
	)
	product, resp, err = channelsClient.Client.ProductsApi.GetProduct(*ctx, supplierID, productID)
	if err != nil {
		return err
	}

	if product.Product == nil || product.Product.Id == "" {
		lgr.Logger().WithField("Product", "not found").WithFields(logrus.Fields{"Product": product, "response status": resp.StatusCode}).Info("Product was not found")
		return ErrorProductIsNotFound
	}

	lgr.Logger().WithField("Product", "found").WithFields(logrus.Fields{"Product": product, "response status": resp.StatusCode}).Info("Product was found")

	channelsClient.SetProductID(product.Product.Id)

	return nil
}

// CheckRateExist is used to check exist supplier by uuid
func CheckRateExist(supplierID, productID, rateID string, channelsClient *channels.ChannelsClient, ctx *context.Context, lgr *logger.LocalLogger) error {
	var (
		rate swclient.ResponseGetRateEnvelope
		resp *http.Response
		err  error
	)
	rate, resp, err = channelsClient.Client.RatesApi.GetRate(*ctx, supplierID, productID, rateID)
	if err != nil {
		return err
	}

	if rate.Rate == nil || rate.Rate.Id == "" {
		return ErrorRateIsNotFound
	}

	lgr.Logger().WithField("Supplier", "found").WithFields(logrus.Fields{"Supplier rate": rate, "response status": resp.StatusCode}).Info("Rate was found")

	//channelsClient.SetSupplierID(supplier.Supplier.Id)

	return nil
}
