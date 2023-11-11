import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getNodes} from "@/requests/api";

export const NodeStatusUnknown = 0
export const NodeStatusReady = 1
export const NodeStatusStarting = 2
export const NodeStatusRunning = 3
export const NodeStatusIdle = 4
export const NodeStatusStopping = 5
export const NodeStatusStopped = 6

export const useNodesStore = defineStore('nodes', () => {
    const nodes = reactive([])

    const GetNodes = () => {
        getNodes().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return
            }
            nodes.splice(0, nodes.length, ...resp.data.data)
        })
    }

    const Count = computed(() => {
        return nodes.length
    })

    const CountActive = computed(() => {
        return nodes.filter(v => v.status === 1).length
    })

    const CountSpider = computed(() => {
        if (nodes.length === 0) {
            return 0
        }

        return nodes.filter(v => v.spider).map(v => v.spider).reduce((a, b) => a + b, 0)
    })

    const CountJob = computed(() => {
        if (nodes.length === 0) {
            return 0
        }

        return nodes.filter(v => v.job).map(v => v.job).reduce((a, b) => a + b, 0)
    })

    const CountTask = computed(() => {
        if (nodes.length === 0) {
            return 0
        }

        return nodes.filter(v => v.task).map(v => v.task).reduce((a, b) => a + b, 0)
    })

    const CountRecord = computed(() => {
        if (nodes.length === 0) {
            return 0
        }

        return nodes.filter(v => v.record).map(v => v.record).reduce((a, b) => a + b, 0)
    })

    return {nodes, GetNodes, Count, CountActive, CountSpider, CountJob, CountTask, CountRecord}
})