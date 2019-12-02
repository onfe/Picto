<template>
  <section>
    <div class="cont">
      <span @click="onSend">Send</span>
      <span @click="onClear">Clear</span>
      <input />
    </div>
  </section>
</template>

<script>
export default {
  methods: {
    onClear() {
      this.$store.dispatch("compose/clear");
    },
    onSend() {
      this.$store.dispatch("compose/send");
    }
  }
};
</script>

<style lang="scss" scoped>
section {
  padding: $spacer;
  padding-top: 0;
  position: relative;
}

.cont {
  border-radius: $spacer;
  padding: $spacer;
  border: 0.5vw solid $grey-l;
  background: $almost-white;
  height: 100%;
}
</style>
