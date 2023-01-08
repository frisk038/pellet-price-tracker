package adapters

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type ButagazClient struct {
	URL string
}

var errorNotFound = fmt.Errorf("price not found in webpage")

func NewButagazClient(url string) *ButagazClient {
	return &ButagazClient{URL: url}
}

func getHtmlPage(webPage string) (string, error) {
	resp, err := http.Get(webPage)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func findTag(attrs []html.Attribute, str string) *string {
	for _, attr := range attrs {
		if attr.Key == str {
			return &attr.Val
		}
	}
	return nil
}

func (b *ButagazClient) GetPrice() (int, error) {
	data, err := getHtmlPage(b.URL)
	if err != nil {
		return 0, err
	}
	z := html.NewTokenizer(strings.NewReader(data))

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return 0, errorNotFound
		case html.StartTagToken:
			t := z.Token()
			if t.Data == "span" {
				priceMl := findTag(t.Attr, "data-price")
				if priceMl != nil {
					price, err := strconv.Atoi(*priceMl)
					if err == nil {
						return price / 100, nil
					}
				}
			}
		}
	}
}
