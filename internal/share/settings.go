package share

// SettingsType contains unbox settings for conf and keep
type SettingsType struct {
	APIPath      string
	APISecret    string
	OutPath      string
	InputPath    string
	TestURL      string
	LimitTimeout int
	DelayMain    int
	DelayBkp     int
	DelayAll     int
	DelaySwitch  int
	SwitchStep   int
	BestN        int
	BackupN      int
	Deduplicate  bool
}

// Settings contains unbox settings for conf and keep
var Settings SettingsType
