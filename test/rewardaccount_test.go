package apparel_test 

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/jabenne/go-ap21/apparel"
)

func TestPostRewardAccount(t *testing.T) {
	config, err := apparel.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	c := apparel.NewClient(config)
	
	payload := apparel.RewardAccount {
		PersonId: "18863417",
		TierId: "101",	
		ProgramId: "101",
	}

	pl, err := xml.Marshal(payload)
	t.Log(string(pl))
	id, err := c.RewardAccount.Post(&payload)
	fmt.Println(id)	
	if err != nil {
		t.Fatal(err)
	}

}

func TestPostPoints(t *testing.T) {
	config, err := apparel.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	c := apparel.NewClient(config)
	pID := "18970471"
	rAID := "9566842"

	p := apparel.PointsPost{
		PersonID: pID,
		RequestID: "273b85ef-abe2-497f-af52-0655c629f580",
		Points: 100,
		Description: "Test",
	}

	pl, err := xml.Marshal(p)
	t.Log(string(pl))

	err = c.RewardAccount.PostPoints(rAID, &p)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPostReward(t *testing.T) {
	config, err := apparel.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	c := apparel.NewClient(config)
	pID := "18970471"
	rAID := "9566842"

	p := apparel.RewardPost{
		PersonID: pID,
		RequestID: "486f0e80-5a92-4c00-b698-22b13eee7258",
		Amount: 100,
		Description: "Test",
	}

	pl, err := xml.Marshal(p)
	t.Log(string(pl))

	err = c.RewardAccount.PostReward(rAID, &p)

	if err != nil {
		t.Fatal(err)
	}
}





