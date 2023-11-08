---
weight: 15
title: Q & A
---

## Question

* In some frameworks, there is a presence of `start_urls`. How is it set up in this framework?

  In this framework, this approach has been removed. It's possible to explicitly create requests within the initial
  method and perform additional processing on those requests, which can actually be more convenient.
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
* What are the ways to improve spider performance?

  To improve the performance of the spider, you can consider disabling some unused middleware or pipelines to reduce
  unnecessary processing and resource consumption. Before disabling any middleware or pipeline, please assess its actual
  impact on the spider's performance. Ensure that disabling any part will not have a negative impact on the
  functionality.

* Why isn't item implemented as a distributed queue?

  The crawler processes its own items, and there is no need to handle items from other crawlers.
  Therefore, while the framework has reserved the architecture for distributed queues, it does not use external queues
  to replace the in-memory queue used by the program.
  If there are performance issues with processing, it is recommended to output the results to a queue.

* How to Set the Request Priority?

  Priorities are allowed to range from 0 to 2147483647.
  Priority 0 is the highest and will be processed first.
  Currently, only Redis-based priority queues are supported.

    ```go
    request.SetPriority(0)
    ```

* When will the crawler end?

  The crawler will end and the program will close when the following conditions are met under normal circumstances:

    1. All requests and parsing methods have been executed.
    2. The item queue is empty.
    3. The request queue is empty.

  When these conditions are fulfilled, the crawler has completed its tasks and will terminate.

* How to prevent the spider from stopping?

  Simply return `pkg.DontStopErr` in the `Stop` method.

    ```go
    package main
    
    import "github.com/lizongying/go-crawler/pkg"
    
    func (s *Spider) Stop(ctx context.Context) (err error) {
        if err = s.Spider.Stop(ctx); err != nil {
            s.logger.Error(err)
            return
        }
    
        err = pkg.DontStopErr
        s.logger.Error(err)
        return
    }

    ```

* Which should be used in the task queue: `request`, `extra`, or `unique_key`?

  Firstly, it should be noted that these three terms are concepts within this framework:
    * `request` contains all the fields of a request, including URL, method, headers, and may have undergone middleware
      processing. The drawback is that it occupies more space, making it somewhat wasteful as a queue value.
    * `extra` is a structured field within the request and, in the framework's design, it contains information that can
      construct a unique request (in most cases). For instance, a list page under a category may include the category ID
      and page number. Similarly, a detail page may include a detail ID. To ensure compatibility with various languages,
      the storage format in the queue is JSON, which is more space-efficient. It's recommended to use this option.
    * `unique_key` is a unique identifier for a request within the framework and is a string. While it can represent
      uniqueness in some cases, it can become cumbersome when requiring a combination of multiple fields to be unique –
      such as in the case of list pages or detail pages involving both a category and an ID. If memory is constrained (
      e.g., in Redis usage), it can be used. However, for greater generality, using `extra` might be more convenient.

  Enqueuing:
    * `YieldExtra` or `MustYieldExtra`

  Dequeuing:
    * `GetExtra` or `MustGetExtra`

* Whether to use `Must[method]`, such as `MustYieldRequest`?

  `Must[method]` is more concise, but it might be less convenient for troubleshooting errors. Whether to use it depends
  on the individual style of the user.
  If there's a need for specific error handling, then regular methods like `YieldRequest` should be used.