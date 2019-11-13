package channels

import (
	"context"
	"encoding/json"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"syreclabs.com/go/faker"
)

func (ch *ChannelsClient) CreateRate(data *[]byte, ctx *context.Context) error {
	var (
		err    error
		rate   swclient.RequestPostRateEnvelope
		prices []string
		rates  []string
	)

	if err = json.Unmarshal(*data, &rate); err != nil {
		return err
	}

	rate.Meta = &swclient.RequestPostCreateChannelEnvelopeMeta{ReqId: uuid.New().String()}

	rate.Rate.Code = faker.App().Author() + faker.App().Name()
	rate.Rate.Id = uuid.New().String()
	rate.Rate.ProductId = ch.productID
	rate.Rate.Name = faker.App().Author() + faker.App().Name()

	for i, _ := range rate.Rate.Prices {
		rate.Rate.Prices[i].Id = uuid.New().String()
		rate.Rate.Prices[i].Name = gofakeit.Sentence(20)
	}

	//rate.Rate.Prices = []swclient.RequestPostRateEnvelopeRatePrices{ch.GenPrice()}

	ch.logger.Logger().WithFields(logrus.Fields{"file data": string(*data),}).Debug(" data from rate json file")
	ch.logger.Logger().WithFields(logrus.Fields{"Rate": rate,}).Debug("rate debug")

	ResponsePostRateEnvelope, resp, err := ch.Client.RatesApi.CreateRate(*ctx, ch.supplierID, ch.productID, rate)

	ch.logger.Logger().WithFields(logrus.Fields{"ResponsePostRateEnvelope": ResponsePostRateEnvelope, "create rate response resp statusCode": resp.StatusCode, "err": err}).Debug("response post rate envelope")

	if err != nil {
		ch.logger.Logger().WithFields(logrus.Fields{"ResponsePostRateEnvelope": ResponsePostRateEnvelope, "response resp statusCode": resp.StatusCode, "create rate body": resp.Body, "err": err}).Error("Channel api create rate error")
		return err
	}

	for _, price := range rate.Rate.Prices {
		prices = append(prices, price.Id)
	}

	rates = append(rates, rate.Rate.Id)

	ch.SetRateIDs(rates)
	ch.SetPriceIDs(prices)

	return nil
}

func (ch *ChannelsClient) GenPrice() swclient.RequestPostRateEnvelopeRatePrices {
	return swclient.RequestPostRateEnvelopeRatePrices{
		Id: uuid.New().String(),
		Labels: []string{
			"A",
			"B",
			"C",
		},
		Name: gofakeit.Sentence(20),
		Net: &swclient.RequestPostRateEnvelopeRateNet{
			Amount:   int64(gofakeit.Number(5, 100)),
			Currency: gofakeit.Currency().Short,
		},
		Retail: &swclient.RequestPostRateEnvelopeRateRetail{
			Amount:   int64(gofakeit.Number(5, 10)),
			Currency: gofakeit.Currency().Short,
		},
		TravelerType: &swclient.RequestPostRateEnvelopeRateTravelerType{
			AgeBand: "ADULT",
			MaxAge:  int32(gofakeit.Number(25, 100)),
			MinAge:  int32(gofakeit.Number(21, 24)),
			Name:    gofakeit.Person().FirstName,
		},
	}
}
