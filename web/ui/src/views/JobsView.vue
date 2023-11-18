<template>
  <a-page-header
      title="Jobs"
      :sub-title="'Total: '+jobsStore.Count"
  >
    <template #extra>
      <a-button key="1" @click="newJob" type="primary">New</a-button>
      <a-switch v-model:checked="checked1" checked-children="auto" un-checked-children="close" @change="changeSwitch"/>
      <a-button key="2" @click="refresh" :disabled="checked1Disable">Refresh</a-button>
    </template>
  </a-page-header>
  <a-table :columns="columns" :data-source="jobsStore.jobs" :scroll="{ x: '100%' }">
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
      <template v-else-if="column.dataIndex === 'status'">
        <span>
          <a-tag
              :key="record.status"
              :color="record.status === JobStatusStopped ? 'volcano' : record.status === JobStatusRunning ? 'green' : 'geekblue'"
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
          <a v-if="record.status === JobStatusStopped" @click="rerun(record.spider, record.id)">Rerun</a>
          <a v-if="record.status === JobStatusRunning" @click="stop(record.spider, record.id)">Stop</a>
          <a-divider type="vertical"/>
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
      <a-tab-pane key="1" tab="Status List">
        <a-list size="small" bordered :data-source="more.status_list">
          <template #renderItem="{ item }">
            <a-list-item>{{ item }}</a-list-item>
          </template>
        </a-list>
      </a-tab-pane>
    </a-tabs>
  </a-drawer>
  <a-modal v-model:open="openJob" title="New Job" width="1000px" @ok="handleJob">
    <a-form
        :label-col="{ span: 4 }"
        :model="formJob"
        :wrapper-col="{ span: 20 }"
        autocomplete="off"
        name="basic"
    >
      <a-form-item
          :rules="[{ required: true, message: 'Please input spider name!' }]"
          label="Spider Name"
          name="name"
      >
        <a-select
            v-model:value="formJob.name"
            show-search
            style="width: 100%"
            :options="spidersStore.SpiderNames"
            :filter-option="filterSpiders"
            placeholder="Select a spider name"
            @change="handleChange"
        ></a-select>
      </a-form-item>
      <a-form-item
          :rules="[{ required: true, message: 'Please input spider func!' }]"
          label="Spider Func"
          name="func"
          :validate-status="func.status"
          :help="func.help"
      >
        <a-select
            v-model:value="formJob.func"
            show-search
            style="width: 100%"
            :options="spidersStore.SpiderFuncs(formJob.name)"
            :filter-option="filterSpiders"
            placeholder="Select a spider func"
            @focus="handleFocus"
        ></a-select>
      </a-form-item>
      <a-form-item
          :rules="[{ required: true, message: 'Please input func args!' }]"
          label="Func Args"
          name="args"
      >
        <a-textarea
            v-model:value="formJob.args"
            placeholder="The arguments should be a json"
            :auto-size="{ minRows: 2, maxRows: 5 }"
        />
      </a-form-item>
      <a-form-item
          :rules="[{ required: true, message: 'Please choose a mode!' }]"
          label="Run Mode"
          name="mode"
      >
        <a-radio-group v-model:value="formJob.mode">
          <a-radio-button value="1">once</a-radio-button>
          <a-radio-button value="2">loop</a-radio-button>
          <a-radio-button value="3">cron</a-radio-button>
        </a-radio-group>
      </a-form-item>
      <a-form-item
          v-if="formJob.mode==='3'"
          label="Spec"
          name="spec"
      >
        <a-input v-model:value="formJob.specValue" addon-before="every" type="number" min="1">
          <template #addonAfter>
            <a-select v-model:value="formJob.specType" style="width: 100px">
              <a-select-option value="s">seconds</a-select-option>
              <a-select-option value="i">minutes</a-select-option>
              <a-select-option value="h">hours</a-select-option>
              <a-select-option value="d">days</a-select-option>
              <a-select-option value="m">months</a-select-option>
              <a-select-option value="w">weeks</a-select-option>
            </a-select>
          </template>
        </a-input>
      </a-form-item>
      <a-form-item
          label="Timeout"
          name="timeout"
      >
        <a-input v-model:value="formJob.timeoutValue" type="number" min="0">
          <template #addonAfter>
            <a-select v-model:value="formJob.timeoutType" style="width: 100px">
              <a-select-option value="1">seconds</a-select-option>
              <a-select-option value="60">minutes</a-select-option>
              <a-select-option value="360">hours</a-select-option>
              <a-select-option value="8640">days</a-select-option>
              <a-select-option value="259200">months</a-select-option>
              <a-select-option value="60480">weeks</a-select-option>
            </a-select>
          </template>
        </a-input>
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script setup>
import {ExclamationCircleOutlined, RightOutlined, SearchOutlined} from "@ant-design/icons-vue";
import {RouterLink, useRoute} from "vue-router";
import {
  JobStatusIdle,
  JobStatusReady,
  JobStatusRunning,
  JobStatusStarting,
  JobStatusStopped,
  JobStatusStopping,
  useJobsStore
} from "@/stores/jobs";
import {formatDuration, formattedDate} from "@/utils/time";
import {sortBigInt, sortInt, sortStr} from "@/utils/sort";
import {computed, createVNode, onBeforeUnmount, reactive, ref} from "vue";
import {message, Modal} from "ant-design-vue";
import {useSpidersStore} from "@/stores/spiders";
import {useNodesStore} from "@/stores/nodes";

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
      sorter: (a, b) => sortStr(a.spider, b.spider),
      width: 200,
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
      title: 'Status',
      dataIndex: 'status',
      width: 100,
      filters: [
        {
          text: 'ready',
          value: JobStatusReady,
        },
        {
          text: 'starting',
          value: JobStatusStarting,
        },
        {
          text: 'running',
          value: JobStatusRunning,
        },
        {
          text: 'idle',
          value: JobStatusIdle,
        },
        {
          text: 'stopping',
          value: JobStatusStopping,
        },
        {
          text: 'stopped',
          value: JobStatusStopped,
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
      width: 200,
      fixed: 'right',
    },
  ];
});

