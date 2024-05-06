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
	port, nolog, https := 8080, false, false
	flag.IntVar(&port, "port", 8080, "Listen port")
	flag.BoolVar(&nolog, "nolog", false, "not print log")
	flag.BoolVar(&https, "https", false, "use https")
	flag.Parse()
	tl := tlnet.NewTlnet()
	var err error
	var filter *tlnet.Filter
	if !nolog {
		filter = newFilter()
	}
	tl.HandleStaticWithFilter("/", "./", filter, nil)
	logging.Info("Gox Start and params:", "port[", port, "]", "nolog[", nolog, "]", "https[", https, "]")
	if https {
		err = tl.HttpStartTlsBytes(fmt.Sprint(":", port), []byte(keystore.ServerCrt), []byte(keystore.ServerKey))
	} else {
		err = tl.HttpStart(fmt.Sprint(":", port))
	}
	logging.Error("Gox Start failed:", err)
}

func newFilter() (f *tlnet.Filter) {
	f = tlnet.NewFilter()
	f.AddIntercept(".*?", func(hc *tlnet.HttpContext) bool {
		logging.Debug(hc.Request().RequestURI)
		return false
	})
	return
}
