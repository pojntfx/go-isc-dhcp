package cmd

const (
	keyPrefix         = "dhclient."
	configFileDefault = ""
	serverHostPortKey = keyPrefix + "serverHostPort"
	configFileKey     = keyPrefix + "configFile"
	deviceKey         = keyPrefix + "device"
)

var (
	serverHostPortFlag string
	configFileFlag     string
)
