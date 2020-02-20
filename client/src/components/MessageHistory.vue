<template lang="html">
  <section>
    <div class="history">
      <div class="message" v-for="msg in history" v-bind:key="msg.id">
        <CanvasMessage v-if="msg.type == 'normal'" v-bind="msg" />
        <Announcement v-else-if="msg.type == 'announcement'" v-bind="msg" />
        <div class="text" v-else>{{ msg.text }}</div>
      </div>
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
  },
  computed: {
    history() {
      return this.$store.state.messages.history;
    }
  }
};
</script>

<style lang="scss" scoped>
.history {
  padding: 1vw;
  overflow-y: scroll;
  display: flex;
  height: 100%;
  flex-direction: column-reverse;
}

section {
  position: relative;
  height: 100%;
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
