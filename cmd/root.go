/*
Copyright © 2019 Nikolay Kuropatkin nictaporuk@gmail.com

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
	"os"

	"github.com/NickTaporuk/channels_booking_clients/booking"
	"github.com/NickTaporuk/channels_booking_clients/channels"
	"github.com/NickTaporuk/channels_booking_clients/config"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/NickTaporuk/channels_booking_clients/utils"
	"github.com/sirupsen/logrus"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var (
		channelsClient *channels.ChannelsClient
		bookingClient  *booking.BookingClient
		ctx            = context.Background()
		err            error
		data           = make(map[string]string)
		lgr            *logger.LocalLogger
		cfg            *config.Configuration
	)

	lgr = &logger.LocalLogger{}
	defer lgr.Close()

	cfg, err = config.NewConfig()
	if err != nil {
		lgr.Logger().WithFields(logrus.Fields{"cfg": cfg, "error": "configuration initialization is failed"}).Error(err)

		os.Exit(1)
	}

	if cfg.Logger.Level == "" {
		lgr.Logger().WithFields(logrus.Fields{"logger level": cfg.Logger.Level, "error": "logger level is empty"}).Error(err)

		os.Exit(1)
	}

	data["level"] = cfg.Logger.Level

	err = lgr.Init(data, &cfg.Logger)
	if err != nil {
		os.Exit(1)
	}

	lgr.Logger().WithFields(logrus.Fields{"config": cfg}).Debug("Debug configuration")

	err = ValidateStopAfterEntity(cfg.StopAfterEntity)
	if err != nil {
		os.Exit(1)
	}

	channelsClient, err = channels.NewChannelClient(cfg.ChannelEnv.XAPIKey, cfg.ChannelEnv.XAPISecret)
	if err != nil {
		os.Exit(1)
	}
	channelsClient.SetLogger(lgr)

	if err = utils.CheckIsUUIDTypeOfFlagValue(cfg.ChannelID); err != nil {
		os.Exit(1)
	} else {
		channelsClient.SetChannelID(cfg.ChannelID)
	}

	supplier := NewSupplierRepository(channelsClient, &ctx, lgr, cfg)
	err = supplier.Execute()
	if err != nil {
		channelsClient.String()
		os.Exit(1)
	}

	if supplier.Name() == cfg.StopAfterEntity {
		channelsClient.String()
		os.Exit(0)
	}

	product := NewProductRepository(channelsClient, &ctx, lgr, cfg)
	err = product.Execute()
	if err != nil {
		channelsClient.String()
		os.Exit(1)
	}

	if product.Name() == cfg.StopAfterEntity {
		channelsClient.String()
		os.Exit(0)
	}

	rate := NewRateRepository(channelsClient, &ctx, lgr, cfg)
	err = rate.Execute()
	if err != nil {
		channelsClient.String()
		os.Exit(1)
	}

	if rate.Name() == cfg.StopAfterEntity {
		channelsClient.String()
		os.Exit(0)
	}

	channelBinding := NewChannelBindingRepository(channelsClient, &ctx, lgr, cfg)
	err = channelBinding.Execute()
	if err != nil {
		channelsClient.String()
		os.Exit(1)
	}

	if channelBinding.Name() == cfg.StopAfterEntity {
		channelsClient.String()
		os.Exit(0)
	}

	bookingClient, err = booking.NewBookingClient(cfg.BookingEnv.XAPIKey, cfg.BookingEnv.XAPISecret)
	if err != nil {
		channelsClient.String()
		lgr.Logger().WithFields(logrus.Fields{"request": channelBinding}).Fatal("Booking client is not initialized")
	}

	bookingClient.SetLogger(lgr)

	booking := NewBookingRepository(bookingClient, channelsClient, &ctx, lgr, cfg)
	err = booking.Execute()
	if err != nil {
		channelsClient.String()
		bookingClient.String()
		os.Exit(1)
	}

	if booking.Name() == cfg.StopAfterEntity {
		channelsClient.String()
		bookingClient.String()
		os.Exit(0)
	}

	hold := NewHoldRepository(bookingClient, channelsClient, &ctx, lgr, cfg)

	err = hold.Execute()
	if err != nil {
		channelsClient.String()
		bookingClient.String()
		lgr.Logger().WithFields(logrus.Fields{"hold": hold, "error": err}).Error(err)
		os.Exit(1)
	}

	if hold.Name() == cfg.StopAfterEntity {
		channelsClient.String()
		bookingClient.String()
		os.Exit(0)
	}
}
