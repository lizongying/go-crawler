import {defineStore} from 'pinia'
import {reactive} from 'vue';

export const useMessageStore = defineStore('message', () => {
    const message = reactive([])

    const GetMessage = async () => {
        if (message.length === 0) {
            message.push({
                level: 'warning',
                title: 'demo使用了自签发证书，如需查看，请安装并信任证书。建议在电脑上查看。',
                content: 'https://github.com/lizongying/go-crawler/blob/main/static/tls/ca.crt',
            })
        }
    }

    return {message, GetMessage}
})