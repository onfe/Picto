<template lang="html">
  <Message :author="msg.author" :colour="msg.colour">
    <canvas v-bind:id="idHash"></canvas>
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
  data() {
    return {
      idHash: null
    };
  },
  beforeMount() {
    this.idHash = Math.random()
      .toString(36)
      .substring(2, 15);
  },
  mounted() {
    const canv = document.getElementById(this.idHash);
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
