package channels

import (
	"context"
	"encoding/json"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"syreclabs.com/go/faker"
)

func (ch *ChannelsClient) CreateSupplier(data *[]byte) error {
	var (
		supplier = new(swclient.RequestPostSupplierEnvelope)
		ctx      = context.Background()
		err      error
	)

	if err = json.Unmarshal([]byte(*data), &supplier); err != nil {
		return err
	}

	supplier.Meta.ReqId = uuid.New().String()
	supplier.Supplier.Id = uuid.New().String()
	supplier.Supplier.Code = faker.App().Author() + faker.App().Name()
	supplier.Supplier.Name = faker.App().Author() + faker.App().Name()

	ch.logger.Logger().WithFields(logrus.Fields{"file data": string(*data),}).Debug(" data from supplier json file")
	ch.logger.Logger().WithFields(logrus.Fields{"Supplier": supplier,}).Debug(supplier)

	ResponsePostSupplierEnvelope, resp, err := ch.Client.SuppliersApi.CreateSupplier(ctx, *supplier)

	ch.logger.Logger().WithFields(logrus.Fields{"ResponsePostSupplierEnvelope": ResponsePostSupplierEnvelope, "create supplier response resp statusCode": resp.StatusCode, "err": err}).Debug("ResponsePostSupplierEnvelope")

	if err != nil {
		ch.logger.Logger().WithFields(logrus.Fields{"ResponsePostSupplierEnvelope": ResponsePostSupplierEnvelope, "response resp statusCode": resp.StatusCode, "create supplier body": resp.Body, "err": err}).Error("Channel api create supplier error")
		return err
	}

	return nil
}
