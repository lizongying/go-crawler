import {defineStore} from 'pinia'
import {reactive} from 'vue';
import {getRecords} from "@/requests/api";

export const useRecordsStore = defineStore('records', () => {
    const records = reactive([])

    const GetRecords = async () => {
        if (records.length === 0) {
            return getRecords()
        }
    }

    return {records, GetRecords}
})