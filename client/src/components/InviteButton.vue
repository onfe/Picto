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

<style lang="scss" scoped></style>
