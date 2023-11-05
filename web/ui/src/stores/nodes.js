import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getNodes} from "@/requests/api";

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

    return {nodes, GetNodes, Count, CountActive, CountTask, CountRecord}
})