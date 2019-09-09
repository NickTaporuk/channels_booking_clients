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
	"github.com/NickTaporuk/channels_booking_clients/config"
	"github.com/NickTaporuk/channels_booking_clients/logger"
	"github.com/sirupsen/logrus"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var (
		//channelsApiHeaders = make(map[string]string)
		//channelsClient     *channels.ChannelsClient
		//bookingClient      *booking.BookingClient
		//ctx                = context.Background()
		//getSupplier        swclient.ResponseGetSupplierEnvelope
		//channelBinding     *swclient.RequestPostCreateChannelEnvelope
		//respChan           swclient.ResponsePostChannelRatesEnvelope
		//respHold           bookingCl.ResponsePostHoldEnvelope
		//resp               *http.Response
		err error
		//rates              []string
		//prices             []string
		data = make(map[string]string)
		lgr  *logger.LocalLogger
		//book               *bookingCl.RequestPostBookingEnvelope
		//respBooking        bookingCl.ResponsePostBookingEnvelope
		cfg *config.Configuration
	)

	lgr = &logger.LocalLogger{}
	defer lgr.Close()

	cfg, err = config.NewConfig()
	if err != nil {
		panic(err)
	}


	if cfg.Logger.Level == "" {
		panic(err)
	}

	data["level"] = cfg.Logger.Level

	err = lgr.Init(data)
	if err != nil {
		panic(err)
	}

		lgr.Logger().WithFields(logrus.Fields{"config": cfg}).Debug("Debug configuration")
}

//func init() {
//	cobra.OnInitialize(initConfig)
//
//	// TODO: need move to cli parameters
//	data["level"] = "debug"
//	lgr = &logger.LocalLogger{}
//
//	defer lgr.Close()
//
//	err = lgr.Init(data)
//
//	if err != nil {
//		panic(err)
//	}
//	//
//	channelsClient, err = channels.NewChannelClient(channelsApiHeaders)
//
//	if err != nil {
//		panic(err)
//	}
//
//	channelsClient.SetLogger(lgr)
//
//	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.channels_booking_clients.yaml)")
//}

//// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		// Search config in home directory with name ".channels_booking_clients" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".cbc")
//	}
//
//	viper.AutomaticEnv() // read in environment variables that match
//
//	// If a config file is found, read it in.
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}
