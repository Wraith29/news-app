<script setup lang="ts">
import { AuthResponse, AuthRequest, authStore } from "@/types/auth";
import { ref } from "vue";

enum Tab {
  login,
  register
}

const selected = ref(Tab.login);
const username = ref("");
const password = ref("");
const isError = ref(false);
const errorMessage = ref("");

function select(t: Tab): void {
  selected.value = t;
}

async function readStream<T>(body: ReadableStream): Promise<T> {
  const reader = body.pipeThrough(new TextDecoderStream()).getReader();

  let messageBody = "";
  while (true) {
    const { done, value } = await reader.read();
    if (done) break;

    messageBody += value;
  }

  const result = JSON.parse(messageBody);

  return result as T;
}

async function submit(): Promise<void> {
  const request = <AuthRequest>{
    username: username.value,
    password: password.value,
  };

  const url = selected.value === Tab.login ? "/api/login" : "/api/register"

  await $fetch(url, {
    method: "POST",
    body: JSON.stringify(request),
    responseType: "stream",
    async onResponseError({ response: { body } }) {
      const result = await readStream<{ data: string }>(body);

      isError.value = true;
      errorMessage.value = result.data;
    },
    async onResponse({ response: { body } }) {
      const result = await readStream<AuthResponse>(body);

      localStorage.setItem("authToken", result.authToken);
      authStore.authToken = result.authToken;
      authStore.loggedIn = true;
    }
  })

}
</script>

<template>
  <div class="auth-popup">
    <div class="tabs">
      <div id="login-tab" class="tab" :class="{ selected: selected === Tab.login }" @click="select(Tab.login)">
        <p>Login</p>
      </div>
      <div id="register-tab" class="tab" :class="{ selected: selected === Tab.register }" @click="select(Tab.register)">
        <p>Register</p>
      </div>
    </div>

    <div class="body">
      <div id="username-input">
        <label for="username">Username</label>
        <input :class="{ error: isError }" v-model="username" type="text" placeholder="Username" />
      </div>

      <div id="password-input">
        <label for="password">Password</label>
        <input :class="{ error: isError }" v-model="password" type="password" placeholder="Password" />
      </div>

      <p class="error-msg" v-show="isError">{{ errorMessage }}</p>

      <div id="submit">
        <button @click="submit">Submit</button>
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped>
.auth-popup {
  z-index: 1;
  border: 1px solid black;
  border-radius: 10px;

  position: absolute;
  top: 15%;
  left: 30%;
  width: 40%;
  height: 50%;

  >.tabs {
    width: calc(100% - 1px);
    height: 10%;
    border: 1px solid red;
    display: flex;
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;

    >.tab {
      padding: 0;
      margin: 0;
      width: 100%;
      text-align: center;
      background-color: beige;

      >p {
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: larger;
        margin: 0;
      }

      &#login-tab {
        border-right: 1px solid red;
        border-top-left-radius: 10px;
      }

      &#register-tab {
        border-top-right-radius: 10px;
      }

      &.selected {
        color: blue;
        background-color: orange;
        text-decoration: underline;
      }
    }
  }

  >.body {
    display: flex;
    flex-direction: column;
    height: calc(90% - 40px);
    padding: 20px;

    >div {
      margin: 50px 0;
      display: flex;
      align-items: center;
      flex-direction: column;

      >label {
        display: none;
      }

      >input {
        height: 20px;
        font-size: large;
        width: 80%;
        padding: 5px;
        background: white;
        border: none;
        border-bottom: 3px solid lightgray;
        border-radius: 5px;

        &.error {
          border-color: red;
        }
      }

      >button {
        position: absolute;
        height: 40px;
        width: 80%;
        bottom: 20px;
        font-size: large;
        background: white;
        border: 1px solid violet;
        border-radius: 5px;

        &:hover {
          background-color: pink;
        }
      }
    }

    >p.error-msg {
      color: red;
      text-align: center;
    }
  }
}
</style>
