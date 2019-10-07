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
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = cobra.Command{
		Use:          "channels_booking_clients",
		Short:        "channel booking generator",
		Long:         "channel booking generator",
		SilenceUsage: true,
	}

	Version    string
	Commit     string
	Date string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, commit, date string) {
	Version = version
	Commit = commit
	Date = date

	RootCmd.AddCommand(infoCmd)
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	//run(version, commit, date)
}

//
//func run(version, commit, date string) {
//	var (
//		ctx  = context.Background()
//		data = make(map[string]string)
//	)
//
//	lgr := &logger.LocalLogger{}
//	defer lgr.Close()
//
//	cfg, err := config.NewConfig()
//	if err != nil {
//		lgr.Logger().WithFields(logrus.Fields{"cfg": cfg, "error": "configuration initialization is failed"}).Error(err)
//
//		os.Exit(1)
//	}
//	// compilation data
//	cfg.Compilation.SetDate(date)
//	cfg.Compilation.SetVersion(version)
//	cfg.Compilation.SetCommit(commit)
//	// logger configuration
//	if cfg.Logger.Level == "" {
//		lgr.Logger().WithFields(logrus.Fields{"logger level": cfg.Logger.Level, "error": "logger level is empty"}).Error(err)
//
//		os.Exit(1)
//	}
//
//	data["level"] = cfg.Logger.Level
//
//	err = lgr.Init(data, &cfg.Logger)
//	if err != nil {
//		os.Exit(1)
//	}
//
//	lgr.Logger().WithFields(logrus.Fields{"config": cfg}).Debug("Debug configuration")
//
//	err = ValidateStopAfterEntity(cfg.StopAfterEntity)
//	if err != nil {
//		os.Exit(1)
//	}
//
//	channelsClient, err := channels.NewChannelClient(cfg.ChannelEnv.XAPIKey, cfg.ChannelEnv.XAPISecret)
//	if err != nil {
//		os.Exit(1)
//	}
//	channelsClient.SetLogger(lgr)
//
//	err = runChannelsSteps(cfg, channelsClient, &ctx, lgr)
//	if err != nil {
//		channelsClient.String()
//		os.Exit(1)
//	}
//
//	bookingClient, err := booking.NewBookingClient(cfg.BookingEnv.XAPIKey, cfg.BookingEnv.XAPISecret)
//	if err != nil {
//		channelsClient.String()
//		lgr.Logger().WithFields(logrus.Fields{"booking client": bookingClient, "error": err}).Fatal("Booking client is not initialized")
//	}
//
//	bookingClient.SetLogger(lgr)
//
//	err = runBookingSteps(cfg, channelsClient, bookingClient, &ctx, lgr)
//	if err != nil {
//		lgr.Logger().WithFields(logrus.Fields{"error": err}).Fatal(err)
//	}
//
//}
//
//// runChannelsSteps - run all steps of channel api
//func runChannelsSteps(cfg *config.Configuration, client *channels.ChannelsClient, ctx *context.Context, lgr *logger.LocalLogger) error {
//
//	if err := utils.CheckIsUUIDTypeOfFlagValue(cfg.ChannelID); err != nil {
//		return err
//	} else {
//		client.SetChannelID(cfg.ChannelID)
//	}
//
//	supplier := NewSupplierRepository(client, ctx, lgr, cfg)
//	err := supplier.Execute()
//	if err != nil {
//		return err
//	}
//
//	if supplier.Name() == cfg.StopAfterEntity {
//		client.String()
//		os.Exit(0)
//	}
//
//	product := NewProductRepository(client, ctx, lgr, cfg)
//	err = product.Execute()
//	if err != nil {
//		return err
//	}
//
//	if product.Name() == cfg.StopAfterEntity {
//		client.String()
//		os.Exit(0)
//	}
//
//	rate := NewRateRepository(client, ctx, lgr, cfg)
//	err = rate.Execute()
//	if err != nil {
//		return err
//	}
//
//	if rate.Name() == cfg.StopAfterEntity {
//		client.String()
//		os.Exit(0)
//	}
//
//	channelBinding := NewChannelBindingRepository(client, ctx, lgr, cfg)
//	err = channelBinding.Execute()
//	if err != nil {
//		return err
//	}
//
//	if channelBinding.Name() == cfg.StopAfterEntity {
//		client.String()
//		os.Exit(0)
//	}
//
//	return nil
//}
//
//// runBookingSteps - run all steps of booking api
//func runBookingSteps(cfg *config.Configuration, channelClient *channels.ChannelsClient, client *booking.BookingClient, ctx *context.Context, lgr *logger.LocalLogger) error {
//
//	bkng := NewBookingRepository(client, channelClient, ctx, lgr, cfg)
//	err := bkng.Execute()
//	if err != nil {
//		return err
//	}
//
//	if bkng.Name() == cfg.StopAfterEntity {
//		channelClient.String()
//		client.String()
//		os.Exit(0)
//	}
//
//	hold := NewHoldRepository(client, channelClient, ctx, lgr, cfg)
//
//	err = hold.Execute()
//	if err != nil {
//		return err
//	}
//
//	if hold.Name() == cfg.StopAfterEntity {
//		channelClient.String()
//		client.String()
//		os.Exit(0)
//	}
//
//	return nil
//}
