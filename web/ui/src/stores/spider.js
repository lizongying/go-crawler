import {defineStore} from 'pinia'
import {reactive} from 'vue';
import {getSpider} from "@/requests/api";

export const SpiderStatusUnknown = 0
export const SpiderStatusReady = 1
export const SpiderStatusStarting = 2
export const SpiderStatusRunning = 3
export const SpiderStatusIdle = 4
export const SpiderStatusStopping = 5
export const SpiderStatusStopped = 6

export const useSpiderStore = defineStore('spider', () => {
    const spider = reactive({})

    const GetSpider = () => {
        getSpider().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return
            }
            spider.name = resp.data.data.name
            spider.funcs = resp.data.data.funcs
        })
    }

    return {spider, GetSpider}
})