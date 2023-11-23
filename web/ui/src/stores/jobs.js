import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getJobs, rerunJob, runJob, stopJob} from "@/requests/api";

export const JobStatusUnknown = 0
export const JobStatusReady = 1
export const JobStatusStarting = 2
export const JobStatusRunning = 3
export const JobStatusIdle = 4
export const JobStatusStopping = 5
export const JobStatusSuccess = 6
export const JobStatusFailure = 7

export const JobStatusName = (status) => {
    switch (status) {
        case JobStatusReady:
            return 'ready'
        case JobStatusStarting:
            return 'starting'
        case JobStatusRunning:
            return 'running'
        case JobStatusIdle:
            return 'idle'
        case JobStatusStopping:
            return 'stopping'
        case JobStatusSuccess:
            return 'success'
        case JobStatusFailure:
            return 'failure'
        default:
            return 'unknown'
    }
}

export const useJobsStore = defineStore('jobs', () => {
    const jobs = reactive([])

    const GetJobs = () => {
        getJobs().then(resp => {
            if (resp.data.data === null) {
                jobs.splice(0, jobs.length)
                return
            }
            jobs.splice(0, jobs.length, ...resp.data.data)
        }).catch(e => {
            console.log(e);
        })
    }

    const Count = computed(() => {
        return jobs.length
    })

    const RunJob = async (data) => {
        await runJob(data).then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return null
            }

            // const jobIdx = jobs.findIndex(v => v.id === data.job_id)
            // if (jobIdx > -1) {
            //     jobs[jobIdx].status = 2
            // }
            return resp.data.data
        }).catch(e => {
            console.log(e)
            return null
        })
    }

    const RerunJob = async (data) => {
        await rerunJob(data).then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return null
            }

            const jobIdx = jobs.findIndex(v => v.id === data.job_id)
            if (jobIdx > -1) {
                jobs[jobIdx].status = JobStatusRunning
            }
            return resp.data.data
        }).catch(e => {
            console.log(e)
            return null
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
                jobs[jobIdx].status = JobStatusSuccess
            }
            return resp.data.data
        }).catch(e => {
            console.log(e)
            return null
        })
    }

    return {
        jobs,
        GetJobs,
        Count,
        StopJob,
        RunJob,
        RerunJob,
    }
})