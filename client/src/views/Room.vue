<template lang="html">
  <main>
    <div class="view">
      <Navbar class="nav" />
      <MessageHistory class="hist" />
    </div>
    <div class="edit">
      <Toolbox class="toolbox" />
      <Compose class="compose" />
      <Footer class="foot" />
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
  width: 100vw;
  height: 100vh;
}

.view {
  flex: 1;
  position: relative;
  display: grid;
  overflow: hidden;
  grid-template: "N H";
}

.nav {
  width: $sidebar-width;
  grid-area: N;
  border-right: 1px solid $almost-white;
}

.hist {
  width: $ratio-width;
  grid-area: H;
  border-bottom: 1px solid $almost-white;
}

.edit {
  flex: 0;
  display: grid;
  grid-template: "T C" "T E";
}

.toolbox {
  grid-area: T;
  width: $sidebar-width;
  border-right: 1px solid $almost-white;
}

.compose {
  grid-area: C;
  width: $ratio-width;
}

.foot {
  grid-area: E;
  width: 100%;
  height: $sidebar-width;
}
</style>
