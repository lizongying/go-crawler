<template>
  <a-layout>
    <a-layout-sider v-if="!isLogin" v-model:collapsed="state.collapsed">
      <div style="text-align: center; height: 60px;">
        <a-typography-title :content="state.collapsed ? 'GO' : 'GO CRAWLER' " :level="3" ellipsis
                            style="color: white; line-height: 60px">
        </a-typography-title>
      </div>
      <a-menu
          v-model:openKeys="state.openKeys"
          v-model:selectedKeys="state.selectedKeys"
          :collapsed="state.collapsed"
          :items="items"
          mode="inline"
          theme="dark"
      ></a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <template v-if="!isLogin">
          <menu-unfold-outlined
              v-if="state.collapsed"
              class="trigger"
              @click="toggleCollapsed"
          />
          <menu-fold-outlined v-else class="trigger" @click="toggleCollapsed"/>
        </template>
        <a-space style="float: right; margin-right: 10px">
          <span style="margin-right: 10px"><a target="_blank" href="https://lizongying.github.io/go-crawler/docs/"><ReadOutlined/>  Docs</a></span>
          <span style="margin-right: 10px" @click="showSetting"><a><SettingOutlined/>  Setting</a></span>
          <a-modal v-model:open="openSetting" title="Setting" width="1000px" @ok="handleSetting">
            <a-form
                :label-col="{ span: 4 }"
                :model="formSetting"
                :wrapper-col="{ span: 20 }"
                autocomplete="off"
                name="basic"
            >
              <a-form-item
                  :rules="[{ required: true, message: 'Please input api host!' }]"
                  label="Api Host"
                  name="apiHost"
              >
                <a-input v-model:value="formSetting.apiHost" placeholder="http://localhost:8090"/>
              </a-form-item>
            </a-form>
          </a-modal>
          <span style="margin-right: 10px" @click="showModal"><a><MailOutlined/>  Message</a></span>
          <a-modal v-model:open="open" title="Message" width="1000px" @ok="handleOk">
            <a-space direction="vertical" style="width: 100%">
              <a-alert
                  v-for="msg in message"
                  :key="msg"
                  :description="msg.content"
                  :message="msg.title"
                  :type="msg.level"
                  show-icon
              />
            </a-space>
          </a-modal>
          <a-dropdown v-if="!isLogin">
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
      <a-layout-content style="height: calc(100vh - 64px);">
        <RouterView/>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script lang='jsx' setup>
import {h, reactive, ref, watch} from 'vue';
import {
  BarsOutlined, CloudDownloadOutlined,
  ClusterOutlined,
  DatabaseOutlined,
  DownOutlined,
  HourglassOutlined,
  MailOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PieChartOutlined,
  ReadOutlined,
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

const isLogin = ref(true);

router.beforeEach((to, from, next) => {
  if (to.name !== 'login' && !userStore.user.token) {
    router.push('/login')
    return
  }
  isLogin.value = to.name === 'login'
  switch (to.name) {
    case '':
      state.selectedKeys = ['1']
      break
    case 'crawlers':
      state.selectedKeys = ['2']
      break
    case 'spiders':
      state.selectedKeys = ['3']
      break
    case 'jobs':
      state.selectedKeys = ['4']
      break
    case 'tasks':
      state.selectedKeys = ['5']
      break
    case 'requests':
      state.selectedKeys = ['6']
      break
    case 'items':
      state.selectedKeys = ['7']
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
    label: <RouterLink to="/crawlers">Crawlers</RouterLink>,
    title: 'Crawlers',
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
    label: <RouterLink to="/jobs">Jobs</RouterLink>,
    title: 'Jobs',
  },
  {
    key: '5',
    icon: () => h(HourglassOutlined),
    label: <RouterLink to="/tasks">Tasks</RouterLink>,
    title: 'Tasks',
  },
  {
    key: '6',
    icon: () => h(CloudDownloadOutlined),
    label: <RouterLink to="/requests">Requests</RouterLink>,
    title: 'Requests',
  },
  {
    key: '7',
    icon: () => h(DatabaseOutlined),
    label: <RouterLink to="/items">Items</RouterLink>,
    title: 'Items',
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
})

const openSetting = ref(false);
const showSetting = () => {
  openSetting.value = true;
};
const handleSetting = _ => {
  openSetting.value = false;
  settingStore.SetApiHost(formSetting.apiHost)
};
</script>
