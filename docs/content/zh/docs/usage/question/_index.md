---
weight: 15
title: 问答
---

## 问答

* 一些框架里都有start_urls，此框架中怎么设置？

  本框架里，去掉了这种方式。可以显式地在初始方法里建立request，可以对request进行额外地处理，实际上可能会更方便些。
    ```go
    startUrls := []string{"/a.html", "/b.html"}
    for _, v:=range startUrls {
		if err = s.YieldRequest(ctx, request.NewRequest().
            SetUrl(fmt.Sprintf("https://a.com%s", v)).
            SetCallBack(s.Parse)); err != nil {
            s.logger.Error(err)
        }
    }

    ```

* 有哪些可以提高爬虫性能的方式？

  要提高爬虫的性能，您可以考虑关闭一些未使用的中间件或Pipeline，以减少不必要的处理和资源消耗。在禁用中间件或Pipeline之前，请评估其对爬虫性能的实际影响。确保禁用的部分不会对功能产生负面影响。

* 为什么item没有实现分布式队列？

  由爬虫处理自己的请求即可，没必要处理其他爬虫的请求。
  所以本框架虽架构上有预留，但不会去用其他外部队列代替本程序内存队列。
  如处理出现性能问题，建议将结果输出到队列。

* 如何设定请求的优先级？

  优先级允许0-2147483647。
  0的优先级最高，最先被处理。
  暂只支持基于redis的优先级队列。
  使用方法

    ```go
    request.SetPriority(0)
    ```

* 爬虫什么时候结束？

  正常情况下，达到以下条件，会判定任务结束，程序关闭：

    1. 请求和解析方法都已执行完毕
    2. item队列为空
    3. request队列为空

* 如何阻止爬虫停止？

  在`Stop`方法中返回`pkg.DontStopErr`即可

  ```go
  package main
  
  import "github.com/lizongying/go-crawler/pkg"
  
  func (s *Spider) Stop(_ pkg.Context) (err error) {
      err = pkg.DontStopErr
      return
  }
  
  ```

* 任务队列使用`request`、`extra`还是`unique_key`?

  首先说明的是，这三个词都是本框架中的概念：
    * `request` 包含了request的所有字段，包括url、method、headers等，甚至经过了中间件处理。缺点是占用空间大，作为队列的值有点浪费。
    * `extra`
      是request中的一个结构体字段，在框架的设计里是包含能够构造唯一的请求（大多数情况下）。比如一个分类下的列表页，可能包含分类id、页码；比如一个详情页，可能包含详情id。为了兼容更多的语言，在队列中的存储形式为json格式，比较节约空间，推荐使用。
    * `unique_key`
      是框架里请求的唯一标识，是一个字符串。在一些情况下，是可以代表唯一的，但在需要多个字段联合唯一的情况下会比较麻烦，比如列表页，比如分类加id的详情页等。如果内存（redis等使用）紧张，可以使用。但为了更加通用，可能使用`extra`
      更加方便。

  入队：
    * `YieldExtra`或`MustYieldExtra`

  出队:
    * `GetExtra`或`MustGetExtra`

* 该不该使用`Must[method]`，如`MustYieldRequest`?

  `Must[method]`更加简洁，但可能对于排查错误不太方便。是不是用，需要看使用者的个人风格。
  如果需要特殊处理err，就需要使用普通的方法了，如`YieldRequest`。

* 其他
  
  * 升级go-crawl
  * 清理缓存