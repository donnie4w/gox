package main

import (
	"flag"
	"fmt"
	"github.com/donnie4w/gofer/keystore"
	"github.com/donnie4w/simplelog/logging"
	"github.com/donnie4w/tlnet"
)

func main() {
	logging.SetFormat(logging.FORMAT_DATE | logging.FORMAT_TIME)
	port := 8080
	nolog := false
	tls := false
	flag.IntVar(&port, "p", 8080, "Listen port")
	flag.BoolVar(&nolog, "nolog", false, "not print log")
	flag.BoolVar(&tls, "tls", false, "tls")
	flag.Parse()
	tl := tlnet.NewTlnet()
	filter := newFilter()
	if nolog {
		filter = nil
	}
	tl.HandleStaticWithFilter("/", "./", filter, nil)
	logging.Debug("params:", "port[", port, "]", "nolog[", nolog, "]", "tls[", tls, "]")
	logging.Debug("Listen :", port)
	if tls && tl.HttpStartTlsBytes(fmt.Sprint(":", port), []byte(keystore.ServerCrt), []byte(keystore.ServerKey)) != nil {
		logging.Error("http Start tls failed")
	} else if err := tl.HttpStart(fmt.Sprint(":", port)); err != nil {
		logging.Error("http Start Error:", err)
	}
}

func newFilter() (f *tlnet.Filter) {
	f = tlnet.NewFilter()
	f.AddIntercept(".*?", func(hc *tlnet.HttpContext) bool {
		logging.Debug(hc.Request().RequestURI)
		return false
	})
	return
}
