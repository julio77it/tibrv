package main

import (
	"flag"
	"github.com/julio77it/tibrv"
	"time"
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

	nettransport.Create(tibrv.Service(service), tibrv.Network(network), tibrv.Daemon(daemon))
	defer nettransport.Destroy()

	if transportType == "net" {
		transport = nettransport
	} else if transportType == "cm" {
		cmtransport.Create(&nettransport)
		transport = cmtransport
	}

	var msg tibrv.RvMessage
	msg.Create()
	defer msg.Destroy()

	msg.SetString("DATA", text)
	msg.SetSendSubject(subject)

	transport.Send(msg)

	time.Sleep(1 * time.Second)
}
