<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  metaInfo: {
    title: "Picto - Draw & chat with friends"
  },
  created() {
    document.addEventListener("sw-status", e => {
      this.$store.dispatch("client/swUpdate", e.detail);
    });
  },
  mounted() {
    if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
      document.getElementsByTagName("html")[0].setAttribute("class", "dark");
    }
  }
};
</script>

<style lang="scss">
#app {
  height: 100%;
  position: relative;
  display: block;
  min-height: 0;
  background: $grey-d;
}
</style>
