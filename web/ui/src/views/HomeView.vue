<template>
  <a-row style="margin: 15px">
    <a-col :span="4">
      <a-statistic title="Total Nodes" :value="nodesStore.Count"/>
      <br>
      <a-statistic title="Active Nodes" :value="nodesStore.CountActive"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Spiders" :value="spidersStore.Count"/>
      <br>
      <a-statistic title="Enable Spiders" :value="spidersStore.CountActive"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Schedules" :value="0"/>
      <br>
      <a-statistic title="Enable Schedules" :value="0"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Tasks" :value="tasksStore.Count"/>
      <br>
      <a-statistic title="Error Tasks" :value="tasksStore.CountError"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Record" :value="recordsStore.Count"/>
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
