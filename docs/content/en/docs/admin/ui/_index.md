---
weight: 2
title: UI
---

## UI

You can directly use https://lizongying.github.io/go-crawler/.

If you want to view the demo, please trust the certificate.
[ca](https://github.com/lizongying/go-crawler/blob/main/static/tls/ca.crt)

develop

```shell
npm run dev --prefix ./web/ui
```

docs develop

```shell
# docs
hugo server --source docs --noBuildLock

```

build

The web server is optional; you can use networking services like Nginx directly.

```shell
# ui
make web_ui

# server
make web_server

```

Run

```shell
./releases/web_server
```