package dns_cache

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

type Ip struct {
	ip   []net.IP
	next net.IP
}
type DnsCache struct {
	cache      map[string]Ip
	mutex      sync.RWMutex
	ttl        time.Duration
	maxRetries int
}

func NewDnsCache(ttl time.Duration, maxRetries int) *DnsCache {
	return &DnsCache{
		cache:      make(map[string]Ip),
		ttl:        ttl,
		maxRetries: maxRetries,
	}
}

func (d *DnsCache) Get(host string) (net.IP, bool) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	ips, found := d.cache[host]
	if found {
		go func() {
			ips.next = ips.ip[rand.Intn(len(ips.ip))]
		}()
		return ips.next, true
	}
	return nil, false
}

func (d *DnsCache) ResolveWithRetry(host string) (net.IP, bool) {
	for i := 0; i < d.maxRetries; i++ {
		ipAddresses, err := net.LookupIP(host)
		var ips []net.IP
		for _, v := range ipAddresses {
			if strings.Contains(v.String(), ".") {
				ips = append(ips, v)
			}
		}
		if err == nil && len(ips) > 0 {
			d.Set(host, ips)
			return d.Get(host)
		}
		fmt.Printf("DNS resolution failed for %s (attempt %d): %v\n", host, i+1, err)
		time.Sleep(time.Second)
	}
	return nil, false
}

func (d *DnsCache) Set(host string, ip []net.IP) {
	if len(ip) == 0 {
		return
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.cache[host] = Ip{ip: ip, next: ip[rand.Intn(len(ip))]}
	time.AfterFunc(d.ttl, func() {
		d.mutex.Lock()
		defer d.mutex.Unlock()
		delete(d.cache, host)
	})
}
