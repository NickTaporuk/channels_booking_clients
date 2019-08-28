package booking

import (
	"time"

	"bitbucket.org/redeam/integration-booking/swclient"
	"github.com/google/uuid"
)

func (b *BookingClient) CreateBooking(priceID, rateID, supplierID string) (*swclient.RequestPostBookingEnvelope, error) {
	var (
		booking   swclient.RequestPostBookingEnvelope
		data      = make(map[string]string)
		t         string
		startTime time.Time
		email     = "mykola.kuropatkin@redeam.com"
		firstName = "Nick"
		lastName  = "Kuropatkin"
		phone     = "073-408-9378"
	)

	t = time.Now().Format("2019-08-27T17:57:36Z")
	startTime = time.Now()

	data["createdAt"] = t
	data["updatedAt"] = t

	var ext interface{} = data

	booking = swclient.RequestPostBookingEnvelope{
		Meta: &swclient.RequestPostBookingEnvelopeMeta{
			ReqId: uuid.New().String(),
		},
		Booking: &swclient.RequestPostBookingEnvelopeBooking{
			Customer: &swclient.RequestPostBookingEnvelopeBookingCustomer{
				Email:     email,
				FirstName: firstName,
				LastName:  lastName,
				Phone:     phone,
			},
			Ext: &ext,
			Id:  uuid.New().String(),
			Items: []swclient.RequestPostBookingEnvelopeBookingItems{
				{
					AvailabilityId: uuid.New().String(),
					Ext:            &ext,
					PriceId:        priceID,
					Quantity:       1,
					RateId:         rateID,
					StartTime:      startTime,
					SupplierId:     supplierID,
					Status:         "STATUS_REDEEMED",
					Traveler: &swclient.RequestPostBookingEnvelopeBookingTraveler{
						Age:       32,
						Country:   "UKR",
						Email:     email,
						FirstName: firstName,
						LastName:  lastName,
						Phone:     phone,
						Type_:     "ADULT",
						Gender:    "Male",
						IsLead:    true,
						Lang:      "eng",
					},
				},
			},
			Status: "UNKNOWN",
		},
	}

	return &booking, nil
}
