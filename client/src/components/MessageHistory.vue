<template lang="html">
  <section id="hist">
    <div class="cont">
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
  },
  watch: {
    history() {
      const el = document.getElementById('hist')
      const els = document.querySelectorAll('.message');
      if (!el && !els) { return };
      const last = els[0];
      if (el.scrollTop + el.clientHeight + last.clientHeight + 10 > el.scrollHeight) {
        setTimeout(() => last.scrollIntoView(), 10);
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.cont {
  padding: 1vw;
  display: flex;
  flex-direction: column-reverse;
  min-height: 100%;
}

section {
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
