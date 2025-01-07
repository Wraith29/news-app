<script setup lang="ts">
const authCookie = useCookie("authToken");
authCookie.value = null;

navigateTo("/auth");
</script>

<template></template>
