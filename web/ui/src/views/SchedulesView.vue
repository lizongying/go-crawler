<template>
  <a-table :columns="columns" :data-source="data" :scroll="{ x: '100%' }">
    <template #headerCell="{ column }">
      <template v-if="['spider', 'schedule'].includes(column.dataIndex)">
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
    sorter: (a, b) => a.schedule > b.schedule,
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: 200,
    fixed: 'right',
  },
];
const data = [
  {
    id: '1',
    spider: 'John Brown',
    schedule: 'every day',
  },
  {
    id: '2',
    spider: 'Jim Green',
    schedule: 'every day',
  },
  {
    id: '3',
    spider: 'Joe Black',
    schedule: 'every day',
  },
];

const schedulesStore = useSchedulesStore()
schedulesStore.GetSchedules()
</script>
<style>
</style>
