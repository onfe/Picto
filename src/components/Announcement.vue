<template lang="html">
  <div class="announcement">
    {{ this.text }}
  </div>
</template>

<script>
export default {
  props: {
    text: {
      type: String,
      default: "Announcement"
    }
  }
};
</script>

<style lang="scss" scoped>
.announcement {
  padding: $spacer;
  font-family: "pixel 5x7";
  line-height: 0.625;
  font-size: $size-pixel;
  color: #fff;
  background: #000;
  border: $spacer / 2 solid var(--secondary);
  border-radius: $spacer;
}
</style>
