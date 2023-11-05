<template>
  <a-table :columns="columns" :data-source="tasksStore.tasks" :scroll="{ x: '100%' }">
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
      <template v-else-if="column.dataIndex === 'spider'">
        <RouterLink :to="'/spiders?name='+record.spider">
          {{ record.spider }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'job'">
        <RouterLink :to="'/jobs?id='+record.job">
          {{ record.job }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'status'">
        <span>
          <a-tag
              :key="record.status"
              :color="record.status === 4 ? 'volcano' : record.status===2 ? 'green' : 'geekblue'"
          >
            {{ taskStatusName(record.status) }}
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
      <template v-else-if="column.dataIndex === 'record'">
        <RouterLink :to="'/records?task='+record.id">
          {{ record.record }}
        </RouterLink>
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
            :closable="false"
            size="large">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="Log"></a-tab-pane>
    </a-tabs>
  </a-drawer>
</template>
<script setup>
import {RightOutlined} from "@ant-design/icons-vue";
import {ref} from "vue";
import {RouterLink} from "vue-router";
import {useTasksStore} from "@/stores/tasks";
import {formatDuration, formattedDate} from "@/utils/time";

const columns = [
  {
    title: 'Id',
    dataIndex: 'id',
    width: 300,
    sorter: (a, b) => a.id - b.id,
  },
  {
    title: 'Node',
    dataIndex: 'node',
    width: 300,
    sorter: (a, b) => a.node - b.node,
  },
  {
    title: 'Spider',
    dataIndex: 'spider',
    sorter: (a, b) => a.spider - b.spider,
    width: 200,
  },
  {
    title: 'Job',
    dataIndex: 'job',
    width: 300,
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
    sorter: (a, b) => a.start_time - b.start_time,
  },
  {
    title: 'Finish Time',
    dataIndex: 'finish_time',
    width: 200,
    sorter: (a, b) => a.finish_time - b.finish_time,
  },
  {
    title: 'Duration',
    dataIndex: 'duration',
    width: 100,
    sorter: (a, b) => a.duration - b.duration,
  },
  {
    title: 'Record',
    dataIndex: 'record',
    width: 150,
    sorter: (a, b) => a.record - b.record,
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: 200,
    fixed: 'right',
  },
];

const tasksStore = useTasksStore();

tasksStore.GetTasks()

const open = ref(false);
const showDrawer = () => {
  open.value = true;
};
const activeKey = ref('1');

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
</script>
<style>
</style>
