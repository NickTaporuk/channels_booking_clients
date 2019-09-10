package booking

import (
	"bitbucket.org/redeam/integration-booking/swclient"
	"github.com/NickTaporuk/channels_booking_clients/logger"
)

const (
	// HeaderKeyXAPIKey
	HeaderKeyXAPIKey = "X-API-Key"
	//
	HeaderKeyXAPISecret = "X-API-Secret"
)

type BookingClient struct {
	client *swclient.APIClient
	logger *logger.LocalLogger
}

func (b *BookingClient) Logger() *logger.LocalLogger {
	return b.logger
}

func (b *BookingClient) SetLogger(logger *logger.LocalLogger) {
	b.logger = logger
}

func (b *BookingClient) Client() *swclient.APIClient {
	return b.client
}

func (b *BookingClient) SetClient(client *swclient.APIClient) {
	b.client = client
}

// NewBookingClient
func NewBookingClient(xAPIKey, xAPISecret string) (*BookingClient, error) {

	var (
		headers = make(map[string]string)
		client  *swclient.APIClient
	)

	headers[HeaderKeyXAPIKey] = xAPIKey
	headers[HeaderKeyXAPISecret] = xAPISecret

	cnf := swclient.NewConfiguration()
	cnf.DefaultHeader = headers

	client = swclient.NewAPIClient(cnf)

	return &BookingClient{client: client}, nil
}
