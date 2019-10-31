<template>
  <div class="home">
    <form @submit.prevent="join">
      <input v-model="name" name="name" />
    </form>
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
      this.$store.dispatch("client/join", this.name);
    }
  }
};
</script>
