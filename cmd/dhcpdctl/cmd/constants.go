package cmd

const (
	keyPrefix         = "dhcpd."
	configFileDefault = ""
	serverHostPortKey = keyPrefix + "serverHostPort"
	configFileKey     = keyPrefix + "configFile"
	deviceKey         = keyPrefix + "device"
	subnetsKey        = keyPrefix + "subnets"
)

var (
	serverHostPortFlag string
	configFileFlag     string
)
