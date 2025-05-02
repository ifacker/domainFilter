package controllor

import (
	"fmt"
	"testing"
)

func TestResolveDomain(t *testing.T) {
	ips, _ := resolveDomain("www.baidu.com")
	for _, ip := range ips {
		fmt.Println(ip.To4().String())
	}
}
