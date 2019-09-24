package utils

import (
	"context"
	"errors"
	"net/http"
	"os"

	bc "bitbucket.org/redeam/integration-booking/swclient"
	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/NickTaporuk/channels_booking_clients/booking"
	"github.com/NickTaporuk/channels_booking_clients/channels"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	ErrorSupplierIsNotFound       = errors.New("supplier isn't found")
	ErrorProductIsNotFound        = errors.New("product isn't found")
	ErrorRateIsNotFound           = errors.New("rate isn't found")
	ErrorChannelBindingIsNotFound = errors.New("channel binding isn't found")
	ErrorBookingIsNotFound        = errors.New("booking isn't found")
	ErrorHoldIsNotFound           = errors.New("hold isn't found")
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
		rate   swclient.ResponseGetRateEnvelope
		resp   *http.Response
		err    error
		prices []string
		rates  []string
	)

	rate, resp, err = channelsClient.Client.RatesApi.GetRate(*ctx, supplierID, productID, rateID)
	if err != nil {
		lgr.Logger().WithField("rate", "not found").WithFields(logrus.Fields{"rate": rate, "response status": resp.StatusCode, "error": err}).Error("rate was not found")
		return err
	}

	if rate.Rate == nil || rate.Rate.Id == "" {
		return ErrorRateIsNotFound
	}

	for _, price := range rate.Rate.Prices {
		prices = append(prices, price.Id)
	}

	rates = append(rates, rate.Rate.Id)

	channelsClient.SetRateIDs(rates)
	channelsClient.SetPriceIDs(prices)
	lgr.Logger().WithField("Rate", "found").WithFields(logrus.Fields{"Rate": rate, "response status": resp.StatusCode}).Info("Rate was found")

	return nil
}

// CheckRateExist is used to check exist supplier by uuid
func CheckChannelBindingExist(channelID string, channelsClient *channels.ChannelsClient, ctx *context.Context, lgr *logger.LocalLogger) error {
	var (
		channelBinding swclient.ResponseGetChannelEnvelope
		resp           *http.Response
		err            error
	)

	channelBinding, resp, err = channelsClient.Client.ChannelsApi.GetChannel(*ctx, channelID)
	if err != nil {
		lgr.Logger().WithField("ChannelBinding", "not found").WithFields(logrus.Fields{"channelBinding": channelBinding, "response status": resp.StatusCode, "error": err}).Error("channelBinding was not found")

		return err
	}

	if channelBinding.Channel == nil || channelBinding.Channel.Id == "" {
		err = ErrorChannelBindingIsNotFound
		lgr.Logger().WithField("ChannelBinding", "not found").WithFields(logrus.Fields{"channelBinding": channelBinding, "response status": resp.StatusCode, "error": err}).Error("channelBinding was not found")

		return err
	}

	lgr.Logger().WithField("ChannelBinding", "found").WithFields(logrus.Fields{"channelBinding": channelBinding, "response status": resp.StatusCode}).Debug("channelBinding was found")

	return nil
}

// CheckRateExist is used to check exist supplier by uuid
func CheckBookingExist(bookingID string, bookingClient *booking.BookingClient, ctx *context.Context, lgr *logger.LocalLogger) error {
	var (
		booking bc.ResponseGetBookingEnvelope
		resp    *http.Response
		err     error
	)

	booking, resp, err = bookingClient.Client().BookingsApi.GetBooking(*ctx, bookingID)
	if err != nil {
		lgr.Logger().WithField("Booking", "not found").WithFields(logrus.Fields{"booking": booking, "response status": resp.StatusCode, "error": err}).Error("Booking was not found")

		return err
	}

	if booking.Booking == nil || booking.Booking.Id == "" {
		err = ErrorBookingIsNotFound
		lgr.Logger().WithField("Booking", "not found").WithFields(logrus.Fields{"booking": booking, "response status": resp.StatusCode, "error": err}).Error("booking was not found")

		return err
	}

	lgr.Logger().WithField("Booking", "found").WithFields(logrus.Fields{"booking": booking, "response status": resp.StatusCode}).Debug("booking was found")

	return nil
}

// CheckHoldExist is used to check exist supplier by uuid
func CheckHoldExist(holdID string, bookingClient *booking.BookingClient, ctx *context.Context, lgr *logger.LocalLogger) error {
	var (
		hold bc.ResponseGetHoldEnvelope
		resp *http.Response
		err  error
	)

	hold, resp, err = bookingClient.Client().HoldsApi.GetHold(*ctx, holdID)
	if err != nil {
		lgr.Logger().WithField("Hold", "not found").WithFields(logrus.Fields{"hold": hold, "response status": resp.StatusCode, "error": err}).Error("hold was not found")

		return err
	}

	if hold.Hold == nil || hold.Hold.Id == "" {
		err = ErrorHoldIsNotFound
		lgr.Logger().WithField("Hold", "not found").WithFields(logrus.Fields{"hold": hold, "response status": resp.StatusCode, "error": err}).Error("hold was not found")

		return err
	}

	lgr.Logger().WithField("Hold", "found").WithFields(logrus.Fields{"hold": hold, "response status": resp.StatusCode}).Debug("hold was found")

	return nil
}
