---
weight: 9
title: Proxy
---

## Proxy

* You can set up tunnel proxies by using tools like [go-proxy](https://github.com/lizongying/go-proxy) to provide random
  proxy switching functionality, transparent to the caller. You can integrate these proxy tools into your spider
  framework to automatically switch proxies when making requests. The random switching tunnel proxy provides convenience
  and ease of use to the caller. In the future, other calling methods may be added, such as maintaining the original
  proxy address, to provide greater flexibility to meet different proxy requirements.

* Proxy Configuration in Spider

  Currently, only random switching of proxies is supported in the spider configuration.