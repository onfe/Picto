<template lang="html">
  <ul>
    <a
      v-for="msg in history"
      v-bind:key="msg.time"
      :style="getStyle(msg)"
      :id="getID(msg)"
      :href="getHref(msg)"
      :class="getType(msg)"
    ></a>
  </ul>
</template>

<script>
export default {
  computed: {
    history() {
      return this.$store.state.messages.history;
    }
  },
  methods: {
    getID(msg) {
      return "pip-" + msg.hash;
    },
    getHref(msg) {
      return "#msg-" + msg.hash;
    },
    getStyle(msg) {
      return "background: " + msg.colour;
    },
    getType(msg) {
      return msg.type + "-pip";
    }
  }
};
</script>

<style lang="scss" scoped>
ul {
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column-reverse;
  overflow: hidden;

  a {
    display: block;
    width: 100%;
    height: $spacer;
    border-radius: $spacer;
    background: #000;
    margin-top: calc($spacer / 2);
    flex-shrink: 0;
  }
}

@media (prefers-color-scheme: dark) {
  :root:not(.dark):not(.light):not(.pink) {
    a.Announcement-pip {
      border: calc($spacer / 4) solid var(--secondary);
    }
  }
}

:root.dark {
  a.Announcement-pip {
    border: calc($spacer / 4) solid var(--secondary);
  }
}
</style>
