<template lang="html">
  <section>
    <div class="container">
      <Message
        :colour="this.$store.state.client.colour"
        :author="this.$store.getters['client/username']"
      >
        <canvas id="sketchpad"></canvas>
      </Message>
    </div>
  </section>
</template>

<script>
import Sketchpad from "@/assets/js/sketchpad.js";
import Message from "@/components/Message.vue";
export default {
  name: "Compose",
  data() {
    return {
      sketchpad: null
    };
  },
  mounted() {
    const canv = document.getElementById("sketchpad");
    this.sketchpad = new Sketchpad(192, 64, canv);
    window._sketch = this.sketchpad;

    window.addEventListener("keypress", this.handleKeys);
    window.addEventListener("keyup", this.handleBack);
  },
  beforeDestroy() {
    window.removeEventListener("keypress", this.handleKeys);
    window.removeEventListener("keyup", this.handleBack);
  },
  methods: {
    handleKeys: function(e) {
      if (e.key == "Enter") {
        this.$store.dispatch("compose/send");
      } else {
        this.$store.dispatch("compose/write", e.key);
      }
      e.preventDefault();
      e.stopPropagation();
    },
    handleBack(e) {
      if (e.which == 8) {
        // desktop backspace
        this.$store.dispatch("compose/backspace");
      }
    }
  },
  components: {
    Message
  }
};
</script>

<style lang="scss" scoped>
.container {
  position: relative;
  height: 100%;
  width: 100%;
  padding: 1vw;
}

canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  image-rendering: pixelated;
  touch-action: none;
}
</style>
