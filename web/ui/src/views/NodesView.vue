<template>
  <a-page-header
      title="Nodes"
      :sub-title="'Total: '+nodesStore.Count"
  >
    <template #extra>
      <a-switch v-model:checked="checked1" checked-children="开" un-checked-children="关" @change="changeSwitch"/>
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

    <template #bodyCell="{ column, record }">
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
                   {{ nodeStatusName(record.status) }}
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.dataIndex === 'enable'">
        <a-switch v-model:checked="record.enable"/>
      </template>
      <template v-else-if="column.dataIndex === 'action'">
        <span>
          <a class="ant-dropdown-link">
            More
            <RightOutlined/>
          </a>
        </span>
      </template>
    </template>
  </a-table>
</template>
<script setup>
import {RightOutlined} from "@ant-design/icons-vue";
import {RouterLink} from "vue-router";
import {
  NodeStatusReady,
  NodeStatusRunning,
  NodeStatusStarting,
  NodeStatusStopped,
  NodeStatusStopping,
  useNodesStore
} from "@/stores/nodes";
import {formatDuration, formattedDate} from "@/utils/time";
import {sortBigInt, sortInt, sortStr} from "@/utils/sort";
import {onBeforeUnmount, ref} from "vue";

const columns = [
  {
    title: 'Id',
    dataIndex: 'id',
    width: 200,
    sorter: (a, b) => sortBigInt(a.id, b.id),
    defaultSortOrder: 'descend',
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
        text: 'stopping',
        value: NodeStatusStopping,
      },
      {
        text: 'stopped',
        value: NodeStatusStopped,
      },
    ],
    onFilter: (value, record) => record.status === value,
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

const nodesStore = useNodesStore();

nodesStore.GetNodes()

const refresh = () => {
  nodesStore.GetNodes()
}
const checked1 = ref(false)
const checked1Disable = ref(false)

let interval = null
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


const nodeStatusName = (status) => {
  switch (status) {
    case NodeStatusReady:
      return 'ready'
    case NodeStatusStarting:
      return 'starting'
    case NodeStatusRunning:
      return 'running'
    case NodeStatusStopping:
      return 'stopping'
    case NodeStatusStopped:
      return 'stopped'
    default:
      return 'unknown'
  }
}
</script>
<style>
</style>
