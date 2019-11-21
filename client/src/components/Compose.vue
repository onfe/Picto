<template lang="html">
  <section>
    <div class="container">
      <Message :author="this.$store.state.client.username">
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
