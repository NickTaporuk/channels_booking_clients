package channels

import (
	"context"
	"encoding/json"

	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"syreclabs.com/go/faker"
)

// CreateProduct is used to create new product of channel api
func (ch *ChannelsClient) CreateProduct(data *[]byte, ctx *context.Context) error {
	var (
		product = new(swclient.RequestPostProductEnvelope)
		err     error
	)

	if err = json.Unmarshal(*data, &product); err != nil {
		return err
	}

	product.Meta = &swclient.RequestPostCreateChannelEnvelopeMeta{ReqId: uuid.New().String()}

	product.Product.Code = faker.App().Author() + faker.App().Name()
	product.Product.Id = uuid.New().String()
	product.Product.Description = faker.App().Author() + faker.App().Name()
	product.Product.Title = faker.App().Author() + faker.App().Name()
	product.Product.Name = faker.App().Author() + faker.App().Name()

	ch.logger.Logger().WithFields(logrus.Fields{"file data": string(*data),}).Debug(" data from product json file")
	ch.logger.Logger().WithFields(logrus.Fields{"Product": product,}).Debug("product debug")

	ResponsePostProductEnvelope, resp, err := ch.Client.ProductsApi.CreateProduct(*ctx, ch.supplierID, *product)

	ch.logger.Logger().WithFields(logrus.Fields{"ResponsePostProductEnvelope": ResponsePostProductEnvelope, "create product response resp statusCode": resp.StatusCode, "err": err}).Debug("response post product envelope")

	if err != nil {
		ch.logger.Logger().WithFields(logrus.Fields{"ResponsePostProductEnvelope": ResponsePostProductEnvelope, "response resp statusCode": resp.StatusCode, "create product body": resp.Body, "err": err}).Error("Channel api create product error")
		return err
	}
	ch.SetProductID(product.Product.Id)

	return nil
}
