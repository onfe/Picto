<template lang="html">
  <section id="hist">
    <div class="cont">
      <div class="message" v-for="msg in history" v-bind:key="msg.id">
        <CanvasMessage v-if="msg.type == 'normal'" v-bind="msg" />
        <Announcement v-else-if="msg.type == 'announcement'" v-bind="msg" />
        <div class="text" v-else>{{ msg.text }}</div>
      </div>
      <div class="spacer">
        🐈
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
      const el = document.getElementById("hist");
      if (!el) {
        return;
      }
      if (el.scrollTop + el.clientHeight + 15 > el.scrollHeight) {
        setTimeout(() => {
          document
            .querySelectorAll(".message")[0]
            .scrollIntoView({ behavior: "smooth" });
        }, 10);
      }
    }
  },
  mounted() {
    const el = document.getElementById("hist");
    el.scrollTop = el.scrollHeight - el.clientHeight;
    setTimeout(() => {
      const el = document.getElementById("hist");
      el.scrollTop = el.scrollHeight - el.clientHeight;
    }, 50);
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

.spacer {
  height: 150vh;
  text-align: center;
  font-size: 1.5em;
  color: $grey-l;
}

.text {
  font-size: 0.85em;
  text-align: center;
  color: $grey;
  font-family: monospace;
}
</style>
