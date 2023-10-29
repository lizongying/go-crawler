<template>
  <a-result status="404" sub-title="Sorry, the page you visited does not exist." title="404">
    <template #extra>
      <RouterLink to="/">
        <a-button type="primary">Back Home</a-button>
      </RouterLink>
    </template>
  </a-result>
</template>
<script setup>
import {RouterLink} from "vue-router";
</script>
<style>
</style>
