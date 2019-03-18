// Package nvg589 creates an API of sorts for the Arris NVG589 RGW
package nvg589

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
)

// BroadBandStatistics - structure to contain various values from the stats
// page
type BroadBandStatistics struct {
	connectionSource        string
	broadbandIPv4           string
	gatewayIPv4             string
	wanMACAddress           string
	primaryDNS              string
	secondaryDNS            string
	maximumTransmissionUnit int16
	dslamVendorID           string
	lineState               [2]string
	downstreamSyncRate      [2]int32
	upstreamSyncrate        [2]int32
	downstreamMaxRate       [2]int32
	upstreamMaxRate         [2]int32
	modulation              [2]string
	dataPath                [2]string
	snMargin                [4]float32
	lineAttenuation         [4]float32
}

const (
	line1idx = 0 // first line in lines
	line2idx = 1 // second line in lines
)

// Connect to the RGW and pulls various stastics
func Connect(rgwInternal string) string {
	ipAddr := net.ParseIP(rgwInternal)

	url, err := url.Parse("http://" + ipAddr.String() + "/cgi-bin/broadbandstatistics.ha")
	if err != nil {
		fmt.Println(err)
	}

	resp, err := http.Get(url.String())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)

	return url.String()

}
