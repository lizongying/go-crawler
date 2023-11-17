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
      <a-tab-pane key="1" tab="Data">
        <pre v-html="more.data"></pre>
      </a-tab-pane>
    </a-tabs>
  </a-drawer>
</template>
<script setup>
import {RightOutlined, SearchOutlined} from "@ant-design/icons-vue";
import {RouterLink, useRoute} from "vue-router";
import {formattedDate} from "@/utils/time";
import {useRecordsStore} from "@/stores/records";
import {computed, onBeforeUnmount, reactive, ref} from "vue";
import {sortBigInt, sortStr} from "@/utils/sort";

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
      title: 'Unique Key',
      dataIndex: 'unique_key',
      width: 200,
      sorter: (a, b) => sortStr(a.unique_key, b.unique_key),
      ellipsis: true,
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
      sorter: (a, b) => sortStr(a.job, b.job),
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
      title: 'Task',
      dataIndex: 'task',
      width: 200,
      sorter: (a, b) => sortBigInt(a.task, b.task),
      customFilterDropdown: true,
      filteredValue: filteredInfo.task || null,
      onFilter: (value, record) =>
          record.task.toString().toLowerCase().includes(value.toLowerCase()),
      onFilterDropdownOpenChange: visible => {
        if (visible) {
          setTimeout(() => {
            searchInput.value.focus();
          }, 100);
        }
      },
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
});

const recordsStore = useRecordsStore();

const open = ref(false);
const more = reactive({})
const showDrawer = record => {
  open.value = true;
  more.data = JSON.stringify(JSON.parse(record.data), null, 2)
      .replace(/(".*?":)/g, '<span class="key">$1</span>')
      .replace(/: "([^"]+)"\n/g, ': "<span class="string">$1</span>":')
      .replace(/: \b(\d+)\b\n/g, ': <span class="number">$1</span>')
      .replace(/: \b(true|false)\b\n/g, ': <span class="boolean">$1</span>')
      .replace(/: \b(null)\b\n/g, ': <span class="null">$1</span>');
};
const activeKey = ref('1');

// auto refresh
const checked1 = ref(true)
const checked1Disable = ref(true)
let interval = 0
const refresh = () => {
  recordsStore.GetRecords()
}
refresh()
if (checked1.value) {
  interval = setInterval(refresh, 1000)
}
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
.key {
  color: blue;
}

.string {
  color: green;
}

.number {
  color: orange;
}

.boolean {
  color: purple;
}

.null {
  color: rgb(128, 128, 128);
}
</style>
