<template lang="html">
  <Message
    @copy="handleCopy"
    :copyable="true"
    :author="msg.author"
    :colour="msg.colour"
  >
    <canvas :id="getID"></canvas>
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
  computed: {
    getID() {
      return "canvas-" + this.msg.hash;
    }
  },
  methods: {
    handleCopy() {
      this.$store.dispatch("compose/copy", this.msg);
    }
  },
  mounted() {
    const canv = document.getElementById(this.getID);
    this.notepad = new Notepad(192, 64, canv);
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
