package main

import (
	"context"
	"fmt"
	"net/http"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/NickTaporuk/channels_booking_clients/channels"
)

func main() {
	var (
		channelsApiHeaders = make(map[string]string)
		channelsClient     *channels.ChannelsClient
		ctx                = context.Background()
		getSupplier        swclient.ResponseGetSupplierEnvelope
		resp               *http.Response
		err                error
	)
	//
	channelsClient, err = channels.NewChannelClient(channelsApiHeaders)

	if err != nil {
		panic(err)
	}

	getSupplier, resp, err = channelsClient.Client.SuppliersApi.GetSupplier(ctx, channels.SupplierID)

	if getSupplier.Supplier == nil {
		panic("Suplier isn't found")
	}

	channelsClient.SetSupplierID(getSupplier.Supplier.Id)

	product, err := channelsClient.CreateProduct()

	respProd, resp, err := channelsClient.Client.ProductsApi.CreateProduct(ctx, channels.SupplierID, *product)
	if err != nil {
		panic(err)
	}

	channelsClient.SetProductID(respProd.Product.Id)

	if respProd.Product == nil {
		panic("product isn't found")
	}

	fmt.Println(respProd, resp, err)

	rate, _ := channelsClient.CreateRate()

	respRate, resp, err := channelsClient.Client.RatesApi.CreateRate(ctx, channels.SupplierID, channelsClient.ProductID(), *rate)
	fmt.Println(respRate, resp, err)
	if err != nil {
		panic(err)
	}

	for _, price := range respRate.Rate.Prices {

	}

	fmt.Println(respRate, resp, err)
}
