<template lang="html">
  <ul class="btn" @click="copy" :class="{ copied: copied }">
    <li>
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
