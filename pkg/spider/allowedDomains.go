package spider

import (
	"github.com/pkg/errors"
	"net/url"
	"path/filepath"
)

func (s *BaseSpider) IsAllowedDomain(Url *url.URL) (ok bool) {
	for domain := range s.allowedDomains {
		matched, err := filepath.Match(domain, Url.Hostname())
		if err != nil {
			s.Logger.Warn(err)
			continue
		}
		if matched {
			ok = true
			return
		}
	}

	return
}

func (s *BaseSpider) GetAllowedDomains() (domains []string) {
	for domain := range s.allowedDomains {
		domains = append(domains, domain)
	}

	return
}

func (s *BaseSpider) ReplaceAllowedDomains(domains []string) (err error) {
	if len(domains) == 0 {
		err = errors.New("domains is empty")
		s.Logger.Error(err)
		return
	}

	domainsMap := make(map[string]struct{})
	for _, domain := range domains {
		if _, ok := domainsMap[domain]; ok {
			err = errors.New("domains duplicate")
			s.Logger.Error(err)
			return
		}
		domainsMap[domain] = struct{}{}
	}

	for _, domain := range domains {
		s.allowedDomains[domain] = struct{}{}
	}

	return
}

func (s *BaseSpider) SetAllowedDomain(domain string) {
	s.allowedDomains[domain] = struct{}{}

	return
}

func (s *BaseSpider) DelAllowedDomain(domain string) (err error) {
	if domain == "*" {
		err = errors.New(`don't allow delete "*"`)
		s.Logger.Error(err)
		return
	}

	if _, ok := s.allowedDomains[domain]; !ok {
		err = errors.New("domain not exists")
		s.Logger.Error(err)
		return
	}

	delete(s.allowedDomains, domain)

	return
}

func (s *BaseSpider) CleanAllowedDomains() {
	s.allowedDomains = s.defaultAllowedDomains

	return
}
