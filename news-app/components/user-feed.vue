<script setup lang="ts">
import { ref } from "vue";
import { UserFeed } from "@/types/user-feed";

type FeedProps = {
  feed: UserFeed;
};

const { feed } = defineProps<FeedProps>();

const expanded = ref(false);
const enabledText = ref(feed.enabled ? "Enabled" : "Disabled");

function toggleExpanded(): void {
  expanded.value = !expanded.value;
}

async function toggleFeedEnabled(): Promise<void> {
  feed.enabled = !feed.enabled;

  enabledText.value = feed.enabled ? "Enabled" : "Disabled";

  await $fetch("/api/feed", {
    method: "PUT",
    body: JSON.stringify({
      feedId: feed.id,
    }),
  });
}

async function leaveFeed(): Promise<void> {
  console.log("Leaving Feed:", feed.id);
}
</script>

<template>
  <div class="user-feed">
    <span class="user-feed-header">
      <div class="user-feed-header-left">
        <p class="user-feed-header-title" @click="toggleExpanded">
          {{ feed.feed_author }}
        </p>
      </div>

      <div class="user-feed-header-right">
        <input
          name="toggleFeed"
          class="user-feed-enabled-toggle"
          type="button"
          :value="enabledText"
          @click="() => toggleFeedEnabled()"
        />

        <input
          name="leaveFeed"
          class="user-feed-leave-btn"
          type="button"
          value="Leave"
          @click="() => leaveFeed()"
        />
      </div>
    </span>

    <div v-if="expanded">
      <p class="user-feed-expanded-data">{{ feed.feed_url }}</p>
    </div>
  </div>
</template>

<style scoped>
div.user-feed {
  border: 1px solid red;
  border-radius: 15px;
  padding: 10px;

  > span.user-feed-header {
    display: flex;
    justify-content: space-between;
    align-items: center;

    > div.user-feed-header-left {
      display: flex;
      align-items: center;

      > p.user-feed-header-title {
        font-size: large;
      }
    }

    > div.user-feed-header-right {
      > input.user-feed-enabled-toggle {
        width: 70px;
        height: 30px;
        background: none;
        border: 1px solid mediumvioletred;
        border-radius: 5px;
        margin-right: 10px;
      }

      > input.user-feed-leave-btn {
        width: 70px;
        height: 30px;
        background: none;
        border: 1px solid mediumvioletred;
        border-radius: 5px;
      }
    }

    > p.user-feed-header-title {
      font-size: large;
      text-decoration: underline;
      margin: 0;
    }
  }
}
</style>
