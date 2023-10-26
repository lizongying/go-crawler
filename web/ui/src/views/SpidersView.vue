<template>
  <a-table :columns="columns" :data-source="spidersStore.spiders" :scroll="{ x: '100%' }">
    <template #headerCell="{ column }">
      <template v-if="['spider', 'last_status', 'last_run_at', 'last_finish_at'].includes(column.dataIndex)">
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
      <template v-else-if="column.dataIndex === 'last_status'">
        <span>
          <a-tag
              :key="record.last_status"
              :color="record.last_status === 4 ? 'volcano' : record.last_status===2 ? 'green' : 'geekblue'"
          >
            {{ statusName(record.last_status) }}
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
      <template v-else-if="column.dataIndex === 'last_run_at'">
        {{ formattedDate(record.last_run_at) }}
      </template>
      <template v-else-if="column.dataIndex === 'last_finish_at'">
        {{ formattedDate(record.last_finish_at) }}
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
    title: 'Name',
    dataIndex: 'spider',
    sorter: (a, b) => a.spider > b.spider,
    width: 200,
  },
  {
    title: 'Last Status',
    dataIndex: 'last_status',
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
    onFilter: (value, record) => record.last_status === value,
  },
  {
    title: 'Last Run At',
    dataIndex: 'last_run_at',
    width: 200,
    sorter: (a, b) => a.last_run_at - b.last_run_at,
  },
  {
    title: 'Last Finish At',
    dataIndex: 'last_finish_at',
    width: 200,
    sorter: (a, b) => a.last_finish_at - b.last_finish_at,
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

const data = reactive([
  // {
  //   id: '1',
  //   spider: 'test1',
  //   last_status: 0,
  //   last_run_at: '1 hour ago',
  //   schedule: 100,
  //   task: 100,
  //   record: 100,
  // },
  // {
  //   id: '2',
  //   spider: 'test2',
  //   last_status: 0,
  //   last_run_at: '1 hour ago',
  //   schedule: 100,
  //   task: 100,
  //   record: 100,
  // },
  // {
  //   id: '3',
  //   spider: 'test3',
  //   last_status: 0,
  //   last_run_at: '1 hour ago',
  //   schedule: 100,
  //   task: 100,
  //   record: 100,
  // },
])


const spidersStore = useSpidersStore();

spidersStore.GetSpiders()

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
