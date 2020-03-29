<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  metaInfo: {
    title: "Picto - Draw & chat with friends"
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
