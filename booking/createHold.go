package booking

import (
	"context"
	"encoding/json"
	"time"

	"bitbucket.org/redeam/integration-booking/swclient"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	// RedeamTimeFormat is local format of redeam api
	RedeamTimeFormat = "2006-01-02T15:04:05Z"
	// AdditionalFieldCreatedAtName is additional field for booking api endpoint /holds of field ext
	AdditionalFieldCreatedAtName = "createdAt"
	// AdditionalFieldUpdatedAtName is additional field for booking api endpoint /holds of field ext
	AdditionalFieldUpdatedAtName = "updatedAt"
)

// CreateHold is create new hold
func (b *BookingClient) CreateHold(rateID, supplierID, priceID string, data *[]byte, ctx *context.Context) error {
	t := time.Now().AddDate(0, 0, 1)

	var hold = new(swclient.RequestPostHoldEnvelope)

	if err := json.Unmarshal(*data, &hold); err != nil {
		return err
	}

	hold.Meta = &swclient.RequestPostBookingEnvelopeMeta{ReqId: uuid.New().String()}
	hold.Hold.Id = uuid.New().String()

	if hold.Hold.Expires.IsZero() {
		hold.Hold.Expires = t
	}

	if hold.Hold.Retrieved.IsZero() {
		hold.Hold.Retrieved = t
	}

	if hold.Hold.Ext == nil {
		var (
			additionalFields = make(map[string]interface{})
			at               = time.Now().UTC()
		)

		additionalFields[AdditionalFieldCreatedAtName] = at.Format(strfmt.RFC3339Millis)
		additionalFields[AdditionalFieldUpdatedAtName] = at.Format(strfmt.RFC3339Millis)

		var ext interface{} = additionalFields

		hold.Hold.Ext = &ext
	}

	for i, _ := range hold.Hold.Items {
		hold.Hold.Items[i].AvailabilityId = uuid.New().String()
		hold.Hold.Items[i].RateId = rateID
		hold.Hold.Items[i].SupplierId = supplierID
		hold.Hold.Items[i].PriceId = priceID
		if hold.Hold.Items[i].At.IsZero() {
			hold.Hold.Items[i].At = t
		}
	}

	ResponsePostHoldEnvelope, resp, err := b.Client().HoldsApi.CreateHold(*ctx, *hold)

	b.logger.Logger().WithFields(logrus.Fields{"ResponsePostHoldEnvelope": ResponsePostHoldEnvelope, "create Hold response resp statusCode": resp.StatusCode, "err": err, "hold data": hold}).Debug("ResponsePostHoldEnvelope")

	if err != nil {
		b.logger.Logger().WithFields(logrus.Fields{"ResponsePostHoldEnvelope": ResponsePostHoldEnvelope, "hold response resp statusCode": resp.StatusCode, "create hold body": resp.Body, "err": err.(swclient.GenericSwaggerError).Model()}).Error("Hold api create hold error")
		return err
	}

	b.SetHoldID(ResponsePostHoldEnvelope.Hold.Id)

	return nil
}
