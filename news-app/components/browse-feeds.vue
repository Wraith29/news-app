<script setup lang="ts">
import Feed from "@/components/feed";

const feeds = await $fetch("/api/all-feeds", {
  method: "GET",
  credentials: "include",
  headers: useRequestHeaders(["cookie"]),
});
</script>

<template>
  <div class="content">
  </div>
</template>

<style scoped>
</style>
