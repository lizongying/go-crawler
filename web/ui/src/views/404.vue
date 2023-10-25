<script setup>
import {RouterLink} from "vue-router";

defineEmits(['router—change'])

</script>

<template>
  <a-result status="404" title="404" sub-title="Sorry, the page you visited does not exist.">
    <template #extra>
      <RouterLink to="/" @click="$emit('router—change','1')">
        <a-button type="primary">Back Home</a-button>
      </RouterLink>
    </template>
  </a-result>
</template>
<style>
</style>
