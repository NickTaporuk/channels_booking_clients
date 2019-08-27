package channels

import (
	"time"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/google/uuid"
)

// CreateProduct is used to create new product of channel api
func (ch *ChannelsClient) CreateProduct() (*swclient.RequestPostProductEnvelope, error) {
	var (
		product = swclient.RequestPostProductEnvelope{}
	)

	t := time.Now()
	tClose := time.Now().AddDate(1, 1, 1) /*.Format("2019-08-02T19:34:49Z")*/

	product = swclient.RequestPostProductEnvelope{
		Product: &swclient.RequestPostProductEnvelopeProduct{
			Code: uuid.New().String(),
			Hours: []swclient.RequestPostProductEnvelopeProductHours{
				{
					DaysOfWeek: []int32{1, 2, 3, 4, 5, 6, 7},
					Times: []swclient.RequestPostProductEnvelopeProductTimes{
						{
							Close: "12:30",
							Open:  "08:00",
						},
					},
					Valid: &swclient.RequestPostProductEnvelopeProductValid{
						From:  t,
						Until: tClose,
					},
				},
			},
			Location: &swclient.RequestPostProductEnvelopeProductLocation{
				Address: &swclient.RequestPostProductEnvelopeProductLocationAddress{
					CountryCode:   "USA",
					Locality:      "City",
					PostalCode:    "postalCode",
					Region:        "Region",
					StreetAddress: "Street 1",
				},
				LongLat: &swclient.RequestPostProductEnvelopeProductLocationLongLat{
					Latitude:  1.00,
					Longitude: 1.00,
				},
				Name:      "Location A",
				Notes:     "",
				UtcOffset: "+01:00",
			},
			Name:   "Nick Kuropatkin Test product 2",
			Status: "UNKNOWN",
			Title:  "Nick Kuropatkin Test product 2",
		},
		Meta: &swclient.RequestPostCreateChannelEnvelopeMeta{
			ReqId: uuid.New().String(),
		},
	}

	return &product, nil
}
