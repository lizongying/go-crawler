import {defineStore} from 'pinia'
import {reactive, watch} from "vue";
import {getUser} from "@/requests/api";

export const useUserStore = defineStore('user', () => {
    const user = reactive({
        username: 'admin',
        password: 'znU2LtswYWW8kbf5',
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
        return await getUser({
            username: user.username,
            password: user.password,
        }).then(resp => {
            console.log(resp.data.data)
            if (resp.data.data === null) {
                return
            }
            setToken(resp.data.data.token)
            setUserInfo({name: resp.data.data.user_info.name, rote: resp.data.data.user_info.rote})
            return user;
        }).catch(e => {
            console.log(e);
            return null
        })
    }
    const Logout = async () => {
        setToken('')
        setUserInfo({})
    }

    return {user, InitUser, Remember, Login, Logout}
})