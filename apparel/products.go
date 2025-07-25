package apparel

import (
	"fmt"
	"time"
)

type ProductsService service

type Product struct {
	ID              string      `xml:"Id,omitemtpy"`
	RowNumber       string      `xml:"RowNumber,omitemtpy"`
	Code            string      `xml:"Code,omitemtpy"`
	Name            string      `xml:"Name,omitemtpy"`
	Description     string      `xml:"Description,omitemtpy"`
	SizeRange       string      `xml:"SizeRange,omitemtpy"`
	RrpText         string      `xml:"RrpText,omitemtpy"`
	UpdateTimeStamp string      `xml:"UpdateTimeStamp,omitemtpy"`
	References      []Reference `xml:"References>Reference,omitemtpy"`
	Clrs            []Clr       `xml:"Clrs>Clr,omitemtpy"`
	CustomData      CustomData  `xml:"CustomData"`
}

type Clr struct {
	Text           string `xml:",chardata,omitemtpy"`
	ID             string `xml:"Id,omitemtpy"`
	ProductId      string `xml:"ProductId,omitemtpy"`
	Sequence       string `xml:"Sequence,omitemtpy"`
	Code           string `xml:"Code,omitemtpy"`
	Name           string `xml:"Name,omitemtpy"`
	TypeCode       string `xml:"TypeCode,omitemtpy"`
	TypeName       string `xml:"TypeName,omitemtpy"`
	MarkdownStatus string `xml:"MarkdownStatus,omitemtpy"`
	SKUs           []SKU  `xml:"SKUs>SKU,omitemtpy"`
}

type SKU struct {
	Text            string `xml:",chardata,omitemtpy"`
	ID              int    `xml:"Id,omitemtpy"`
	ClrId           string `xml:"ClrId,omitemtpy"`
	Sequence        string `xml:"Sequence,omitemtpy"`
	SizeCode        string `xml:"SizeCode,omitemtpy"`
	SizeDescription string `xml:"SizeDescription,omitemtpy"`
	OriginalPrice   string `xml:"OriginalPrice,omitemtpy"`
	Price           string `xml:"Price,omitemtpy"`
	RetailPrice     string `xml:"RetailPrice,omitemtpy"`
	FreeStock       string `xml:"FreeStock,omitemtpy"`
	NextAvailable   string `xml:"NextAvailable,omitemtpy"`
	Barcode         string `xml:"Barcode,omitemtpy"`
	GlobalSequence  string `xml:"GlobalSequence,omitemtpy"`
}

type products struct {
	Products []Product `xml:"Product"`
}

type GetProductsOpts struct {
	CustomData   bool       `url:"customData"`
	ExtendedRefs bool       `url:"extendedRefs"`
	UpdatedAfter *time.Time `url:"updatedAfter"`
}

func (o *GetProductsOpts) generateParams() map[string]string {
	params := make(map[string]string)

	params["customData"] = fmt.Sprintf("%t", o.CustomData)
	params["extendedRefs"] = fmt.Sprintf("%t", o.ExtendedRefs)

	if o.UpdatedAfter != nil {
		params["updatedAfter"] = o.UpdatedAfter.Format("2006-01-02T15:04:05")
	}

	return params

}

func (s *ProductsService) Get(opts GetProductsOpts) ([]Product, error) {
	var resSucc products
	var resErr APIError

	_, err := s.client.R().
		SetErrorResult(&resErr).
		SetQueryParams(opts.generateParams()).
		SetSuccessResult(&resSucc).
		Get("/products")

	if err != nil {
		return nil, err
	}

	return resSucc.Products, nil
}
