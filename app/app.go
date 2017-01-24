package app

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/yudai/gotty/backends"
	"github.com/yudai/gotty/utils"

	"github.com/braintree/manners"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/websocket"
	"github.com/yudai/umutex"
)

type InitMessage struct {
	Arguments string `json:"Arguments,omitempty"`
	AuthToken string `json:"AuthToken,omitempty"`
}

type App struct {
	manager backends.ClientContextManager
	options *Options

	upgrader *websocket.Upgrader
	server   *manners.GracefulServer

	onceMutex *umutex.UnblockingMutex
	timer     *time.Timer

	// clientContext writes concurrently
	// Use atomic operations.
	connections *int64
}

type Options struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	Address          string                 `hcl:"address"`
	Port             string                 `hcl:"port"`
	PermitWrite      bool                   `hcl:"permit_write"`
	EnableBasicAuth  bool                   `hcl:"enable_basic_auth"`
	Credential       string                 `hcl:"credential"`
	EnableRandomUrl  bool                   `hcl:"enable_random_url"`
	RandomUrlLength  int                    `hcl:"random_url_length"`
	IndexFile        string                 `hcl:"index_file"`
	EnableTLS        bool                   `hcl:"enable_tls"`
	TLSCrtFile       string                 `hcl:"tls_crt_file"`
	TLSKeyFile       string                 `hcl:"tls_key_file"`
	VerifyClientCert bool                   `hcl:"verify_client_cert"`
	ClientCAs        []string               `hcl:"client_cas"`
	TitleFormat      string                 `hcl:"title_format"`
	EnableReconnect  bool                   `hcl:"enable_reconnect"`
	ReconnectTime    int                    `hcl:"reconnect_time"`
	Once             bool                   `hcl:"once"`
	Preferences      map[string]interface{} `hcl:"preferences"`
=======
	Address                             string                 `hcl:"address"`
	Port                                string                 `hcl:"port"`
	PermitWrite                         bool                   `hcl:"permit_write"`
	EnableBasicAuth                     bool                   `hcl:"enable_basic_auth"`
	Credential                          string                 `hcl:"credential"`
	EnableRandomUrl                     bool                   `hcl:"enable_random_url"`
	RandomUrlLength                     int                    `hcl:"random_url_length"`
	IndexFile                           string                 `hcl:"index_file"`
	EnableTLS                           bool                   `hcl:"enable_tls"`
	TLSCrtFile                          string                 `hcl:"tls_crt_file"`
	TLSKeyFile                          string                 `hcl:"tls_key_file"`
	EnableClientCertificate             bool                   `hcl:"enable_client_certificate"`
	ClientCAFile                        string                 `hcl:"client_ca_file"`
	EnableClientCertificateVerification bool                   `hcl:"enable_client_certificate_verification"`
	TitleFormat                         string                 `hcl:"title_format"`
	EnableReconnect                     bool                   `hcl:"enable_reconnect"`
	ReconnectTime                       int                    `hcl:"reconnect_time"`
	Once                                bool                   `hcl:"once"`
	Preferences                         map[string]interface{} `hcl:"preferences"`
>>>>>>> 7321b43... Add client certificate fields to the configuration struct
=======
	Address             string                 `hcl:"address"`
	Port                string                 `hcl:"port"`
	Path								string								 `hcl:"path"`
	PermitWrite         bool                   `hcl:"permit_write"`
	EnableBasicAuth     bool                   `hcl:"enable_basic_auth"`
	Credential          string                 `hcl:"credential"`
	EnableRandomUrl     bool                   `hcl:"enable_random_url"`
	RandomUrlLength     int                    `hcl:"random_url_length"`
	IndexFile           string                 `hcl:"index_file"`
	EnableTLS           bool                   `hcl:"enable_tls"`
	TLSCrtFile          string                 `hcl:"tls_crt_file"`
	TLSKeyFile          string                 `hcl:"tls_key_file"`
	EnableTLSClientAuth bool                   `hcl:"enable_tls_client_auth"`
	TLSCACrtFile        string                 `hcl:"tls_ca_crt_file"`
	TitleFormat         string                 `hcl:"title_format"`
	EnableReconnect     bool                   `hcl:"enable_reconnect"`
	ReconnectTime       int                    `hcl:"reconnect_time"`
	MaxConnection       int                    `hcl:"max_connection"`
	Once                bool                   `hcl:"once"`
	Timeout             int                    `hcl:"timeout"`
	PermitArguments     bool                   `hcl:"permit_arguments"`
