package booking

import (
	"bitbucket.org/redeam/integration-booking/swclient"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	// RFC3339Millis represents a ISO8601 format to millis instead of to nanos
	RFC3339Millis = "2006-01-02T15:04:05.000Z07:00"
)

// CreateHold is create new hold
func (b *BookingClient) CreateHold(rateID, supplierID, priceID string, data *[]byte, ctx *context.Context) error {

	var (
		err  error
		hold = new(swclient.RequestPostHoldEnvelope)
	)

	if err = json.Unmarshal(*data, &hold); err != nil {
		return err
	}

	hold.Meta = &swclient.RequestPostBookingEnvelopeMeta{ReqId: uuid.New().String()}
	hold.Hold.Id = uuid.New().String()

	for i, _ := range hold.Hold.Items {
		hold.Hold.Items[i].AvailabilityId = uuid.New().String()
		hold.Hold.Items[i].RateId = rateID
		hold.Hold.Items[i].SupplierId = supplierID
		hold.Hold.Items[i].PriceId = priceID
	}

	ResponsePostHoldEnvelope, resp, err := b.Client().HoldsApi.CreateHold(*ctx, *hold)

	b.logger.Logger().WithFields(logrus.Fields{"ResponsePostHoldEnvelope": ResponsePostHoldEnvelope, "create Hold response resp statusCode": resp.StatusCode, "err": err}).Debug("ResponsePostHoldEnvelope")

	if err != nil {
		b.logger.Logger().WithFields(logrus.Fields{"ResponsePostHoldEnvelope": ResponsePostHoldEnvelope, "hold response resp statusCode": resp.StatusCode, "create hold body": resp.Body, "err": err}).Error("Hold api create hold error")
		return err
	}

	return nil
}
