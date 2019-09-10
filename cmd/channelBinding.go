/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	ChannelBindingCommandName = "channelBinding"
)

var (
	ErrorChannelBindingFileIsNotFound = errors.New("file is not found by path")
)

// ChannelBindingRepository
type ChannelBindingRepository struct {
	client        *channels.ChannelsClient
	ctx           *context.Context
	logger        *logger.LocalLogger
	configuration *config.Configuration
	name          string
}

func NewChannelBindingRepository(client *channels.ChannelsClient, ctx *context.Context, logger *logger.LocalLogger, configuration *config.Configuration) *ChannelBindingRepository {
	return &ChannelBindingRepository{client: client, ctx: ctx, logger: logger, configuration: configuration}
}

func (s *ChannelBindingRepository) Name() string {
	return ChannelBindingCommandName
}

func (s *ChannelBindingRepository) Configuration() *config.Configuration {
	return s.configuration
}

func (s *ChannelBindingRepository) SetConfiguration(configuration *config.Configuration) {
	s.configuration = configuration
}

func (s *ChannelBindingRepository) Logger() *logger.LocalLogger {
	return s.logger
}

func (s *ChannelBindingRepository) SetLogger(logger *logger.LocalLogger) {
	s.logger = logger
}

func (s *ChannelBindingRepository) Ctx() *context.Context {
	return s.ctx
}

func (s *ChannelBindingRepository) SetCtx(ctx *context.Context) {
	s.ctx = ctx
}

func (s *ChannelBindingRepository) Client() *channels.ChannelsClient {
	return s.client
}

func (s *ChannelBindingRepository) SetClient(client *channels.ChannelsClient) {
	s.client = client
}

// Execute represents the supplier command
func (s *ChannelBindingRepository) Execute() error {
	var (
		err error
	)

	if s.configuration.ChannelBinding.ID != "" {
		if err = utils.CheckIsUUIDTypeOfFlagValue(s.configuration.ChannelBinding.ID); err == nil {
			if err = utils.CheckChannelBindingExist(s.configuration.ChannelBinding.ID, s.client, s.ctx, s.logger); err != nil {
				s.logger.Logger().WithFields(logrus.Fields{"ChannelBinding": s, "error": err}).Error(err)
				return err
			}
		}
	} else if s.configuration.ChannelBinding.Path != "" {
		err = utils.CheckIsPathTypeOfFlagValue(s.configuration.ChannelBinding.Path)
		if err == nil {

			data, err := ioutil.ReadFile(s.configuration.ChannelBinding.Path)
			if err != nil {
				s.logger.Logger().WithField("ChannelBinding", "error").WithFields(logrus.Fields{"error": ErrorChannelBindingFileIsNotFound,}).Error(err)
				return err
			}

			err = s.client.CreateChannelBinding(&data, s.Ctx())
			if err != nil {
				return err
			}
		}
	}

	return err
}
