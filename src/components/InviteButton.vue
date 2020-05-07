<template lang="html">
  <ul
    title="Copy invite link"
    class="btn"
    @click="copy"
    :class="{ copied: copied }"
  >
    <div v-if="copied" class="notif">Copied Join Link!</div>
    <li class="btn">
      <font-awesome-icon class="icn" icon="user-plus" />
    </li>
  </ul>
</template>

<script>
export default {
  methods: {
    copy() {
      if (navigator.canShare) {
        navigator.share({
          title: "Picto",
          text: "Join me on Picto!",
          url: this.$store.getters["room/invite"]
        });
      } else {
        navigator.clipboard
          .writeText(this.$store.getters["room/invite"])
          .then(() => {
            this.copied = true;
            setTimeout(() => (this.copied = false), 1000);
          });
      }
    }
  },
  data() {
    return {
      copied: false
    };
  }
};
</script>

<style lang="scss" scoped>
ul {
  overflow: visible;
}
.notif {
  position: absolute;
  display: block;
  white-space: nowrap;
  left: calc(100% + #{$spacer} * 2);
  margin-top: calc(#{$spacer}/ 2);

  font-family: "pixel 5x7";
  font-size: $size-pixel;

  border-radius: $spacer;
  border: $spacer / 2 solid var(--secondary);
  padding: 0 $spacer;

  background: var(--subtle);

  z-index: 1000;
}

// bit jank but ensures animation is consistent and immediate.
.icn {
  animation: hi-fade-active 400ms;
  &:active {
    animation: none;
    animation: hi-fade-hover 400ms;
  }
}

@include highlightFade("hi-fade-hover", "background", $grey-l, $almost-white);
@include highlightFade("hi-fade-active", "background", $grey-l, $almost-white);
</style>
