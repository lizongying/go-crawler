import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getTasks} from "@/requests/api";

export const useTasksStore = defineStore('tasks', () => {
    const tasks = reactive([])

    const GetTasks = () => {
        getTasks().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
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