package middleware

import (
	"fmt"
	"net"
	"net/http"
	"net/netip"
)

func SubNetFilter(trustedSubnet string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, err := getIP(r)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			if !isAllowed(ip, trustedSubnet) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func getIP(r *http.Request) (netip.Addr, error) {
	const ipHeaderName = "X-Real-IP"
	ip := r.Header.Get(ipHeaderName)
	if ip == "" {
		return netip.Addr{}, fmt.Errorf("can't get ip address from header [%s]", ipHeaderName)
	}
	return netip.ParseAddr(ip)
}

func isAllowed(ipAddr netip.Addr, cidr string) bool {
	if !ipAddr.IsValid() {
		return false
	}
	_, ipSub, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}
	return ipSub.Contains(ipAddr.AsSlice())
}
