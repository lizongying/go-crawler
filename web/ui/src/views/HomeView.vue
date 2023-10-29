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
      <a-statistic :value="schedulesStore.Count" title="Total Schedules"/>
      <br>
      <a-statistic :value="schedulesStore.CountEnable" title="Enable Schedules"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="tasksStore.Count" title="Total Tasks"/>
      <br>
      <a-statistic :value="tasksStore.CountError" title="Error Tasks"/>
    </a-col>
    <a-col :span="4">
      <a-statistic :value="recordsStore.Count" title="Total Record"/>
    </a-col>
  </a-row>
</template>
<script setup>
import {useNodesStore} from "@/stores/nodes";
import {useSpidersStore} from "@/stores/spiders";
import {useRecordsStore} from "@/stores/records";
import {useTasksStore} from "@/stores/tasks";
import {useSchedulesStore} from "@/stores/schedules";

const nodesStore = useNodesStore()
nodesStore.GetNodes()

const schedulesStore = useSchedulesStore()
schedulesStore.GetSchedules()

const spidersStore = useSpidersStore()
spidersStore.GetSpiders()

const tasksStore = useTasksStore()
tasksStore.GetTasks()

const recordsStore = useRecordsStore()
recordsStore.GetRecords()
</script>
<style>
</style>
