package channels

import (
	"bitbucket.org/redeam/integration-channel/swclient"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (ch *ChannelsClient) CreateChannelBinding(data *[]byte, ctx *context.Context) error {
	var (
		err            error
		channelBinding = new(swclient.RequestPostCreateChannelEnvelope)
	)

	if err = json.Unmarshal([]byte(*data), &channelBinding); err != nil {
		return err
	}

	channelBinding.Meta = &swclient.RequestPostCreateChannelEnvelopeMeta{ReqId: uuid.New().String()}
	channelBinding.RateIds = ch.RateIDs()

	ch.logger.Logger().WithFields(logrus.Fields{"file data": string(*data),}).Debug(" data from product json file")
	ch.logger.Logger().WithFields(logrus.Fields{"channelBinding": channelBinding,}).Debug("channelBinding debug")

	ResponsePostChannelBindingEnvelope, resp, err := ch.Client.ChannelsApi.CreateChannelBinding(*ctx, ch.ChannelID(), ch.productID, ch.supplierID, *channelBinding)

	ch.logger.Logger().WithFields(logrus.Fields{"ResponsePostChannelBindingEnvelope": ResponsePostChannelBindingEnvelope, "create channelBinding response resp statusCode": resp.StatusCode, "err": err}).Debug("response post channelBinding envelope")

	if err != nil {
		ch.logger.Logger().WithFields(logrus.Fields{"ResponsePostChannelBindingEnvelope": ResponsePostChannelBindingEnvelope, "response resp statusCode": resp.StatusCode, "create channelBinding body": resp.Body, "err": err}).Error("Channel api create channelBinding error")
		return err
	}

	return nil
}
