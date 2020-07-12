<template lang="html">
  <section>
    <div class="container">
      <Message
        id="compose"
        :colour="this.$store.state.client.colour"
        :author="this.$store.state.client.username"
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
      sketchpad: null,
      firstkey: true
    };
  },
  mounted() {
    setTimeout(
      function() {
        const canv = document.getElementById("sketchpad");
        const nametag = document
          .getElementById("compose")
          .querySelector(".author");
        const perc = nametag.clientWidth / canv.clientWidth;

        this.sketchpad = new Sketchpad(192, 64, canv, perc);
        window._sketch = this.sketchpad;

        document.addEventListener("keypress", this.handleKeys);
        document.addEventListener("keydown", this.handleBack);
      }.bind(this),
      50
    );
  },
  beforeDestroy() {
    document.removeEventListener("keypress", this.handleKeys);
    document.removeEventListener("keydown", this.handleBack);
  },
  methods: {
    handleKeys: function(e) {
      if (e.key == "Enter") {
        this.$store.dispatch("compose/send");
      } else {
        this.$store.dispatch("compose/write", e.key);
      }
      e.stopPropagation();
    },
    handleBack(e) {
      if (e.which == 8) {
        // desktop backspace
        this.$store.dispatch("compose/backspace");
        e.stopPropagation();
        e.preventDefault();
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
  padding: $spacer;
}

canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  image-rendering: pixelated;
  touch-action: none;
  background-image: linear-gradient(
    rgba(0, 0, 0, 0.1) #{$spacer / 4},
    transparent 0
  );
  background-size: 100% 19%;
  cursor: crosshair;
}
</style>
