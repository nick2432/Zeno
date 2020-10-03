package config

type Flags struct {
	Job       string
	Workers   int
	MaxHops   uint
	WARC      bool
	Headless  bool
	Seencheck bool
	JSON      bool
	Debug     bool
}

type Application struct {
	Flags Flags
}

var App *Application

func init() {
	App = &Application{}
}
