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

    return {nodes, GetNodes, Count, CountActive}
})