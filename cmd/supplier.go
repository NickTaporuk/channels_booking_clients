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

	"github.com/NickTaporuk/channels_booking_clients/channels"
	"github.com/NickTaporuk/channels_booking_clients/config"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/NickTaporuk/channels_booking_clients/utils"
	"github.com/sirupsen/logrus"
)

const (
	SupplierCommandName = "supplier"
)

var (
	ErrorSupplierFileIsNotFound = errors.New("file is not found by path")
)

// SupplierRepository
type SupplierRepository struct {
	client        *channels.ChannelsClient
	ctx           *context.Context
	logger        *logger.LocalLogger
	configuration *config.Configuration
	name          string
}

func NewSupplierRepository(client *channels.ChannelsClient, ctx *context.Context, logger *logger.LocalLogger, configuration *config.Configuration) *SupplierRepository {
	return &SupplierRepository{client: client, ctx: ctx, logger: logger, configuration: configuration}
}

func (s *SupplierRepository) Name() string {
	return SupplierCommandName
}

func (s *SupplierRepository) Configuration() *config.Configuration {
	return s.configuration
}

func (s *SupplierRepository) SetConfiguration(configuration *config.Configuration) {
	s.configuration = configuration
}

func (s *SupplierRepository) Logger() *logger.LocalLogger {
	return s.logger
}

func (s *SupplierRepository) SetLogger(logger *logger.LocalLogger) {
	s.logger = logger
}

func (s *SupplierRepository) Ctx() *context.Context {
	return s.ctx
}

func (s *SupplierRepository) SetCtx(ctx *context.Context) {
	s.ctx = ctx
}

func (s *SupplierRepository) Client() *channels.ChannelsClient {
	return s.client
}

func (s *SupplierRepository) SetClient(client *channels.ChannelsClient) {
	s.client = client
}

// Execute represents the supplier command
func (s *SupplierRepository) Execute() error {
	var (
		err error
	)

	if s.configuration.Supplier.ID != "" {
		if err = utils.CheckIsUUIDTypeOfFlagValue(s.configuration.Supplier.ID); err == nil {
			if err = utils.CheckSupplierExist(s.configuration.Supplier.ID, s.client, s.ctx, s.logger); err != nil {
				s.logger.Logger().WithFields(logrus.Fields{"supplier": s, "error": err}).Error(err)
				return err
			}
		}
	} else if s.configuration.Supplier.Path != "" {
		err = utils.CheckIsPathTypeOfFlagValue(s.configuration.Supplier.Path)
		if err == nil {

			data, err := ioutil.ReadFile(s.configuration.Supplier.Path)
			if err != nil {
				s.logger.Logger().WithField("Supplier", "error").WithFields(logrus.Fields{"error": ErrorSupplierFileIsNotFound,}).Error(err)
				return err
			}

			err = s.client.CreateSupplier(&data, s.Ctx())
			if err != nil {
				return err
			}
		}
	}

	return err
}
