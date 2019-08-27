package channels

import (
	"bitbucket.org/redeam/integration-channel/swclient"
)

const (
	XAPIKey    = "key-redeam-qa-4f9c3z68"
	XAPISecret = "secret-redeam-qa-v8z73c9x"
	SupplierID = "b0105adc-3693-4e42-8905-cbbd97f80bbb"
)

type (
	ChannelsClient struct {
		Client     *swclient.APIClient
		productID  string
		supplierID string
		rateIDs []string
	}
)

func (ch *ChannelsClient) SetRateIDs(rateIDs []string) {
	ch.rateIDs = rateIDs
}

func (ch *ChannelsClient) RateIDs() []string {
	return ch.rateIDs
}

func (ch *ChannelsClient) SetRateID(rateID string) {
	ch.rateIDs = append(ch.rateIDs,rateID)
}

func (ch *ChannelsClient) ProductID() string {
	return ch.productID
}

func (ch *ChannelsClient) SetProductID(productID string) {
	ch.productID = productID
}

func (ch *ChannelsClient) SupplierID() string {
	return ch.supplierID
}

func (ch *ChannelsClient) SetSupplierID(supplierID string) {
	ch.supplierID = supplierID
}

// NewChannelClient is constructor for channels api
func NewChannelClient(headers map[string]string) (*ChannelsClient, error) {
	var (
		chClient *swclient.APIClient
	)

	headers["X-API-Key"] = XAPIKey
	headers["X-API-Secret"] = XAPISecret

	cnf := swclient.NewConfiguration()
	cnf.DefaultHeader = headers

	chClient = swclient.NewAPIClient(cnf)

	return &ChannelsClient{
		Client: chClient,
	}, nil
}
