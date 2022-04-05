package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/julio77it/tibrv"
)

func main() {
	var service, network, daemon, subject, transportType, format string
	flag.StringVar(&service, "service", "", "Tibco RendezVous Bus service")
	flag.StringVar(&network, "network", "", "Tibco RendezVous Bus network")
	flag.StringVar(&daemon, "daemon", "", "Tibco RendezVous Bus daemon")
	flag.StringVar(&subject, "subject", ">", "Tibco RendezVous listening subject")
	flag.StringVar(&transportType, "type", "net", "Tibco RendezVous transport type [net,vect,cm,dq,ft]")
	flag.StringVar(&format, "format", "msg", "choose message format to print [msg,json]")
	flag.Parse()

	log.Printf("RVD %s %s %s", service, network, daemon)

	var queue tibrv.RvQueue
	var nettransport tibrv.RvNetTransport
	var cmtransport tibrv.RvCmTransport
	var dqtransport tibrv.RvDqTransport
	var netlistener tibrv.RvListener
	var vectlistener tibrv.RvVectListener
	var cmlistener tibrv.RvCmListener
	var dqlistener tibrv.RvDqListener
	var ftmember tibrv.FtMember

	if err := queue.Create(); err != nil {
		panic(err)
	}
	defer queue.Destroy()

	if err := nettransport.Create(tibrv.Service(service), tibrv.Network(network), tibrv.Daemon(daemon)); err != nil {
		panic(err)
	}
	defer nettransport.Destroy()

	callback := func(m *tibrv.RvMessage) {
		sendSubject, _ := m.GetSendSubject()
		replySubject, _ := m.GetReplySubject()

		switch format {
		case "msg":
			log.Printf("%s - %s - %s\n", sendSubject, replySubject, *m)

		case "json":
			j, _ := m.JSON()
			log.Printf("%s - %s - %s\n", sendSubject, replySubject, j)
		}
	}

	if transportType == "net" {
		if err := netlistener.Create(queue, callback, nettransport, subject); err != nil {
			panic(err)
		}
		fmt.Println("tibrv.RvNetTransport")
	} else if transportType == "vect" {
		if err := vectlistener.Create(queue, callback, nettransport, subject); err != nil {
			panic(err)
		}
		fmt.Println("tibrv.RvNetTransport")
	} else if transportType == "cm" {
		if err := cmtransport.Create(&nettransport); err != nil {
			panic(err)
		}
		if err := cmlistener.Create(queue, callback, cmtransport, subject); err != nil {
			panic(err)
		}
		fmt.Println("tibrv.RvCmTransport")
	} else if transportType == "dq" {
		if err := dqtransport.Create(
			&nettransport,
			tibrv.Name("rvlistener"),
		); err != nil {
			panic(err)
		}
		if err := dqlistener.Create(queue, callback, dqtransport, subject); err != nil {
			panic(err)
		}
		fmt.Println("tibrv.RvDqTransport")
	} else if transportType == "ft" {
		ftcallback := func() func(groupName string, ftAction uint) {
			return func(groupName string, ftAction uint) {
				if ftAction == tibrv.FtActivate {
					fmt.Println("tibrv.FtMember Activate")
					if err := netlistener.Create(queue, callback, nettransport, subject); err != nil {
						panic(err)
					}
				} else if ftAction == tibrv.FtDeactivate {
					fmt.Println("tibrv.FtMember Deactivate")
					if err := netlistener.Destroy(); err != nil {
						panic(err)
					}
				} else if ftAction == tibrv.FtPrepareToActivate {
					fmt.Println("tibrv.FtMember PrepareToActivate")
				}
			}
		}()
		if err := ftmember.Create(
			queue,
			ftcallback,
			nettransport,
			"rvlistener",
			2, 1, 1, 0, 3,
		); err != nil {
			panic(err)
		}
		fmt.Println("tibrv.FtMember")
	} else {
		fmt.Println("no transport type supported : ", transportType)
		return
	}

	for {
		if err := queue.Dispatch(); err != nil {
			panic(err)
		}
	}
}
