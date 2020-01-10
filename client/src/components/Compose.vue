<template lang="html">
  <section>
    <div class="container">
      <Message
        id="compose"
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

        this.sketchpad = new Sketchpad(384, 128, canv, perc);
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
      e.preventDefault();
      e.stopPropagation();
    },
    handleBack(e) {
      if (e.which == 8) {
        // desktop backspace
        this.$store.dispatch("compose/backspace");
        e.preventDefault();
        e.stopPropagation();
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
  background-image: linear-gradient($grey-l 0.25vw, transparent 0);
  background-size: 100% 19%;
}
</style>
