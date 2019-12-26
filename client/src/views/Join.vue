<template>
  <div class="home">
    <form @submit.prevent="join">
      <input v-model="name" name="name" />
    </form>
    <button @click="join" type="button" name="button">Join</button>
  </div>
</template>

<script>
export default {
  name: "home",
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
