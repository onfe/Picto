<template lang="html">
  <section :class="{ firefox: isFirefox() }">
    <div class="history">
      <div
        class="message"
        v-for="msg in history"
        v-bind:key="msg.time"
        :id="getID(msg)"
      >
        <CanvasMessage v-if="msg.type == 'Message'" :msg="msg" />
        <Announcement v-else-if="msg.type == 'Announcement'" v-bind="msg" />
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
  },
  methods: {
    getID(msg) {
      return "msg-" + msg.hash;
    },
    isFirefox() {
      // Browser detection of firefox
      return typeof InstallTrigger !== "undefined";
    }
  }
};
</script>

<style lang="scss" scoped>
.history {
  padding: $spacer;
  overflow-y: scroll;
  overflow-x: hidden;
  display: flex;
  height: 100%;
  flex-direction: column-reverse;
}

.message {
  margin-top: $spacer;
}

.firefox {
  .history {
    flex-direction: column !important;
  }

  .history,
  .message {
    transform: scaleY(-1);
  }
}

.text {
  font-size: 0.85em;
  text-align: center;
  color: $grey;
  font-family: monospace;
}
</style>
