package uspsgo

import (
	"fmt"

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

func main() {
	ctx := Context()
	c := New(rt.MustGetEnv("USPS_KEY"), rt.MustGetEnv("USPS_SECRET"))
	info, err := c.Address(ctx, testAddress2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", info)
	ciyState, err := c.CityState(ctx, testAddress2.ZIPCode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", ciyState)
	addr, err := c.ZIPCode(ctx, testAddress2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(addr.String())
}
