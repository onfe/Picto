<template lang="html">
  <section>
    <div
      class="message"
      v-for="(msg, key) in this.$store.state.messages.history"
      v-bind:key="key"
    >
      <Message />
    </div>
  </section>
</template>

<script>
import Message from "@/components/Message.vue"
export default {
  components: {
    Message
  }
}
</script>

<style lang="scss" scoped>
section {
  padding: 1vw;
  display: flex;
  flex-direction: column-reverse;
}

.message {
  margin-top: 4px;
}
</style>
