<template lang="html">
  <main>
    <div class="view">
      <Navbar class="nav" />
      <div class="feed">
        <RoomInfo class="info" />
        <MessageHistory class="hist" />
      </div>
    </div>
    <div class="edit">
      <Toolbox @keyboard="handleKeys" class="toolbox" />
      <Compose class="compose" />
      <Footer class="foot" />
    </div>
    <Keyboard :show="showKeyboard" class="keyboard"></Keyboard>
  </main>
</template>

<script>
import Compose from "@/components/Compose.vue";
import MessageHistory from "@/components/MessageHistory.vue";
import Footer from "@/components/Footer.vue";
import Toolbox from "@/components/Toolbox.vue";
import Navbar from "@/components/Navbar.vue";
import Keyboard from "@/components/Keyboard.vue";
import RoomInfo from "@/components/RoomInfo.vue";

export default {
  components: {
    Compose,
    MessageHistory,
    Footer,
    Toolbox,
    Navbar,
    Keyboard,
    RoomInfo
  },
  data() {
    return {
      showKeyboard: false
    };
  },
  methods: {
    handleKeys() {
      this.showKeyboard = !this.showKeyboard;
    }
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
  display: flex;
  flex-direction: column;
}

.view {
  position: relative;
  flex: 1;
  display: flex;
  flex-direction: row;
  min-height: 0;
}

.nav {
  width: $sidebar-width;
  height: 100%;
  border-right: 1px solid $almost-white;
}

.feed {
  position: relative;
  flex: 0;
  width: $ratio-width;
  min-height: 0;
}

.info {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
}

.hist {
  width: $ratio-width;
  border-bottom: 1px solid $almost-white;
  height: calc(100% - #{$sidebar-width});
  margin-top: $sidebar-width;
}

.edit {
  display: grid;
  grid-template: "T C" "T E";
  flex: 0;
}

.keyboard {
  flex: 1 auto auto;
  min-height: 0;
  border-top: 1px solid $almost-white;
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
