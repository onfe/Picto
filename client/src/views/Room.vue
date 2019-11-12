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
import MessageHistory from "@/components/MessageHistory.vue";

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
  width: $ratio-width;
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
  height: $ratio-height;
  background: #ccc;
}

.compose {
  grid-area: C;
  width: $ratio-width;
  height: $ratio-height;
}

.text {
  grid-area: E;
  width: 100%;
  height: 2em;
  background: orange;
}
</style>
