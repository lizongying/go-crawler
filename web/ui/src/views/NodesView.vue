<template>
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
              :color="record.status === 2 ? 'volcano' : 'green'"
          >
            {{ record.status === 2 ? 'OFFLINE' : 'ONLINE' }}
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
import {useNodesStore} from "@/stores/nodes";
import {formatDuration, formattedDate} from "@/utils/time";

const columns = [
  {
    title: 'Id',
    dataIndex: 'id',
    width: 300,
    sorter: (a, b) => a.id - b.id,
  },
  {
    title: 'Hostname',
    dataIndex: 'hostname',
    width: 200,
    ellipsis: true,
    sorter: (a, b) => a.hostname - b.hostname,
  },
  {
    title: 'Ip',
    dataIndex: 'ip',
    width: 200,
    ellipsis: true,
    sorter: (a, b) => a.ip - b.ip,
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
        text: 'online',
        value: 'online',
      },
      {
        text: 'offline',
        value: 'offline',
      },
    ],
    onFilter: (value, record) => record.status === value,
  },
  {
    title: 'Spider',
    dataIndex: 'spider',
    width: 100,
    sorter: (a, b) => a.spider - b.spider,
  },
  {
    title: 'Job',
    dataIndex: 'job',
    width: 100,
    sorter: (a, b) => a.job - b.job,
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

const nodesStore = useNodesStore();

nodesStore.GetNodes()
</script>
<style>
</style>
