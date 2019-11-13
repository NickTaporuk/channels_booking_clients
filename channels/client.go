package channels

import (
	"bitbucket.org/redeam/integration-channel/swclient"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/jedib0t/go-pretty/table"
	"os"
)

const (
	// HeaderKeyXAPIKey
	HeaderKeyXAPIKey = "X-API-Key"
	//
	HeaderKeyXAPISecret = "X-API-Secret"
)

type (
	ChannelsClient struct {
		Client     *swclient.APIClient
		productID  string
		supplierID string
		channelID  string
		rateIDs    []string
		priceIDs   []string
		logger     *logger.LocalLogger
	}
)

func (ch *ChannelsClient) ChannelID() string {
	return ch.channelID
}

func (ch *ChannelsClient) SetChannelID(channelID string) {
	ch.channelID = channelID
}

func (ch *ChannelsClient) Logger() *logger.LocalLogger {
	return ch.logger
}

func (ch *ChannelsClient) SetLogger(logger *logger.LocalLogger) {
	ch.logger = logger
}

func (ch *ChannelsClient) PriceIDs() []string {
	return ch.priceIDs
}

func (ch *ChannelsClient) SetPriceIDs(priceIDs []string) {
	ch.priceIDs = priceIDs
}

func (ch *ChannelsClient) SetRateIDs(rateIDs []string) {
	ch.rateIDs = rateIDs
}

func (ch *ChannelsClient) RateIDs() []string {
	return ch.rateIDs
}

func (ch *ChannelsClient) SetRateID(rateID string) {
	ch.rateIDs = append(ch.rateIDs, rateID)
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

func (ch ChannelsClient) String() {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Entity Name", "UUID",})
	t.AppendRows([]table.Row{
		{1, "supplier ID", ch.supplierID},
		{2, "product  ID", ch.productID},
		{3, "rate     IDs", ch.rateIDs},
		{4, "price    IDs", ch.priceIDs},
	})
	t.Render()

	//ch.logger.Logger().WithFields(logrus.Fields{"supplier ID":ch.supplierID,"product ID":ch.productID, "rate ID":ch.rateIDs, "price ID":ch.priceIDs}).Info("channel client data")
}

// NewChannelClient is constructor for channels api
func NewChannelClient(xAPIKey, xAPISecretKey string) (*ChannelsClient, error) {
	var (
		headers  = make(map[string]string)
		chClient *swclient.APIClient
	)

	headers[HeaderKeyXAPIKey] = xAPIKey
	headers[HeaderKeyXAPISecret] = xAPISecretKey

	cnf := swclient.NewConfiguration()
	cnf.DefaultHeader = headers

	chClient = swclient.NewAPIClient(cnf)

	return &ChannelsClient{
		Client: chClient,
	}, nil
}