const jobsStore = useJobsStore()

const jobStatusName = (status) => {
  switch (status) {
    case JobStatusReady:
      return 'ready'
    case JobStatusStarting:
      return 'starting'
    case JobStatusRunning:
      return 'running'
    case JobStatusIdle:
      return 'idle'
    case JobStatusStopping:
      return 'stopping'
    case JobStatusStopped:
      return 'stopped'
    default:
      return 'unknown'
  }
}

// auto refresh
const checked1 = ref(true)
const checked1Disable = ref(true)
let interval = 0
const refresh = () => {
  jobsStore.GetJobs()
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

// rerun confirm
function rerun(spiderName, jobId) {
  Modal.confirm({
    title: 'Do you want to rerun the job?',
    icon: createVNode(ExclamationCircleOutlined),
    content: 'When clicked the OK button, the job will be rerun.',
    async onOk() {
      try {
        const res = await jobsStore.RerunJob({spider_name: spiderName, job_id: jobId})
        console.log(spiderName, jobId, res)
      } catch {
        return console.log('Oops errors!');
      }
    },
    onCancel() {
    },
  });
}

// stop confirm
function stop(spiderName, jobId) {
  Modal.confirm({
    title: 'Do you want to stop the job?',
    icon: createVNode(ExclamationCircleOutlined),
    content: 'When clicked the OK button, the job will be stop.',
    async onOk() {
      try {
        const res = await jobsStore.StopJob({spider_name: spiderName, job_id: jobId})
        console.log(spiderName, jobId, res)
      } catch {
        return console.log('Oops errors!');
      }
    },
    onCancel() {
    },
  });
}

// more
const open = ref(false);
const more = reactive({})
const showDrawer = record => {
  open.value = true;
  more.status_list = Object.entries(record.status_list).map(([k, v]) => `${formattedDate(k / 1000000000)} ${jobStatusName(v)}`).reverse();
};
// status list
const activeKey = ref('1');

// new job
const formJob = reactive({
  name: '',
  func: '',
  args: '{}',
  mode: '1',
  specType: 'h',
  specValue: 1,
  timeoutType: 'h',
  timeoutValue: 0,
})
const openJob = ref(false);
const newJob = () => {
  openJob.value = true;
};
const handleJob = () => {
  if (formJob.name === '') {
    message.error('Spider name empty');
    return
  }
  if (formJob.func === '') {
    message.error('Spider func empty');
    return
  }
  if (formJob.args === '') {
    message.error('Func args empty');
    return
  }
  try {
    let js = JSON.parse(formJob.args)
    formJob.args = JSON.stringify(js)

    openJob.value = false;
    const data = {
      "timeout": 0,
      "name": formJob.name,
      "func": formJob.func,
      "args": formJob.args,
      "mode": parseInt(formJob.mode)
    }
    if (formJob.mode === '3') {
      data.spec = formJob.specValue + formJob.specType
    }
    if (formJob.timeoutValue > 0) {
      data.timeout = formJob.timeoutValue * parseInt(formJob.timeoutType)
    }
    jobsStore.RunJob(data)
  } catch (e) {
    console.log(e)
    message.error('Argument error');
  }
};

const spidersStore = useSpidersStore()
spidersStore.GetSpiders()
const filterSpiders = (input, option) => {
  return option.label.toLowerCase().indexOf(input.toLowerCase()) >= 0;
};

const func = reactive({
  status: '',
  help: '',
})
const handleFocus = () => {
  if (formJob.name === '') {
    func.status = 'error'
    func.help = 'Should be input a spider name'
  }
};
const handleChange = () => {
  if (formJob.name !== '') {
    func.status = ''
    func.help = ''
  }
  formJob.func = ''
};

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
</style>
