/*
Copyright © 2019 Nickolay Kuropatkin nictaporuk@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"errors"
	"io/ioutil"

	"github.com/NickTaporuk/channels_booking_clients/booking"
	"github.com/NickTaporuk/channels_booking_clients/channels"
	"github.com/NickTaporuk/channels_booking_clients/config"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/NickTaporuk/channels_booking_clients/utils"
	"github.com/sirupsen/logrus"
)

const (
	BookingCommandName = "booking"
)

var (
	ErrorBookingFileIsNotFound = errors.New("booking.json file was not found by path")
)

// BookingRepository
type BookingRepository struct {
	client        *booking.BookingClient
	channelClient *channels.ChannelsClient
	ctx           *context.Context
	logger        *logger.LocalLogger
	configuration *config.Configuration
	name          string
}

func (s *BookingRepository) ChannelClient() *channels.ChannelsClient {
	return s.channelClient
}

func (s *BookingRepository) SetChannelClient(channelClient *channels.ChannelsClient) {
	s.channelClient = channelClient
}

func NewBookingRepository(client *booking.BookingClient, channelClient *channels.ChannelsClient, ctx *context.Context, logger *logger.LocalLogger, configuration *config.Configuration) *BookingRepository {
	return &BookingRepository{client: client, channelClient: channelClient, ctx: ctx, logger: logger, configuration: configuration}
}

func (s *BookingRepository) Name() string {
	return BookingCommandName
}

func (s *BookingRepository) Configuration() *config.Configuration {
	return s.configuration
}

func (s *BookingRepository) SetConfiguration(configuration *config.Configuration) {
	s.configuration = configuration
}

func (s *BookingRepository) Logger() *logger.LocalLogger {
	return s.logger
}

func (s *BookingRepository) SetLogger(logger *logger.LocalLogger) {
	s.logger = logger
}

func (s *BookingRepository) Ctx() *context.Context {
	return s.ctx
}

func (s *BookingRepository) SetCtx(ctx *context.Context) {
	s.ctx = ctx
}

func (s *BookingRepository) Client() *booking.BookingClient {
	return s.client
}

func (s *BookingRepository) SetClient(client *booking.BookingClient) {
	s.client = client
}

// Execute represents the supplier command
func (s *BookingRepository) Execute() error {
	var (
		err error
	)

	if s.configuration.Booking.ID != "" {
		if err = utils.CheckIsUUIDTypeOfFlagValue(s.configuration.Booking.ID); err == nil {
			if err = utils.CheckBookingExist(s.configuration.Booking.ID, s.client, s.ctx, s.logger); err != nil {
				s.logger.Logger().WithFields(logrus.Fields{"supplier": s, "error": err}).Error(err)
				return err
			}
		}
	} else if s.configuration.Booking.Path != "" {
		err = utils.CheckIsPathTypeOfFlagValue(s.configuration.Booking.Path)
		if err == nil {

			data, err := ioutil.ReadFile(s.configuration.Booking.Path)
			if err != nil {
				s.logger.Logger().WithField("Booking", "error").WithFields(logrus.Fields{"error": ErrorBookingFileIsNotFound,}).Error(err)
				return err
			}

			priceID := s.channelClient.PriceIDs()[0]
			rateID := s.channelClient.RateIDs()[0]
			supplierID := s.channelClient.SupplierID()
			err = s.client.CreateBooking(priceID, rateID, supplierID, &data, s.Ctx())
			if err != nil {
				return err
			}
		}
	}

	return err
}
