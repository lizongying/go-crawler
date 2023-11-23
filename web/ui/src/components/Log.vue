<template>
  <div ref="logDiv" style="height: calc(100vh - 144px); overflow: auto;">
    <a-typography-paragraph v-for="log in logs" :key="log.key" :class="wrap ? '' : 'break'">{{
        log.data
      }}
    </a-typography-paragraph>
  </div>
  <a-space direction="horizontal" style="margin-top:10px;">
    <a-switch v-model:checked="keepBottom" checked-children="keep bottom open"
              un-checked-children="keep bottom close"/>
    <a-switch v-model:checked="wrap" checked-children="word wrap open" un-checked-children="word wrap close"/>
  </a-space>
</template>

<script setup>
import {getLog} from "@/requests/api";
import {nextTick, onBeforeUnmount, onMounted, reactive, ref, watch} from "vue";

const props = defineProps({
  taskId: {
    type: String,
    default: '',
  },
  maxRow: {
    type: Number,
    default: 100,
  },
  wrap: {
    type: Boolean,
    default: true,
  }
})

const logs = reactive([])

let eventSource = null

const logDiv = ref(null);

const wrap = ref(true);
const keepBottom = ref(true);

onMounted(async () => {
  eventSource = await getLog(props.taskId);
  let index = 0
  eventSource.addEventListener("message", (event) => {
    index++
    logs.push({key: index, data: event.data});
    if (logs.length > props.maxRow) {
      logs.splice(0, 1)
    }
  });

  eventSource.addEventListener("error", (error) => {
    console.error("Error with SSE:", error);
    eventSource.close();
  });

  let scrollElem = logDiv.value;
  watch(logs, _ => {
    if (!keepBottom.value) {
      return
    }
    nextTick(() => {
      scrollElem.scrollTo({top: scrollElem.scrollHeight, behavior: "smooth"});
    })
  })
});
onBeforeUnmount(() => {
  eventSource.close();
})
</script>

<style scoped>
.break {
  overflow-x: auto;
  white-space: nowrap;
}
</style>
