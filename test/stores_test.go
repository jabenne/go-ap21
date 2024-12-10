package apparel_test

import (
	"testing"

	"github.com/jabenne/go-ap21/apparel"
)

func TestGetAllStores(t *testing.T) {
	config, err := apparel.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	c := apparel.NewClient(config)
	s, err := c.Stores.Get()
	if err != nil { 
		t.Fatal(err)
	}
	t.Log(s)
}

func TestGetStoreById(t *testing.T) {
	config, err := apparel.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	c := apparel.NewClient(config)
	s, err := c.Stores.GetById("100912")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

