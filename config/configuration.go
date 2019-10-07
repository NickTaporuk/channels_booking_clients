package config

// Configuration is general structure of app configuration
type Configuration struct {
	Logger          LoggerConfiguration
	ChannelID       string
	ChannelEnv      ChannelEnvConfiguration
	BookingEnv      BookingEnvConfiguration
	Supplier        SupplierConfiguration
	Product         ProductConfiguration
	Rate            RateConfiguration
	ChannelBinding  ChannelBindingConfiguration
	Booking         BookingConfiguration
	Hold            HoldConfiguration
	Compilation     CompilationConfiguration
	StopAfterEntity string
}
