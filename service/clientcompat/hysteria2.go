package clientcompat

import (
	"fmt"
	"net/url"

	"trojan-panel/model/constant"
)

// Hysteria2ShareServer adapts the official multi-port authority syntax for
// clients such as v2rayN that require a numeric URI port and an mport query.
func Hysteria2ShareServer(domain string, port uint, portHopping string, client string) (string, string) {
	server := fmt.Sprintf("%s:%d", domain, port)
	if portHopping == "" {
		return server, ""
	}
	if client == constant.ClientV2Ray {
		return server, "&mport=" + url.QueryEscape(portHopping)
	}
	return fmt.Sprintf("%s:%s", domain, portHopping), ""
}
