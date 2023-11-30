<template>
  <a-page-header
      title="Home"
  >
    <template #extra>
      <a-switch v-model:checked="checked1" checked-children="开" un-checked-children="关" @change="changeSwitch"/>
      <a-button key="2" @click="refresh" :disabled="checked1Disable">Refresh</a-button>
    </template>
  </a-page-header>
  <a-row style="margin: 15px">
    <a-col :span="4">
      <a-statistic :value="crawlersStore.Count" title="Total Crawlers"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="crawlersStore.CountSpider" title="Total Spiders"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="crawlersStore.CountJob" title="Total Jobs"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="crawlersStore.CountTask" title="Total Tasks"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="crawlersStore.CountItem" title="Total Item"/>
    </a-col>
  </a-row>
</template>
<script setup>
import {useCrawlersStore} from "@/stores/crawlers";
import {onBeforeUnmount, ref} from "vue";

const crawlersStore = useCrawlersStore()

// auto refresh
const checked1 = ref(true)
const checked1Disable = ref(true)
let interval = 0
const refresh = () => {
  crawlersStore.GetCrawlers()
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
</script>
<style>
</style>
