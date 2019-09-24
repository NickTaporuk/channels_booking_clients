package booking

import (
	"context"
	"encoding/json"
	"net/http"

	"bitbucket.org/redeam/integration-booking/swclient"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (b *BookingClient) CreateBooking(priceID, rateID, supplierID string, data *[]byte, ctx *context.Context) error {
	var (
		err     error
		booking = new(swclient.RequestPostBookingEnvelope)
	)

	if err = json.Unmarshal(*data, &booking); err != nil {
		return err
	}

	booking.Meta = &swclient.RequestPostBookingEnvelopeMeta{ReqId: uuid.New().String()}
	booking.Booking.Id = uuid.New().String()

	for i, _ := range booking.Booking.Items {
		booking.Booking.Items[i].PriceId = priceID
		booking.Booking.Items[i].SupplierId = supplierID
		booking.Booking.Items[i].RateId = rateID
		booking.Booking.Items[i].AvailabilityId = uuid.New().String()
		booking.Booking.Items[i].Id = uuid.New().String()
	}

	b.logger.Logger().WithFields(logrus.Fields{"file data": string(*data),}).Debug(" data from supplier json file")
	b.logger.Logger().WithFields(logrus.Fields{"Booking": booking,}).Debug(booking)

	ResponsePostBookingEnvelope, resp, err := b.Client().BookingsApi.CreateBooking(*ctx, *booking)

    var statusCode	int
	if resp == nil {
		statusCode = http.StatusBadGateway
	} else {
		statusCode = resp.StatusCode
	}


	if err != nil {
		b.logger.Logger().WithFields(logrus.Fields{"ResponsePostBookingEnvelope": ResponsePostBookingEnvelope, "booking response resp statusCode": statusCode, "create booking body": resp.Body, "err": err.(swclient.GenericSwaggerError).Model()}).Error("Booking api create booking error")
		return err
	} else {
		b.logger.Logger().WithFields(logrus.Fields{"ResponsePostBookingEnvelope": ResponsePostBookingEnvelope, "create Booking response resp statusCode":  statusCode, "err": err}).Debug("ResponsePostBookingEnvelope")
	}

	return nil
}
