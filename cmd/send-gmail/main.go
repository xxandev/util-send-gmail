package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"util-send-gmail/internal/gmail"
	"util-send-gmail/internal/utils"
)

var (
	config         Config
	to, subj, body string
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("[SEND-GMAIL] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)

	if ok, err := config.Init(); ok && err != nil {
		log.Printf("warning read config file: %v", err)
	}

	example := flag.Bool("example", false, "create example config (*.yaml)")
	mute := flag.Bool("mute", false, "mute log")
	login := flag.String("login", "", "gmail account login")
	pass := flag.String("pass", "", "gmail account password")
	flag.StringVar(&to, "to", "", "message/caption text")
	flag.StringVar(&subj, "subj", "", "photo/image path")
	flag.StringVar(&body, "body", "", "photo/image path")
	flag.Parse()

	if *example {
		config.Example()
		content, err := config.Marshal(YAML)
		if err != nil {
			log.Fatalf("error marshal example-config: %v\n", err)
		}
		if err := utils.CreateFile("send-gmail.yaml", content); err != nil {
			log.Fatalf("error create example-config: %v\n", err)
		}
		log.Println("example-config created")
		os.Exit(0)
	}

	config.Gmail.Set(*login, *pass)
	if err := config.Check(); err != nil {
		log.Fatalf("error check config: %v\n", err)
	}

	if *mute {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	mail := gmail.New(&config.Gmail)
	if err := mail.Send(to, subj, body); err != nil {
		log.Fatalf("error send mail: %v", err)
	}
	log.Println("send mail ok")
}
