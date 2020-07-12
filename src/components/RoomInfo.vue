<template lang="html">
  <aside :class="{ expand: expand }">
    <div class="wrap">
      <RoomTitle :edit="expand" />
      <div class="body">
        <UserList class="userList" />
        <ThemeButton class="themeButton" />
      </div>
    </div>
  </aside>
</template>

<script>
import UserList from "@/components/UserList.vue";
import RoomTitle from "@/components/RoomTitle.vue";
import ThemeButton from "@/components/ThemeButton.vue";
export default {
  components: {
    UserList,
    RoomTitle,
    ThemeButton
  },
  computed: {
    expand() {
      return this.$store.state.room.showInfo;
    }
  }
};
</script>

<style lang="scss" scoped>
aside {
  z-index: 100;
  border-bottom: $border-subtle;
  overflow: hidden;
  max-height: $sidebar-width;

  transition: max-height 400ms ease-in-out;

  &.expand {
    max-height: calc(#{$sidebar-width} + 10em);
  }
}

.wrap {
  position: relative;
  min-height: $sidebar-width;
  background: var(--background);
  padding: $spacer;
}

.body {
  font-family: monospace;
  display: flex;
  flex-wrap: nowrap;
  justify-content: flex-end;
  align-items: center;
}

.userList {
  flex-grow: 1;
}

.themeButton {
  flex-shrink: 0;
  height: $sidebar-width - (2 * $spacer);
  width: $sidebar-width - (2 * $spacer);
  margin: 0 $spacer;
}
</style>
