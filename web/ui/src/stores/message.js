import {defineStore} from 'pinia'
import {reactive} from 'vue';

export const useMessageStore = defineStore('message', () => {
    const message = reactive([])

    const GetMessage = async () => {
        if (message.length === 0) {
            // message.push({
            //     level: 'info',
            //     title: 'title',
            //     content: 'content',
            // })
        }
    }

    return {message, GetMessage}
})