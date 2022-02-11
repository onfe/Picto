<template lang="html">
  <div :class="'ratio ' + (this.hidden ? 'hidden' : '')">
    <div class="inner" v-bind:style="{ borderColor: this.borderCol }">
      <div
        v-if="author"
        class="author"
        v-bind:style="{ background: this.colour, borderColor: this.borderCol }"
      >
        <span>{{ author }}</span>
        <span
          v-if="copyable && !hidden"
          class="copy"
          :title="'Copy this message' + (this.hidden ? ' (disabled)' : '')"
        >
          <font-awesome-icon
            @click="!hidden ? $emit('copy') : null"
            class="icn"
            icon="copy"
          />
        </span>
        <span
          v-if="copyable"
          class="hide"
          :title="(this.hidden ? 'Show' : 'Hide') + ' this message'"
        >
          <font-awesome-icon
            @click="$emit('hide')"
            class="icn"
            :icon="this.hidden ? 'eye' : 'eye-slash'"
          />
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
    },
    hidden: {
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
.ratio.hidden {
  padding-bottom: 6.5 * $spacer;
}

.inner {
  z-index: 0;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border-radius: $spacer;
  border: calc($spacer / 2) solid desaturate(darken(#e97777, 20%), 20%);
  overflow: hidden;
}

.content {
  position: relative;
  z-index: -1;
  width: 100%;
  height: 100%;
  background: var(--message-background);

  &:hover {
    opacity: 1;
  }
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
  border-bottom: calc($spacer / 2) solid transparent;
  border-right: calc($spacer / 2) solid transparent;
  border-color: inherit;
  border-bottom-right-radius: $spacer;
  color: #fff;
  font-family: "pixel 5x7";
  user-select: none;
  text-shadow: $shadow-access;

  .copy,
  .hide {
    max-width: 0;
    opacity: 0;
    overflow: hidden;
    display: inline-block;
    height: 3 * $spacer;
    transition: all 400ms ease-in-out;
    cursor: pointer;

    .icn {
      display: block;
      margin-left: $spacer;
      padding: calc($spacer / 4);
      width: 3 * $spacer;
      height: 3 * $spacer;
    }
  }

  &:hover > .copy,
  &:hover > .hide {
    max-width: 4vw;
    opacity: 1;
  }
}
</style>
