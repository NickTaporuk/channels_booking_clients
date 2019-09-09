package main

import (
	"context"

	"net/http"

	bookingCl "bitbucket.org/redeam/integration-booking/swclient"
	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/NickTaporuk/channels_booking_clients/booking"
	"github.com/NickTaporuk/channels_booking_clients/channels"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/sirupsen/logrus"
	_ "github.com/spf13/viper"
)

func Run() {
	var (
		channelsApiHeaders = make(map[string]string)
		channelsClient     *channels.ChannelsClient
		bookingClient      *booking.BookingClient
		ctx                = context.Background()
		getSupplier        swclient.ResponseGetSupplierEnvelope
		channelBinding     *swclient.RequestPostCreateChannelEnvelope
		respChan           swclient.ResponsePostChannelRatesEnvelope
		respHold           bookingCl.ResponsePostHoldEnvelope
		resp               *http.Response
		err                error
		rates              []string
		prices             []string
		data               = make(map[string]string)
		lgr                *logger.LocalLogger
		book               *bookingCl.RequestPostBookingEnvelope
		respBooking        bookingCl.ResponsePostBookingEnvelope
	)

	data["level"] = "debug"
	lgr = &logger.LocalLogger{}
	defer lgr.Close()

	err = lgr.Init(data)

	if err != nil {
		panic(err)
	}
	//
	channelsClient, err = channels.NewChannelClient(channelsApiHeaders)

	if err != nil {
		panic(err)
	}

	getSupplier, resp, err = channelsClient.Client.SuppliersApi.GetSupplier(ctx, channels.SupplierID)

	if getSupplier.Supplier == nil {
		panic("Supplier isn't found")
	}

	lgr.Logger().WithField("Supplier", "done").WithFields(logrus.Fields{"Supplier response": getSupplier, "response status": resp.StatusCode}).Info("Supplier creation was done")

	channelsClient.SetSupplierID(getSupplier.Supplier.Id)
	lgr.Logger().Println("<==============================================================================================================================================================>")

	product, err := channelsClient.CreateProduct()

	respProd, resp, err := channelsClient.Client.ProductsApi.CreateProduct(ctx, channels.SupplierID, *product)
	if err != nil {
		panic(err)
	}

	channelsClient.SetProductID(respProd.Product.Id)

	if respProd.Product == nil {
		lgr.Logger().WithFields(logrus.Fields{"request": product}).Fatal(err)

	}

	lgr.Logger().WithField("Product", "done").WithFields(logrus.Fields{"Product response": respProd}).Info("Product creation was done")
	lgr.Logger().Println("<==============================================================================================================================================================>")

	rate, _ := channelsClient.CreateRate()

	respRate, resp, err := channelsClient.Client.RatesApi.CreateRate(ctx, channels.SupplierID, channelsClient.ProductID(), *rate)
	if err != nil {
		lgr.Logger().WithFields(logrus.Fields{"request": rate}).Fatal(err)
	}
	lgr.Logger().WithField("Rate", "done").WithFields(logrus.Fields{"Rate response": respRate}).Info("Rate creation was done")
	lgr.Logger().Println("<==============================================================================================================================================================>")

	if respRate.Rate.Id == "" {
		panic("Rate id is empty")
	}

	for _, price := range respRate.Rate.Prices {
		prices = append(prices, price.Id)
	}

	rates = append(rates, respRate.Rate.Id)

	channelsClient.SetRateIDs(rates)
	channelsClient.SetPriceIDs(prices)

	channelBinding, err = channelsClient.CreateChannelBinding()

	respChan, resp, err = channelsClient.Client.ChannelsApi.CreateChannelBinding(ctx, channels.ChannelId, channelsClient.ProductID(), channelsClient.SupplierID(), *channelBinding)
	lgr.Logger().WithField("Channel binding", "done").WithFields(logrus.Fields{"channel response": respChan}).Info("channel binding creation was done")
	lgr.Logger().Println("<==============================================================================================================================================================>")

	bookingClient, err = booking.NewBookingClient(channelsApiHeaders)
	if err != nil {
		lgr.Logger().WithFields(logrus.Fields{"request": channelBinding}).Fatal("Booking client is not initialized")
	}

	priceID := channelsClient.PriceIDs()[0]
	rateID := channelsClient.RateIDs()[0]

	book, err = bookingClient.CreateBooking(priceID, rateID, channels.SupplierID)
	if err != nil {
		lgr.Logger().Fatal(err)
	}

	respBooking, resp, err = bookingClient.Client().BookingsApi.CreateBooking(ctx, *book)
	if err != nil {
		lgr.Logger().WithFields(logrus.Fields{"booking response": resp.Body, "err": err}).Fatal(err)
	}
	lgr.Logger().WithField("Booking", "done").WithFields(logrus.Fields{"Booking response": respBooking}).Info("Booking creation was done")
	lgr.Logger().Println("<==============================================================================================================================================================>")

	hold, err := bookingClient.CreateHold(rateID, channels.SupplierID, priceID);
	respHold, resp, err = bookingClient.Client().HoldsApi.CreateHold(ctx, hold)
	if err != nil {
		lgr.Logger().WithFields(logrus.Fields{"request": hold, "resp body": resp.Body}).Fatal(err)
	}
	lgr.Logger().WithField("Hold", "done").WithFields(logrus.Fields{"Hold response": respHold, "request": hold}).Info("Hold creation was done")
	lgr.Logger().Println("<==============================================================================================================================================================>")
}

//func main() {
//	Run()
//}
