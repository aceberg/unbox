package share

// Colors contains ANSI escape codes for colored terminal output
type Colors struct {
	Err   string
	Bkp   string
	Ok    string
	Warn  string
	Main  string
	Reset string
}

// Col provides the default ANSI color codes for terminal output
var Col = Colors{
	Err:   "\033[31m", // red
	Bkp:   "\033[32m", // green
	Ok:    "\033[92m", // bright green
	Warn:  "\033[33m", // yellow
	Main:  "\033[36m", // cyan
	Reset: "\033[0m",  // reset
}
