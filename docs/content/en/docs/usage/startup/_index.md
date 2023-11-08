---
weight: 13
title: Startup
---

## Startup

By configuring environment variables or parameters, you can start the crawler more flexibly, including selecting the
configuration file, specifying the spider's name, defining the initial method, passing additional parameters, and
setting the startup mode.

project Structure

* It is recommended to have one spider for each website (or sub-website) or each specific business. You don't need to
  split it too finely, nor do you need to include all websites and businesses in one spider.
* You can package each spider
  separately or combine multiple spiders together to reduce the number of files. However, during execution, only one
  spider can be started.

```go
app.NewApp(NewExample1Spider, NewExample2Spider).Run()
```

```shell
spider -c example.yml -n example -f TestOk -m once
```

* Configuration file path, must be configured. It is recommended to use different configuration files for different
  environments.
    * Environment variable `CRAWLER_CONFIG_FILE`
    * Startup parameter `-c`
* Spider name, must be configured.
    * Environment variable `CRAWLER_NAME`
    * Startup parameter `-n`
* Initial method, default is "Test". Please note that the case must be consistent.
    * Environment variable `CRAWLER_FUNC`
    * Startup parameter `-f`
* Additional parameters, this parameter is optional. It is recommended to use a JSON string. The parameters will be
  passed to the initial method.
    * Environment variable `CRAWLER_ARGS`
    * Startup parameter `-a`
* Startup mode, default is 0(manual). You can use different modes as needed
    * Environment variable `CRAWLER_MODE`
    * Startup parameter `-m`
    * You can use different modes as needed:
    * Optional values
        * 0: manual. Executes manually (default is no execution); can be managed through the API.
        * 1: once. Executes only once.
        * 2: loop. Executes repeatedly.
        * 3: cron. Executes at scheduled intervals.
* Scheduled task. This configuration is only applied when the mode is set to "cron", such as "1s/2i/3h/4d/5m/6w"
    * Environment variable `CRAWLER_SPEC`
    * Startup parameter `-s`