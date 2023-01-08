package main

import (
	"context"
	"fmt"
	"htmlparser/adapters"
	"htmlparser/business"
	"os"

	"github.com/robfig/cron"
)

func main() {
	URL := os.Getenv("BUTAGAZ_URL")
	btgz := adapters.NewButagazClient(URL)
	tlgm, err := adapters.NewTelegramClient()
	if err != nil {
		fmt.Println(err)
	}
	ckrh, err := adapters.NewPGClient()
	if err != nil {
		fmt.Println(err)
	}

	bn := business.NewBusinessNotifier(tlgm, btgz, ckrh, URL)

	cr := cron.New()
	cr.AddFunc("0 30 * * * *", func() {
		err := bn.SendPriceUpdate(context.Background())
		if err != nil {
			fmt.Println(err)
		}
	})
	cr.Start()

	select {}
}
