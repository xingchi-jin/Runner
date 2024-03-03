package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func Command() {
	app := kingpin.New("Runner", "Harness Runner to execute tasks")
	app.HelpFlag.Short('h')
	server.Register(app)
	certs.Register(app)
	client.Register(app)

	kingpin.MustParse(app.Parse(os.Args[1:]))
}