# pitt
Go package manager

## Install
```bash
$ go get github.com/redwhitelab/pitt
```

## Usage
```bash
NAME:
   pitt - go project managment tool

USAGE:
   pitt [global options] command [command options] [arguments...]

VERSION:
   dev

COMMANDS:
   new, n	create new Go project
   search, s	search packages
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```

### Search Packages

```bash
$ pitt search http2
search query:  http2
github.com/bradfitz/http2 - old repo for HTTP/2 support for Go (see README for new home) (1640 stars)
github.com/jxck/http2 - http2 implementation in Go (69 stars)
github.com/c0nrad/http2fuzz - HTTP/2 fuzzer written in Golang (49 stars)
github.com/xyproto/http2check - Utility for checking if a web server supports HTTP/2 (16 stars)
github.com/bradfitz/talk-http2go - HTTP/2 Go Talk (10 stars)
github.com/google/http2preload - A Go package for manipulating HTTP/2 preload, supporting Google App Engine. (5 stars)
github.com/moul/http2curl - :triangular_ruler: Convert Golang's http.Request to CURL command line (4 stars)
github.com/aleasoluciones/http2amqp -  (3 stars)
github.com/xtaci/http2tun - openvpn tunnel (2 stars)
github.com/nightgunner5/http2 - golang net/http helpers (1 stars)
github.com/nekolunar/http2 - HTTP/2 for The Go programming language (1 stars)
github.com/zoncoen/minimum_http2 -  (1 stars)
github.com/tildedave/go-http2-impl - Learning Go, implementing HTTP 2.0 (1 stars)
github.com/tatsuhiro-t/go-http2-hpack - HTTP/2 HPACK implementation in golang (1 stars)
github.com/dougwatson/http2amqp - A high performance HTTP to AMQP (rabbitmq) gateway (1 stars)
github.com/aleasoluciones/http2snmp -  (1 stars)
github.com/rollbackup/http2 - net/http extensions (0 stars)
github.com/honsty/http2 -  (0 stars)
github.com/mrsavinov/http2 - http2 server (0 stars)
github.com/citysir/http2proxy - http2 proxy  (0 stars)
github.com/wellington/playground - Play on wellington http2! (0 stars)
github.com/sdd330/gohttp2_dockerfile - Docker file for http2 server (0 stars)
github.com/tkyshm/apns2 - APNs provider server by http2. (0 stars)
github.com/k2wanko/go-http2-example -  (0 stars)
github.com/channelmeter/http2-client -  (0 stars)
github.com/whyrusleeping/bufpipe - a buffered pipe implementation from bradfitz/http2 (0 stars)
github.com/ami-gs/simplehttp2 - Simple implementation of HTTP2 in Golang (0 stars)
github.com/common-benchmarks/golang-func-vs-rpc-vs-protobufhttp2 - Benchmark of function vs RPC vs Protobuf RPC with HTTP2 (0 stars)
github.com/stuart-warren/serveit - simple static http2 webserver/proxy (inspired by Caddy) (0 stars)
github.com/ops-console/http200ok -  (0 stars)
```