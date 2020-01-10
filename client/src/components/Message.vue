<template lang="html">
  <div class="ratio">
    <div class="inner" v-bind:style="{ borderColor: this.borderCol }">
      <span
        class="author"
        v-bind:style="{ background: this.colour, borderColor: this.borderCol }"
      >
        {{ author }}
      </span>
      <div class="content">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
import Color from "color";
export default {
  props: {
    author: {
      type: String,
      default: ""
    },
    colour: {
      type: String,
      default: "#e97777"
    }
  },
  computed: {
    borderCol() {
      return Color(this.colour)
        .darken(0.2)
        .hex();
    }
  }
};
</script>

<style lang="scss" scoped>
$perc-spacer: 1%;

.ratio {
  position: relative;
  padding-bottom: 33.333%;
}

.inner {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: #fff;
  border-radius: $spacer;
  border: 0.5vw solid desaturate(darken(#e97777, 20%), 20%);
  overflow: hidden;
}

.content {
  width: 100%;
  height: 100%;
}

.author {
  z-index: 5;
  position: absolute;
  top: 0;
  left: 0;
  padding: $perc-spacer;
  font-size: $size-pixel;
  line-height: 0.625;
  background: #e97777;
  border-bottom: 0.5vw solid transparent;
  border-right: 0.5vw solid transparent;
  border-color: inherit;
  border-bottom-right-radius: $spacer;
  color: #fff;
  font-family: "pixel 5x7";
  user-select: none;
}
</style>
