import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getSchedules} from "@/requests/api";

export const useSchedulesStore = defineStore('schedules', () => {
    const schedules = reactive([])

    const GetSchedules = () => {
        getSchedules().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return
            }
            schedules.splice(0, schedules.length, ...resp.data.data)
        })
    }

    const Count = computed(() => {
        return schedules.length
    })

    const CountActive = computed(() => {
        return schedules.length
    })

    return {schedules, GetSchedules, Count, CountActive}
})