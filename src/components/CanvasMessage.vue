<template lang="html">
  <Message
    @copy="handleCopy"
    @hide="handleHide"
    :copyable="true"
    :author="msg.author"
    :colour="msg.colour"
    :hidden="msg.hidden"
  >
    <canvas v-show="!msg.hidden" :id="getID"></canvas>
    <div v-if="msg.hidden" class="canvasCover">Hidden</div>
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
    },
    handleHide() {
      this.$store.dispatch("messages/toggleHidden", this.msg);
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
$perc-spacer: 1%;

canvas,
.canvasCover {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  image-rendering: pixelated;
}
.canvasCover {
  font-family: "pixel 5x7";
  padding: $perc-spacer;
  font-size: $size-pixel;
  line-height: 0.625;
  text-align: center;
  color: #000;
}
</style>
