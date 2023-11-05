import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getJobs} from "@/requests/api";

export const useJobsStore = defineStore('jobs', () => {
    const jobs = reactive([])

    const GetJobs = () => {
        getJobs().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return
            }
            jobs.splice(0, jobs.length, ...resp.data.data)
        })
    }

    const Count = computed(() => {
        return jobs.length
    })

    const CountEnable = computed(() => {
        return jobs.filter(v => v.enable).length
    })

    return {jobs, GetJobs, Count, CountEnable}
})