import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getRecords} from "@/requests/api";

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