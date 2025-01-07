<script setup lang="ts">
import { Feed } from "@/types/feed";

type FeedProps = {
  feed: Feed;
};

const { feed } = defineProps<Feed>();
</script>

<template>
  <div class="content">
  </div>
</template>

<style scoped>
</style>
