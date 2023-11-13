import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getTasks} from "@/requests/api";

export const TaskStatusUnknown = 0
export const TaskStatusPending = 1
export const TaskStatusRunning = 2
export const TaskStatusSuccess = 3
export const TaskStatusError = 4

export const useTasksStore = defineStore('tasks', () => {
    const tasks = reactive([])

    const GetTasks = () => {
        getTasks().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                tasks.splice(0, tasks.length)
                return
            }
            tasks.splice(0, tasks.length, ...resp.data.data)
        })
    }

    const Count = computed(() => {
        return tasks.length
    })

    const CountError = computed(() => {
        return tasks.filter(v => v.status === 4).length
    })

    return {tasks, GetTasks, Count, CountError}
})