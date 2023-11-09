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
      <template v-else-if="column.dataIndex === 'status'">
        <span>
          <a-tag
              :key="record.status"
              :color="record.status === 2 ? 'volcano' : record.status === 1 ? 'green' : 'geekblue'"
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
          <a v-if="record.status === 2" @click="rerun(record.spider, record.id)">Rerun</a>
          <a v-if="record.status === 1" @click="stop(record.spider, record.id)">Stop</a>
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
import {sortBigInt, sortInt, sortStr} from "@/utils/sort";

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
    case 1:
      return 'started'
    case 2:
      return 'stopped'
    default:
      return 'unknown'
  }
}

const stop = async (spiderName, jobId) => {
  const res = await jobsStore.StopJob({spider_name: spiderName, job_id: jobId})
  console.log(spiderName, jobId, res)
}
const rerun = async (spiderName, jobId) => {
  const res = await jobsStore.RerunJob({spider_name: spiderName, job_id: jobId})
  console.log(spiderName, jobId, res)
}
</script>
<style>
</style>
