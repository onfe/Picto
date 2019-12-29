<template>
  <div class="join-form">
    <p v-if="$route.params.id">
      You're joining room: <br /><strong>{{ $route.params.id }}</strong>
    </p>
    <p v-else>Enter a nickname to create a room</p>

    <form @submit.prevent="join">
      <input
        v-model="name"
        class="font-pixel"
        name="name"
        placeholder="Type a nickname"
        autocorrect="off"
        autocapitalize="off"
        autofocus
      />

      <button type="submit" name="button">
        <svg
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
      </button>
    </form>
  </div>
</template>

<script>
export default {
  name: "join-form",
  components: {},
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
  beforeCreate() {
    // check the room exists that we're about to join.
    const id = this.$route.params.id;
    if (!id) {
      // if we're not trying to join a room we don't need to check it exists.
      return;
    }
    const url =
      window.location.origin + `/api/?method=room_exists&room_id=${id}`;
    const options = {
      method: "GET"
    };

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
  }
};
</script>

<style lang="scss" scoped>
form {
  display: flex;
}

input {
  flex: 1;

  width: 100%;
  min-width: 0;

  font-size: 2rem;
  border: 5px solid $grey-d;
  background-color: transparent;

  transition: background-color 0.15s ease-in-out;

  padding: 0.25rem 0.5rem;
  box-sizing: border-box;

  border-radius: 0;
  -webkit-appearance: none;
  -moz-appearance: none;

  &:focus {
    background-color: white;
    outline: none;
  }
}

button {
  width: 20%;

  background-color: $grey-d;
  border: none;
  color: white;

  // margin: 0;
  vertical-align: middle;

  svg {
    fill: $almost-white;
    height: 25px;
  }

  &:hover,
  &:active {
    svg {
      transform: scale(1.1);
    }
  }
}
</style>
