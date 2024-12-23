<script setup lang="ts">
const articles = await $fetch("/api/articles", {
  method: "GET",
  credentials: "include",
  headers: useRequestHeaders(["cookie"]),
});
</script>

<template>
  <div id="background">
  
  <div id="content">
    <ul>
      <li v-for="article in articles">
        {{ article.title }}
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
  background-color: white;
}
</style>
