import {defineStore} from 'pinia'
import {reactive, watch} from "vue";

export const useUserStore = defineStore('user', () => {
    const user = reactive({
        username: '',
        password: '',
        remember: true,
        token: '',
        userInfo: {},
    })

    watch(
        user,
        (user) => {
            localStorage.setItem('user', JSON.stringify(user))
        },
        {deep: true}
    )

    const setToken = token => {
        user.token = token
    }

    const setUserInfo = userInfo => {
        user.userInfo = userInfo
    }

    const InitUser = () => {
        const userStr = localStorage.getItem('user')
        if (userStr !== '') {
            const userObj = JSON.parse(userStr)
            if (userObj) {
                Remember(userObj.username, userObj.password, userObj.remember)
                setToken(userObj.token)
                setUserInfo(userObj.userInfo)
            }
        }
    }

    const Remember = (username, password, remember) => {
        user.username = username
        user.password = password
        user.remember = remember
    }

    const Login = async () => {
        setToken('admin')
        setUserInfo({name: 'Admin'})
    }
    const Logout = async () => {
        setToken('')
        setUserInfo({})
    }

    return {user, InitUser, Remember, Login, Logout}
})