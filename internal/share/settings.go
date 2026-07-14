package share

// SettingsType contains unbox settings for conf and keep
type SettingsType struct {
	APIPath     string
	APISecret   string
	OutPath     string
	InputPath   string
	TestURL     string
	DelayMain   uint
	DelayBkp    uint
	DelayAll    uint
	DelaySwitch uint
	Deduplicate bool
}

// Settings contains unbox settings for conf and keep
var Settings SettingsType
