<template lang="html">
  <Message :author="author" :colour="colour">
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
    data: {
      type: Object,
      default: () => {
        return {
          span: 384,
          data: Array(384 * 128).fill(0)
        };
      }
    },
    author: {
      type: String,
      default: "Author"
    },
    colour: {
      type: String,
      default: "#e97777"
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
    this.notepad.loadImageData(this.data.raw());
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
