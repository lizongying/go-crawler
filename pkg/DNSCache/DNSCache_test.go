package DNSCache

import (
	"testing"
	"time"
)

func TestNewDNSCache(t *testing.T) {
	dnsCache := NewDNSCache(3*time.Second, 3)

	host := "baidu.com"
	ip, found := dnsCache.Get(host)
	if found {
		t.Logf("DNS Cache Hit: %s -> %s\n", host, ip)
	} else {
		t.Logf("DNS Cache Miss: %s\n", host)
		_, ok := dnsCache.ResolveWithRetry(host)
		if ok {
			t.Logf("Resolved %s and added to cache\n", host)
			ticker := time.NewTicker(time.Second)
		out:
			for {
				select {
				case <-ticker.C:
					ip, found = dnsCache.Get(host)
					if found {
						t.Logf("DNS Cache Hit: %s -> %s\n", host, ip)
					} else {
						t.Logf("DNS Cache Expired: %s\n", host)
						ticker.Stop()
						break out
					}
				}
			}
		} else {
			t.Logf("Failed to resolve %s\n", host)
		}
	}
}
