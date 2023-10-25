<template>
  <a-layout>
    <a-layout-sider v-if="!isLogin" v-model:collapsed="state.collapsed">
      <div style="text-align: center; height: 60px;">
        <a-typography-title :content="state.collapsed ? 'GO' : 'GO CRAWLER' " ellipsis :level="3"
                            style="color: white; line-height: 60px">
        </a-typography-title>
      </div>
      <a-menu
          v-model:openKeys="state.openKeys"
          v-model:selectedKeys="state.selectedKeys"
          mode="inline"
          theme="dark"
          :collapsed="state.collapsed"
          :items="items"
      ></a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header v-if="!isLogin" style="background: #fff; padding: 0">
        <menu-unfold-outlined
            v-if="state.collapsed"
            class="trigger"
            @click="toggleCollapsed"
        />
        <menu-fold-outlined v-else class="trigger" @click="toggleCollapsed"/>
        <a-space style="float: right; margin-right: 10px">
          <span @click="showModal" style="margin-right: 10px"><a><MailOutlined/>  Message</a></span>
          <a-modal v-model:open="open" width="1000px" title="Message" @ok="handleOk">
            <a-space direction="vertical" style="width: 100%">
              <a-alert
                  v-for="msg in message"
                  :key="msg"
                  :message="msg.title"
                  :description="msg.content"
                  :type="msg.level === 'info' ? 'info': 'success'"
                  show-icon
              />
            </a-space>
          </a-modal>
          <span @click="showSetting" style="margin-right: 10px"><a><SettingOutlined/>  Setting</a></span>
          <a-modal v-model:open="openSetting" width="1000px" title="Setting" @ok="handleSetting">
            <a-form
                :model="formSetting"
                name="basic"
                :label-col="{ span: 4 }"
                :wrapper-col="{ span: 20 }"
                autocomplete="off"
            >
              <a-form-item
                  label="Api Host"
                  name="apiHost"
                  :rules="[{ required: true, message: 'Please input api host!' }]"
              >
                <a-input v-model:value="formSetting.apiHost" placeholder="http://localhost:8090"/>
              </a-form-item>
              <a-form-item
                  label="Api Access Key"
                  name="apiAccessKey"
                  :rules="[{ required: true, message: 'Please input api access key!' }]"
              >
                <a-input v-model:value="formSetting.apiAccessKey"
                         placeholder="8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"/>
              </a-form-item>
            </a-form>
          </a-modal>
          <a-dropdown>
            <a class="ant-dropdown-link" @click.prevent>
              <UserOutlined/>
              {{ userStore.user.userInfo.name }}
              <DownOutlined/>
            </a>
            <template #overlay>
              <a-menu>
                <a-menu-item>
                  <a href="javascript:" @click="logout">Logout</a>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </a-space>
      </a-layout-header>
      <a-layout-content style="height: 100vh;">
        <RouterView @router—change="routerChange"/>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script lang='jsx' setup>
import {h, reactive, ref, watch} from 'vue';
import {
  BarsOutlined,
  ClusterOutlined,
  DatabaseOutlined,
  DownOutlined,
  HourglassOutlined,
  MailOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PieChartOutlined,
  ScheduleOutlined,
  SettingOutlined,
  UserOutlined,
} from '@ant-design/icons-vue';
import router from "./router";

import {useUserStore} from './stores/user'
import {useSettingStore} from "@/stores/setting";
import {useMessageStore} from "@/stores/message";

const userStore = useUserStore();

userStore.InitUser()

const logout = () => {
  userStore.Logout()
  router.push('/login')
}

const isLogin = ref(false);

router.beforeEach((to, from, next) => {
  isLogin.value = to.path === '/login'
  if (to.path !== '/login' && !userStore.user.token) {
    router.push('/login')
    return
  }
  switch (to.path) {
    case '/':
      state.selectedKeys = ['1']
      break
    case '/nodes':
      state.selectedKeys = ['2']
      break
    case '/spiders':
      state.selectedKeys = ['3']
      break
    case '/schedules':
      state.selectedKeys = ['4']
      break
    case '/tasks':
      state.selectedKeys = ['5']
      break
    case '/records':
      state.selectedKeys = ['6']
      break
    default:
      state.selectedKeys = []
  }
  next();
});

const state = reactive({
  collapsed: false,
  selectedKeys: ['1'],
  openKeys: [],
  preOpenKeys: [],
});

const items = reactive([
  {
    key: '1',
    icon: () => h(PieChartOutlined),
    label: <RouterLink to="/">Home</RouterLink>,
    title: 'Home',
  },
  {
    key: '2',
    icon: () => h(ClusterOutlined),
    label: <RouterLink to="/nodes">Nodes</RouterLink>,
    title: 'Nodes',
  },
  {
    key: '3',
    icon: () => h(BarsOutlined),
    label: <RouterLink to="/spiders">Spiders</RouterLink>,
    title: 'Spiders',
  },
  {
    key: '4',
    icon: () => h(ScheduleOutlined),
    label: <RouterLink to="/schedules">Schedules</RouterLink>,
    title: 'Schedules',
  },
  {
    key: '5',
    icon: () => h(HourglassOutlined),
    label: <RouterLink to="/tasks">Tasks</RouterLink>,
    title: 'Tasks',
  },
  {
    key: '6',
    icon: () => h(DatabaseOutlined),
    label: <RouterLink to="/records">Records</RouterLink>,
    title: 'Records',
  }
]);
watch(
    () => state.openKeys,
    (_val, oldVal) => {
      state.preOpenKeys = oldVal;
    },
);
const toggleCollapsed = () => {
  state.collapsed = !state.collapsed;
  state.openKeys = state.collapsed ? [] : state.preOpenKeys;
};

const routerChange = index => {
  state.selectedKeys = [index]
}

// message
const messageStore = useMessageStore();

messageStore.GetMessage()

const message = reactive(messageStore.message)

const open = ref(false);
const showModal = () => {
  open.value = true;
};
const handleOk = e => {
  open.value = false;
};

// setting
const settingStore = useSettingStore();

settingStore.InitSetting()

const formSetting = reactive({
  apiHost: settingStore.setting.apiHost,
  apiAccessKey: settingStore.setting.apiAccessKey,
})

const openSetting = ref(false);
const showSetting = () => {
  openSetting.value = true;
};
const handleSetting = _ => {
  openSetting.value = false;
  settingStore.SetApiHost(formSetting.apiHost)
  settingStore.SetApiAccessKey(formSetting.apiAccessKey)
};
</script>
