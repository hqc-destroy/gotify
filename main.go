package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/codegangsta/cli"

	"github.com/yudai/gotty/app"
	"github.com/yudai/gotty/backends/ptycommand"
	"github.com/yudai/gotty/utils"
)

func main() {
	cmd := cli.NewApp()
	cmd.Name = "gotty"
	cmd.Version = app.Version
	cmd.Usage = "Share your terminal as a web application"
	cmd.HideHelp = true
	cli.AppHelpTemplate = helpTemplate

<<<<<<< HEAD
<<<<<<< HEAD
	flags := []flag{
		flag{"address", "a", "IP address to listen"},
		flag{"port", "p", "Port number to listen"},
		flag{"permit-write", "w", "Permit clients to write to the TTY (BE CAREFUL)"},
		flag{"credential", "c", "Credential for Basic Authentication (ex: user:pass, default disabled)"},
		flag{"random-url", "r", "Add a random string to the URL"},
		flag{"random-url-length", "", "Random URL length"},
		flag{"tls", "t", "Enable TLS/SSL"},
		flag{"tls-crt", "", "TLS/SSL crt file path"},
		flag{"tls-key", "", "TLS/SSL key file path"},
		flag{"client", "C", "Enable Client Certificate"},
		flag{"client-ca-file", "", "Client CA certificate file"},
		flag{"client-verify", "", "Enable verification of client certificate"},
		flag{"index", "", "Custom index.html file"},
		flag{"title-format", "", "Title format of browser window"},
		flag{"reconnect", "", "Enable reconnection"},
		flag{"reconnect-time", "", "Time to reconnect"},
		flag{"timeout", "", "Timeout seconds for waiting a client (0 to disable)"},
		flag{"max-connection", "", "Maximum connection to gotty, 0(default) means no limit"},
		flag{"once", "", "Accept only one client and exit on disconnection"},
		flag{"permit-arguments", "", "Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB)"},
		flag{"close-signal", "", "Signal sent to the command process when gotty close it (default: SIGHUP)"},
		flag{"width", "", "Static width of the screen, 0(default) means dynamically resize"},
		flag{"height", "", "Static height of the screen, 0(default) means dynamically resize"},
	}

	mappingHint := map[string]string{
		"index":          "IndexFile",
		"tls":            "EnableTLS",
		"tls-crt":        "TLSCrtFile",
		"tls-key":        "TLSKeyFile",
		"client":         "EnableClientCertificate",
		"client-ca-file": "ClientCAFile",
		"client-verify":  "EnableClientCertificateVerification",
		"random-url":     "EnableRandomUrl",
		"reconnect":      "EnableReconnect",
	}

	cliFlags, err := generateFlags(flags, mappingHint)
=======
	options := &app.Options{}
	if err := utils.ApplyDefaultValues(options); err != nil {
		exit(err, 1)
	}

	cliFlags, flagMappings, err := utils.GenerateFlags(options)
>>>>>>> d71e2fc... generate falgs based on struct options instead of defining them externally
=======
	appOptions := &app.Options{}
	if err := utils.ApplyDefaultValues(appOptions); err != nil {
		exit(err, 1)
	}
	backendOptions := &ptycommand.Options{}
	if err := utils.ApplyDefaultValues(backendOptions); err != nil {
		exit(err, 1)
	}

	cliFlags, flagMappings, err := utils.GenerateFlags(appOptions, backendOptions)
>>>>>>> 496ef86... refactor: decouple gotty app with terminal backends
	if err != nil {
		exit(err, 3)
	}

	cmd.Flags = append(
		cliFlags,
		cli.StringFlag{
			Name:   "config",
			Value:  "~/.gotty",
			Usage:  "Config file path",
			EnvVar: "GOTTY_CONFIG",
		},
	)

	cmd.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			msg := "Error: No command given."
			cli.ShowAppHelp(c)
			exit(fmt.Errorf(msg), 1)
		}

		configFile := c.String("config")
		_, err := os.Stat(utils.ExpandHomeDir(configFile))
		if configFile != "~/.gotty" || !os.IsNotExist(err) {
			if err := utils.ApplyConfigFile(configFile, appOptions, backendOptions); err != nil {
				exit(err, 2)
			}
		}

		utils.ApplyFlags(cliFlags, flagMappings, c, appOptions, backendOptions)

<<<<<<< HEAD
<<<<<<< HEAD
		if c.IsSet("credential") {
			options.EnableBasicAuth = true
=======
		options.EnableBasicAuth = c.IsSet("credential")
		options.EnableTLSClientAuth = c.IsSet("tls-ca-crt")
=======
		appOptions.EnableBasicAuth = c.IsSet("credential")
		appOptions.EnableTLSClientAuth = c.IsSet("tls-ca-crt")
>>>>>>> 496ef86... refactor: decouple gotty app with terminal backends

		if err := app.CheckConfig(appOptions); err != nil {
			exit(err, 6)
>>>>>>> d71e2fc... generate falgs based on struct options instead of defining them externally
		}

		manager, err := ptycommand.NewCommandClientContextManager(c.Args(), backendOptions)
		if err != nil {
			exit(err, 3)
		}
		app, err := app.New(manager, appOptions)
		if err != nil {
			exit(err, 3)
		}

		registerSignals(app)

		err = app.Run()
		if err != nil {
			exit(err, 4)
		}
	}
	cmd.Run(os.Args)
}

func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func registerSignals(app *app.App) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		for {
			s := <-sigChan
			switch s {
			case syscall.SIGINT, syscall.SIGTERM:
				if app.Exit() {
					fmt.Println("Send ^C to force exit.")
				} else {
					os.Exit(5)
				}
			}
		}
	}()
}
