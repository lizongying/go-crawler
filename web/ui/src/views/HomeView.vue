<script setup>
</script>

<template>
  <a-row style="margin: 10px">
    <a-col :span="4">
      <a-statistic title="Total Nodes" :value="200" style="margin-right: 50px"/>
      <br>
      <a-statistic title="Active Nodes" :value="200" style="margin-right: 50px"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Spiders" :value="200" style="margin-right: 50px"/>
      <br>
      <a-statistic title="Enable Spiders" :value="200" style="margin-right: 50px"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Schedules" :value="200" style="margin-right: 50px"/>
      <br>
      <a-statistic title="Enable Schedules" :value="200" style="margin-right: 50px"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Tasks" :precision="2" :value="1128"/>
      <br>
      <a-statistic title="Error Tasks" :precision="2" :value="1128"/>
    </a-col>
    <a-col :span="4">
      <a-statistic title="Total Record" :precision="2" :value="112893"/>
      <br>
      <a-statistic title="Current Record" :precision="2" :value="112893"/>
    </a-col>
  </a-row>
</template>
<style>
</style>
