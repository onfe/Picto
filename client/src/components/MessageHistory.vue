<template lang="html">
  <section>
    <div
      class="message"
      v-for="msg in this.$store.state.messages.history"
      v-bind:key="msg.id"
    >
      <CanvasMessage v-if="msg.type == 'normal'" v-bind="msg" />
      <Announcement v-if="msg.type == 'announcement'" v-bind="msg" />
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
</style>
