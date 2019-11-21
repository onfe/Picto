<template>
  <div>
    <span @click="onSend">Send</span>
    <span @click="onClear">Clear</span>
  </div>
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

<style lang="css" scoped></style>
