package main

import (
	"flag"
	"fmt"
	"github.com/julio77it/tibrv"
	"log"
)

func main() {
	var service, network, daemon, subject, transportType string
	flag.StringVar(&service, "service", "", "Tibco RendezVous Bus service")
	flag.StringVar(&network, "network", "", "Tibco RendezVous Bus network")
	flag.StringVar(&daemon, "daemon", "", "Tibco RendezVous Bus daemon")
	flag.StringVar(&subject, "subject", ">", "Tibco RendezVous listening subject")
	flag.StringVar(&transportType, "type", "net", "Tibco RendezVous transport type [net,vect,cm,dq,ft]")
	flag.Parse()

	var queue tibrv.RvQueue
	var nettransport tibrv.RvNetTransport
	var cmtransport tibrv.RvCmTransport
	var dqtransport tibrv.RvDqTransport
	var netlistener tibrv.RvListener
	var vectlistener tibrv.RvVectListener
	var cmlistener tibrv.RvCmListener
	var dqlistener tibrv.RvDqListener
	var ftmember tibrv.FtMember

	queue.Create()
	defer queue.Destroy()

	nettransport.Create(tibrv.Service(service), tibrv.Network(network), tibrv.Daemon(daemon))
	defer nettransport.Destroy()

	callback := func(m *tibrv.RvMessage) {
		sendSubject, _ := m.GetSendSubject()
		replySubject, _ := m.GetReplySubject()

		log.Printf("%s - %s - %s\n", sendSubject, replySubject, *m)
	}

	if transportType == "net" {
		netlistener.Create(queue, callback, nettransport, subject)
		fmt.Println("tibrv.RvNetTransport")
	} else if transportType == "vect" {
		vectlistener.Create(queue, callback, nettransport, subject)
		fmt.Println("tibrv.RvNetTransport")
	} else if transportType == "cm" {
		cmtransport.Create(&nettransport)
		cmlistener.Create(queue, callback, cmtransport, subject)
		fmt.Println("tibrv.RvCmTransport")
	} else if transportType == "dq" {
		dqtransport.Create(
			&nettransport,
			tibrv.Name("rvlistener"),
		)
		dqlistener.Create(queue, callback, dqtransport, subject)
		fmt.Println("tibrv.RvDqTransport")
	} else if transportType == "ft" {
		ftcallback := func() func(groupName string, ftAction uint) {
			return func(groupName string, ftAction uint) {
				if ftAction == tibrv.FtActivate {
					fmt.Println("tibrv.FtMember Activate")
					netlistener.Create(queue, callback, nettransport, subject)
				} else if ftAction == tibrv.FtDeactivate {
					fmt.Println("tibrv.FtMember Deactivate")
					netlistener.Destroy()
				} else if ftAction == tibrv.FtPrepareToActivate {
					fmt.Println("tibrv.FtMember PrepareToActivate")
				}
			}
		}()
		ftmember.Create(
			queue,
			ftcallback,
			nettransport,
			"rvlistener",
			2, 1, 1, 0, 3,
		)
		fmt.Println("tibrv.FtMember")
	} else {
		fmt.Println("no transport type supported : ", transportType)
		return
	}

	for {
		queue.Dispatch()
	}
}
