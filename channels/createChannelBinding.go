package channels

import (
	"time"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/google/uuid"
)

func (ch *ChannelsClient) CreateChannelBinding() (*swclient.RequestPostCreateChannelEnvelope, error) {
	var (
		channelBinding swclient.RequestPostCreateChannelEnvelope
	)

	t := time.Now()
	tClose := time.Now().AddDate(1, 1, 1)

	channelBinding = swclient.RequestPostCreateChannelEnvelope{
		Meta: &swclient.RequestPostCreateChannelEnvelopeMeta{
			ReqId: uuid.New().String(),
		},
		PriceAt: "PRICE_AT_SALE",
		RateIds: ch.RateIDs(),
		Valid: &swclient.RequestPostCreateChannelEnvelopeValid{
			From:  t,
			Until: tClose,
		},
		PriceTags:[]string{"A", "B", "C"},
	}

	return &channelBinding, nil
}
