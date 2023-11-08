---
weight: 13
title: 启动
---

## 启动

通过配置环境变量或参数，您可以更灵活地启动爬虫，包括选择配置文件、指定爬虫名称、指定初始方法、传递额外参数以及设定启动模式。

```shell
spider -c example.yml -n example -f TestOk -m once
```

* 配置文件路径，必须进行配置。建议不同环境使用不同的配置文件。
    * 环境变量 `CRAWLER_CONFIG_FILE`
    * 启动参数 `-c`
* 爬虫名称，必须进行配置。
    * 环境变量 `CRAWLER_NAME`
    * 启动参数 `-n`
* 初始方法名称，默认Test，注意大小写需一致。
    * 环境变量 `CRAWLER_FUNC`
    * 启动参数 `-f`
* 额外的参数，该参数是非必须项。建议使用JSON字符串。参数会被传递到初始方法中。
    * 环境变量 `CRAWLER_ARGS`
    * 启动参数 `-a`
* 模式，默认为0(manual)。您可以根据需要使用不同的模式。
    * 环境变量 `CRAWLER_MODE`
    * 启动参数 `-m`
    * 可选值
        * 0: manual 手动执行，默认不执行，可以通过api进行管理。
        * 1: once 只执行一次
        * 2: loop 一直重复执行
        * 3: cron 定时执行
* 定时任务。只有在模式为cron下，才会应用此配置。如"1s/2i/3h/4d/5m/6w"
    * 环境变量 `CRAWLER_SPEC`
    * 启动参数 `-s`