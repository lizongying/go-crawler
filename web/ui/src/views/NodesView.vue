<template>
  <a-table :columns="columns" :data-source="data">
    <template #headerCell="{ column }">
      <template v-if="['hostname', 'ip', 'enable', 'status', 'start_time'].includes(column.dataIndex)">
        <span style="font-weight: bold">
          {{ column.title }}
        </span>
      </template>
    </template>

    <template #bodyCell="{ column, record }">
      <template v-if="column.dataIndex === 'spider'">
        <RouterLink :to="'/spiders?node_id='+record.id" @click="$emit('router—change','3')">
          {{ record.spider }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'schedule'">
        <RouterLink :to="'/schedules?node_id='+record.id" @click="$emit('router—change','4')">
          {{ record.schedule }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'task'">
        <RouterLink :to="'/tasks?node_id='+record.id" @click="$emit('router—change','5')">
          {{ record.task }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'start_time'">
          {{ formattedDate(record.start_time) }}
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
import {reactive} from "vue";
import {formattedDate} from "../utils/time";

defineEmits(['router—change'])

const columns = [
  {
    title: 'Hostname',
    dataIndex: 'hostname',
    width: 200,
    ellipsis: true,
    sorter: (a, b) => a.hostname > b.hostname,
  },
  {
    title: 'Ip',
    dataIndex: 'ip',
    width: 200,
    ellipsis: true,
    sorter: (a, b) => a.ip > b.ip,
  },
  {
    title: 'Start Time',
    dataIndex: 'start_time',
    width: 200,
    ellipsis: true,
    sorter: (a, b) => a.ip > b.ip,
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
  },
  {
    title: 'Schedule',
    dataIndex: 'schedule',
    width: 100,
  },
  {
    title: 'Task',
    dataIndex: 'task',
    width: 100,
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: 200,
    fixed: 'right',
  },
];

const data = reactive([
  // {
  //   id: '1',
  //   hostname: 'localhost',
  //   ip: '127.0.0.1:9999',
  //   status: 1,
  //   enable: true,
  //   spider: 32,
  //   schedule: 10,
  //   task: 100,
  // },
  // {
  //   id: '2',
  //   hostname: 'localhost',
  //   ip: '127.0.0.1:9999',
  //   status: 1,
  //   enable: true,
  //   spider: 42,
  //   schedule: 10,
  //   task: 100,
  // },
  // {
  //   id: '3',
  //   hostname: 'localhost',
  //   ip: '127.0.0.1:9999',
  //   status: 2,
  //   enable: false,
  //   spider: 32,
  //   schedule: 10,
  //   task: 100,
  // },
])

const nodesStore = useNodesStore();

nodesStore.GetNodes().then(resp => {
  resp.data.data.forEach(
      v => {
        data.push(v)
      }
  )
})

</script>
<style>
</style>
