<template>
  <a-table :columns="columns" :data-source="tasksStore.tasks" :scroll="{ x: '100%' }">
    <template #headerCell="{ column }">
      <template v-if="['id', 'spider', 'status', 'schedule'].includes(column.dataIndex)">
        <span style="font-weight: bold">
          {{ column.title }}
        </span>
      </template>
    </template>

    <template #bodyCell="{ column, record }">
      <template v-if="column.dataIndex === 'spider'">
        <RouterLink :to="'/spiders?name='+record.spider" @click="$emit('router—change','3')">
          {{ record.spider }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'schedule'">
        <RouterLink :to="'/schedules?schedule='+record.schedule" @click="$emit('router—change','4')">
          {{ record.schedule }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'status'">
        <span>
          <a-tag
              :key="record.status"
              :color="record.status === 4 ? 'volcano' : record.status===2 ? 'green' : 'geekblue'"
          >
            {{ statusName(record.status) }}
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
      <template v-else-if="column.dataIndex === 'action'">
        <span>
          <a>Run</a>
          <a-divider type="vertical"/>
          <a>Delete</a>
          <a-divider type="vertical"/>
          <a class="ant-dropdown-link" @click="showDrawer">
            More
            <RightOutlined/>
          </a>
        </span>
      </template>
    </template>
  </a-table>
  <a-drawer v-model:open="open"
            size="large"
            :closable="false">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="Tab 1">Content of Tab Pane 1</a-tab-pane>
      <a-tab-pane key="2" tab="Tab 2" force-render>Content of Tab Pane 2</a-tab-pane>
      <a-tab-pane key="3" tab="Tab 3">Content of Tab Pane 3</a-tab-pane>
    </a-tabs>
  </a-drawer>
</template>
<script setup>
import {RightOutlined} from "@ant-design/icons-vue";
import {ref} from "vue";
import {RouterLink} from "vue-router";
import {useTasksStore} from "@/stores/tasks";
import {formatDuration, formattedDate} from "@/utils/time";

defineEmits(['router—change'])

const columns = [
  {
    title: 'Spider',
    dataIndex: 'spider',
    sorter: (a, b) => a.spider > b.spider,
    width: 200,
  },
  {
    title: 'Schedule',
    dataIndex: 'schedule',
    width: 100,
  },
  {
    title: 'Node',
    dataIndex: 'node',
    width: 300,
  },
  {
    title: 'Command',
    dataIndex: 'command',
    width: 150,
  },
  {
    title: 'Status',
    dataIndex: 'status',
    width: 100,
    filters: [
      {
        text: 'pending',
        value: 1,
      },
      {
        text: 'running',
        value: 2,
      },
      {
        text: 'success',
        value: 3,
      },
      {
        text: 'error',
        value: 4,
      },
    ],
    onFilter: (value, record) => record.status === value,
  },
  {
    title: 'Start Time',
    dataIndex: 'start_time',
    width: 200,
  },
  {
    title: 'Finish Time',
    dataIndex: 'finish_time',
    width: 200,
  },
  {
    title: 'Duration',
    dataIndex: 'duration',
    width: 100,
  },
  {
    title: 'Record',
    dataIndex: 'record',
    width: 100,
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: 200,
    fixed: 'right',
  },
];
// const data = [
//   {
//     id: '1',
//     spider: 'John Brown',
//     schedule: 'every day',
//     status: 'success',
//     command: 'go-crawler',
//     node: 'localhost',
//     started: '1 hour ago',
//     finished: '1 second ago',
//     duration: '1h',
//     records: 100,
//   },
//   {
//     id: '2',
//     spider: 'Jim Green',
//     schedule: 'every day',
//     status: 'running',
//     command: 'go-crawler',
//     node: 'localhost',
//     started: '1 hour ago',
//     finished: '1 second ago',
//     duration: '1h',
//     records: 100,
//   },
//   {
//     id: '3',
//     spider: 'Joe Black',
//     schedule: 'every day',
//     status: 'error',
//     command: 'go-crawler',
//     node: 'localhost',
//     started: '1 hour ago',
//     finished: '1 second ago',
//     duration: '1h',
//     records: 100,
//   },
// ];

const tasksStore = useTasksStore();

tasksStore.GetTasks()

const open = ref(false);
const showDrawer = () => {
  open.value = true;
};
const activeKey = ref('1');


const statusName = (status) => {
  switch (status) {
    case 1:
      return 'starting'
    case 2:
      return 'started'
    case 3:
      return 'stopping'
    case 4:
      return 'stopped'
    default:
      return 'unknown'
  }
}
</script>
<style>
</style>
