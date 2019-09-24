package cmd

import (
	"github.com/NickTaporuk/channels_booking_clients/config"
	"github.com/NickTaporuk/channels_booking_clients/logger"
)

// EntityRepository is interface of repositories
type EntityRepository interface {
	Execute() error
	Name() string
	Configuration() *config.Configuration
	Logger() *logger.LocalLogger
}
