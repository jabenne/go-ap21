package apparel

import (
	"fmt"
)

type RewardAccountService service

type RewardAccount struct {
	ID        string `xml:"Id,omitempty"`
	ProgramId string `xml:"ProgramId"`
	TierId    string `xml:"TierId"`
	PersonId  string `xml:"PersonId"`
}

func (s *RewardAccountService) Post(b *RewardAccount) (*RewardAccount, error) {
	var rE APIError
	var rS RewardAccount
	res, err := s.client.R().
		SetBody(b).
		SetErrorResult(&rE).
		SetSuccessResult(&rS).
		Post("/rewards/Accounts")
	if err != nil {
		return nil, err
	}

	switch s := res.StatusCode; s {
	case 201:
		return &rS, nil
	case 400:
		return nil, fmt.Errorf(rE.ErrorCode, rE.Description)
	case 403:
		return nil, fmt.Errorf(rE.ErrorCode, rE.Description)
	default:
		return nil, fmt.Errorf("Unhandled Status Code: %d", s)
	}
}

type PointsPost struct {
	RequestID   string `xml:"RequestId,omitempty"`
	PersonID    string `xml:"PersonId,omitempty"`
	Points      int    `xml:"Points,omitempty"`
	Description string `xml:"Description,omitempty"`
	Reference   string `xml:"Reference,omitempty"`
	ExpiryDate  string `xml:"ExpiryDate,omitempty"`
}

func (s *RewardAccountService) PostPoints(raID string, b *PointsPost) error {
	var rE APIError
	res, err := s.client.R().
		SetBody(b).
		SetPathParam("raID", raID).
		SetErrorResult(&rE).
		Post("/Rewards/Accounts/{raID}/Points")
	if err != nil {
		return err
	}

	switch s := res.StatusCode; s {
	case 201:
		return nil
	case 400:
		return rE
	case 403:
		return rE
	default:
		return fmt.Errorf("Unhandled Status Code")
	}
}

type PointsRedemption struct {
	AutoConfirm bool   `xml:"AutoConfirm"`
	RequestID   string `xml:"RequestId,omitempty"`
	PersonID    string `xml:"PersonId,omitempty"`
	Points      int    `xml:"Points,omitempty"`
	Description string `xml:"Description,omitempty"`
	Reference   string `xml:"Reference,omitempty"`
	ExpiryDate  string `xml:"ExpiryDate,omitempty"`
}

func (s *RewardAccountService) PointsRedemption(raID string, b *PointsRedemption) error {
	var rE APIError
	res, err := s.client.R().
		SetBody(b).
		SetPathParam("raID", raID).
		SetErrorResult(&rE).
		Post("/Rewards/Accounts/{raID}/Points/Redemptions")
	if err != nil {
		return err
	}

	switch s := res.StatusCode; s {
	case 200:
		return nil
	case 400:
		return rE
	case 403:
		return rE
	default:
		return fmt.Errorf("Unhandled Status Code")
	}
}

type RewardPost struct {
	RequestID   string             `xml:"RequestId,omitempty"`
	PersonID    string             `xml:"PersonId,omitempty"`
	Amount      int                `xml:"Amount,omitempty"`
	Description string             `xml:"Description,omitempty"`
	Reference   string             `xml:"Reference,omitempty"`
	ExpiryDate  string             `xml:"ExpiryDate,omitempty"`
	IssueReason *RewardIssueReason `xml:"IssueReason,omitempty"`
}

type RewardIssueReason struct {
	ID string `xml:"Id,omitempty"`
}

func (s *RewardAccountService) PostReward(raID string, b *RewardPost) error {
	var rE APIError
	res, err := s.client.R().
		SetBody(b).
		SetPathParam("raID", raID).
		SetErrorResult(&rE).
		Post("/Rewards/Accounts/{raID}/Rewards")
	if err != nil {
		return err
	}

	switch s := res.StatusCode; s {
	case 201:
		return nil
	case 400:
		return rE
	case 403:
		return rE
	default:
		return fmt.Errorf("Unhandled Status Code: %d", s)
	}
}

type RewardRedemption struct {
	AutoConfirm bool			   `xml:"AutoConfirm"`
	RequestID   string             `xml:"RequestId,omitempty"`
	PersonID    string             `xml:"PersonId,omitempty"`
	Amount      int                `xml:"Amount,omitempty"`
	Description string             `xml:"Description,omitempty"`
	Reference   string             `xml:"Reference,omitempty"`
	ExpiryDate  string             `xml:"ExpiryDate,omitempty"`
}

func (s *RewardAccountService) RewardRedemption(raID string, b *RewardRedemption) error {
	var rE APIError
	res, err := s.client.R().
		SetBody(b).
		SetPathParam("raID", raID).
		SetErrorResult(&rE).
		Post("/Rewards/Accounts/{raID}/Rewards/Redemptions")
	if err != nil {
		return err
	}

	switch s := res.StatusCode; s {
	case 200:
		return nil
	case 400:
		return rE
	case 403:
		return rE
	default:
		return fmt.Errorf("Unhandled Status Code: %d", s)
	}
}