<<<<<<< HEAD
<<<<<<< HEAD
	Preferences         map[string]interface{} `hcl:"preferences"`
>>>>>>> a4e77b2... Added handling of —permit-arguments option
=======
=======
	CloseSignal         int                    `hcl:"close_signal"`
>>>>>>> 888fe87... Add configuration to modify signal sent to child process when close it
	Preferences         HtermPrefernces        `hcl:"preferences"`
	RawPreferences      map[string]interface{} `hcl:"preferences"`
<<<<<<< HEAD
>>>>>>> 589ec6b... Handle hterm preferences with better care
=======
	Width               int                    `hcl:"width"`
	Height              int                    `hcl:"height"`
>>>>>>> 8fd09cd... Add an option to disable client window resizes
=======
	Address             string                 `hcl:"address" flagName:"address" flagSName:"a" flagDescribe:"IP address to listen" default:""`
	Port                string                 `hcl:"port" flagName:"port" flagSName:"p" flagDescribe:"Port number to liten" default:"8080"`
	PermitWrite         bool                   `hcl:"permit_write" flagName:"permit-write" flagSName:"w" flagDescribe:"Permit clients to write to the TTY (BE CAREFUL)" default:"false"`
	EnableBasicAuth     bool                   `hcl:"enable_basic_auth" default:"false"`
	Credential          string                 `hcl:"credential" flagName:"credential" flagSName:"c" flagDescribe:"Credential for Basic Authentication (ex: user:pass, default disabled)" default:""`
	EnableRandomUrl     bool                   `hcl:"enable_random_url flagName:"random-url" flagSName:"r" flagDescribe:"Add a random string to the URL"" default:"false"`
	RandomUrlLength     int                    `hcl:"random_url_length" flagName:"random-url-length" flagDescribe:"Random URL length" default:"8"`
	IndexFile           string                 `hcl:"index_file" flagName:"index" flagDescribe:"Custom index.html file" default:""`
	EnableTLS           bool                   `hcl:"enable_tls" flagName:"tls" flagSName:"t" flagDescribe:"Enable TLS/SSL" default:"false"`
	TLSCrtFile          string                 `hcl:"tls_crt_file" flagName:"tls-crt" flagDescribe:"TLS/SSL certificate file path" default:"~/.gotty.crt"`
	TLSKeyFile          string                 `hcl:"tls_key_file" flagName:"tls-key" flagDescribe:"TLS/SSL key file path" default:"~/.gotty.key"`
	EnableTLSClientAuth bool                   `hcl:"enable_tls_client_auth" default:"false"`
	TLSCACrtFile        string                 `hcl:"tls_ca_crt_file" flagName:"tls-ca-crt" flagDescribe:"TLS/SSL CA certificate file for client certifications" default:"~/.gotty.ca.crt"`
	EnableReconnect     bool                   `hcl:"enable_reconnect" flagName:"reconnect" flagDescribe:"Enable reconnection" default:"false"`
	ReconnectTime       int                    `hcl:"reconnect_time" flagName:"reconnect-time" flagDescribe:"Time to reconnect" default:"10"`
	MaxConnection       int                    `hcl:"max_connection" flagName:"max-connection" flagDescribe:"Maximum connection to gotty" default:"0"`
	Once                bool                   `hcl:"once" flagName:"once" flagDescribe:"Accept only one client and exit on disconnection" default:"false"`
	Timeout             int                    `hcl:"timeout" flagName:"timeout" flagDescribe:"Timeout seconds for waiting a client(0 to disable)" default:"0"`
	PermitArguments     bool                   `hcl:"permit_arguments" flagName:"permit-arguments" flagDescribe:"Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB)" default:"true"`
	Preferences         HtermPrefernces        `hcl:"preferences"`
	RawPreferences      map[string]interface{} `hcl:"preferences"`
	Width               int                    `hcl:"width" flagName:"width" flagDescribe:"Static width of the screen, 0(default) means dynamically resize" default:"0"`
	Height              int                    `hcl:"height" flagName:"height" flagDescribe:"Static height of the screen, 0(default) means dynamically resize" default:"0"`
