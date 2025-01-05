<script setup lang="ts">
import { ref } from "vue";
import { Feed } from "@/types/feed";

type FeedProps = {
  feed: Feed;
};

const { feed } = defineProps<FeedProps>();

const expanded = ref(false);
</script>

<template>
  <div v-if="expanded">
    
  </div>
  <div v-else>

  </div>
</template>
