package main

import (
	"context"
	"fmt"
	"htmlparser/adapters"
	"htmlparser/business"
	"net/http"
	"os"
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

	http.HandleFunc("/update", func(w http.ResponseWriter, req *http.Request) {
		err = bn.SendPriceUpdate(context.Background())
		if err != nil {
			fmt.Fprint(w, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)

	select {}
}
