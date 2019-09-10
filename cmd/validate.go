package cmd

import (
	"errors"
)

var (
	ErrorStopAfterEntityIsUnknown = errors.New("value of stopAfterEntity is unknown")
)

// ValidateStopAfterEntity validate value of configuration field stopAfterEntity
func ValidateStopAfterEntity(stopAfterEntity string) error {

	switch stopAfterEntity {
	case
		SupplierCommandName,
		ProductCommandName,
		RateCommandName,
		ChannelBindingCommandName,
		BookingCommandName,
		HoldCommandName:
		return nil
	}

	return ErrorStopAfterEntityIsUnknown
}
