package uspsgo

import (
	"testing"

	"github.com/james-darko/uspsgo/rt"
)

var testAddress = &Address{
	Firm:          "DLV Roofing Systems",
	StreetAddress: "2554 US-290",
	City:          "DRIPPING SPRINGS",
	State:         "TX",
	ZIPCode:       "78620",
}

var testAddress2 = &Address{
	Firm:          "DLV Roofing Systems",
	StreetAddress: "490 JESSEN LN",
	City:          "CHARLESTON",
	State:         "SC",
	ZIPCode:       "29492",
}

func TestZIPCode(t *testing.T) {
	readEnv()
	ctx := Context()
	c := New(rt.MustGetEnv("USPS_KEY"), rt.MustGetEnv("USPS_SECRET"))

	info, err := c.Address(ctx, testAddress2)
	if err != nil {
		t.Fatalf("Address lookup failed: %v", err)
	}
	t.Logf("Address Info: %+v\n", info)

	cityState, err := c.CityState(ctx, testAddress2.ZIPCode)
	if err != nil {
		t.Fatalf("CityState lookup failed: %v", err)
	}
	t.Logf("CityState Info: %+v\n", cityState)

	addr, err := c.ZIPCode(ctx, testAddress2)
	if err != nil {
		t.Fatalf("ZIPCode lookup failed: %v", err)
	}
	t.Logf("ZIPCode Info: %s\n", addr.String())
}
