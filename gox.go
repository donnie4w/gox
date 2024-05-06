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
	port, nolog, https, root := 8080, false, false, "./"
	flag.IntVar(&port, "port", 8080, "Listen port")
	flag.BoolVar(&nolog, "nolog", false, "not print log")
	flag.BoolVar(&https, "https", false, "use https")
	flag.StringVar(&root, "root", "./", "root path")
	flag.Parse()
	tl := tlnet.NewTlnet()
	tl.HandleStatic("/", root, func(hc *tlnet.HttpContext) {
		if !nolog {
			logging.Debug(hc.Request().RequestURI)
		}
	})
	logging.Info("Gox Start and params:", "port[", port, "]", "nolog[", nolog, "]", "https[", https, "]", "root[", root, "]")
	var err error
	if https {
		err = tl.HttpStartTlsBytes(fmt.Sprint(":", port), []byte(keystore.ServerCrt), []byte(keystore.ServerKey))
	} else {
		err = tl.HttpStart(fmt.Sprint(":", port))
	}
	logging.Error("Gox Start failed:", err)
}
