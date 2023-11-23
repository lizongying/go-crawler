<template>
  <a-page-header
      title="Tasks"
      :sub-title="'Total: '+tasksStore.Count"
  >
    <template #extra>
      <a-switch v-model:checked="checked1" checked-children="auto" un-checked-children="close" @change="changeSwitch"/>
      <a-button key="2" @click="refresh" :disabled="checked1Disable">Refresh</a-button>
    </template>
  </a-page-header>
  <a-table :columns="columns" :data-source="tasksStore.tasks" :scroll="{ x: '100%' }">
    <template #headerCell="{ column }">
      <template v-if="column.dataIndex !== ''">
        <span style="font-weight: bold">
          {{ column.title }}
        </span>
      </template>
    </template>
    <template
        #customFilterDropdown="{ setSelectedKeys, selectedKeys, confirm, clearFilters, column }"
    >
      <div style="padding: 8px">
        <a-input
            ref="searchInput"
            :placeholder="`Search ${column.dataIndex}`"
            :value="selectedKeys[0]"
            style="width: 188px; margin-bottom: 8px; display: block"
            @change="e => setSelectedKeys(e.target.value ? [e.target.value] : [])"
            @pressEnter="handleSearch(selectedKeys, confirm, column.dataIndex)"
        />
        <a-button
            type="primary"
            size="small"
            style="width: 90px; margin-right: 8px"
            @click="handleSearch(selectedKeys, confirm, column.dataIndex)"
        >
          <template #icon>
            <SearchOutlined/>
          </template>
          Search
        </a-button>
        <a-button size="small" style="width: 90px" @click="handleReset(clearFilters)">
          Reset
        </a-button>
      </div>
    </template>
    <template #customFilterIcon="{ filtered }">
      <search-outlined :style="{ color: filtered ? '#108ee9' : undefined }"/>
    </template>
    <template #bodyCell="{ text, column, record }">
      <template v-if="column.dataIndex === 'node'">
        <RouterLink :to="'/nodes?id='+record.node">
          {{ record.node }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'spider'">
        <RouterLink :to="'/spiders?name='+record.spider">
          {{ record.spider }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'job'">
        <RouterLink :to="'/jobs?id='+record.job">
          {{ record.job }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'status'">
        <span>
          <a-tag
              :key="record.status"
              :color="record.status === TaskStatusFailure ? 'volcano' : record.status===TaskStatusSuccess ? 'green' : 'geekblue'"
          >
            {{ TaskStatusName(record.status) }}
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
      <template v-else-if="column.dataIndex === 'record'">
        <RouterLink :to="'/records?task='+record.id">
          {{ record.record }}
        </RouterLink>
      </template>
      <template v-else-if="column.dataIndex === 'action'">
        <span>
          <a class="ant-dropdown-link" @click="showDrawer(record)">
            More
            <RightOutlined/>
          </a>
        </span>
      </template>
      <span v-if="state.searchText && state.searchedColumn === column.dataIndex">
        <template
            v-for="(fragment, i) in text
            .toString()
            .split(new RegExp(`(?<=${state.searchText})|(?=${state.searchText})`, 'i'))"
        >
          <mark
              v-if="fragment.toLowerCase() === state.searchText.toLowerCase()"
              :key="i"
              class="highlight"
          >
            {{ fragment }}
          </mark>
          <template v-else>{{ fragment }}</template>
        </template>
      </span>
    </template>
  </a-table>
  <a-drawer v-model:open="open"
            :closable="false"
            size="large">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="Log">
        <Log :taskId="taskId"></Log>
      </a-tab-pane>
    </a-tabs>
  </a-drawer>
</template>
<script setup>
import {RightOutlined, SearchOutlined} from "@ant-design/icons-vue";
import {computed, onBeforeUnmount, reactive, ref} from "vue";
import {RouterLink, useRoute} from "vue-router";
import {
  TaskStatusFailure,
  TaskStatusName,
  TaskStatusPending,
  TaskStatusRunning,
  TaskStatusSuccess,
  useTasksStore
} from "@/stores/tasks";
import {formatDuration, formattedDate} from "@/utils/time";
import {sortBigInt, sortInt, sortStr} from "@/utils/sort";
import Log from "@/components/Log.vue";

const filteredInfo = reactive({});
const {query} = useRoute();
Object.entries(query).forEach(([k, v]) => {
  filteredInfo[k] = [v]
});
const columns = computed(() => {
  return [
    {
      title: 'Id',
      dataIndex: 'id',
      width: 200,
      sorter: (a, b) => sortBigInt(a.id, b.id),
      defaultSortOrder: 'descend',
      customFilterDropdown: true,
      filteredValue: filteredInfo.id || null,
      onFilter: (value, record) =>
          record.id.toString().toLowerCase().includes(value.toLowerCase()),
      onFilterDropdownOpenChange: visible => {
        if (visible) {
          setTimeout(() => {
            searchInput.value.focus();
          }, 100);
        }
      },
    },
    {
      title: 'Node',
      dataIndex: 'node',
      width: 200,
      sorter: (a, b) => sortBigInt(a.node, b.node),
      customFilterDropdown: true,
      filteredValue: filteredInfo.node || null,
      onFilter: (value, record) =>
          record.node.toString().toLowerCase().includes(value.toLowerCase()),
      onFilterDropdownOpenChange: visible => {
        if (visible) {
          setTimeout(() => {
            searchInput.value.focus();
          }, 100);
        }
      },
    },
    {
      title: 'Spider',
      dataIndex: 'spider',
      width: 200,
      sorter: (a, b) => sortStr(a.spider, b.spider),
      customFilterDropdown: true,
      filteredValue: filteredInfo.spider || null,
      onFilter: (value, record) =>
          record.spider.toString().toLowerCase().includes(value.toLowerCase()),
      onFilterDropdownOpenChange: visible => {
        if (visible) {
          setTimeout(() => {
            searchInput.value.focus();
          }, 100);
        }
      },
    },
    {
      title: 'Job',
      dataIndex: 'job',
      width: 200,
      sorter: (a, b) => sortBigInt(a.job, b.job),
      customFilterDropdown: true,
      filteredValue: filteredInfo.job || null,
      onFilter: (value, record) =>
          record.job.toString().toLowerCase().includes(value.toLowerCase()),
      onFilterDropdownOpenChange: visible => {
        if (visible) {
          setTimeout(() => {
            searchInput.value.focus();
          }, 100);
        }
      },
    },
    {
      title: 'Status',
      dataIndex: 'status',
      width: 100,
      filters: [
        {
          text: 'pending',
          value: TaskStatusPending,
        },
        {
          text: 'running',
          value: TaskStatusRunning,
        },
        {
          text: 'success',
          value: TaskStatusSuccess,
        },
        {
          text: 'failure',
          value: TaskStatusFailure,
        },
      ],
      onFilter: (value, record) => record.status === value,
      filteredValue: null,
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
      title: 'Stop Reason',
      dataIndex: 'stop_reason',
      sorter: (a, b) => sortStr(a.stop_reason, b.stop_reason),
      width: 200,
      ellipsis: true,
      customFilterDropdown: true,
      filteredValue: filteredInfo.stop_reason || null,
      onFilter: (value, record) =>
          record.stop_reason.toString().toLowerCase().includes(value.toLowerCase()),
      onFilterDropdownOpenChange: visible => {
        if (visible) {
          setTimeout(() => {
            searchInput.value.focus();
          }, 100);
        }
      },
    },
    {
      title: 'Record',
      dataIndex: 'record',
      width: 100,
      sorter: (a, b) => sortInt(a.record, b.record),
    },
    {
      title: 'Action',
      dataIndex: 'action',
      width: 200,
      fixed: 'right',
    },
  ];
});

const tasksStore = useTasksStore();

const open = ref(false);

const taskId = ref('')
const showDrawer = (record) => {
  open.value = true;
  taskId.value = record.id
};
const activeKey = ref('1');

// auto refresh
const checked1 = ref(true)
const checked1Disable = ref(true)
let interval = 0
const refresh = () => {
  tasksStore.GetTasks()
}
refresh()
if (checked1.value) {
  interval = setInterval(refresh, 1000)
}
const changeSwitch = () => {
  if (checked1.value) {
    if (!checked1Disable.value) {
      interval = setInterval(refresh, 1000)
    }
    checked1Disable.value = true
  } else {
    clearInterval(interval)
    checked1Disable.value = false
  }
}
onBeforeUnmount(() => {
  clearInterval(interval)
})

// search
const state = reactive({
  searchText: '',
  searchedColumn: '',
});
const searchInput = ref();
const handleSearch = (selectedKeys, confirm, dataIndex) => {
  filteredInfo[dataIndex] = selectedKeys
  confirm();
  state.searchText = selectedKeys[0];
  state.searchedColumn = dataIndex;
};
const handleReset = clearFilters => {
  Object.keys(filteredInfo).forEach(key => {
    delete filteredInfo[key];
  });
  clearFilters({
    confirm: true,
  });
  state.searchText = '';
};
</script>
<style>
.highlight {
  background-color: rgb(255, 192, 105);
  padding: 0;
}
</style>