>>>>>>> d71e2fc... generate falgs based on struct options instead of defining them externally
}

var Version = "1.0.0"

<<<<<<< HEAD
<<<<<<< HEAD
var DefaultOptions = Options{
<<<<<<< HEAD
<<<<<<< HEAD
	Address:          "",
	Port:             "8080",
	PermitWrite:      false,
	EnableBasicAuth:  false,
	Credential:       "",
	EnableRandomUrl:  false,
	RandomUrlLength:  8,
	IndexFile:        "",
	EnableTLS:        false,
	TLSCrtFile:       "~/.gotty.crt",
	TLSKeyFile:       "~/.gotty.key",
	VerifyClientCert: false,
	ClientCAs:        []string{},
	TitleFormat:      "GoTTY - {{ .Command }} ({{ .Hostname }})",
	EnableReconnect:  false,
	ReconnectTime:    10,
	Once:             false,
	Preferences:      make(map[string]interface{}),
=======
	Address:                             "",
	Port:                                "8080",
	PermitWrite:                         false,
	EnableBasicAuth:                     false,
	Credential:                          "",
	EnableRandomUrl:                     false,
	RandomUrlLength:                     8,
	IndexFile:                           "",
	EnableTLS:                           false,
	TLSCrtFile:                          "~/.gotty.crt",
	TLSKeyFile:                          "~/.gotty.key",
	EnableClientCertificate:             false,
	ClientCAFile:                        "~/.gotty.ca.crt",
	EnableClientCertificateVerification: false,
	TitleFormat:                         "GoTTY - {{ .Command }} ({{ .Hostname }})",
	EnableReconnect:                     false,
	ReconnectTime:                       10,
	Once:                                false,
	Preferences:                         make(map[string]interface{}),
>>>>>>> 7321b43... Add client certificate fields to the configuration struct
=======
	Address:             "",
	Port:                "8080",
	Path:								 "",
	PermitWrite:         false,
	EnableBasicAuth:     false,
	Credential:          "",
	EnableRandomUrl:     false,
	RandomUrlLength:     8,
	IndexFile:           "",
	EnableTLS:           false,
	TLSCrtFile:          "~/.gotty.crt",
	TLSKeyFile:          "~/.gotty.key",
	EnableTLSClientAuth: false,
	TLSCACrtFile:        "~/.gotty.ca.crt",
	TitleFormat:         "GoTTY - {{ .Command }} ({{ .Hostname }})",
	EnableReconnect:     false,
	ReconnectTime:       10,
	MaxConnection:       0,
	Once:                false,
	CloseSignal:         1, // syscall.SIGHUP
	Preferences:         HtermPrefernces{},
<<<<<<< HEAD
>>>>>>> 589ec6b... Handle hterm preferences with better care
=======
	Width:               0,
	Height:              0,
>>>>>>> 8fd09cd... Add an option to disable client window resizes
}

