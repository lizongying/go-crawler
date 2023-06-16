package crawler

import (
	"github.com/pkg/errors"
	"net/url"
	"path/filepath"
)

func (c *Crawler) IsAllowedDomain(Url *url.URL) (ok bool) {
	for domain := range c.allowedDomains {
		matched, err := filepath.Match(domain, Url.Hostname())
		if err != nil {
			c.logger.Warn(err)
			continue
		}
		if matched {
			ok = true
			return
		}
	}

	return
}

func (c *Crawler) GetAllowedDomains() (domains []string) {
	for domain := range c.allowedDomains {
		domains = append(domains, domain)
	}

	return
}

func (c *Crawler) ReplaceAllowedDomains(domains []string) (err error) {
	if len(domains) == 0 {
		err = errors.New("domains is empty")
		c.logger.Error(err)
		return
	}

	domainsMap := make(map[string]struct{})
	for _, domain := range domains {
		if _, ok := domainsMap[domain]; ok {
			err = errors.New("domains duplicate")
			c.logger.Error(err)
			return
		}
		domainsMap[domain] = struct{}{}
	}

	for _, domain := range domains {
		c.allowedDomains[domain] = struct{}{}
	}

	return
}

func (c *Crawler) SetAllowedDomain(domain string) {
	c.allowedDomains[domain] = struct{}{}

	return
}

func (c *Crawler) DelAllowedDomain(domain string) (err error) {
	if domain == "*" {
		err = errors.New(`don't allow delete "*"`)
		c.logger.Error(err)
		return
	}

	if _, ok := c.allowedDomains[domain]; !ok {
		err = errors.New("domain not exists")
		c.logger.Error(err)
		return
	}

	delete(c.allowedDomains, domain)

	return
}

func (c *Crawler) CleanAllowedDomains() {
	c.allowedDomains = c.defaultAllowedDomains

	return
}
