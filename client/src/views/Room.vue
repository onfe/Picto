<template lang="html">
  <main>
    <div class="view">
      <Navbar class="nav" />
      <MessageHistory class="hist" />
    </div>
    <div class="edit">
      <Toolbox class="toolbox" />
      <Compose class="compose" />
      <TextEntry class="text" />
    </div>
  </main>
</template>

<script>
import Compose from "@/components/Compose.vue";
import MessageHistory from "@/components/MessageHistory.vue"

export default {
  components: {
    Compose,
    MessageHistory
  }
};
</script>

<style lang="scss" scoped>
main {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
}

$sidebar-width: 8vw;
$main-width: calc(100vw - #{$sidebar-width});
$main-height: calc(#{$main-width} / 3);

.view {
  flex: 1;
  position: relative;
  display: grid;
  grid-template: "N H";
}

.nav {
  width: $sidebar-width;
  height: 100%;
  grid-area: N;
  background: purple;
}

.hist {
  width: $main-width;
  height: 100%;
  grid-area: H;
}

.edit {
  flex: 0;
  display: grid;
  grid-template: "T C" "E E";
}

.toolbox {
  grid-area: T;
  width: $sidebar-width;
  height: $main-height;
  background: #ccc;
}

.compose {
  grid-area: C;
  width: $main-width;
  height: $main-height;
  background: pink;
}

.text {
  grid-area: E;
  width: 100%;
  height: $sidebar-width;
  background: orange;
}
</style>
