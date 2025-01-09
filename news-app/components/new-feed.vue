<script setup lang="ts"></script>

<template>
  <div id="new-feed-box">
    <p id="new-feed-title">New Feed</p>
  </div>
</template>

<style scoped>
div#new-feed-parent {
  margin: 20px;
  padding: 20px;
  border: 3px solid darkblue;
  border-radius: 25px;

  > p#new-feed-title {
    margin: 0;
    margin-bottom: 10px;
    font-size: x-large;
    text-decoration: underline;
  }
}
</style>
