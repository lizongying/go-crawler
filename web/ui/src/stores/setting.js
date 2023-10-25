import {defineStore} from 'pinia'
import {reactive, watch} from 'vue';

export const useSettingStore = defineStore('setting', () => {
    const setting = reactive({
        apiHost: 'http://localhost:8090',
        apiAccessKey: '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918',
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

    const SetApiAccessKey = apiAccessKey => {
        setting.apiAccessKey = apiAccessKey
    }

    const InitSetting = () => {
        const settingStr = localStorage.getItem('setting')
        if (settingStr !== '') {
            const settingObj = JSON.parse(settingStr)
            if (settingObj) {
                SetApiHost(settingObj.apiHost)
                SetApiAccessKey(settingObj.apiAccessKey)
            }
        }
    }

    return {setting, InitSetting, SetApiHost, SetApiAccessKey}
})