=======
>>>>>>> d71e2fc... generate falgs based on struct options instead of defining them externally
func New(command []string, options *Options) (*App, error) {
	titleTemplate, err := template.New("title").Parse(options.TitleFormat)
	if err != nil {
		return nil, errors.New("Title format string syntax error")
	}

=======
func New(manager backends.ClientContextManager, options *Options) (*App, error) {
>>>>>>> 496ef86... refactor: decouple gotty app with terminal backends
	connections := int64(0)
	return &App{
		manager: manager,
		options: options,

		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			Subprotocols:    []string{"gotty"},
		},
		onceMutex:   umutex.New(),
		connections: &connections,
	}, nil
}

<<<<<<< HEAD
func ApplyConfigFile(options *Options, filePath string) error {
	filePath = ExpandHomeDir(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	fileString := []byte{}
	log.Printf("Loading config file at: %s", filePath)
	fileString, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := hcl.Decode(options, string(fileString)); err != nil {
		return err
	}

=======
func CheckConfig(options *Options) error {
	if options.EnableTLSClientAuth && !options.EnableTLS {
		return errors.New("TLS client authentication is enabled, but TLS is not enabled")
	}
>>>>>>> d71e2fc... generate falgs based on struct options instead of defining them externally
	return nil
}

func (app *App) Run() error {
	if app.options.PermitWrite {
		log.Printf("Permitting clients to write input to the PTY.")
	}

	if app.options.Once {
		log.Printf("Once option is provided, accepting only one client")
	}

	path := app.options.Path
	if app.options.EnableRandomUrl {
		path += "/" + generateRandomString(app.options.RandomUrlLength)
	}

	endpoint := net.JoinHostPort(app.options.Address, app.options.Port)

	wsHandler := http.HandlerFunc(app.handleWS)
	customIndexHandler := http.HandlerFunc(app.handleCustomIndex)
	authTokenHandler := http.HandlerFunc(app.handleAuthToken)
	staticHandler := http.FileServer(
		&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "static"},
	)

	var siteMux = http.NewServeMux()

	if app.options.IndexFile != "" {
		log.Printf("Using index file at " + app.options.IndexFile)
		siteMux.Handle(path+"/", customIndexHandler)
	} else {
		siteMux.Handle(path+"/", http.StripPrefix(path+"/", staticHandler))
	}
	siteMux.Handle(path+"/auth_token.js", authTokenHandler)
	siteMux.Handle(path+"/js/", http.StripPrefix(path+"/", staticHandler))
	siteMux.Handle(path+"/favicon.png", http.StripPrefix(path+"/", staticHandler))

	siteHandler := http.Handler(siteMux)

	if app.options.EnableBasicAuth {
		log.Printf("Using Basic Authentication")
		siteHandler = wrapBasicAuth(siteHandler, app.options.Credential)
	}

	siteHandler = wrapHeaders(siteHandler)

	wsMux := http.NewServeMux()
	wsMux.Handle("/", siteHandler)
	wsMux.Handle(path+"/ws", wsHandler)
	siteHandler = (http.Handler(wsMux))

	siteHandler = wrapLogger(siteHandler)

	scheme := "http"
	if app.options.EnableTLS {
		scheme = "https"
	}
	if app.options.Address != "" {
		log.Printf(
			"URL: %s",
			(&url.URL{Scheme: scheme, Host: endpoint, Path: path + "/"}).String(),
		)
	} else {
		for _, address := range listAddresses() {
			log.Printf(
				"URL: %s",
				(&url.URL{
					Scheme: scheme,
					Host:   net.JoinHostPort(address, app.options.Port),
					Path:   path + "/",
				}).String(),
			)
		}
	}

	serverMaker := func() *http.Server {
		return &http.Server{
			Addr:    endpoint,
			Handler: siteHandler}
	}
	if app.options.VerifyClientCert && app.options.EnableTLS {
		serverMaker = func() *http.Server {
			clientCaPool := x509.NewCertPool()
			for _, path := range app.options.ClientCAs {
				pem, err := ioutil.ReadFile(path)
				if err != nil {
					log.Printf("Could not read pem file at: " + path)
					return nil
				}
				if clientCaPool.AppendCertsFromPEM(pem) {
					log.Printf("Could not parse pem file at: " + path)
					return nil
				}
			}
			return &http.Server{
				Addr:    endpoint,
				Handler: siteHandler,
				TLSConfig: &tls.Config{
					ClientAuth:               tls.RequireAndVerifyClientCert,
					ClientCAs:                clientCaPool,
					PreferServerCipherSuites: true}}
		}
	}

	server := serverMaker()
	if server == nil {
		log.Printf("Failed to build server.")
		return errors.New("Failed to build server.")
	}

	var err error
	app.server = manners.NewWithServer(
		server,
	)
<<<<<<< HEAD
=======

	if app.options.Timeout > 0 {
		app.timer = time.NewTimer(time.Duration(app.options.Timeout) * time.Second)
		go func() {
			<-app.timer.C
			app.Exit()
		}()
	}

>>>>>>> 8c9433f... Add timeout option
	if app.options.EnableTLS {
		crtFile := utils.ExpandHomeDir(app.options.TLSCrtFile)
		keyFile := utils.ExpandHomeDir(app.options.TLSKeyFile)
		log.Printf("TLS crt file: " + crtFile)
		log.Printf("TLS key file: " + keyFile)
		if app.options.EnableClientCertificate {
			caFile := ExpandHomeDir(app.options.ClientCAFile)
			log.Printf("Client CA file: " + caFile)
			caCert, err := ioutil.ReadFile(caFile)
			if err != nil {
				return errors.New("Cannot open CA file " + caFile)
			}
			caCertPool := x509.NewCertPool()
			if !caCertPool.AppendCertsFromPEM(caCert) {
				return errors.New("Cannot parse CA file data in " + caFile)
			}
			tlsVerifyPolicy := tls.RequireAnyClientCert
			if app.options.EnableClientCertificateVerification {
				log.Print("Enabling verification of client certificate")
				tlsVerifyPolicy = tls.RequireAndVerifyClientCert
			}
			tlsConfig := &tls.Config{
				ClientCAs:  caCertPool,
				ClientAuth: tlsVerifyPolicy,
			}
			app.server.TLSConfig = tlsConfig
		}
		err = app.server.ListenAndServeTLS(crtFile, keyFile)
	} else {
		err = app.server.ListenAndServe()
	}
	if err != nil {
		return err
	}

	log.Printf("Exiting...")

	return nil
}

<<<<<<< HEAD
=======
func (app *App) makeServer(addr string, handler *http.Handler) (*http.Server, error) {
	server := &http.Server{
		Addr:    addr,
		Handler: *handler,
	}

	if app.options.EnableTLSClientAuth {
		caFile := utils.ExpandHomeDir(app.options.TLSCACrtFile)
		log.Printf("CA file: " + caFile)
		caCert, err := ioutil.ReadFile(caFile)
		if err != nil {
			return nil, errors.New("Could not open CA crt file " + caFile)
		}
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			return nil, errors.New("Could not parse CA crt file data in " + caFile)
		}
		tlsConfig := &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
		server.TLSConfig = tlsConfig
	}

	return server, nil
}

