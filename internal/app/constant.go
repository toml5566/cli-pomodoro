package app

const PROJECT_NAME = "CLI Pomodoro"

// Timer states
type SessionType int

const (
	Focus SessionType = iota
	ShortBreak
	LongBreak
)

// Session durations (in seconds for easier testing, change to minutes * 60 for production)
const (
	FocusDuration      = 25 * 60 // 25 minutes
	ShortBreakDuration = 5 * 60  // 5 minutes
	LongBreakDuration  = 15 * 60 // 15 minutes
)

// background color codes
const (
	FocusBgColor      = "#1a1a2e" // Dark blue
	ShortBreakBgColor = "#16537e" // Green-blue
	LongBreakBgColor  = "#0f3460" // Darker blue-green
)

// ASCII art for large numbers
var asciiNumbers = map[string][]string{
	"0": {
		" ███████ ",
		"██     ██",
		"██     ██",
		"██     ██",
		"██     ██",
		"██     ██",
		" ███████ ",
	},
	"1": {
		"    ██   ",
		"  ████   ",
		"    ██   ",
		"    ██   ",
		"    ██   ",
		"    ██   ",
		"  ██████ ",
	},
	"2": {
		" ███████ ",
		"██     ██",
		"       ██",
		" ███████ ",
		"██       ",
		"██       ",
		"█████████",
	},
	"3": {
		" ███████ ",
		"██     ██",
		"       ██",
		" ███████ ",
		"       ██",
		"██     ██",
		" ███████ ",
	},
	"4": {
		"██     ██",
		"██     ██",
		"██     ██",
		"█████████",
		"       ██",
		"       ██",
		"       ██",
	},
	"5": {
		"█████████",
		"██       ",
		"██       ",
		"████████ ",
		"       ██",
		"██     ██",
		" ███████ ",
	},
	"6": {
		" ███████ ",
		"██     ██",
		"██       ",
		"████████ ",
		"██     ██",
		"██     ██",
		" ███████ ",
	},
	"7": {
		"█████████",
		"       ██",
		"      ██ ",
		"     ██  ",
		"    ██   ",
		"   ██    ",
		"  ██     ",
	},
	"8": {
		" ███████ ",
		"██     ██",
		"██     ██",
		" ███████ ",
		"██     ██",
		"██     ██",
		" ███████ ",
	},
	"9": {
		" ███████ ",
		"██     ██",
		"██     ██",
		" ████████",
		"       ██",
		"██     ██",
		" ███████ ",
	},
	":": {
		"         ",
		"   ███   ",
		"   ███   ",
		"         ",
		"   ███   ",
		"   ███   ",
		"         ",
	},
}
