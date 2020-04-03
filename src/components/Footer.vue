<template>
  <section>
    <div class="cont">
      <a title="Clear canvas" @click="onClear">Clear</a>
      <a title="Copy last message to canvas" @click="onCopy">Copy</a>
      <a title="Send message" @click="onSend">Send</a>
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
  background: var(--button);
  height: 100%;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;

  a {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    text-align: center;
    font-family: "pixel 5x7";
    font-size: $spacer * 4;
    border-right: $spacer solid var(--background);
    user-select: none;

    &:last-of-type {
      border: 0;
    }
  }
}
</style>
