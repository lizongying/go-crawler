import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getJobs, stopJob} from "@/requests/api";

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

    const StopJob = async (data) => {
        await stopJob(data).then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return null
            }

            const jobIdx = jobs.findIndex(v => v.id === data.job_id)
            if (jobIdx > -1) {
                jobs[jobIdx].status = 2
            }
            return resp.data.data
        }).catch(e => {
            console.log(e)
            return null
        })
    }

    const Count = computed(() => {
        return jobs.length
    })

    const CountEnable = computed(() => {
        return jobs.filter(v => v.enable).length
    })

    return {jobs, GetJobs, StopJob, Count, CountEnable}
})