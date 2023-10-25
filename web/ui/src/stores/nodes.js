import {defineStore} from 'pinia'
import {reactive} from 'vue';
import {getNodes} from "@/requests/api";

export const useNodesStore = defineStore('nodes', () => {
    const nodes = reactive([])

    const GetNodes = async () => {
        if (nodes.length === 0) {
            return getNodes()
        }
    }

    return {nodes, GetNodes}
})