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
}
</script>

<style lang="scss" scoped>
</style>
