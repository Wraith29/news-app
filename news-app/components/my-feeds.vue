<script setup lang="ts">
import { ref } from "vue";
import type { UserFeed } from "@/types/user-feed";

type MyFeedProps = {
  feeds: UserFeed[];
  refreshPage: () => Promise<void>;
};

const { feeds, refreshPage } = defineProps<MyFeedProps>();
</script>

<template>
  <div id="my-feeds-parent">
    <p id="my-feeds-title">My Feeds</p>
    <UserFeed v-for="feed in feeds" :feed="feed" :refreshPage="refreshPage" />
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
