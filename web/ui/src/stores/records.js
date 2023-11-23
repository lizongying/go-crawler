import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getRecords} from "@/requests/api";

export const ItemStatusUnknown = 0
export const ItemStatusPending = 1
export const ItemStatusRunning = 2
export const ItemStatusSuccess = 3
export const ItemStatusFailure = 4

export const ItemStatusName = (status) => {
    switch (status) {
        case ItemStatusPending:
            return 'pending'
        case ItemStatusRunning:
            return 'running'
        case ItemStatusSuccess:
            return 'success'
        case ItemStatusFailure:
            return 'failure'
        default:
            return 'unknown'
    }
}
export const useRecordsStore = defineStore('records', () => {
    const records = reactive([])

    const GetRecords = () => {
        getRecords().then(resp => {
            if (resp.data.data === null) {
                records.splice(0, records.length)
                return
            }
            records.splice(0, records.length, ...resp.data.data)
        }).catch(e => {
            console.log(e);
        })
    }

    const Count = computed(() => {
        return records.length
    })

    return {records, GetRecords, Count}
})