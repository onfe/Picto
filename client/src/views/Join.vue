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
  created() {
    console.log(this.$route);
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
  }
};
</script>
