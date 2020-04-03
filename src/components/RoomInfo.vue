<template lang="html">
  <aside :class="{ expand: expand }">
    <div class="wrap">
      <RoomTitle :edit="expand" />
      <div class="body">
        <UserList />
      </div>
    </div>
  </aside>
</template>

<script>
import UserList from "@/components/UserList.vue";
import RoomTitle from "@/components/RoomTitle.vue";
export default {
  components: {
    UserList,
    RoomTitle
  },
  computed: {
    expand() {
      return this.$store.state.client.showInfo;
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
}
</style>
