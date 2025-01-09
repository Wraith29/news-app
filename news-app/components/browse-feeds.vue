<script setup lang="ts">
import type { Feed } from "@/types/feed";

type BrowseFeedProps = {
  feeds: Feed[];
  refreshPage: () => Promise<void>;
};

const { feeds, refreshPage } = defineProps<BrowseFeedProps>();
</script>

<template>
  <div id="browse-feeds-box">
    <p id="browse-feeds-title">Browse Feeds</p>
    <Feed v-for="feed in feeds" :feed="feed" :refreshPage="refreshPage" />
  </div>
</template>

<style scoped>
div#browse-feeds-box {
  margin: 20px;
  padding: 20px;
  border: 3px solid darkred;
  border-radius: 25px;

  > p#browse-feeds-title {
    margin: 0;
    margin-bottom: 10px;
    font-size: x-large;
    text-decoration: underline;
  }
}
</style>
