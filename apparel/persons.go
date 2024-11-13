package apparel

import (
	"fmt"
	"strings"
)

type PersonsService service

type CustomData struct {}

type Addresses struct {
	Billing Address `xml:"Billing,omitempty"`
}

type Address struct {
	ContactName  string `xml:"ContactName,omitempty"`
	AddressLine1 string `xml:"AddressLine1,omitempty"`
	AddressLine2 string `xml:"AddressLine2,omitempty"`
	City         string `xml:"City,omitempty"`
	State        string `xml:"State,omitempty"`
	Postcode     string `xml:"Postcode,omitempty"`
	Country      string `xml:"Country,omitempty"`
}

type Contact struct {
	Email  string `xml:"Email,omitempty"`
	Phones Phones `xml:"Phones,omitempty"`
} 

type Phones struct {
	Home   string `xml:"Home,omitempty"`
	Mobile string `xml:"Mobile,omitempty"`
	Work   string `xml:"Work,omitempty"`
} 

type RewardsAccount struct {
	ID          string `xml:"Id,omitempty"`
	ProgramId   string `xml:"ProgramId,omitempty"`
	ProgramName string `xml:"ProgramName,omitempty"`
	TierId      string `xml:"TierId,omitempty"`
	TierName    string `xml:"TierName,omitempty"`
}

type Currency struct { 
	Code   string `xml:"Code,omitempty"`
	Format string `xml:"Format,omitempty"`
} 

type Loyalty struct {
	ID            string `xml:"Id,omitempty"`
	LoyaltyTypeId string `xml:"LoyaltyTypeId,omitempty"`
	LoyaltyType   string `xml:"LoyaltyType,omitempty"`
	CardNo        string `xml:"CardNo,omitempty"`
	Expiry        string `xml:"Expiry,omitempty"`
	Balance       string `xml:"Balance,omitempty"`
	CreditStatus  string `xml:"CreditStatus,omitempty"`
	Message       string `xml:"Message,omitempty"`
	JoinDate      string `xml:"JoinDate,omitempty"`
	StatusId      string `xml:"StatusId,omitempty"`
}

type Reference struct {
	ReferenceTypeId string `xml:"ReferenceTypeId,omitempty"`
	ID              string `xml:"Id,omitempty"`
}

type Person struct {
	ID              string   `xml:"Id,omitempty"`
	Code            string   `xml:"Code,omitempty"`
	Title           string   `xml:"Title,omitempty"`
	Initials        string   `xml:"Initials,omitempty"`
	Firstname       string   `xml:"Firstname,omitempty"`
	Surname         string   `xml:"Surname,omitempty"`
	Sex             string   `xml:"Sex,omitempty"`
	DateOfBirth     string   `xml:"DateOfBirth,omitempty"`
	StartDate       string   `xml:"StartDate,omitempty"`
	JobTitle        string   `xml:"JobTitle,omitempty"`
	Privacy         string   `xml:"Privacy,omitempty"`
	UpdateTimeStamp string   `xml:"UpdateTimeStamp,omitempty"`
	References      []Reference `xml:"References,omitempty"`
	IsAgent   string `xml:"IsAgent,omitempty"`
	Addresses *Addresses `xml:"Addresses,omitempty"`
	Contacts *Contact`xml:"Contacts,omitempty"`
	Currency *Currency `xml:"Currency,omitempty"`
	Loyalties *struct {
		Loyalty Loyalty `xml:"Loyalty,omitempty"`
	} `xml:"Loyalties,omitempty"`
	RewardsAccounts *struct {
		Account RewardsAccount`xml:"Account,omitempty"`
	} `xml:"RewardsAccounts,omitempty"`
}


type PersonsGetOpts struct {
	Surname string
	Firstname string
	Email string
	Phone string
	Code string
	Password string
	LoyaltyOnly bool
	UpdatedAfter string
}

func (o *PersonsGetOpts) BuildParams() map[string]string {
	params := make(map[string]string)

	if o.LoyaltyOnly {
		params["loyaltyonly"] = "true" 
	}
	if o.UpdatedAfter != "" {
		params["updatedafter"] = o.UpdatedAfter
	}
	if o.Password != "" {
		params["password"] = o.Password
	}
	if o.Code != "" {
		params["code"] = o.Code
	}
	if o.Phone != "" {
		params["phone"] = o.Phone
	}
	if o.Email != "" {
		params["email"] = o.Email
	}
	if o.Firstname != "" {
		params["firstname"] = o.Firstname
	}
	if o.Surname != "" {
		params["surname"] = o.Surname
	}

	return params
}

func (s *PersonsService) Get(p *PersonsGetOpts) (*Person, error) {
	var resSucc Person 
	var resErr APIError

	_, err := s.client.R().
		SetQueryParams(p.BuildParams()).
		SetErrorResult(&resErr).
		SetSuccessResult(&resSucc).
		Get("/persons")

	if err != nil {
		return nil, err
	}

	return &resSucc, nil
}

func (s *PersonsService) GetById(id string) (*Person, error) {
	var resSucc Person 
	var resErr APIError

	_, err := s.client.R().
		SetPathParam("id", id).
		SetSuccessResult(resSucc).
		SetErrorResult(resErr).
		Get("/persons/{id}")

	if err != nil {
		return nil, err
	}

	return &resSucc, nil
}

func (s *PersonsService) Post(b *Person) (string, error) {
	var apiErr APIError 
	res, err := s.client.R().
		SetBody(b).
		SetErrorResult(&apiErr).
		Post("/persons")
	if err != nil {
		return "", err
	}
	
	switch s := res.StatusCode; s {
		case 201:
			location := res.GetHeader("location")
			lastPath := strings.Split(location, "/")[5]
			return lastPath[:strings.IndexByte(lastPath, '?')], nil
		case 400:
			return "", apiErr 
		case 403:
			return "", apiErr
		default:
			return "", fmt.Errorf("Unhandled Status Code")
	}
}
