import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getCrawlers} from "@/requests/api";

export const CrawlerStatusUnknown = 0
export const CrawlerStatusReady = 1
export const CrawlerStatusStarting = 2
export const CrawlerStatusRunning = 3
export const CrawlerStatusIdle = 4
export const CrawlerStatusStopping = 5
export const CrawlerStatusStopped = 6
export const CrawlerStatusName = (status) => {
    switch (status) {
        case CrawlerStatusReady:
            return 'ready'
        case CrawlerStatusStarting:
            return 'starting'
        case CrawlerStatusRunning:
            return 'running'
        case CrawlerStatusIdle:
            return 'idle'
        case CrawlerStatusStopping:
            return 'stopping'
        case CrawlerStatusStopped:
            return 'stopped'
        default:
            return 'unknown'
    }
}

export const useCrawlersStore = defineStore('crawlers', () => {
    const crawlers = reactive([])

    const GetCrawlers = () => {
        getCrawlers().then(resp => {
            if (resp.data.data === null) {
                crawlers.splice(0, crawlers.length)
                return
            }
            crawlers.splice(0, crawlers.length, ...resp.data.data)
        }).catch(e => {
            console.log(e);
        })
    }

    const Count = computed(() => {
        return crawlers.length
    })

    const CountSpider = computed(() => {
        if (crawlers.length === 0) {
            return 0
        }

        return crawlers.filter(v => v.spider).map(v => v.spider).reduce((a, b) => a + b, 0)
    })

    const CountJob = computed(() => {
        if (crawlers.length === 0) {
            return 0
        }

        return crawlers.filter(v => v.job).map(v => v.job).reduce((a, b) => a + b, 0)
    })

    const CountTask = computed(() => {
        if (crawlers.length === 0) {
            return 0
        }

        return crawlers.filter(v => v.task).map(v => v.task).reduce((a, b) => a + b, 0)
    })

    const CountItem = computed(() => {
        if (crawlers.length === 0) {
            return 0
        }

        return crawlers.filter(v => v.item).map(v => v.item).reduce((a, b) => a + b, 0)
    })

    return {crawlers, GetCrawlers, Count, CountSpider, CountJob, CountTask, CountItem}
})