import axios from 'axios'

const api = async () => {
    const useSettingStore = () => import('@/stores/setting')
    const userUserStore = () => import('@/stores/user')
    const settingStore = await useSettingStore()
    const userStore = await userUserStore()
    return {
        host: settingStore.useSettingStore().setting.apiHost,
        config: {
            headers: {
                'Content-Type': 'application/json',
                'X-API-Key': userStore.useUserStore().user.token,
            },
            timeout: 10000
        }
    }
}

const getLog = async data => {
    const {host, config} = await api()
    return new EventSource(`${host}/log?X-API-Key=${config.headers['X-API-Key']}&task_id=${data}`);
};

const getUser = async data => {
    const {host, config} = await api()
    return axios.post(host + '/user', data, config);
};

const getCrawlers = async data => {
    const {host, config} = await api()
    return axios.post(host + '/crawlers', data, config);
};

const getSpiders = async data => {
    const {host, config} = await api()
    return axios.post(host + '/spiders', data, config);
};

const getJobs = async data => {
    const {host, config} = await api()
    return axios.post(host + '/jobs', data, config);
};

const runJob = async data => {
    const {host, config} = await api()
    return axios.post(host + '/job/run', data, config);
};

const rerunJob = async data => {
    const {host, config} = await api()
    return axios.post(host + '/job/rerun', data, config);
};

const stopJob = async data => {
    const {host, config} = await api()
    return axios.post(host + '/job/stop', data, config);
};

const getTasks = async data => {
    const {host, config} = await api()
    return axios.post(host + '/tasks', data, config);
};

const getRequests = async data => {
    const {host, config} = await api()
    return axios.post(host + '/requests', data, config);
};

const getItems = async data => {
    const {host, config} = await api()
    return axios.post(host + '/items', data, config);
};

const getSpider = async data => {
    const {host, config} = await api()
    return axios.post(host + '/spider', data, config);
};

export {
    getLog,
    getUser,
    getCrawlers,
    getSpiders,
    getJobs,
    runJob,
    rerunJob,
    stopJob,
    getTasks,
    getRequests,
    getItems,
    getSpider
}