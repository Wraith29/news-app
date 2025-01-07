<script setup lang="ts">
</script>

<template>
  <div class="background">
    <div class="content">
      <Header :page="'feeds'" />

      <MyFeeds />
      <BrowseFeeds />
      <NewFeed />
    </div>
  </div>
</template>

<style scoped>
div.background {
  width: 100%;
  height: 100%;
  display: flex;
}

div.content {
  margin: 8px;
  width: calc(100vw - 16px);
  height: calc(100vh - 16px);
  display: flex;
  flex-direction: column;
  background-color: white;
}
</style>
