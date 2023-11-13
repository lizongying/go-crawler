import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getSpiders} from "@/requests/api";

export const SpiderStatusUnknown = 0
export const SpiderStatusReady = 1
export const SpiderStatusStarting = 2
export const SpiderStatusRunning = 3
export const SpiderStatusIdle = 4
export const SpiderStatusStopping = 5
export const SpiderStatusStopped = 6

export const useSpidersStore = defineStore('spiders', () => {
    const spiders = reactive([])

    const GetSpiders = () => {
        getSpiders().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                spiders.splice(0, spiders.length)
                return
            }
            spiders.splice(0, spiders.length, ...resp.data.data)
        })
    }

    const SpiderNames = computed(() => {
        return spiders.filter(v => v.spider).map(v => {
            return {value: v.spider, label: v.spider}
        })
    })

    const SpiderFuncs = (id) => {
        if (!id) {
            return
        }
        const spider = spiders.find(v => v.spider === id)
        if (spider === null) {
            return
        }
        console.log(spider.funcs.map(v => {
            return {value: v, label: v}
        }))
        return spider.funcs.map(v => {
            return {value: v, label: v}
        })
    }

    const Count = computed(() => {
        return spiders.length
    })

    const CountActive = computed(() => {
        return spiders.length
    })

    return {spiders, GetSpiders, SpiderNames, SpiderFuncs, Count, CountActive}
})