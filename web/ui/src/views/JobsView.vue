<template>
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
import {useJobsStore} from "@/stores/jobs";
import {formatDuration, formattedDate} from "@/utils/time";

const columns = [
  {
    title: 'Id',
    dataIndex: 'id',
    width: 300,
    sorter: (a, b) => a.id - b.id,
  },
  {
    title: 'Schedule',
    dataIndex: 'schedule',
    width: 150,
    sorter: (a, b) => a.schedule - b.schedule,
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
    width: 150,
    sorter: (a, b) => a.duration - b.duration,
  },
  {
    title: 'Task',
    dataIndex: 'task',
    sorter: (a, b) => a.task - b.task,
    width: 100,
  },
  {
    title: 'Record',
    dataIndex: 'record',
    sorter: (a, b) => a.record - b.record,
    width: 100,
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: 200,
    fixed: 'right',
  },
];

const jobsStore = useJobsStore()
jobsStore.GetJobs()
</script>
<style>
</style>
