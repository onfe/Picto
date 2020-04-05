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
      document.addEventListener(
        "visibilitychange",
        function() {
          this.$store.dispatch("messages/read");
        }.bind(this)
      );
    }
  },
  beforeDestroy() {
    this.$store.dispatch("client/leave");
  },
  metaInfo() {
    if (this.$route.params.id) {
      var unread = this.$store.state.messages.unread;
      return {
        title:
          (unread > 0 ? "(" + unread + " Unread) " : "") +
          this.$store.getters["client/roomTitle"] +
          " - Picto"
      };
    }
  }
};
</script>

<style lang="scss" scoped>
main {
  width: 100vmin;
  height: 100%;
  margin: 0 auto;
  background: #fff;
  min-height: 0;
  display: flex;
  flex-direction: column;
  background: var(--background);
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
  border-right: $border-subtle;
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
  border-bottom: $border-subtle;
  height: calc(100% - #{$sidebar-width});
  margin-top: $sidebar-width;
}

.edit {
  display: grid;
  position: relative;
  grid-template: "T C" "T E";
  grid-auto-columns: min-content;
  grid-auto-rows: min-content;
  flex: 0;
}

.keyboard {
  flex: 1 auto auto;
  min-height: 0;
  border-top: $border-subtle;
}

.toolbox {
  grid-area: T;
  width: $sidebar-width;
  border-right: $border-subtle;
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
