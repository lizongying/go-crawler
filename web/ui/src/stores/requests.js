import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getRequests} from "@/requests/api";

export const useRequestsStore = defineStore('requests', () => {
    const requests = reactive([])

    const GetRequests = () => {
        getRequests().then(resp => {
            if (resp.data.data === null) {
                requests.splice(0, requests.length)
                return
            }
            requests.splice(0, requests.length, ...resp.data.data)
        }).catch(e => {
            console.log(e);
        })
    }

    const Count = computed(() => {
        return requests.length
    })

    return {requests, GetRequests, Count}
})