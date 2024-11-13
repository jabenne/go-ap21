package apparel_test

import (
	"encoding/xml"
	"testing"

	"github.com/jabenne/go-ap21/apparel"
)

func TestGetPersonsById(t *testing.T) {
	config, err := apparel.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	c := apparel.NewClient(config)
	_, err = c.Persons.GetById("18863410")
	if err != nil { 
		t.Fatal(err)
	}
	
}

func TestGetPersons(t *testing.T) {
	config, err := apparel.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	c := apparel.NewClient(config)
	opts := apparel.PersonsGetOpts{}
	_, err = c.Persons.Get(&opts)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPostPerson(t *testing.T) {
	config, err := apparel.NewConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}

	c := apparel.NewClient(config)

	email := "apitest12@jbennett.com"
	mobile := "0461430888"
	
	addressLine1 := "10 Test Street"
	addressLine2 := "Unit 20"
	city := "TestCity"
	state := "NSW"
	postcode := "2000"
	country := "Australia"

	bAddress := apparel.Address{
		ContactName: "Test User",
		AddressLine1: addressLine1,
		AddressLine2: addressLine2,
		City: city,
		State: state,
		Postcode: postcode ,
		Country: country,
	}

	contact := apparel.Contact{
		Email: email,
		Phones: apparel.Phones {
			Mobile: mobile, 
		},
	}

	payload := apparel.Person{
		Firstname: "Test",
		Surname: "User",
		Contacts: &contact,
		Addresses: &apparel.Addresses{
			Billing: bAddress,
		},
	}
	pl, err := xml.Marshal(payload)
	t.Log(string(pl))
	id, err := c.Persons.Post(&payload)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(id)

}


