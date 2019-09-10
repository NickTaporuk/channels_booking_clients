package booking

import (
	"context"
	"encoding/json"

	"bitbucket.org/redeam/integration-booking/swclient"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (b *BookingClient) CreateBooking(priceID, rateID, supplierID string, data *[]byte, ctx *context.Context) error {
	var (
		err     error
		booking swclient.RequestPostBookingEnvelope
	)

	if err = json.Unmarshal(*data, &booking); err != nil {
		return err
	}

	booking.Meta = &swclient.RequestPostBookingEnvelopeMeta{ReqId: uuid.New().String()}
	booking.Booking.Id = uuid.New().String()
	b.UpdateBookingItems(priceID, rateID, supplierID, booking.Booking.Items)

	b.logger.Logger().WithFields(logrus.Fields{"file data": string(*data),}).Debug(" data from supplier json file")
	b.logger.Logger().WithFields(logrus.Fields{"Supplier": booking,}).Debug(booking)

	ResponsePostBookingEnvelope, resp, err := b.Client().BookingsApi.CreateBooking(*ctx, booking)

	b.logger.Logger().WithFields(logrus.Fields{"ResponsePostBookingEnvelope": ResponsePostBookingEnvelope, "create Booking response resp statusCode": resp.StatusCode, "err": err}).Debug("ResponsePostBookingEnvelope")

	if err != nil {
		b.logger.Logger().WithFields(logrus.Fields{"ResponsePostBookingEnvelope": ResponsePostBookingEnvelope, "booking response resp statusCode": resp.StatusCode, "create booking body": resp.Body, "err": err}).Error("Booking api create booking error")
		return err
	}

	return nil
}

func (b *BookingClient) UpdateBookingItems(priceID, rateID, supplierID string, items []swclient.RequestPostBookingEnvelopeBookingItems) {
	var (
		item swclient.RequestPostBookingEnvelopeBookingItems
	)

	for _, item = range items {
		item.PriceId = priceID
		item.SupplierId = supplierID
		item.RateId = rateID
		item.AvailabilityId = uuid.New().String()
	}
}
