<template lang="html">
  <aside :class="{ expand: expand }">
    <div class="wrap">
      <h1 class="room-title">
        {{ $store.getters["client/roomTitle"] }}
      </h1>
      <div class="body">
        <UserList />
      </div>
    </div>
  </aside>
</template>

<script>
import UserList from "@/components/UserList.vue"
export default {
  components: {
    UserList
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
  border-bottom: 1px solid $almost-white;
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
  background: #fff;
  padding: $spacer;
}

.body {
  font-family: monospace;
}

.room-title {
  height: $sidebar-width - $spacer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: "pixel 5x7";
  font-size: 6vw;
  margin: 0;
  padding-bottom: $spacer;
  font-weight: normal;
}
</style>
