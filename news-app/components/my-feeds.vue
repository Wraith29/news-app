<script setup lang="ts">
const feeds = await $fetch("/api/my-feeds", {
  method: "GET",
  credentials: "include",
  headers: useRequestHeaders(["cookie"]),
});
</script>

<template>
  <div id="my-feeds-parent">
    <p id="my-feeds-title">My Feeds</p>
    <UserFeed v-for="feed in feeds" :feed="feed" />
  </div>
</template>

<style scoped>
div#my-feeds-parent {
  margin: 20px;
  padding: 20px;
  border: 1px solid darkgreen;
  border-radius: 25px;

  > p#my-feeds-title {
    margin: 0;
    margin-bottom: 10px;
    font-size: x-large;
    text-decoration: underline;
  }
}
</style>
