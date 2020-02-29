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
  },
  beforeDestroy() {
    this.$store.dispatch("client/leave");
  }
};
</script>

<style lang="scss" scoped>
main {
  width: 100%;
  height: 100%;
  min-height: 0;
}

.view {
  position: relative;
  overflow: hidden;
  height: $height-upper;
}

.nav {
  width: $sidebar-width;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  display: inline-block;
  border-right: 1px solid $almost-white;
}

.hist {
  width: $ratio-width;
  margin-left: $sidebar-width;
  border-bottom: 1px solid $almost-white;
}

.edit {
  display: grid;
  grid-template: "T C" "T E";
  height: $height-lower;
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
