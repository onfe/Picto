<template lang="html">
  <section>
    <div
      class="message"
      v-for="msg in this.$store.state.messages.history"
      v-bind:key="msg.id"
    >
      <CanvasMessage v-if="msg.type == 'normal'" v-bind="msg" />
      <Announcement v-else-if="msg.type == 'announcement'" v-bind="msg" />
      <div class="text" v-else>{{ msg.text }}</div>
    </div>
  </section>
</template>

<script>
import CanvasMessage from "@/components/CanvasMessage.vue";
import Announcement from "@/components/Announcement.vue";

export default {
  components: {
    CanvasMessage,
    Announcement
  }
};
</script>

<style lang="scss" scoped>
section {
  padding: 1vw;
  display: flex;
  flex-direction: column-reverse;
  overflow-y: scroll;
}

.message {
  margin-top: 1vw;
}

.text {
  font-size: 0.85em;
  text-align: center;
  color: $grey;
  font-family: monospace;
}
</style>
