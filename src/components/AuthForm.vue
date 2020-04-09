<template>
  <section class="auth-form" id="core">
    <p>Enter your auth token...</p>

    <div class="authbox font-pixel" :class="{ error: error }">
      <form @submit.prevent="authenticate">
        <input
          v-model="token"
          type="password"
          name="token"
          id="token"
          class="font-pixel"
          placeholder="Type your auth token"
          autocorrect="off"
          autocapitalize="off"
          autofocus
          :disabled="loading"
        />
        <label for="token" class="sr-only">Token</label>
        <button
          type="submit"
          name="button"
          title="Submit token"
          aria-label="Submit token"
          :disabled="loading"
        >
          <svg
            v-if="!loading"
            xmlns="http://www.w3.org/2000/svg"
            width="21"
            height="35"
            viewBox="0 0 21 35"
          >
            <path
              d="M551,25 L551,18 L544,18 L544,25 L551,25 Z M558,32 L558,25 L551,25 L551,32 L558,32 Z M565,39 L565,32 L558,32 L558,39 L565,39 Z M558,46 L558,39 L551,39 L551,46 L558,46 Z M551,53 L551,46 L544,46 L544,53 L551,53 Z"
              transform="translate(-544 -18)"
            />
          </svg>
          <Spinner v-else />
        </button>
      </form>
      <div class="error-message" v-if="error">
        {{ error }}
      </div>
    </div>
  </section>
</template>

<script>
import Spinner from "@/components/Spinner.vue";

export default {
  name: "auth-form",
  components: {
    Spinner
  },
  data() {
    return {
      token: "",
      error: null,
      loading: false
    };
  },
  methods: {
    authenticate() {
      this.loading = true;

      const token = this.token;
      if (token.length == 0) {
        this.error = "no token provided";
        this.loading = false;
        return;
      }

      const url = window.location.origin + `/api/?method=auth&token=${token}`;
      const options = {
        method: "GET"
      };

      setTimeout(() => {
        fetch(url, options)
          .then(resp => {
            return resp.text();
          })
          .then(result => {
            this.loading = false;
            if (result == "true") {
              // the token if valid!
              this.$emit("authenticated", token);
            } else {
              this.error = "invalid token";
            }
          });
      }, 250);
    }
  }
};
</script>

<style lang="scss" scoped>
.authbox {
  background: $grey-d;
  transition: background 200ms ease-in-out;
  padding: 0.5rem;

  &.error {
    background: $red-d;
  }
}

.error-message {
  padding-top: 0.75rem;
  color: $almost-white;
  font-size: 1.5rem;
}

form {
  display: flex;
}

input {
  flex: 1;

  width: 100%;
  min-width: 0;

  font-size: 2rem;
  background-color: transparent;

  padding: 0.25rem 0.5rem;
  border: 0;
  background: #fff;

  border-radius: 0;
  -webkit-appearance: none;
  -moz-appearance: none;

  &:focus {
    background-color: white;
    outline: none;
  }
}

button {
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;

  max-width: 2em;
  min-width: 1em;
  width: 15%;
  margin-left: 0.5rem;

  background-color: transparent;
  border: none;
  color: white;

  font-size: 2.5rem;

  svg {
    fill: $almost-white;
    height: 25px;
    transition: transform 400ms ease-in-out;
  }

  &:hover,
  &:active {
    svg {
      transform: translateX(0.125em);
    }
  }
}

p {
  line-height: 1.2;
}

.auth-form {
  color: var(--primary-join);
}

.home {
  line-height: 1.2;
  display: block;
  color: var(--secondary-join);
  font-size: 0.85rem;
  margin-block-start: 1em;
  margin-block-end: 1em;
}
</style>
