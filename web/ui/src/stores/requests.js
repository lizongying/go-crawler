import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getRequests} from "@/requests/api";

export const RequestStatusUnknown = 0
export const RequestStatusPending = 1
export const RequestStatusRunning = 2
export const RequestStatusSuccess = 3
export const RequestStatusFailure = 4

export const RequestStatusName = (status) => {
    switch (status) {
        case RequestStatusPending:
            return 'pending'
        case RequestStatusRunning:
            return 'running'
        case RequestStatusSuccess:
            return 'success'
        case RequestStatusFailure:
            return 'failure'
        default:
            return 'unknown'
    }
}
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