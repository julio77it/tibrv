package main

import (
	"flag"
	"time"

	"github.com/julio77it/tibrv"
)

func main() {
	var service, network, daemon, subject, text, transportType string
	flag.StringVar(&service, "service", "", "Tibco RendezVous Bus service")
	flag.StringVar(&network, "network", "", "Tibco RendezVous Bus network")
	flag.StringVar(&daemon, "daemon", "", "Tibco RendezVous Bus daemon")
	flag.StringVar(&subject, "subject", "PING", "Tibco RendezVous listening subject")
	flag.StringVar(&transportType, "type", "net", "Tibco RendezVous transport type [net,cm]")
	flag.StringVar(&text, "text", "ping", "Text message")
	flag.Parse()

	var transport tibrv.RvTransport
	var nettransport tibrv.RvNetTransport
	var cmtransport tibrv.RvCmTransport

	if err := nettransport.Create(tibrv.Service(service), tibrv.Network(network), tibrv.Daemon(daemon)); err != nil {
		panic(err)
	}
	defer nettransport.Destroy()

	if transportType == "net" {
		transport = nettransport
	} else if transportType == "cm" {
		if err := cmtransport.Create(&nettransport); err != nil {
			panic(err)
		}
		transport = cmtransport
	}

	var msg tibrv.RvMessage
	if err := msg.Create(); err != nil {
		panic(err)
	}
	defer msg.Destroy()

	if err := msg.SetString("DATA", text); err != nil {
		panic(err)
	}
	if err := msg.SetSendSubject(subject); err != nil {
		panic(err)
	}

	if err := transport.Send(msg); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
}
