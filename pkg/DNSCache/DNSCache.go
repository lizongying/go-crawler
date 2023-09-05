package DNSCache

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
type DNSCache struct {
	cache      map[string]Ip
	mutex      sync.RWMutex
	ttl        time.Duration
	maxRetries int
}

func NewDNSCache(ttl time.Duration, maxRetries int) *DNSCache {
	return &DNSCache{
		cache:      make(map[string]Ip),
		ttl:        ttl,
		maxRetries: maxRetries,
	}
}

func (c *DNSCache) Get(host string) (net.IP, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	ips, found := c.cache[host]
	if found {
		go func() {
			ips.next = ips.ip[rand.Intn(len(ips.ip))]
		}()
		return ips.next, true
	}
	return nil, false
}

func (c *DNSCache) ResolveWithRetry(host string) (net.IP, bool) {
	for i := 0; i < c.maxRetries; i++ {
		ipAddresses, err := net.LookupIP(host)
		var ips []net.IP
		for _, v := range ipAddresses {
			if strings.Contains(v.String(), ".") {
				ips = append(ips, v)
			}
		}
		if err == nil && len(ips) > 0 {
			c.Set(host, ips)
			return c.Get(host)
		}
		fmt.Printf("DNS resolution failed for %s (attempt %d): %v\n", host, i+1, err)
		time.Sleep(time.Second)
	}
	return nil, false
}

func (c *DNSCache) Set(host string, ip []net.IP) {
	if len(ip) == 0 {
		return
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[host] = Ip{ip: ip, next: ip[rand.Intn(len(ip))]}
	time.AfterFunc(c.ttl, func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		delete(c.cache, host)
	})
}
