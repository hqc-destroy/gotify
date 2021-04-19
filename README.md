# ![](https://raw.githubusercontent.com/sorenisanerd/gotty/master/resources/favicon.png) GoTTY - Share your terminal as a web application

[![GitHub release](http://img.shields.io/github/release/sorenisanerd/gotty.svg?style=flat-square)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/sorenisanerd/gotty/releases
[license]: https://github.com/sorenisanerd/gotty/blob/master/LICENSE

GoTTY is a simple command line tool that turns your CLI tools into web applications.

![Screenshot](https://raw.githubusercontent.com/sorenisanerd/gotty/master/screenshot.gif)

# Installation

<<<<<<< HEAD
<<<<<<< HEAD
Download the latest stable binary file from the [Releases](https://github.com/sorenisanerd/gotty/releases) page. Note that the release marked `Pre-release` is built for testing purpose, which can include unstable or breaking changes. Download a release marked [Latest release](https://github.com/sorenisanerd/gotty/releases/latest) for a stable build.
=======
Download the latest stable binary file from the [Releases](https://github.com/yudai/gotty/releases) page. Note that the release marked `Pre-release` is built for testing purpose, which can include unstable or breaking changes. Download a release marked [Latest release](https://github.com/yudai/gotty/releases/latest) for a stable build.
>>>>>>> 4876619... Fix typo in readme
=======
## From release page

You can download the latest stable binary file from the [Releases](https://github.com/sorenisanerd/gotty/releases) page. Note that the release marked `Pre-release` is built for testing purpose, which can include unstable or breaking changes. Download a release marked [Latest release](https://github.com/sorenisanerd/gotty/releases/latest) for a stable build.
>>>>>>> 8c95f33... Update documentation

(Files named with `darwin_amd64` are for Mac OS X users)

## Homebrew Installation

You can install GoTTY with [Homebrew](http://brew.sh/) as well.

```sh
$ brew install sorenisanerd/gotty/gotty
```

## `go get` Installation (Development)

If you have a Go language environment, you can install GoTTY with the `go get` command. However, this command builds a binary file from the latest master branch, which can include unstable or breaking changes. GoTTY requires go1.9 or later.

```sh
$ go get github.com/sorenisanerd/gotty
```

# Usage

```
Usage: gotty [options] <command> [<arguments...>]
```

Run `gotty` with your preferred command as its arguments (e.g. `gotty top`).

By default, GoTTY starts a web server at port 8080. Open the URL on your web browser and you can see the running command as if it were running on your terminal.

## Options

```
<<<<<<< HEAD
<<<<<<< HEAD
--address, -a                                                IP address to listen [$GOTTY_ADDRESS]
--port, -p "8080"                                            Port number to listen [$GOTTY_PORT]
--path, -m ""                                                Base path
--permit-write, -w                                           Permit clients to write to the TTY (BE CAREFUL) [$GOTTY_PERMIT_WRITE]
--credential, -c                                             Credential for Basic Authentication (ex: user:pass, default disabled) [$GOTTY_CREDENTIAL]
--random-url, -r                                             Add a random string to the URL [$GOTTY_RANDOM_URL]
--random-url-length "8"                                      Random URL length [$GOTTY_RANDOM_URL_LENGTH]
--tls, -t                                                    Enable TLS/SSL [$GOTTY_TLS]
<<<<<<< HEAD
--tls-crt "~/.gotty.key"                                     TLS/SSL crt file path [$GOTTY_TLS_CRT]
--tls-key "~/.gotty.crt"                                     TLS/SSL key file path [$GOTTY_TLS_KEY]
--client, -C                                                 Enable Client Certificate [$GOTTY_CLIENT]
--client-ca-file "~/.gotty.ca.crt"                           Client CA certificate file [$GOTTY_CLIENT_CA_FILE]
--client-verify                                              Enable verification of client certificate [$GOTTY_CLIENT_VERIFY]
--index                                                      Custom index file [$GOTTY_INDEX]
=======
--tls-crt "~/.gotty.crt"                                     TLS/SSL certificate file path [$GOTTY_TLS_CRT]
--tls-key "~/.gotty.key"                                     TLS/SSL key file path [$GOTTY_TLS_KEY]
--tls-ca-crt "~/.gotty.ca.crt"                               TLS/SSL CA certificate file for client certifications [$GOTTY_TLS_CA_CRT]
--index                                                      Custom index.html file [$GOTTY_INDEX]
>>>>>>> 6a6d0e1... Add `--close-signal` option to README
--title-format "GoTTY - {{ .Command }} ({{ .Hostname }})"    Title format of browser window [$GOTTY_TITLE_FORMAT]
--reconnect                                                  Enable reconnection [$GOTTY_RECONNECT]
--reconnect-time "10"                                        Time to reconnect [$GOTTY_RECONNECT_TIME]
--timeout "0"                                                Timeout seconds for waiting a client (0 to disable) [$GOTTY_TIMEOUT]
--max-connection "0"                                         Set the maximum number of simultaneous connections (0 to disable)
--once                                                       Accept only one client and exit on disconnection [$GOTTY_ONCE]
--permit-arguments                                           Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB) [$GOTTY_PERMIT_ARGUMENTS]
--close-signal "1"                                           Signal sent to the command process when gotty close it (default: SIGHUP) [$GOTTY_CLOSE_SIGNAL]
--config "~/.gotty"                                          Config file path [$GOTTY_CONFIG]
--version, -v                                                print the version
=======
--address value, -a value     IP address to listen (default: "0.0.0.0") [$GOTTY_ADDRESS]
--port value, -p value        Port number to listen (default: "8080") [$GOTTY_PORT]
--permit-write, -w            Permit clients to write to the TTY (BE CAREFUL) [$GOTTY_PERMIT_WRITE]
--credential value, -c value  Credential for Basic Authentication (ex: user:pass, default disabled) [$GOTTY_CREDENTIAL]
--random-url, -r              Add a random string to the URL [$GOTTY_RANDOM_URL]
--random-url-length value     Random URL length (default: 8) [$GOTTY_RANDOM_URL_LENGTH]
--tls, -t                     Enable TLS/SSL [$GOTTY_TLS]
--tls-crt value               TLS/SSL certificate file path (default: "~/.gotty.crt") [$GOTTY_TLS_CRT]
--tls-key value               TLS/SSL key file path (default: "~/.gotty.key") [$GOTTY_TLS_KEY]
--tls-ca-crt value            TLS/SSL CA certificate file for client certifications (default: "~/.gotty.ca.crt") [$GOTTY_TLS_CA_CRT]
--index value                 Custom index.html file [$GOTTY_INDEX]
--title-format value          Title format of browser window (default: "{{ .command }}@{{ .hostname }}") [$GOTTY_TITLE_FORMAT]
--reconnect                   Enable reconnection [$GOTTY_RECONNECT]
--reconnect-time value        Time to reconnect (default: 10) [$GOTTY_RECONNECT_TIME]
--max-connection value        Maximum connection to gotty (default: 0) [$GOTTY_MAX_CONNECTION]
--once                        Accept only one client and exit on disconnection [$GOTTY_ONCE]
--timeout value               Timeout seconds for waiting a client(0 to disable) (default: 0) [$GOTTY_TIMEOUT]
--permit-arguments            Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB) [$GOTTY_PERMIT_ARGUMENTS]
--width value                 Static width of the screen, 0(default) means dynamically resize (default: 0) [$GOTTY_WIDTH]
--height value                Static height of the screen, 0(default) means dynamically resize (default: 0) [$GOTTY_HEIGHT]
--ws-origin value             A regular expression that matches origin URLs to be accepted by WebSocket. No cross origin requests are acceptable by default [$GOTTY_WS_ORIGIN]
--term value                  Terminal name to use on the browser, one of xterm or hterm. (default: "xterm") [$GOTTY_TERM]
--close-signal value          Signal sent to the command process when gotty close it (default: SIGHUP) (default: 1) [$GOTTY_CLOSE_SIGNAL]
--close-timeout value         Time in seconds to force kill process after client is disconnected (default: -1) (default: -1) [$GOTTY_CLOSE_TIMEOUT]
--config value                Config file path (default: "~/.gotty") [$GOTTY_CONFIG]
--version, -v                 print the version
>>>>>>> 79e1328... Update README.md
=======
   --address value, -a value     IP address to listen (default: "0.0.0.0") [$GOTTY_ADDRESS]
   --port value, -p value        Port number to liten (default: "8080") [$GOTTY_PORT]
   --path value, -m value        Base path (default: "/") [$GOTTY_PATH]
   --permit-write, -w            Permit clients to write to the TTY (BE CAREFUL) (default: false) [$GOTTY_PERMIT_WRITE]
   --credential value, -c value  Credential for Basic Authentication (ex: user:pass, default disabled) [$GOTTY_CREDENTIAL]
   --random-url, -r              Add a random string to the URL (default: false) [$GOTTY_RANDOM_URL]
   --random-url-length value     Random URL length (default: 8) [$GOTTY_RANDOM_URL_LENGTH]
   --tls, -t                     Enable TLS/SSL (default: false) [$GOTTY_TLS]
   --tls-crt value               TLS/SSL certificate file path (default: "~/.gotty.crt") [$GOTTY_TLS_CRT]
   --tls-key value               TLS/SSL key file path (default: "~/.gotty.key") [$GOTTY_TLS_KEY]
   --tls-ca-crt value            TLS/SSL CA certificate file for client certifications (default: "~/.gotty.ca.crt") [$GOTTY_TLS_CA_CRT]
   --index value                 Custom index.html file [$GOTTY_INDEX]
   --title-format value          Title format of browser window (default: "{{ .command }}@{{ .hostname }}") [$GOTTY_TITLE_FORMAT]
   --reconnect                   Enable reconnection (default: false) [$GOTTY_RECONNECT]
   --reconnect-time value        Time to reconnect (default: 10) [$GOTTY_RECONNECT_TIME]
   --max-connection value        Maximum connection to gotty (default: 0) [$GOTTY_MAX_CONNECTION]
   --once                        Accept only one client and exit on disconnection (default: false) [$GOTTY_ONCE]
   --timeout value               Timeout seconds for waiting a client(0 to disable) (default: 0) [$GOTTY_TIMEOUT]
   --permit-arguments            Permit clients to send command line arguments in URL (e.g. http://example.com:8080/?arg=AAA&arg=BBB) (default: false) [$GOTTY_PERMIT_ARGUMENTS]
   --width value                 Static width of the screen, 0(default) means dynamically resize (default: 0) [$GOTTY_WIDTH]
   --height value                Static height of the screen, 0(default) means dynamically resize (default: 0) [$GOTTY_HEIGHT]
   --ws-origin value             A regular expression that matches origin URLs to be accepted by WebSocket. No cross origin requests are acceptable by default [$GOTTY_WS_ORIGIN]
   --term value                  Terminal name to use on the browser, one of xterm or hterm. (default: "hterm") [$GOTTY_TERM]
   --close-signal value          Signal sent to the command process when gotty close it (default: SIGHUP) (default: 1) [$GOTTY_CLOSE_SIGNAL]
   --close-timeout value         Time in seconds to force kill process after client is disconnected (default: -1) (default: -1) [$GOTTY_CLOSE_TIMEOUT]
   --config value                Config file path (default: "~/.gotty") [$GOTTY_CONFIG]
   --help, -h                    show help (default: false)
   --version, -v                 print the version (default: false)
>>>>>>> f01cf51... Automate options documentation
```

### Config File

You can customize default options and your terminal (hterm) by providing a config file to the `gotty` command. GoTTY loads a profile file at `~/.gotty` by default when it exists.

```
// Listen at port 9000 by default
port = "9000"

// Enable TSL/SSL by default
enable_tls = true

// hterm preferences
// Smaller font and a little bit bluer background color
preferences {
    font_size = 5
    background_color = "rgb(16, 16, 32)"
}
```

See the [`.gotty`](https://github.com/sorenisanerd/gotty/blob/master/.gotty) file in this repository for the list of configuration options.

### Security Options

By default, GoTTY doesn't allow clients to send any keystrokes or commands except terminal window resizing. When you want to permit clients to write input to the TTY, add the `-w` option. However, accepting input from remote clients is dangerous for most commands. When you need interaction with the TTY for some reasons, consider starting GoTTY with tmux or GNU Screen and run your command on it (see "Sharing with Multiple Clients" section for detail).

<<<<<<< HEAD
<<<<<<< HEAD
To restrict client access, you can use the `-c` option to enable the basic authentication. With this option, clients need to input the specified username and password to connect to the GoTTY server. The `-r` option is a little bit casualer way to restrict access. With this option, GoTTY generates a random URL so that only people who know the URL can get access to the server. Note that the credentical will be transmitted between the server and clients in plain text.
=======
To restrict client access, you can use the `-c` option to enable the basic authentication. With this option, clients need to input the specified username and password to connect to the GoTTY server. Note that the credentials will be transmitted between the server and clients in plain text. For more strict authentication, consider the SSL/TLS client certificate authentication described below.
=======
To restrict client access, you can use the `-c` option to enable the basic authentication. With this option, clients need to input the specified username and password to connect to the GoTTY server. Note that the credential will be transmitted between the server and clients in plain text. For more strict authentication, consider the SSL/TLS client certificate authentication described below.
>>>>>>> 7e2a6e7... docs: fix typo

<<<<<<< HEAD
The `-r` option is a little bit casualer way to restrict access. With this option, GoTTY generates a random URL so that only people who know the URL can get access to the server.  
>>>>>>> 2397fb0... fixed small error
=======
The `-r` option is a little bit more casual way to restrict access. With this option, GoTTY generates a random URL so that only people who know the URL can get access to the server.
>>>>>>> 0c3bf61... Document new -m option in README

All traffic between the server and clients are NOT encrypted by default. When you send secret information through GoTTY, we strongly recommend you use the `-t` option which enables TLS/SSL on the session. By default, GoTTY loads the crt and key files placed at `~/.gotty.crt` and `~/.gotty.key`. You can overwrite these file paths with the `--tls-crt` and `--tls-key` options. When you need to generate a self-signed certification file, you can use the `openssl` command.

```sh
openssl req -x509 -nodes -days 9999 -newkey rsa:2048 -keyout ~/.gotty.key -out ~/.gotty.crt
```

For added security you can use an SSL/TLS client certificate by enabling it with the `-C` option (this requires the `-t` or `--tls` flag to be set). This requires all client connecting to provide a valid certificate that can be validated (use the `--client-verify` option to make verification mandatory) against the CA file that is provided via the `--client-ca-file` option.

(NOTE: For Safari uses, see [how to enable self-signed certificates for WebSockets](http://blog.marcon.me/post/24874118286/secure-websockets-safari) when use self-signed certificates)

## Sharing with Multiple Clients

GoTTY starts a new process with the given command when a new client connects to the server. This means users cannot share a single terminal with others by default. However, you can use terminal multiplexers for sharing a single process with multiple clients.
### Screen
After installing GNU screen, start a new session with `screen -S name-for-session` and connect to it with gotty in another terminal window/tab through `screen -x name-for-session`. All commands and activities being done in the first terminal tab/window will now be broadcasted by gotty.
### Tmux
For example, you can start a new tmux session named `gotty` with `top` command by the command below.

```sh
$ gotty tmux new -A -s gotty top
```

This command doesn't allow clients to send keystrokes, however, you can attach the session from your local terminal and run operations like switching the mode of the `top` command. To connect to the tmux session from your terminal, you can use following command.

```sh
$ tmux new -A -s gotty
```

By using terminal multiplexers, you can have the control of your terminal and allow clients to just see your screen.

### Quick Sharing on tmux

To share your current session with others by a shortcut key, you can add a line like below to your `.tmux.conf`.

```
# Start GoTTY in a new window with C-t
bind-key C-t new-window "gotty tmux attach -t `tmux display -p '#S'`"
```

## Playing with Docker

When you want to create a jailed environment for each client, you can use Docker containers like following:

```sh
$ gotty -w docker run -it --rm busybox
```

## Development

You can build a binary using the following commands. Windows is not supported now. go1.9 is required.

```sh
# Install tools
go get github.com/jteeuwen/go-bindata/...
go get github.com/tools/godep

# Build
make
```

To build the frontend part (JS files and other static files), you need `npm`.

## Architecture

GoTTY uses [xterm.js](https://xtermjs.org/) and [hterm](https://groups.google.com/a/chromium.org/forum/#!forum/chromium-hterm) to run a JavaScript based terminal on web browsers. GoTTY itself provides a websocket server that simply relays output from the TTY to clients and receives input from clients and forwards it to the TTY. This hterm + websocket idea is inspired by [Wetty](https://github.com/krishnasrinivas/wetty).

## Alternatives

### Command line client

* [gotty-client](https://github.com/moul/gotty-client): If you want to connect to GoTTY server from your terminal

### Terminal/SSH on Web Browsers

* [Secure Shell (Chrome App)](https://chrome.google.com/webstore/detail/secure-shell/pnhechapfaindjhompbnflcldabbghjo): If you are a chrome user and need a "real" SSH client on your web browser, perhaps the Secure Shell app is what you want
* [Wetty](https://github.com/krishnasrinivas/wetty): Node based web terminal (SSH/login)
* [ttyd](https://tsl0922.github.io/ttyd): C port of GoTTY with CJK and IME support

### Terminal Sharing

* [tmate](http://tmate.io/): Forked-Tmux based Terminal-Terminal sharing
* [termshare](https://termsha.re): Terminal-Terminal sharing through a HTTP server
* [tmux](https://tmux.github.io/): Tmux itself also supports TTY sharing through SSH)

# License

The MIT License

# Contributors

## Original author

 * Iwasaki Yudai

## Maintainer

 * Søren L. Hansen

## Contributors

 * 0xflotus
 * Anand Patil
 * Andrea Lusuardi - uovobw
 * Andy Skelton
 * Artem Medvedev
 * Blake Jennings
 * Christian Jensen
 * Christopher Wilkinson
 * Cyrus
 * David Horsley
 * Fazal Majid
 * freakhill
 * fredster33
 * Jan-Willem Korver
 * Jason Cooke
 * Johan Gall
 * Korenevskiy Denis
 * Lin
 * Manfred Touron
 * Massimiliano Stucchi
 * mattn
 * Mikhail f. Shiryaev
 * Quentin Perez
 * Richard Metzler
 * Robert Bittle
 * Sebastian Haas
 * shingt
 * Shoji Ihara
 * Shuanglei Tao
 * Stephan van Ellewee
 * The Gitter Badger
 * Xinyun Zhou
 * Yifa Zhang
 * yogesh singh
 * zlji
