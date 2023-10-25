import {defineStore} from 'pinia'
import {reactive} from 'vue';
import {getSpiders} from "@/requests/api";

export const useSpidersStore = defineStore('spiders', () => {
    const spiders = reactive([])

    const GetSpiders = async () => {
        if (spiders.length === 0) {
            return getSpiders()
        }
    }

    return {spiders, GetSpiders}
})