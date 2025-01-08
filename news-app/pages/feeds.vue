<script setup lang="ts"></script>

<template>
  <div class="background">
    <div class="feeds-content">
      <Header :page="'feeds'" />

      <div id="feed-selectors">
        <MyFeeds class="feed-selector" />
        <BrowseFeeds class="feed-selector" />
      </div>

      <NewFeed id="new-feed-parent" />
    </div>
  </div>
</template>

<style scoped>
div.background {
  width: 100%;
  height: 100%;
  display: flex;
}

div.feeds-content {
  margin: 8px;
  width: calc(100vw - 16px);
  height: calc(100vh - 16px);
  display: flex;
  flex-direction: column;
  background-color: white;
  overflow: hidden;

  > div#feed-selectors {
    display: flex;
    height: 80%;

    > .feed-selector {
      width: 100%;
    }
  }

  > #new-feed-parent {
    margin: 20px;
    height: 30%;
  }
}
</style>
