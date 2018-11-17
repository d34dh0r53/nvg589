// Package nvg589 creates an API of sorts for the Arris NVG589 RGW
package nvg589

import (
	"fmt"
	"net"
	"net/url"
)

// Connect to the RGW and pulls various stastics
func Connect(rgwInternal string) string {
	ipAddr := net.ParseIP(rgwInternal)

	u, err := url.Parse("http://" + ipAddr.String() + "/cgi-bin/broadbandstatistics.ha")

	if err != nil {
		fmt.Println(err)
	}

	return u.String()

}
