<template>
  <a-spin :spinning="loading" tip="Loading...">
    <a-layout style="align-items: center;">
      <a-layout-header style="margin-top: 250px; width: 300px; text-align:center; background: #f5f5f5"><b>Login</b>
      </a-layout-header>
      <a-layout-content style="width: 300px">
        <a-form
            :model="formState"
            class="login-form"
            labelAlign="right"
            name="normal_login"
            @finish="onFinish"
            @finishFailed="onFinishFailed"
        >
          <a-form-item
              :rules="[{ required: true, message: 'Please input your username!' }]"
              name="username"
          >
            <a-input v-model:value="formState.username" placeholder="admin">
              <template #prefix>
                <UserOutlined class="site-form-item-icon"/>
              </template>
            </a-input>
          </a-form-item>

          <a-form-item
              :rules="[{ required: true, message: 'Please input your password!' }]"
              name="password"
          >
            <a-input-password v-model:value="formState.password" placeholder="admin">
              <template #prefix>
                <LockOutlined class="site-form-item-icon"/>
              </template>
            </a-input-password>
          </a-form-item>

          <a-form-item>
            <a-form-item name="remember" no-style>
              <a-checkbox v-model:checked="formState.remember">Remember me</a-checkbox>
            </a-form-item>
          </a-form-item>

          <a-form-item>
            <a-button :disabled="disabled" class="login-form-button" html-type="submit" type="primary">
              Log in
            </a-button>
          </a-form-item>
        </a-form>
      </a-layout-content>
    </a-layout>
  </a-spin>
</template>
<script setup>
import {computed, reactive, ref} from 'vue';
import {LockOutlined, UserOutlined} from "@ant-design/icons-vue";
import {useUserStore} from '@/stores/user'
import router from "../router";
import {message} from "ant-design-vue";

const userStore = useUserStore();

const formState = reactive({
  username: userStore.user.username,
  password: userStore.user.password,
  remember: userStore.user.remember,
});

const onFinish = values => {
  (async () => {
    if (values.username !== '' && values.password !== '') {
      loading.value = true
      const user = await userStore.Login()
      loading.value = false
      if (user === null) {
        message.error('Request Error');
        return
      }
      if (user.userInfo) {
        if (values.remember) {
          userStore.Remember(values.username, values.password, values.remember)
        } else {
          userStore.Remember('', '', values.remember)
        }
        await router.push('/')
      }
    } else {
      message.error('username or password error');
    }
  })()
};
const onFinishFailed = errorInfo => {
  console.log('Failed:', errorInfo);
};
const disabled = computed(() => {
  return !(formState.username && formState.password);
});

const loading = ref(false)
</script>
<style>
</style>
