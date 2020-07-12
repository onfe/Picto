<template>
  <section class="join-form" id="core">
    <p v-if="$route.params.id">
      You're joining room: <br /><strong>{{ $route.params.id }}</strong>
    </p>
    <p v-else>Enter a nickname to create a room...</p>

    <div class="joinbox font-pixel" :class="{ error: error }">
      <form @submit.prevent="join">
        <input
          v-model="name"
          name="name"
          id="name"
          class="font-pixel"
          placeholder="Type a nickname"
          autocorrect="off"
          autocapitalize="off"
          autofocus
          :disabled="loading"
        />
        <label for="name" class="sr-only">Nickname</label>
        <button
          type="submit"
          name="button"
          :title="
            $route.params.id ? `Join room (${$route.params.id})` : 'Create room'
          "
          :aria-label="
            $route.params.id ? `Join room (${$route.params.id})` : 'Create room'
          "
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
    <router-link v-if="$route.params.id" class="home" to="/">
      I don't want to join this room
    </router-link>
  </section>
</template>

<script>
import Spinner from "@/components/Spinner.vue";

export default {
  name: "join-form",
  components: {
    Spinner
  },
  data() {
    return {
      name: ""
    };
  },
  methods: {
    join() {
      const room = this.$route.params.id;
      const name = this.name;
      if (room) {
        this.$store.dispatch("client/join", { name, room });
      } else {
        this.$store.dispatch("client/join", { name });
      }
    }
  },
  computed: {
    loading() {
      return this.$store.state.client.status === "connecting";
    },
    error() {
      return this.$store.state.client.errorMessage;
    }
  },
  mounted() {
    // check the room exists that we're about to join.
    const id = this.$route.params.id;
    if (!id) {
      // if we're not trying to join a room we don't need to check it exists.
      return;
    }
    const url = window.location.origin + `/api/?method=room_exists&id=${id}`;
    const options = {
      method: "GET"
    };

    setTimeout(() => {
      fetch(url, options)
        .then(resp => {
          return resp.text();
        })
        .then(result => {
          if (result != "true") {
            // room doesn't exist anymore! redirect.
            this.$router.push("/");
          }
        });
    }, 250);
  }
};
</script>

<style lang="scss" scoped>
.joinbox {
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

.join-form {
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
