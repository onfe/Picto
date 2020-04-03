<template lang="html">
  <div class="ratio">
    <div class="inner" v-bind:style="{ borderColor: this.borderCol }">
      <div
        class="author"
        v-bind:style="{ background: this.colour, borderColor: this.borderCol }"
      >
        <span>{{ author }}</span>
        <span v-if="copyable" class="copy" title="Copy this message">
          <font-awesome-icon @click="$emit('copy')" class="icn" icon="copy" />
        </span>
      </div>
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
    },
    copyable: {
      type: Boolean,
      default: false
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
  z-index: 0;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: #fff;
  border-radius: $spacer;
  border: $spacer / 2 solid desaturate(darken(#e97777, 20%), 20%);
  overflow: hidden;
}

.content {
  position: relative;
  z-index: -1;
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
  border-bottom: $spacer / 2 solid transparent;
  border-right: $spacer / 2 solid transparent;
  border-color: inherit;
  border-bottom-right-radius: $spacer;
  color: #fff;
  font-family: "pixel 5x7";
  user-select: none;
  text-shadow: $shadow-access;

  .copy {
    max-width: 0;
    opacity: 0;
    overflow: hidden;
    display: inline-block;
    height: 3 * $spacer;
    transition: all 200ms ease-in-out;
    cursor: pointer;

    .icn {
      display: block;
      margin-left: $spacer;
      padding: $spacer / 4;
      width: 3 * $spacer;
      height: 3 * $spacer;
    }
  }

  &:hover > .copy {
    max-width: 4vw;
    opacity: 1;
  }
}
</style>