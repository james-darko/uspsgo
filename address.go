package uspsgo

import (
	"context"
	"strings"
)

var addressEndpoint = endpointBase + "/addresses/v3/address"
var zipToCityStateEndpoint = endpointBase + "/addresses/v3/city-state"
var zipOfAddressEndpoint = endpointBase + "/addresses/v3/zipcode"

func (c Client) Address(ctx context.Context, addr *Address) (*AddressInfo, error) {
	resp := &AddressInfo{}
	err := c.request(ctx, addressEndpoint, addr, resp)
	if err != nil {
		return nil, err
	}
	resp.Address.Firm = addr.Firm
	return resp, nil
}

type AdditionalInfo struct {
	DeliveryPoint        string `json:"deliveryPoint"`
	CarrierRoute         string `json:"carrierRoute"`
	DPVConfirmation      string `json:"DPVConfirmation"`
	DPVCMRA              string `json:"DPVCMRA"`
	Business             string `json:"business"`
	CentralDeliveryPoint string `json:"centralDeliveryPoint"`
	Vacant               string `json:"vacant"`
}

type Address struct {
	Firm                      string `json:"firm"`
	StreetAddress             string `json:"streetAddress"`
	StreetAddressAbbreviation string `json:"streetAddressAbbreviation"`
	SecondaryAddress          string `json:"secondaryAddress"`
	CityAbbreviation          string `json:"cityAbbreviation"`
	City                      string `json:"city"`
	State                     string `json:"state"`
	ZIPCode                   string `json:"ZIPCode"`
	ZIPPlus4                  string `json:"ZIPPlus4"`
}

func (a *Address) String() string {
	var buf strings.Builder
	if a.Firm != "" {
		buf.WriteString(a.Firm)
		buf.WriteString(" ")
	}
	buf.WriteString(a.StreetAddress)
	if a.SecondaryAddress != "" {
		buf.WriteString(" ")
		buf.WriteString(a.SecondaryAddress)
	}
	buf.WriteString(", ")
	buf.WriteString(a.City)
	buf.WriteString(", ")
	buf.WriteString(a.State)
	buf.WriteString(", ")
	buf.WriteString(a.ZIPCode)
	if a.ZIPPlus4 != "" {
		buf.WriteString("-")
		buf.WriteString(a.ZIPPlus4)
	}
	return buf.String()
}

func (a *Address) StringAbbrivated() string {
	var buf strings.Builder
	if a.Firm != "" {
		buf.WriteString(a.Firm)
		buf.WriteString(" ")
	}
	buf.WriteString(a.StreetAddressAbbreviation)
	if a.SecondaryAddress != "" {
		buf.WriteString(" ")
		buf.WriteString(a.SecondaryAddress)
	}
	buf.WriteString(", ")
	buf.WriteString(a.CityAbbreviation)
	buf.WriteString(", ")
	buf.WriteString(a.State)
	buf.WriteString(", ")
	buf.WriteString(a.ZIPCode)
	if a.ZIPPlus4 != "" {
		buf.WriteString("-")
		buf.WriteString(a.ZIPPlus4)
	}
	return buf.String()
}

type Correction struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type Match struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type AddressInfo struct {
	Firm           string         `json:"firm"`
	Address        Address        `json:"address"`
	AdditionalInfo AdditionalInfo `json:"additionalInfo"`
	Corrections    []Correction   `json:"corrections"`
	Matches        []Match        `json:"matches"`
	Warnings       []string       `json:"warnings"`
}

type zipCodeParam struct {
	ZIPCode string `json:"ZIPCode"`
}

func (c Client) CityState(ctx context.Context, zipCode string) (*CityState, error) {
	resp := &CityState{}
	err := c.request(ctx, zipToCityStateEndpoint, zipCodeParam{zipCode}, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type CityState struct {
	City  string `json:"city"`
	State string `json:"state"`
	ZIP5  string `json:"ZIPCode"`
}

func (c Client) ZIPCode(ctx context.Context, addr *Address) (*Address, error) {
	var resp addressZip
	err := c.request(ctx, zipOfAddressEndpoint, addr, &resp)
	if err != nil {
		return nil, err
	}
	r := &resp.Address
	r.Firm = resp.Firm
	return r, nil
}

type addressZip struct {
	Firm    string  `json:"firm"`
	Address Address `json:"address"`
}
