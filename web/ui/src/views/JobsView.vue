<template>
  <a-page-header
      title="Jobs"
      :sub-title="'Total: '+jobsStore.Count"
  >
    <template #extra>
      <a-button key="1" type="primary">New</a-button>
      <a-switch v-model:checked="checked1" checked-children="开" un-checked-children="关" @change="changeSwitch"/>
      <a-button key="2" @click="refresh" :disabled="checked1Disable">Refresh</a-button>
    </template>
  </a-page-header>
  <a-table :columns="columns" :data-source="jobsStore.jobs" :scroll="{ x: '100%' }">
    <template #headerCell="{ column }">
      <template v-if="column.dataIndex !== ''">
        <span style="font-weight: bold">
          {{ column.title }}
        </span>
      </template>
    </template>

    <template #bodyCell="{ column, record }">
      <template v-if="column.dataIndex === 'node'">
        <RouterLink :to="'/nodes?id='+record.node">
          {{ record.node }}
        </RouterLink>
      </template>
      <template v-if="column.dataIndex === 'spider'">
        <RouterLink :to="'/spiders?name='+record.spider">
          {{ record.spider }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'status'">
        <span>
          <a-tag
              :key="record.status"
              :color="record.status === JobStatusStopped ? 'volcano' : record.status === JobStatusRunning ? 'green' : 'geekblue'"
          >
            {{ jobStatusName(record.status) }}
          </a-tag>
        </span>
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
      <template v-else-if="column.dataIndex === 'task'">
        <RouterLink :to="'/tasks?job='+record.id">
          {{ record.task }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'record'">
        <RouterLink :to="'/records?job='+record.id">
          {{ record.record }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'action'">
        <span>
          <a v-if="record.status === JobStatusStopped" @click="rerun(record.spider, record.id)">Rerun</a>
          <a v-if="record.status === JobStatusRunning" @click="stop(record.spider, record.id)">Stop</a>
          <a-divider type="vertical"/>
          <a>Delete</a>
          <a-divider type="vertical"/>
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
import {ExclamationCircleOutlined, RightOutlined} from "@ant-design/icons-vue";
import {RouterLink} from "vue-router";
import {
  JobStatusIdle,
  JobStatusReady,
  JobStatusRunning,
  JobStatusStarting,
  JobStatusStopped,
  JobStatusStopping,
  useJobsStore
} from "@/stores/jobs";
import {formatDuration, formattedDate} from "@/utils/time";
import {sortBigInt, sortInt, sortStr} from "@/utils/sort";
import {createVNode, onBeforeUnmount, reactive, ref} from "vue";
import {Modal} from "ant-design-vue";

const columns = [
  {
    title: 'Id',
    dataIndex: 'id',
    width: 200,
    sorter: (a, b) => sortBigInt(a.id, b.id),
    defaultSortOrder: 'descend',
  },
  {
    title: 'Schedule',
    dataIndex: 'schedule',
    width: 150,
    sorter: (a, b) => sortStr(a.schedule, b.schedule),
  },
  {
    title: 'Command',
    dataIndex: 'command',
    width: 350,
    ellipsis: true,
  },
  {
    title: 'Node',
    dataIndex: 'node',
    width: 200,
    sorter: (a, b) => sortBigInt(a.node, b.node),
  },
  {
    title: 'Spider',
    dataIndex: 'spider',
    sorter: (a, b) => sortStr(a.spider, b.spider),
    width: 200,
  },
  {
    title: 'Status',
    dataIndex: 'status',
    width: 100,
    filters: [
      {
        text: 'ready',
        value: JobStatusReady,
      },
      {
        text: 'starting',
        value: JobStatusStarting,
      },
      {
        text: 'running',
        value: JobStatusRunning,
      },
      {
        text: 'idle',
        value: JobStatusIdle,
      },
      {
        text: 'stopping',
        value: JobStatusStopping,
      },
      {
        text: 'stopped',
        value: JobStatusStopped,
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
    title: 'Task',
    dataIndex: 'task',
    sorter: (a, b) => sortInt(a.task, b.task),
    width: 100,
  },
  {
    title: 'Record',
    dataIndex: 'record',
    sorter: (a, b) => sortInt(a.record, b.record),
    width: 100,
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: 250,
    fixed: 'right',
  },
];

const jobsStore = useJobsStore()
jobsStore.GetJobs()

const jobStatusName = (status) => {
  switch (status) {
    case JobStatusReady:
      return 'ready'
    case JobStatusStarting:
      return 'starting'
    case JobStatusRunning:
      return 'running'
    case JobStatusIdle:
      return 'idle'
    case JobStatusStopping:
      return 'stopping'
    case JobStatusStopped:
      return 'stopped'
    default:
      return 'unknown'
  }
}

const refresh = () => {
  jobsStore.GetJobs()
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

// rerun confirm
function rerun(spiderName, jobId) {
  Modal.confirm({
    title: 'Do you want to rerun the job?',
    icon: createVNode(ExclamationCircleOutlined),
    content: 'When clicked the OK button, the job will be rerun.',
    async onOk() {
      try {
        const res = await jobsStore.RerunJob({spider_name: spiderName, job_id: jobId})
        console.log(spiderName, jobId, res)
      } catch {
        return console.log('Oops errors!');
      }
    },
    onCancel() {
    },
  });
}

// stop confirm
function stop(spiderName, jobId) {
  Modal.confirm({
    title: 'Do you want to stop the job?',
    icon: createVNode(ExclamationCircleOutlined),
    content: 'When clicked the OK button, the job will be stop.',
    async onOk() {
      try {
        const res = await jobsStore.StopJob({spider_name: spiderName, job_id: jobId})
        console.log(spiderName, jobId, res)
      } catch {
        return console.log('Oops errors!');
      }
    },
    onCancel() {
    },
  });
}

const open = ref(false);
const more = reactive({})
const showDrawer = record => {
  open.value = true;
  more.status_list = Object.entries(record.status_list).map(([k, v]) => `${formattedDate(k / 1000000000)} ${jobStatusName(v)}`).reverse();
};
// status list
const activeKey = ref('1');

</script>
<style>
</style>
