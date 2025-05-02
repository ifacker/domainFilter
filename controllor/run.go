package controllor

import (
	"domainFilter/public"
	"strings"

	"github.com/gookit/color"
)

func Run() {
	pool := NewPool(5, 20)
	defer pool.Close()

	// 动态调整协程数量
	pool.SetMaxWorkers(10)

	// 把域名解析为IP地址
	for _, domain := range public.Targets {
		if strings.TrimSpace(domain) == "" {
			continue
		}
		pool.Submit(func() {
			ips, err := resolveDomain(domain)
			if err != nil {
				if !public.Options.NoErrDomain {
					color.C256(250).Printf("[-] %s\n", err)
				}
				return
			}
			for _, ip := range ips {
				if ip.To4() != nil {
					ipv4 := ip.To4().String()
					public.IpMapMutex.Lock()
					// 检查键是否存在，如果不存在则创建一个新的切片
					if _, ok := public.IpDomainMap[ipv4]; !ok {
						public.IpDomainMap[ipv4] = []string{}
					}
					public.IpDomainMap[ipv4] = append(public.IpDomainMap[ipv4], domain)
					public.IpMapMutex.Unlock()
					if public.Options.ViewPrint {
						color.C256(49).Printf("[+] %s -> %v\n", ipv4, domain)
					}
				}
			}
		})
	}
}
