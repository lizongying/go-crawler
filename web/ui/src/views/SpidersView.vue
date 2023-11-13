<template>
  <a-page-header
      title="Spiders"
      :sub-title="'Total: '+spidersStore.Count"
  >
    <template #extra>
      <a-switch v-model:checked="checked1" checked-children="auto" un-checked-children="close" @change="changeSwitch"/>
      <a-button key="2" @click="refresh" :disabled="checked1Disable">Refresh</a-button>
    </template>
  </a-page-header>
  <a-table :columns="columns" :data-source="spidersStore.spiders" :scroll="{ x: '100%' }">
    <template #headerCell="{ column }">
      <template
          v-if="column.dataIndex !== ''">
        <span style="font-weight: bold">
          {{ column.title }}
        </span>
      </template>
    </template>

    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'spider'">
        <a>
          {{ record.spider }}
        </a>
      </template>
      <template v-else-if="column.dataIndex === 'status'">
        <span>
          <a-tag
              :key="record.status"
              :color="record.status === SpiderStatusStopped ? 'volcano' : record.status===SpiderStatusRunning ? 'green' : 'geekblue'"
          >
            {{ spiderStatusName(record.status) }}
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.dataIndex === 'last_task_status'">
        <span>
          <a-tag
              :key="record.last_task_status"
              :color="record.last_task_status === TaskStatusError ? 'volcano' : record.last_task_status===TaskStatusRunning ? 'green' : 'geekblue'"
          >
            {{ taskStatusName(record.last_task_status) }}
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.dataIndex === 'node'">
        <RouterLink :to="'/nodes?id='+record.node">
          {{ record.node }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'job'">
        <RouterLink :to="'/jobs?spider='+record.spider">
          {{ record.job }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'task'">
        <RouterLink :to="'/tasks?spider='+record.spider">
          {{ record.task }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'record'">
        <RouterLink :to="'/records?spider='+record.spider">
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
      <template v-else-if="column.dataIndex === 'last_task_start_time'">
        {{ formattedDate(record.last_task_start_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'last_task_finish_time'">
        {{ formattedDate(record.last_task_finish_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'action'">
        <span>
          <template v-if="record.status === SpiderStatusStopped">
          <a>Run</a>
          <a-divider type="vertical"/>
          </template>
          <a class="ant-dropdown-link" @click="showDrawer(record)">
            More
            <RightOutlined/>
          </a>
        </span>
      </template>
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
import {RightOutlined} from "@ant-design/icons-vue";
import {RouterLink} from "vue-router";
import {
  SpiderStatusIdle,
  SpiderStatusReady,
  SpiderStatusRunning,
  SpiderStatusStarting,
  SpiderStatusStopped,
  SpiderStatusStopping,
  useSpidersStore
} from "@/stores/spiders";
import {formatDuration, formattedDate} from "@/utils/time";
import {sortBigInt, sortInt, sortStr} from "@/utils/sort";
import {onBeforeUnmount, reactive, ref} from "vue";
import {TaskStatusError, TaskStatusPending, TaskStatusRunning, TaskStatusSuccess} from "@/stores/tasks";

const columns = [
  {
    title: 'Id',
    dataIndex: 'id',
    width: 200,
    sorter: (a, b) => sortBigInt(a.id, b.id),
    defaultSortOrder: 'descend',
  },
  {
    title: 'Name',
    dataIndex: 'spider',
    width: 200,
    sorter: (a, b) => sortStr(a.spider, b.spider),
  },
  {
    title: 'Node',
    dataIndex: 'node',
    width: 200,
    sorter: (a, b) => sortBigInt(a.node, b.node),
  },
  {
    title: 'Status',
    dataIndex: 'status',
    width: 100,
    filters: [
      {
        text: 'ready',
        value: SpiderStatusReady,
      },
      {
        text: 'starting',
        value: SpiderStatusStarting,
      },
      {
        text: 'running',
        value: SpiderStatusRunning,
      },
      {
        text: 'idle',
        value: SpiderStatusIdle,
      },
      {
        text: 'stopping',
        value: SpiderStatusStopping,
      },
      {
        text: 'stopped',
        value: SpiderStatusStopped,
      },
    ],
    onFilter: (value, record) => record.status === value,
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
    title: 'Last Task Status',
    dataIndex: 'last_task_status',
    width: 200,
    filters: [
      {
        text: 'pending',
        value: TaskStatusPending,
      },
      {
        text: 'running',
        value: TaskStatusRunning,
      },
      {
        text: 'success',
        value: TaskStatusSuccess,
      },
      {
        text: 'error',
        value: TaskStatusError,
      },
    ],
    onFilter: (value, record) => record.last_task_status === value,
  },
  {
    title: 'Last Task Start Time',
    dataIndex: 'last_task_start_time',
    width: 200,
    sorter: (a, b) => a.last_task_start_time - b.last_task_start_time,
  },
  {
    title: 'Last Task Finish Time',
    dataIndex: 'last_task_finish_time',
    width: 200,
    sorter: (a, b) => {
      if (a.last_task_finish_time === b.last_task_finish_time) {
        return 0
      }
      const a_finish_time = a.last_task_finish_time !== 0 ? a.last_task_finish_time : Math.floor(Date.now() / 1000)
      const b_finish_time = b.last_task_finish_time !== 0 ? b.last_task_finish_time : Math.floor(Date.now() / 1000)
      return a_finish_time - b_finish_time
    },
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

const spidersStore = useSpidersStore();

spidersStore.GetSpiders()

const spiderStatusName = (status) => {
  switch (status) {
    case 1:
      return 'ready'
    case 2:
      return 'starting'
    case 3:
      return 'started'
    case 4:
      return 'idle'
    case 5:
      return 'stopping'
    case 6:
      return 'stopped'
    default:
      return 'unknown'
  }
}

const taskStatusName = (status) => {
  switch (status) {
    case 1:
      return 'pending'
    case 2:
      return 'running'
    case 3:
      return 'success'
    case 4:
      return 'error'
    default:
      return 'unknown'
  }
}
const refresh = () => {
  spidersStore.GetSpiders()
}
const checked1 = ref(true)
const checked1Disable = ref(true)

let interval = setInterval(refresh, 1000)
const changeSwitch = () => {
  if (checked1.value) {
    interval = setInterval(refresh, 1000)
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
  more.status_list = Object.entries(record.status_list).map(([k, v]) => `${formattedDate(k / 1000000000)} ${spiderStatusName(v)}`).reverse();
};
// status list
const activeKey = ref('1');

</script>
<style>
</style>
