package booking

import (
	"bitbucket.org/redeam/integration-booking/swclient"
)

const (
	XAPIKey    = "key-redeam-qa-4f9c3z68"
	XAPISecret = "secret-redeam-qa-v8z73c9x"
	SupplierID = "b0105adc-3693-4e42-8905-cbbd97f80bbb"
)

type BookingClient struct {
	Client     *swclient.APIClient
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

	return &BookingClient{Client: client}, nil
}
