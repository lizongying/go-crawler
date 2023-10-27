import {defineStore} from 'pinia'
import {computed, reactive} from 'vue';
import {getSpiders} from "@/requests/api";

export const useSpidersStore = defineStore('spiders', () => {
    const spiders = reactive([])

    const GetSpiders = () => {
        getSpiders().then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return
            }
            spiders.splice(0, spiders.length, ...resp.data.data)
        })
    }

    const Count = computed(() => {
        return spiders.length
    })

    const CountActive = computed(() => {
        return spiders.length
    })

    return {spiders, GetSpiders, Count, CountActive}
})