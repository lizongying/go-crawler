<template>
  <a-page-header
      title="Records"
      :sub-title="'Total: '+recordsStore.Count"
  >
    <template #extra>
      <a-switch v-model:checked="checked1" checked-children="auto" un-checked-children="close" @change="changeSwitch"/>
      <a-button key="2" @click="refresh" :disabled="checked1Disable">Refresh</a-button>
    </template>
  </a-page-header>
  <a-table :columns="columns" :data-source="recordsStore.records" :scroll="{ x: '100%' }">
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
      <template v-else-if="column.dataIndex === 'job'">
        <RouterLink :to="'/jobs?id='+record.job">
          {{ record.job }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'task'">
        <RouterLink :to="'/tasks?id='+record.task">
          {{ record.task }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'save_time'">
        {{ formattedDate(record.save_time) }}
      </template>
      <template v-else-if="column.dataIndex === 'action'">
        <span>
          <a class="ant-dropdown-link" @click="showDrawer(record)">
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
      <a-tab-pane key="1" tab="Data">
        <pre>{{ more.data }}</pre>
      </a-tab-pane>
    </a-tabs>
  </a-drawer>
</template>
<script setup>
import {RightOutlined} from "@ant-design/icons-vue";
import {RouterLink} from "vue-router";
import {formattedDate} from "@/utils/time";
import {useRecordsStore} from "@/stores/records";
import {onBeforeUnmount, reactive, ref} from "vue";
import {sortBigInt, sortStr} from "@/utils/sort";

const columns = [
  {
    title: 'Id',
    dataIndex: 'id',
    width: 200,
    sorter: (a, b) => sortBigInt(a.id, b.id),
    defaultSortOrder: 'descend',
  },
  {
    title: 'Unique Key',
    dataIndex: 'unique_key',
    width: 200,
    sorter: (a, b) => sortStr(a.unique_key, b.unique_key),
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
    width: 200,
    sorter: (a, b) => sortStr(a.spider, b.spider),
  },
  {
    title: 'Job',
    dataIndex: 'job',
    width: 200,
    sorter: (a, b) => sortStr(a.job, b.job),
  },
  {
    title: 'Task',
    dataIndex: 'task',
    width: 200,
    sorter: (a, b) => sortBigInt(a.task, b.task),
  },
  {
    title: 'Meta',
    dataIndex: 'meta',
    width: 200,
    ellipsis: true,
  },
  {
    title: 'Save Time',
    dataIndex: 'save_time',
    width: 200,
    sorter: (a, b) => a.save_time - b.save_time,
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: 200,
    fixed: 'right',
  },
];

const recordsStore = useRecordsStore();

recordsStore.GetRecords()

const open = ref(false);
const more = reactive({})
const showDrawer = record => {
  open.value = true;
  more.data = record.data
};
const activeKey = ref('1');

const refresh = () => {
  recordsStore.GetRecords()
}
const checked1 = ref(true)
const checked1Disable = ref(true)

let interval = setInterval(refresh, 1000)
const changeSwitch = () => {
  if (checked1.value) {
    interval = setInterval(refresh, 1000)
    checked1Disable.value = true
  } else {
    clearInterval(interval)
    checked1Disable.value = false
  }
}
onBeforeUnmount(() => {
  clearInterval(interval)
})
</script>
<style>
</style>
