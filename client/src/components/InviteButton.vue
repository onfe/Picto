<template lang="html">
  <ul class="btn" @click="copy" :class="{ copied: copied }">
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
      navigator.clipboard
        .writeText(this.$store.getters["client/inviteLink"])
        .then(() => {
          this.copied = true;
          setTimeout(() => (this.copied = false), 1000);
        });
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
  width: 22.5vw;
  left: calc(100% + #{$spacer} * 2);
  margin-top: calc(#{$spacer}/ 2);

  font-family: "pixel 5x7";
  font-size: 4vw;

  border-radius: $spacer;
  border: 0.5vw solid $grey-l;
  padding: 0 $spacer;

  background: $almost-white;

  z-index: 1000;
}

// bit jank but ensures animation is consistent and immediate.
.icn {
  background: $almost-white;
  animation: hi-fade-active 400ms;
  &:active {
    animation: none;
    animation: hi-fade-hover 400ms;
  }
}

@include highlightFade("hi-fade-hover", "background", $grey-l, $almost-white);
@include highlightFade("hi-fade-active", "background", $grey-l, $almost-white);
</style>
