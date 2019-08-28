package booking

import (
	"time"

	"bitbucket.org/redeam/integration-booking/swclient"
	"github.com/google/uuid"
)

func (b *BookingClient) CreateHold(rateID, supplierID, priceID string) (swclient.RequestPostHoldEnvelope, error) {

	var (
		hold             swclient.RequestPostHoldEnvelope
		additionalFields = make(map[string]string)
	)

	expires := time.Now().AddDate(0, 0, 1)
	at := time.Now()

	additionalFields["createdAt"] = at.Format(time.RFC3339)
	additionalFields["updatedAt"] = at.Format(time.RFC3339)

	var ext interface{} = additionalFields

	hold = swclient.RequestPostHoldEnvelope{
		Meta: &swclient.RequestPostBookingEnvelopeMeta{
			ReqId: uuid.New().String(),
		},
		Hold: &swclient.RequestPostHoldEnvelopeHold{
			Expires: expires,
			Id:      uuid.New().String(),
			Items: []swclient.RequestPostHoldEnvelopeHoldItems{
				{
					At:             at,
					AvailabilityId: uuid.New().String(),
					Ext:            &ext,
					Quantity:       1,
					RateId:         rateID,
					SupplierId:     supplierID,
					PriceId:        priceID,
					TravelerType:   "ADULT",
				},
			},
		},
	}

	return hold, nil
}
