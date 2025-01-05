<script setup lang="ts">
import Feed from "@/components/feed";

const feeds = await $fetch("/api/my-feeds", {
  method: "GET",
  credentials: "include",
  headers: useRequestHeaders(["cookie"]),
});
</script>

<template>
  <div id="background">
    <div id="content">
      <p id="feeds-title">My Feeds</p>
      <ul id="my-feeds">
        <li v-for="feed in feeds">
          <Feed :feed="feed" />
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
div#background {
  width: 100%;
  height: 100%;
  display: flex;
}

div#content {
  margin: 8px;
  width: calc(100vw - 16px);
  height: calc(100vh - 16px);
  display: flex;
  flex-direction: column;
  background-color: white;
}
</style>
