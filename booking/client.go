package booking

import (
	"bitbucket.org/redeam/integration-booking/swclient"
	"github.com/sirupsen/logrus"
)

const (
	XAPIKey    = "key-redeam-qa-4f9c3z68"
	XAPISecret = "secret-redeam-qa-v8z73c9x"
	SupplierID = "b0105adc-3693-4e42-8905-cbbd97f80bbb"
)

type BookingClient struct {
	client *swclient.APIClient
	logger *logrus.Logger
}

func (b *BookingClient) Client() *swclient.APIClient {
	return b.client
}

func (b *BookingClient) SetClient(client *swclient.APIClient) {
	b.client = client
}

func (b *BookingClient) Logger() *logrus.Logger {
	return b.logger
}

func (b *BookingClient) SetLogger(logger *logrus.Logger) {
	b.logger = logger
}

func NewBookingClient(headers map[string]string) (*BookingClient, error) {

	var (
		client *swclient.APIClient
	)

	headers["X-API-Key"] = XAPIKey
	headers["X-API-Secret"] = XAPISecret

	cnf := swclient.NewConfiguration()
	cnf.DefaultHeader = headers

	client = swclient.NewAPIClient(cnf)

	return &BookingClient{client: client}, nil
}
