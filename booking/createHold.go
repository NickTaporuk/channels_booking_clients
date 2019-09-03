package booking

import(
	"time"

	"bitbucket.org/redeam/integration-booking/swclient"
	"github.com/google/uuid"
)

const (
	// RFC3339Millis represents a ISO8601 format to millis instead of to nanos
	RFC3339Millis = "2006-01-02T15:04:05.000Z07:00"
)

func (b *BookingClient) CreateHold(rateID, supplierID, priceID string) (swclient.RequestPostHoldEnvelope, error) {

	var (
		hold             swclient.RequestPostHoldEnvelope
		additionalFields = make(map[string]string)
	)

	expires := time.Now().AddDate(0, 0, 1)
	at := time.Now().UTC()

	additionalFields["createdAt"] = at.Format(RFC3339Millis)
	additionalFields["updatedAt"] = at.Format(RFC3339Millis)

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
