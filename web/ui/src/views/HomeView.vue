<template>
  <a-row style="margin: 15px">
    <a-col :span="4">
      <a-statistic :value="nodesStore.Count" title="Total Nodes"/>
      <br>
      <a-statistic :value="nodesStore.CountActive" title="Active Nodes"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="spidersStore.Count" title="Total Spiders"/>
      <br>
      <a-statistic :value="spidersStore.CountActive" title="Enable Spiders"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="schedulesStore.Count" title="Total Jobs"/>
      <br>
      <a-statistic :value="schedulesStore.CountEnable" title="Enable Jobs"/>
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
import {useSpidersStore} from "@/stores/spiders";
import {useJobsStore} from "@/stores/jobs";

const nodesStore = useNodesStore()
nodesStore.GetNodes()

const schedulesStore = useJobsStore()
schedulesStore.GetJobs()

const spidersStore = useSpidersStore()
spidersStore.GetSpiders()
</script>
<style>
</style>
