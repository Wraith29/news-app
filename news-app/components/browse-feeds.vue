<script setup lang="ts">
const feeds = await $fetch("/api/feed/all", {
  method: "GET",
  credentials: "include",
  headers: useRequestHeaders(["cookie"]),
});
</script>

<template>
  <div id="browse-feeds-box">
    <p id="browse-feeds-title">Browse Feeds</p>
    <Feed v-for="feed in feeds" :feed="feed" />
  </div>
</template>

<style scoped>
div#browse-feeds-box {
  margin: 20px;
  padding: 20px;
  border: 1px solid green;

  > p#browse-feeds-title {
    margin: 0;
    margin-bottom: 10px;
    font-size: x-large;
    text-decoration: underline;
  }
}
</style>
