<template>
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
        <RouterLink :to="'/nodes?id='+record.node" @click="$emit('router—change','1')">
          {{ record.node }}
        </RouterLink>
      </template>
      <template v-if="column.dataIndex === 'spider'">
        <RouterLink :to="'/spiders?name='+record.spider" @click="$emit('router—change','3')">
          {{ record.spider }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'task_id'">
        <RouterLink :to="'/tasks?id='+record.task_id" @click="$emit('router—change','5')">
          {{ record.task_id }}
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
            size="large"
            :closable="false">
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
import {reactive, ref} from "vue";

defineEmits(['router—change'])

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
    width: 200,
    sorter: (a, b) => a.spider - b.spider,
  },
  {
    title: 'Task',
    dataIndex: 'task',
    width: 300,
    sorter: (a, b) => a.task - b.task,
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
</script>
<style>
</style>
