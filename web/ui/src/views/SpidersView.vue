<template>
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
              :color="record.status === 4 ? 'volcano' : record.status===2 ? 'green' : 'geekblue'"
          >
            {{ spiderStatusName(record.status) }}
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.dataIndex === 'last_task_status'">
        <span>
          <a-tag
              :key="record.last_task_status"
              :color="record.last_task_status === 4 ? 'volcano' : record.last_task_status===2 ? 'green' : 'geekblue'"
          >
            {{ taskStatusName(record.last_task_status) }}
          </a-tag>
        </span>
      </template>
      <template v-else-if="column.dataIndex === 'schedule'">
        <RouterLink :to="'/schedules?spider_name='+record.name" @click="$emit('router—change','4')">
          {{ record.schedule }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'task'">
        <RouterLink :to="'/tasks?spider_name='+record.name" @click="$emit('router—change','5')">
          {{ record.task }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'record'">
        <RouterLink :to="'/records?spider_name='+record.name" @click="$emit('router—change','6')">
          {{ record.record }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'start_time'">
        {{ formattedDate(record.start_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'finish_time'">
        {{ formattedDate(record.finish_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'last_task_start_time'">
        {{ formattedDate(record.last_task_start_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'last_task_finish_time'">
        {{ formattedDate(record.last_task_finish_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'action'">
        <span>
          <a>Run</a>
          <a-divider type="vertical"/>
          <a>Delete</a>
          <a-divider type="vertical"/>
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
import {useSpidersStore} from "@/stores/spiders";
import {reactive} from "vue";
import {formattedDate} from "@/utils/time";

defineEmits(['router—change'])

const columns = [
  {
    title: 'Node',
    dataIndex: 'node',
    width: 300,
    sorter: (a, b) => a.node - b.node,
  },
  {
    title: 'Name',
    dataIndex: 'spider',
    width: 200,
    sorter: (a, b) => a.spider - b.spider,
  },
  {
    title: 'Status',
    dataIndex: 'status',
    width: 200,
    filters: [
      {
        text: 'starting',
        value: 1,
      },
      {
        text: 'started',
        value: 2,
      },
      {
        text: 'stopping',
        value: 3,
      },
      {
        text: 'stopped',
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
    title: 'Last Task Status',
    dataIndex: 'last_task_status',
    width: 200,
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
    sorter: (a, b) => a.last_task_finish_time - b.last_task_finish_time,
  },
  {
    title: 'Schedule',
    dataIndex: 'schedule',
    width: 100,
    sorter: (a, b) => a.schedule - b.schedule,
  },
  {
    title: 'Task',
    dataIndex: 'task',
    width: 100,
    sorter: (a, b) => a.task - b.task,
  },
  {
    title: 'Record',
    dataIndex: 'record',
    width: 100,
    sorter: (a, b) => a.record - b.record,
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
