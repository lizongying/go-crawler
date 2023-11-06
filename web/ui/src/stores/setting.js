import {defineStore} from 'pinia'
import {reactive, watch} from 'vue';

export const useSettingStore = defineStore('setting', () => {
    const setting = reactive({
        apiHost: 'http://localhost:8090',
    })

    watch(
        setting,
        (setting) => {
            localStorage.setItem('setting', JSON.stringify(setting))
        },
        {deep: true}
    )

    const SetApiHost = apiHost => {
        setting.apiHost = apiHost
    }

    const InitSetting = () => {
        const settingStr = localStorage.getItem('setting')
        if (settingStr !== '') {
            const settingObj = JSON.parse(settingStr)
            if (settingObj) {
                SetApiHost(settingObj.apiHost)
            }
        }
    }

    return {setting, InitSetting, SetApiHost}
})