<template>
  <a-page-header
      title="Nodes"
      :sub-title="'Total: '+nodesStore.Count"
  >
    <template #extra>
      <a-switch v-model:checked="checked1" checked-children="auto" un-checked-children="close" @change="changeSwitch"/>
      <a-button key="2" @click="refresh" :disabled="checked1Disable">Refresh</a-button>
    </template>
  </a-page-header>
  <a-table :columns="columns" :data-source="nodesStore.nodes" :scroll="{ x: '100%' }">
    <template #headerCell="{ column }">
      <template v-if="column.dataIndex !== ''">
        <span style="font-weight: bold">
          {{ column.title }}
        </span>
      </template>
    </template>
    <template
        #customFilterDropdown="{ setSelectedKeys, selectedKeys, confirm, clearFilters, column }"
    >
      <div style="padding: 8px">
        <a-input
            ref="searchInput"
            :placeholder="`Search ${column.dataIndex}`"
            :value="selectedKeys[0]"
            style="width: 188px; margin-bottom: 8px; display: block"
            @change="e => setSelectedKeys(e.target.value ? [e.target.value] : [])"
            @pressEnter="handleSearch(selectedKeys, confirm, column.dataIndex)"
        />
        <a-button
            type="primary"
            size="small"
            style="width: 90px; margin-right: 8px"
            @click="handleSearch(selectedKeys, confirm, column.dataIndex)"
        >
          <template #icon>
            <SearchOutlined/>
          </template>
          Search
        </a-button>
        <a-button size="small" style="width: 90px" @click="handleReset(clearFilters)">
          Reset
        </a-button>
      </div>
    </template>
    <template #customFilterIcon="{ filtered }">
      <search-outlined :style="{ color: filtered ? '#108ee9' : undefined }"/>
    </template>
    <template #bodyCell="{ text, column, record }">
      <template v-if="column.dataIndex === 'spider'">
        <RouterLink :to="'/spiders?node='+record.id">
          {{ record.spider }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'job'">
        <RouterLink :to="'/jobs?node='+record.id">
          {{ record.job }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'task'">
        <RouterLink :to="'/tasks?node='+record.id">
          {{ record.task }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'record'">
        <RouterLink :to="'/records?node='+record.id">
          {{ record.record }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'start_time'">
        {{ formattedDate(record.start_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'finish_time'">
        {{ formattedDate(record.finish_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'duration'">
        {{ formatDuration(record.finish_time - record.start_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'status'">
        <span>
          <a-tag
              :key="record.status"
              :color="record.status === NodeStatusStopped ? 'volcano' : record.status === NodeStatusRunning ? 'green' : 'geekblue'"
          >
                   {{ NodeStatusName(record.status) }}
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.dataIndex === 'enable'">
        <a-switch v-model:checked="record.enable"/>
      </template>
      <template v-else-if="column.dataIndex === 'action'">
        <span>
          <a class="ant-dropdown-link" @click="showDrawer(record)">
            More
            <RightOutlined/>
          </a>
        </span>
      </template>
      <span v-if="state.searchText && state.searchedColumn === column.dataIndex">
        <template
            v-for="(fragment, i) in text
            .toString()
            .split(new RegExp(`(?<=${state.searchText})|(?=${state.searchText})`, 'i'))"
        >
          <mark
              v-if="fragment.toLowerCase() === state.searchText.toLowerCase()"
              :key="i"
              class="highlight"
          >
            {{ fragment }}
          </mark>
          <template v-else>{{ fragment }}</template>
        </template>
      </span>
    </template>
  </a-table>
  <a-drawer v-model:open="open"
            :closable="false"
            size="large">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="Status List">
        <a-list size="small" bordered :data-source="more.status_list">
          <template #renderItem="{ item }">
            <a-list-item>{{ item }}</a-list-item>
          </template>
        </a-list>
      </a-tab-pane>
    </a-tabs>
  </a-drawer>
</template>
<script setup>
import {RightOutlined, SearchOutlined} from "@ant-design/icons-vue";
import {RouterLink, useRoute} from "vue-router";
import {
  NodeStatusIdle,
  NodeStatusName,
  NodeStatusReady,
  NodeStatusRunning,
  NodeStatusStarting,
  NodeStatusStopped,
  NodeStatusStopping,
  useNodesStore
} from "@/stores/nodes";
import {formatDuration, formattedDate} from "@/utils/time";
import {sortBigInt, sortInt, sortStr} from "@/utils/sort";
import {computed, onBeforeUnmount, reactive, ref} from "vue";

const filteredInfo = reactive({});
const {query} = useRoute();
Object.entries(query).forEach(([k, v]) => {
  filteredInfo[k] = [v]
});
const columns = computed(() => {
  return [
    {
      title: 'Id',
      dataIndex: 'id',
      width: 200,
      sorter: (a, b) => sortBigInt(a.id, b.id),
      defaultSortOrder: 'descend',
      customFilterDropdown: true,
      filteredValue: filteredInfo.id || null,
      onFilter: (value, record) =>
          record.id.toString().toLowerCase().includes(value.toLowerCase()),
      onFilterDropdownOpenChange: visible => {
        if (visible) {
          setTimeout(() => {
            searchInput.value.focus();
          }, 100);
        }
      },
    },
    {
      title: 'Hostname',
      dataIndex: 'hostname',
      width: 200,
      ellipsis: true,
      sorter: (a, b) => sortStr(a.hostname, b.hostname),
    },
    {
      title: 'Ip',
      dataIndex: 'ip',
      width: 200,
      ellipsis: true,
      sorter: (a, b) => sortStr(a.ip, b.ip),
    },
    {
      title: 'Start Time',
      dataIndex: 'start_time',
      width: 200,
      sorter: (a, b) => a.start_time - b.start_time,
    },
    {
      title: 'Finish Time',
      dataIndex: 'finish_time',
      width: 200,
      sorter: (a, b) => {
        if (a.finish_time === b.finish_time) {
          return 0
        }
        const a_finish_time = a.finish_time !== 0 ? a.finish_time : Math.floor(Date.now() / 1000)
        const b_finish_time = b.finish_time !== 0 ? b.finish_time : Math.floor(Date.now() / 1000)
        return a_finish_time - b_finish_time
      },
    },
    {
      title: 'Duration',
      dataIndex: 'duration',
      width: 150,
      sorter: (a, b) => {
        let a_finish_time = a.finish_time
        if (a.start_time === 0 && a.finish_time === 0) {
          a_finish_time = Math.floor(Date.now() / 1000)
        }
        let b_finish_time = b.finish_time
        if (b.start_time === 0 && b.finish_time === 0) {
          b_finish_time = Math.floor(Date.now() / 1000)
        }
        return (a_finish_time - a.start_time) - (b_finish_time - b.start_time)
      },
    },
    {
      title: 'Enable',
      dataIndex: 'enable',
      width: 100,
      filters: [
        {
          text: 'enable',
          value: true,
        },
        {
          text: 'disable',
          value: false,
        },
      ],
      onFilter: (value, record) => record.enable === value,
      filteredValue: null,
    },
    {
      title: 'Status',
      dataIndex: 'status',
      width: 100,
      filters: [
        {
          text: 'ready',
          value: NodeStatusReady,
        },
        {
          text: 'starting',
          value: NodeStatusStarting,
        },
        {
          text: 'running',
          value: NodeStatusRunning,
        },
        {
          text: 'idle',
          value: NodeStatusIdle,
        },
        {
          text: 'stopping',
          value: NodeStatusStopping,
        },
        {
          text: 'stopped',
          value: NodeStatusStopped,
        },
      ],
      onFilter: (value, record) => record.status === value,
      filteredValue: null,
    },
    {
      title: 'Spider',
      dataIndex: 'spider',
      width: 100,
      sorter: (a, b) => sortInt(a.spider, b.spider),
    },
    {
      title: 'Job',
      dataIndex: 'job',
      width: 100,
      sorter: (a, b) => sortInt(a.job, b.job),
    },
    {
      title: 'Task',
      dataIndex: 'task',
      width: 100,
      sorter: (a, b) => sortInt(a.task, b.task),
    },
    {
      title: 'Record',
      dataIndex: 'record',
      width: 100,
      sorter: (a, b) => sortInt(a.record, b.record),
    },
    {
      title: 'Action',
      dataIndex: 'action',
      width: 200,
      fixed: 'right',
    },
  ];
});

const nodesStore = useNodesStore();

// auto refresh
const checked1 = ref(true)
const checked1Disable = ref(true)
let interval = 0
const refresh = () => {
  nodesStore.GetNodes()
}
refresh()
if (checked1.value) {
  interval = setInterval(refresh, 1000)
}
const changeSwitch = () => {
  if (checked1.value) {
    if (!checked1Disable.value) {
      interval = setInterval(refresh, 1000)
    }
    checked1Disable.value = true
  } else {
    clearInterval(interval)
    checked1Disable.value = false
  }
}
onBeforeUnmount(() => {
  clearInterval(interval)
})

const open = ref(false);
const more = reactive({})
const showDrawer = record => {
  open.value = true;
  more.status_list = Object.entries(record.status_list).map(([k, v]) => `${formattedDate(k / 1000000000)} ${NodeStatusName(v)}`).reverse();
};
// status list
const activeKey = ref('1');

// search
const state = reactive({
  searchText: '',
  searchedColumn: '',
});
const searchInput = ref();
const handleSearch = (selectedKeys, confirm, dataIndex) => {
  filteredInfo[dataIndex] = selectedKeys
  confirm();
  state.searchText = selectedKeys[0];
  state.searchedColumn = dataIndex;
};
const handleReset = clearFilters => {
  Object.keys(filteredInfo).forEach(key => {
    delete filteredInfo[key];
  });
  clearFilters({
    confirm: true,
  });
  state.searchText = '';
};
</script>
<style>
</style>
