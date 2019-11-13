package booking

import (
	"bitbucket.org/redeam/integration-booking/swclient"
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

type BookingClient struct {
	client    *swclient.APIClient
	logger    *logger.LocalLogger
	holdID    string
	bookingID string
}

func (b *BookingClient) BookingID() string {
	return b.bookingID
}

func (b *BookingClient) SetBookingID(bookingID string) {
	b.bookingID = bookingID
}

func (b *BookingClient) HoldID() string {
	return b.holdID
}

func (b *BookingClient) SetHoldID(holdID string) {
	b.holdID = holdID
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

func (b *BookingClient) String() {
	t := table.NewWriter()
	t.SetStyle(table.StyleLight)
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Entity Name", "UUID",})
	t.AppendRows([]table.Row{
		{1, "booking ID", b.bookingID},
		{2, "hold ID", b.holdID},
	})
	t.Render()

	//b.logger.Logger().WithFields(logrus.Fields{"booking ID": b.bookingID, "hold ID": b.holdID}).Info("booking client data")
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
