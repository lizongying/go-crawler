import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getItems} from "@/requests/api";

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
export const useItemsStore = defineStore('items', () => {
    const items = reactive([])

    const GetItems = () => {
        getItems().then(resp => {
            if (resp.data.data === null) {
                items.splice(0, items.length)
                return
            }
            items.splice(0, items.length, ...resp.data.data)
        }).catch(e => {
            console.log(e);
        })
    }

    const Count = computed(() => {
        return items.length
    })

    return {items, GetItems, Count}
})