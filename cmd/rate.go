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
	"github.com/NickTaporuk/channels_booking_clients/config"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/NickTaporuk/channels_booking_clients/utils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

const (
	RateCommandName = "rate"
)

var (
	ErrorRateFileIsNotFound = errors.New("file is not found by path")
)

// RateRepository
type RateRepository struct {
	client        *channels.ChannelsClient
	ctx           *context.Context
	logger        *logger.LocalLogger
	configuration *config.Configuration
	name          string
}

func NewRateRepository(client *channels.ChannelsClient, ctx *context.Context, logger *logger.LocalLogger, configuration *config.Configuration) *RateRepository {
	return &RateRepository{client: client, ctx: ctx, logger: logger, configuration: configuration}
}

func (s *RateRepository) Name() string {
	return RateCommandName
}

func (s *RateRepository) Configuration() *config.Configuration {
	return s.configuration
}

func (s *RateRepository) SetConfiguration(configuration *config.Configuration) {
	s.configuration = configuration
}

func (s *RateRepository) Logger() *logger.LocalLogger {
	return s.logger
}

func (s *RateRepository) SetLogger(logger *logger.LocalLogger) {
	s.logger = logger
}

func (s *RateRepository) Ctx() *context.Context {
	return s.ctx
}

func (s *RateRepository) SetCtx(ctx *context.Context) {
	s.ctx = ctx
}

func (s *RateRepository) Client() *channels.ChannelsClient {
	return s.client
}

func (s *RateRepository) SetClient(client *channels.ChannelsClient) {
	s.client = client
}

// Execute represents the supplier command
func (s *RateRepository) Execute() error {
	var (
		err error
	)

	if s.configuration.Rate.ID != "" {
		if err = utils.CheckIsUUIDTypeOfFlagValue(s.configuration.Rate.ID); err == nil {
			if err = utils.CheckRateExist(s.configuration.Supplier.ID, s.configuration.Product.ID, s.configuration.Rate.ID, s.client, s.ctx, s.logger); err != nil {
				s.logger.Logger().WithFields(logrus.Fields{"rate": s, "error": err}).Error(err)
				return err
			}
		}
	} else if s.configuration.Rate.Path != "" {
		err = utils.CheckIsPathTypeOfFlagValue(s.configuration.Rate.Path)
		if err == nil {

			data, err := ioutil.ReadFile(s.configuration.Rate.Path)
			if err != nil {
				s.logger.Logger().WithField("Rate", "error").WithFields(logrus.Fields{"error": ErrorRateFileIsNotFound, "data": data}).Error(err)
				return err
			}

			err = s.client.CreateRate(&data, s.Ctx())
			if err != nil {
				return err
			}
		}
	}

	return err
}
