package apparel_test

import (
	"testing"

	"github.com/jabenne/go-ap21/apparel"
)

func TestNew(t *testing.T) {
	config, err:= apparel.NewConfigFromEnv() 
	if err != nil {
		t.Fatal(err)
	}
	_ = apparel.NewClient(config)
}
