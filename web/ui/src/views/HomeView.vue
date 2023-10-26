<template>
  <a-row style="margin: 10px">
    <a-col :span="4">
      <a-statistic title="Total Nodes" :value="nodesStore.Count" style="margin-right: 50px"/>
      <br>
      <a-statistic title="Active Nodes" :value="nodesStore.CountActive" style="margin-right: 50px"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Spiders" :value="spidersStore.Count" style="margin-right: 50px"/>
      <br>
      <a-statistic title="Enable Spiders" :value="spidersStore.CountActive" style="margin-right: 50px"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Schedules" :value="200" style="margin-right: 50px"/>
      <br>
      <a-statistic title="Enable Schedules" :value="200" style="margin-right: 50px"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Tasks" :value="1128"/>
      <br>
      <a-statistic title="Error Tasks" :value="1128"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Record" :value="recordsStore.Count"/>
      <br>
      <a-statistic title="Current Record" :value="recordsStore.CountActive"/>
    </a-col>
  </a-row>
</template>
<script setup>
import {useNodesStore} from "@/stores/nodes";
import {useSpidersStore} from "@/stores/spiders";
import {useRecordsStore} from "@/stores/records";

const nodesStore = useNodesStore();

nodesStore.GetNodes()

const spidersStore = useSpidersStore();

spidersStore.GetSpiders()

const recordsStore = useRecordsStore();

recordsStore.GetRecords()
</script>

<style>
</style>
