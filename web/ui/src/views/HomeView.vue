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
      <a-statistic :value="nodesStore.Count" title="Total Nodes"/>
      <br>
      <a-statistic :value="nodesStore.CountActive" title="Active Nodes"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="nodesStore.CountSpider" title="Total Spiders"/>
      <br>
      <a-statistic :value="nodesStore.CountSpider" title="Enable Spiders"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="nodesStore.CountJob" title="Total Jobs"/>
      <br>
      <a-statistic :value="nodesStore.CountJob" title="Enable Jobs"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="nodesStore.CountTask" title="Total Tasks"/>
      <br>
      <a-statistic :value="nodesStore.CountTask" title="Success Tasks"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="nodesStore.CountRecord" title="Total Record"/>
    </a-col>
  </a-row>
</template>
<script setup>
import {useNodesStore} from "@/stores/nodes";
import {onBeforeUnmount, ref} from "vue";

const nodesStore = useNodesStore()
nodesStore.GetNodes()

const refresh = () => {
  nodesStore.GetNodes()
}
const checked1 = ref(false)
const checked1Disable = ref(false)

let interval = null
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
