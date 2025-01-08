<script setup lang="ts">
import { Feed } from "@/types/feed";

type FeedProps = {
  feed: Feed;
};

const { feed } = defineProps<FeedProps>();
</script>

<template>
  <div class="feed-content">
    <p>{{ feed.feed_author }}</p>
  </div>
</template>

<style scoped>
div.feed {
  border: 1px solid red;
  padding: 10px;

  > span.feed-header {
    display: flex;
    align-items: center;

    > input.feed-enabled-toggle {
      width: 20px;
      height: 20px;
      border: none;
      border-radius: 5px;
    }

    > p.feed-header-title {
      font-size: large;
      text-decoration: underline;
      margin: 0;
    }
  }
}
</style>
