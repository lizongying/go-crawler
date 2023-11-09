import {defineStore} from 'pinia'
import {reactive} from 'vue';
import {getJobs, rerunJob, runJob, stopJob} from "@/requests/api";

export const JobStatusUnknown = 0
export const JobStatusReady = 1
export const JobStatusStarting = 2
export const JobStatusRunning = 3
export const JobStatusStopping = 4
export const JobStatusStopped = 5

export const useJobsStore = defineStore('jobs', () => {
    const jobs = reactive([])

    // const JobStatusUnknown = 0
    // const JobStatusReady = 1
    // const JobStatusStarting = 2
    // const JobStatusRunning = 3
    // const JobStatusStopping = 4
    // const JobStatusStopped = 5

    const GetJobs = () => {
        getJobs().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return
            }
            jobs.splice(0, jobs.length, ...resp.data.data)
        })
    }

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
                jobs[jobIdx].status = JobStatusStopped
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
        StopJob,
        RunJob,
        RerunJob,
    }
})