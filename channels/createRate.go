package channels

import (
	"fmt"
	"time"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

func (ch *ChannelsClient) CreateRate() (*swclient.RequestPostRateEnvelope, error) {
	var (
		rate     swclient.RequestPostRateEnvelope
		rateName string
	)

	rateName = fmt.Sprintf(`Test-%s`, ch.ProductID())
	t := time.Now()
	tClose := time.Now().AddDate(1, 1, 1)

	rate = swclient.RequestPostRateEnvelope{
		Rate: &swclient.RequestPostRateEnvelopeRate{
			Cancelable:     true,
			Code:           rateName,
			Cutoff:         1,
			Holdable:       true,
			HoldablePeriod: 864000,
			Hours: []swclient.RequestPostProductEnvelopeProductHours{
				{
					DaysOfWeek: []int32{
						1, 2, 3, 4, 5, 6, 7,
					},
					Times: []swclient.RequestPostProductEnvelopeProductTimes{
						{
							Close: "17:59",
							Open:  "17:58",
						},
					},
					Valid: &swclient.RequestPostProductEnvelopeProductValid{
						From:  t,
						Until: tClose,
					},
				},
			},
			Id:           uuid.New().String(),
			MaxTravelers: 3,
			MinTravelers: 1,
			Name:         rateName,
			Prices: []swclient.RequestPostRateEnvelopeRatePrices{
				ch.GenPrice(),
				ch.GenPrice(),
			},
			ProductId: ch.ProductID(),
			Title:     gofakeit.Quote(),
			Valid: &swclient.RequestPostRateEnvelopeRateValid{
				From:  t,
				Until: tClose,
			},
			Type_: "RESERVED",
		},
		Meta: &swclient.RequestPostCreateChannelEnvelopeMeta{
			ReqId: uuid.New().String(),
		},
	}

	return &rate, nil
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
