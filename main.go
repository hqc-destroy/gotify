package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	cli "github.com/urfave/cli/v2"

	"github.com/sorenisanerd/gotty/backend/localcommand"
	"github.com/sorenisanerd/gotty/pkg/homedir"
	"github.com/sorenisanerd/gotty/server"
	"github.com/sorenisanerd/gotty/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "gotty"
	app.Version = Version + "+" + CommitID
	app.Usage = "Share your terminal as a web application"
<<<<<<< HEAD
	app.HideHelp = true
	cli.AppHelpTemplate = helpTemplate

<<<<<<< HEAD
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
=======
	appOptions := &server.Options{}
>>>>>>> a6133f3... Refactor
=======
	app.HideHelpCommand = true
	appOptions := &server.Options{}

>>>>>>> c3c670b... Switch to urfave/cli/v2... again.
	if err := utils.ApplyDefaultValues(appOptions); err != nil {
		exit(err, 1)
	}
	backendOptions := &localcommand.Options{}
	if err := utils.ApplyDefaultValues(backendOptions); err != nil {
		exit(err, 1)
	}

	cliFlags, flagMappings, err := utils.GenerateFlags(appOptions, backendOptions)
>>>>>>> 496ef86... refactor: decouple gotty app with terminal backends
	if err != nil {
		exit(err, 3)
	}

	app.Flags = append(
		cliFlags,
		&cli.StringFlag{
			Name:    "config",
			Value:   "~/.gotty",
			Usage:   "Config file path",
			EnvVars: []string{"GOTTY_CONFIG"},
		},
	)

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			msg := "Error: No command given."
			cli.ShowAppHelp(c)
			exit(fmt.Errorf(msg), 1)
		}

		configFile := c.String("config")
		_, err := os.Stat(homedir.Expand(configFile))
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

		err = appOptions.Validate()
		if err != nil {
			exit(err, 6)
>>>>>>> d71e2fc... generate falgs based on struct options instead of defining them externally
		}

		args := c.Args()
		factory, err := localcommand.NewFactory(args.First(), args.Tail(), backendOptions)
		if err != nil {
			exit(err, 3)
		}

		hostname, _ := os.Hostname()
		appOptions.TitleVariables = map[string]interface{}{
			"command":  args.First(),
			"argv":     args.Tail(),
			"hostname": hostname,
		}

		srv, err := server.New(factory, appOptions)
		if err != nil {
			exit(err, 3)
		}

		ctx, cancel := context.WithCancel(context.Background())
		gCtx, gCancel := context.WithCancel(context.Background())

		log.Printf("GoTTY is starting with command: %s", strings.Join(args.Slice(), " "))

		errs := make(chan error, 1)
		go func() {
			errs <- srv.Run(ctx, server.WithGracefullContext(gCtx))
		}()
		err = waitSignals(errs, cancel, gCancel)

		if err != nil && err != context.Canceled {
			fmt.Printf("Error: %s\n", err)
			exit(err, 8)
		}

		return nil
	}
	app.Run(os.Args)
}

func exit(err error, code int) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func waitSignals(errs chan error, cancel context.CancelFunc, gracefullCancel context.CancelFunc) error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	select {
	case err := <-errs:
		return err

	case s := <-sigChan:
		switch s {
		case syscall.SIGINT:
			gracefullCancel()
			fmt.Println("C-C to force close")
			select {
			case err := <-errs:
				return err
			case <-sigChan:
				fmt.Println("Force closing...")
				cancel()
				return <-errs
			}
		default:
			cancel()
			return <-errs
		}
	}
}
