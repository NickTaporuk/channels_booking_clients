package channels

import "bitbucket.org/redeam/integration-channel/swclient"

func (ch *ChannelsClient) CreateChannelBinding() (*swclient.RequestPostCreateChannelEnvelope, error) {
	var (
		channelBinding swclient.RequestPostCreateChannelEnvelope
	)

	channelBinding = swclient.RequestPostCreateChannelEnvelope{}

	return &channelBinding, nil
}
