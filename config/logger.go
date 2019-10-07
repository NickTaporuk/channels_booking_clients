package config

// LoggerConfiguration is used to
type LoggerConfiguration struct {
	Level string
	PrettyPrint bool
	File bool
	StdOut bool
	FilePath string
}
