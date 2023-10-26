import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getRecords} from "@/requests/api";

export const useRecordsStore = defineStore('records', () => {
    const records = reactive([])

    const GetRecords = () => {
        getRecords().then(resp => {
            console.log(resp.data.data)
            records.splice(0, records.length, ...resp.data.data)
        })
    }

    const Count = computed(() => {
        return records.length
    })

    const CountActive = computed(() => {
        return records.length
    })

    return {records, GetRecords, Count, CountActive}
})