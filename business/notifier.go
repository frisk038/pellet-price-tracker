package business

import (
	"context"
	"fmt"
	"htmlparser/adapters"
)

type BusinessNotifier struct {
	t   *adapters.TelgramClient
	b   *adapters.ButagazClient
	p   *adapters.PGClient
	url string
}

func NewBusinessNotifier(t *adapters.TelgramClient, b *adapters.ButagazClient, pg *adapters.PGClient, url string) *BusinessNotifier {
	return &BusinessNotifier{t: t, b: b, p: pg, url: url}
}

func (b *BusinessNotifier) SendPriceUpdate(ctx context.Context) error {
	webPrice, err := b.b.GetPrice()
	if err != nil {
		return err
	}

	dbPrice, err := b.p.GetPrice(ctx)
	if err != nil {
		return err
	}

	if webPrice != dbPrice {
		err = b.t.SendToGroup(fmt.Sprintf("The current price of pellet is %d€ \n%s", webPrice, b.url))
		if err == nil {
			return b.p.InsertPrice(ctx, webPrice)
		}
	}
	return nil
}
