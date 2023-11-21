import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getTasks} from "@/requests/api";

export const TaskStatusUnknown = 0
export const TaskStatusPending = 1
export const TaskStatusRunning = 2
export const TaskStatusSuccess = 3
export const TaskStatusFailure = 4

export const TaskStatusName = (status) => {
    switch (status) {
        case TaskStatusPending:
            return 'pending'
        case TaskStatusRunning:
            return 'running'
        case TaskStatusSuccess:
            return 'success'
        case TaskStatusFailure:
            return 'failure'
        default:
            return 'unknown'
    }
}

export const useTasksStore = defineStore('tasks', () => {
    const tasks = reactive([])

    const GetTasks = () => {
        getTasks().then(resp => {
            if (resp.data.data === null) {
                tasks.splice(0, tasks.length)
                return
            }
            tasks.splice(0, tasks.length, ...resp.data.data)
        }).catch(e => {
            console.log(e);
        })
    }

    const Count = computed(() => {
        return tasks.length
    })

    return {tasks, GetTasks, Count}
})