func (app *App) stopTimer() {
	if app.options.Timeout > 0 {
		app.timer.Stop()
	}
}

func (app *App) restartTimer() {
	if app.options.Timeout > 0 {
		app.timer.Reset(time.Duration(app.options.Timeout) * time.Second)
	}
}

>>>>>>> 8c9433f... Add timeout option
func (app *App) handleWS(w http.ResponseWriter, r *http.Request) {
	app.stopTimer()
	connections := atomic.AddInt64(app.connections, 1)
	defer func() {
		connections := atomic.AddInt64(app.connections, -1)

		if app.options.MaxConnection != 0 {
			log.Printf("Connection closed: %s, connections: %d/%d",
				r.RemoteAddr, connections, app.options.MaxConnection)
		} else {
			log.Printf("Connection closed: %s, connections: %d",
				r.RemoteAddr, connections)
		}

		if connections == 0 {
			app.restartTimer()
		}
	}()

	if int64(app.options.MaxConnection) != 0 {
		if connections > int64(app.options.MaxConnection) {
			log.Printf("Reached max connection: %d", app.options.MaxConnection)
			return
		}
	}
	log.Printf("New client connected: %s", r.RemoteAddr)

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	conn, err := app.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Failed to upgrade connection: " + err.Error())
		return
	}
	defer conn.Close()

	_, stream, err := conn.ReadMessage()
	if err != nil {
		log.Print("Failed to authenticate websocket connection")
		return
	}
	var init InitMessage

	err = json.Unmarshal(stream, &init)
	if err != nil {
		log.Printf("Failed to parse init message %v", err)
		return
	}
	if init.AuthToken != app.options.Credential {
		log.Print("Failed to authenticate websocket connection")
		return
	}

	var queryPath string
	if app.options.PermitArguments && init.Arguments != "" {
		queryPath = init.Arguments
	} else {
		queryPath = "?"
	}

	query, err := url.Parse(queryPath)
	if err != nil {
		log.Print("Failed to parse arguments")
		return
	}
	params := query.Query()
	ctx, err := app.manager.New(params)
	if err != nil {
		log.Printf("Failed to new client context %v", err)
		return
	}

	app.server.StartRoutine()
	defer app.server.FinishRoutine()

	if app.options.Once {
		if app.onceMutex.TryLock() { // no unlock required, it will die soon
			log.Printf("Last client accepted, closing the listener.")
			app.server.Close()
		} else {
			log.Printf("Server is already closing.")
			conn.Close()
			return
		}
	}

	context := &clientContext{app: app, connection: conn, writeMutex: &sync.Mutex{}, ClientContext: ctx}
	context.goHandleClient()
}

