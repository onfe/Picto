<template>
  <section>
    <div class="cont">
      <a @click="onClear">Clear</a>
      <a @click="onCopy">Copy</a>
      <a @click="onSend">Send</a>
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
    },
    onCopy() {
      this.$store.dispatch("compose/copy");
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
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;

  a {
    flex: 1;
    text-align: center;
    font-family: "pixel 5x7";
    font-size: 4vw;
    border-right: 0.5vw solid $grey-l;
    user-select: none;

    &:last-of-type {
      border: 0;
    }
  }
}
</style>
