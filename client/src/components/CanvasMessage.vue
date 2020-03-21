<template lang="html">
  <Message :author="msg.author" :colour="msg.colour">
    <canvas :id="getID(msg)"></canvas>
  </Message>
</template>

<script>
import Notepad from "@/assets/js/notepad.js";
import Message from "@/components/Message.vue";

export default {
  components: {
    Message
  },
  props: {
    msg: {
      type: Object,
      default: () => new Message()
    }
  },
  methods: {
    getID(msg) {
      return "canvas-" + msg.hash;
    }
  },
  mounted() {
    const canv = document.getElementById(this.getID(this.msg));
    this.notepad = new Notepad(384, 128, canv);
    this.notepad.loadImageData(this.msg.raw());
  }
};
</script>

<style lang="scss" scoped>
canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  image-rendering: pixelated;
}
</style>