func (app *App) handleCustomIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, utils.ExpandHomeDir(app.options.IndexFile))
}

func (app *App) handleAuthToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte("var gotty_auth_token = '" + app.options.Credential + "';"))
}

func (app *App) Exit() (firstCall bool) {
	if app.server != nil {
		firstCall = app.server.Close()
		if firstCall {
			log.Printf("Received Exit command, waiting for all clients to close sessions...")
		}
		return firstCall
	}
	return true
}

func wrapLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWrapper{w, 200}
		handler.ServeHTTP(rw, r)
		log.Printf("%s %d %s %s", r.RemoteAddr, rw.status, r.Method, r.URL.Path)
	})
}

func wrapHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "GoTTY/"+Version)
		handler.ServeHTTP(w, r)
	})
}

func wrapBasicAuth(handler http.Handler, credential string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(token) != 2 || strings.ToLower(token[0]) != "basic" {
			w.Header().Set("WWW-Authenticate", `Basic realm="GoTTY"`)
			http.Error(w, "Bad Request", http.StatusUnauthorized)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(token[1])
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if credential != string(payload) {
			w.Header().Set("WWW-Authenticate", `Basic realm="GoTTY"`)
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		log.Printf("Basic Authentication Succeeded: %s", r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}

func generateRandomString(length int) string {
	const base = 36
	size := big.NewInt(base)
	n := make([]byte, length)
	for i, _ := range n {
		c, _ := rand.Int(rand.Reader, size)
		n[i] = strconv.FormatInt(c.Int64(), base)[0]
	}
	return string(n)
}

func listAddresses() (addresses []string) {
	ifaces, _ := net.Interfaces()

	addresses = make([]string, 0, len(ifaces))

	for _, iface := range ifaces {
		ifAddrs, _ := iface.Addrs()
		for _, ifAddr := range ifAddrs {
			switch v := ifAddr.(type) {
			case *net.IPNet:
				addresses = append(addresses, v.IP.String())
			case *net.IPAddr:
				addresses = append(addresses, v.IP.String())
			}
		}
	}

	return
}
