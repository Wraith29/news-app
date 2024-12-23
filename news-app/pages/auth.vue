<script setup lang="ts">
import { ref } from "vue";
import { AuthRequest } from "@/types/auth";

const username = ref("");
const password = ref("");
const error = ref("");

async function submit(path: string): Promise<void> {
  const request = <AuthRequest>{
    username: username.value,
    password: password.value,
  };

  const url = `/api/${path}`;

  await $fetch(url, {
    method: "POST",
    body: JSON.stringify(request),
    responseType: "json",
    credentials: "include",
    async onResponse({ response }) {
    console.log(response);
      switch (response.status) {
        case 401:
        case 500:
          error.value = response._data.data;
          return;
        case 200:
          error.value = "";
          navigateTo("/");
          return;
      }
    },
  });
}
</script>

<template>
  <div id="wrapper">
    <div id="body">
      <p id="header">Login or Register</p>

      <div class="box inputs">
        <div class="form-input">
          <label for="username">Username</label>
          <input
            :class="{ error: error.length !== 0 }"
            name="username"
            placeholder="Username"
            type="text"
            v-model="username"
          />
        </div>

        <div class="form-input">
          <label>Password</label>
          <input
            :class="{ error: error.length !== 0 }"
            name="password"
            placeholder="Password"
            type="password"
            v-model="password"
          />
        </div>
      </div>

      <p class="error" v-if="error.length !== 0">{{ error }}</p>

      <div class="box buttons">
        <input
          name="login"
          value="Login"
          type="button"
          @click="() => submit('login')"
        />

        <input
          name="register"
          value="Register"
          type="button"
          @click="() => submit('register')"
        />
      </div>
    </div>
    <div id="page-split"></div>
  </div>
</template>

<style scoped>
div#wrapper {
  width: 100%;
  height: 100%;
  display: flex;
}

div#body {
  background-color: white;
  margin: 8px 0 8px 8px;
  height: calc(100vh - 16px);
  width: calc(50% - 8px);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  > p#header {
    font-size: xx-large;
    font-weight: bold;
    text-decoration: underline;
  }

  > div.inputs {
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    height: 300px;
    margin-bottom: 25px;

    > div.form-input {
      > input {
        font-size: large;
        width: 100%;
        padding: 10px 0;
        padding-left: 10px;
        border: none;
        border-bottom: 3px solid lightgray;
        border-radius: 3px;

        &.error {
          border-color: red;
        }
      }

      > label {
        display: none;
      }
    }
  }

  p.error {
    color: red;
    font-size: large;
  }

  > div.buttons {
    width: 30%;
    display: flex;
    flex-direction: column;

    > input {
      width: 100%;
      font-size: large;
      padding: 10px;
      background-image: linear-gradient(
        to right,
        mediumvioletred,
        mediumpurple
      );
      color: white;
      border: none;
      border-radius: 5px;
      margin: 3px;
      cursor: pointer;

      &:hover {
        background-image: linear-gradient(
          to left,
          mediumvioletred,
          mediumpurple
        );
      }
    }
  }
}

div#page-split {
  background-color: purple;
  height: 100vh;
  width: 50%;
}
</style>
