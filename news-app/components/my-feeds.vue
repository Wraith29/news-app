<script setup lang="ts">
const feeds = await $fetch("/api/my-feeds", {
  method: "GET",
  credentials: "include",
  headers: useRequestHeaders(["cookie"]),
});
</script>

<template>
  <div class="content">
    <UserFeed v-for="feed in feeds" :feed="feed" />
  </div>
</template>

<style scoped>
</style>
