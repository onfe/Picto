<template lang="html">
  <main>
    <div class="view">
      <Navbar class="nav" />
      <MessageHistory class="hist" />
    </div>
    <div class="edit">
      <Toolbox class="toolbox" />
      <Compose class="compose" />
      <Footer class="text" />
    </div>
  </main>
</template>

<script>
import Compose from "@/components/Compose.vue";
import MessageHistory from "@/components/MessageHistory.vue";
import Footer from "@/components/Footer.vue";
import Toolbox from "@/components/Toolbox.vue";
import Navbar from "@/components/Navbar.vue";

export default {
  components: {
    Compose,
    MessageHistory,
    Footer,
    Toolbox,
    Navbar
  },
  mounted() {
    if (this.$store.state.client.room == null) {
      this.$router.replace(`/join/${this.$route.params.id}`);
    }
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
  grid-template: "T C" "T E";
}

.toolbox {
  grid-area: T;
  width: $sidebar-width;
}

.compose {
  grid-area: C;
  width: $ratio-width;
}

.text {
  grid-area: E;
  width: 100%;
  height: 2em;
  background: orange;
}
</style>
