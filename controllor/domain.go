package controllor

import "net"

// resolveDomain 函数用于将域名解析为 IP 地址
func resolveDomain(domain string) ([]net.IP, error) {
	// 调用 net.LookupIP 函数进行 DNS 解析
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
