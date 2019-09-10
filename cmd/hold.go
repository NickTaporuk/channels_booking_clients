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
	"github.com/NickTaporuk/channels_booking_clients/channels"
	"io/ioutil"

	"github.com/NickTaporuk/channels_booking_clients/booking"
	"github.com/NickTaporuk/channels_booking_clients/config"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/NickTaporuk/channels_booking_clients/utils"
	"github.com/sirupsen/logrus"
)

const (
	HoldCommandName = "hold"
)

var (
	ErrorHoldFileIsNotFound = errors.New("hold.json file was not found by path")
)

// HoldRepository
type HoldRepository struct {
	client        *booking.BookingClient
	channelClient *channels.ChannelsClient
	ctx           *context.Context
	logger        *logger.LocalLogger
	configuration *config.Configuration
}

func NewHoldRepository(client *booking.BookingClient, channelClient *channels.ChannelsClient, ctx *context.Context, logger *logger.LocalLogger, configuration *config.Configuration) *HoldRepository {
	return &HoldRepository{client: client, channelClient: channelClient, ctx: ctx, logger: logger, configuration: configuration}
}

func (s *HoldRepository) ChannelsClient() *channels.ChannelsClient {
	return s.channelClient
}

func (s *HoldRepository) SetChannelsClient(channelClient *channels.ChannelsClient) {
	s.channelClient = channelClient
}

func (s *HoldRepository) Name() string {
	return HoldCommandName
}

func (s *HoldRepository) Configuration() *config.Configuration {
	return s.configuration
}

func (s *HoldRepository) SetConfiguration(configuration *config.Configuration) {
	s.configuration = configuration
}

func (s *HoldRepository) Logger() *logger.LocalLogger {
	return s.logger
}

func (s *HoldRepository) SetLogger(logger *logger.LocalLogger) {
	s.logger = logger
}

func (s *HoldRepository) Ctx() *context.Context {
	return s.ctx
}

func (s *HoldRepository) SetCtx(ctx *context.Context) {
	s.ctx = ctx
}

func (s *HoldRepository) Client() *booking.BookingClient {
	return s.client
}

func (s *HoldRepository) SetClient(client *booking.BookingClient) {
	s.client = client
}

// Execute represents the supplier command
func (s *HoldRepository) Execute() error {
	var (
		err error
	)

	if s.configuration.Hold.ID != "" {
		if err = utils.CheckIsUUIDTypeOfFlagValue(s.configuration.Hold.ID); err == nil {
			if err = utils.CheckHoldExist(s.configuration.Hold.ID, s.client, s.ctx, s.logger); err != nil {
				s.logger.Logger().WithFields(logrus.Fields{"hold": s, "error": err}).Error(err)
				return err
			}
		}
	} else if s.configuration.Hold.Path != "" {
		err = utils.CheckIsPathTypeOfFlagValue(s.configuration.Hold.Path)
		if err == nil {

			data, err := ioutil.ReadFile(s.configuration.Hold.Path)
			if err != nil {
				s.logger.Logger().WithField("Hold", "error").WithFields(logrus.Fields{"error": ErrorHoldFileIsNotFound,}).Error(err)
				return err
			}

			priceID := s.channelClient.PriceIDs()[0]
			rateID := s.channelClient.RateIDs()[0]
			supplierID := s.channelClient.SupplierID()

			err = s.client.CreateHold(rateID, supplierID, priceID, &data, s.Ctx())
			if err != nil {
				return err
			}
		}
	}

	return err
}
