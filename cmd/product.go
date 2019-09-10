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
	ProductCommandName = "product"
)

var (
	ErrorProductFileIsNotFound = errors.New("file is not found by path")
)

// ProductRepository
type ProductRepository struct {
	client        *channels.ChannelsClient
	ctx           *context.Context
	logger        *logger.LocalLogger
	configuration *config.Configuration
	name          string
}

func NewProductRepository(client *channels.ChannelsClient, ctx *context.Context, logger *logger.LocalLogger, configuration *config.Configuration) *ProductRepository {
	return &ProductRepository{client: client, ctx: ctx, logger: logger, configuration: configuration}
}

func (p *ProductRepository) Name() string {
	return ProductCommandName
}

func (p *ProductRepository) Configuration() *config.Configuration {
	return p.configuration
}

func (p *ProductRepository) SetConfiguration(configuration *config.Configuration) {
	p.configuration = configuration
}

func (p *ProductRepository) Logger() *logger.LocalLogger {
	return p.logger
}

func (p *ProductRepository) SetLogger(logger *logger.LocalLogger) {
	p.logger = logger
}

func (p *ProductRepository) Ctx() *context.Context {
	return p.ctx
}

func (p *ProductRepository) SetCtx(ctx *context.Context) {
	p.ctx = ctx
}

func (p *ProductRepository) Client() *channels.ChannelsClient {
	return p.client
}

func (p *ProductRepository) SetClient(client *channels.ChannelsClient) {
	p.client = client
}

// Execute represents the supplier command
func (p *ProductRepository) Execute() error {
	var (
		err error
	)

	if p.configuration.Product.ID != "" {
		if err = utils.CheckIsUUIDTypeOfFlagValue(p.configuration.Product.ID); err == nil {
			if err = utils.CheckProductExist(p.configuration.Supplier.ID,p.configuration.Product.ID, p.client, p.ctx, p.logger); err != nil {
				p.logger.Logger().WithFields(logrus.Fields{"product": p, "error": err}).Error(err)
				return err
			}
		}
	} else if p.configuration.Product.Path != "" {
		err = utils.CheckIsPathTypeOfFlagValue(p.configuration.Product.Path)
		if err == nil {

			data, err := ioutil.ReadFile(p.configuration.Product.Path)
			if err != nil {
				p.logger.Logger().WithField("Product", "error").WithFields(logrus.Fields{"error": ErrorProductFileIsNotFound,}).Error(err)
				return err
			}

			err = p.client.CreateProduct(&data, p.Ctx())
			if err != nil {
				return err
			}
		}
	}

	return err
}
