import axios from "axios";
import {useSettingStore} from "@/stores/setting";

// setting
const settingStore = useSettingStore();

const config = {
    headers: {
        'Content-Type': 'application/json',
        'X-API-Key': '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918'
    }
}

const getNodes = data => {
    return axios.post(settingStore.setting.apiHost + '/nodes', data, config);
};

const getSpiders = data => {
    return axios.post(settingStore.setting.apiHost + '/spiders', data, config);
};

const getJobs = data => {
    return axios.post(settingStore.setting.apiHost + '/jobs', data, config);
};

const getTasks = data => {
    return axios.post(settingStore.setting.apiHost + '/tasks', data, config);
};

const getRecords = data => {
    return axios.post(settingStore.setting.apiHost + '/records', data, config);
};

export {getNodes, getSpiders, getJobs, getTasks, getRecords}