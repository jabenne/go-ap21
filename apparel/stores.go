package apparel

type StoresService service

type Store struct {
	StoreId  string   `xml:"StoreId,omitempty"`
	Code     string   `xml:"Code,omitempty"`
	StoreNo  string   `xml:"StoreNo,omitempty"`
	Name     string   `xml:"Name,omitempty"`
	Address1 string   `xml:"Address1,omitempty"`
	Address2 string   `xml:"Address2,omitempty"`
	City     string   `xml:"City,omitempty"`
	State    string   `xml:"State,omitempty"`
	Postcode string   `xml:"Postcode,omitempty"`
	Country  string   `xml:"Country,omitempty"`
	Email    string   `xml:"Email,omitempty"`
	PhoneNum string   `xml:"PhoneNum,omitempty"`
}

type stores struct {
	Stores []Store	`xml:"Store"` 
}

func (s *StoresService) Get() ([]Store, error) {
	var resSucc stores 
	var resErr APIError

	_, err := s.client.R().
		SetErrorResult(&resErr).
		SetSuccessResult(&resSucc).
		Get("/stores")

	if err != nil {
		return nil, err
	}

	return resSucc.Stores, nil
}

func (s *StoresService) GetById(id string) (*Store, error) {
	var resSucc Store 
	var resErr APIError

	_, err := s.client.R().
		SetPathParam("id", id).
		SetSuccessResult(&resSucc).
		SetErrorResult(&resErr).
		Get("/stores/{id}")

	if err != nil {
		return nil, err
	}

	return &resSucc, nil

}
