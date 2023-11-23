<template>
  <div ref="logDiv" style="height: calc(100vh - 110px); overflow: auto;">
    <a-typography-paragraph v-for="log in logs" :key="log.key" :class="wrap ? '' : 'break'">{{
        log.data
      }}
    </a-typography-paragraph>
  </div>
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
