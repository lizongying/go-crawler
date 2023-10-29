<template>
  <a-table :columns="columns" :data-source="schedulesStore.schedules" :scroll="{ x: '100%' }">
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
      <template v-else-if="column.dataIndex === 'task'">
        <RouterLink :to="'/tasks?schedule='+record.id">
          {{ record.task }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'record'">
        <RouterLink :to="'/records?schedule='+record.id">
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
import {useSchedulesStore} from "@/stores/schedules";

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
    width: 200,
    sorter: (a, b) => a.schedule - b.schedule,
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

const schedulesStore = useSchedulesStore()
schedulesStore.GetSchedules()
</script>
<style>
</style>
