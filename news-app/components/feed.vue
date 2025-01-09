<script setup lang="ts">
import type { Feed } from "@/types/feed";

type FeedProps = {
  feed: Feed;
  refreshPage: () => Promise<void>;
};

const { feed, refreshPage } = defineProps<FeedProps>();

async function joinFeed(): Promise<void> {
  await $fetch("/api/feed/join", {
    method: "PUT",
    body: JSON.stringify({
      feedId: feed.id,
    }),
  });

  await refreshPage();
}
</script>

<template>
  <div class="feed">
    <span class="feed-header">
      <div class="feed-header-left">
        <span class="feed-header-title-wrapper">
          <p class="feed-header-title">
            {{ feed.feedAuthor }}
          </p>
          <sub class="feed-sub-count">({{ feed.subscribers }} Subs)</sub>
        </span>

        <p class="feed-url">
          <a :href=feed.feedUrl target="_blank">{{ feed.feedUrl }}</a>
        </p>
      </div>

      <div class="feed-header-right">
        <input name="joinFeed" class="feed-join-btn active" v-if="feed.isJoined" type="button" value="Joined" />
        <input name="joinFeed" class="feed-join-btn" v-else type="button" value="Join" @click="() => joinFeed()" />
      </div>
    </span>
  </div>
</template>

<style scoped>
p {
  margin: 0;
}

div.feed {
  border: 1px solid mediumvioletred;
  border-radius: 15px;
  padding: 10px;

  >span.feed-header {
    display: flex;
    justify-content: space-between;
    align-items: center;

    >div.feed-header-left {
      display: flex;
      flex-direction: column;

      >span.feed-header-title-wrapper {
        display: flex;
        align-items: center;

        >p.feed-header-title {
          font-size: large;
          text-decoration: underline;
        }

        >sub.feed-sub-count {
        padding-left: 10px;
          text-decoration: none;
        }
      }
    }

    >div.feed-header-right {
      >input.feed-join-btn{
        width: 70px;
        height: 30px;
        background: none;
        border: 1px solid mediumvioletred;
        border-radius: 5px;
        cursor: pointer;
        text-align: center;
        display: flex;
        justify-content: center;
        align-items: center;

        &.active {
          border: 1px solid darkgreen;
        }
      }
    }
  }
}
</style>
