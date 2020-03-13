<template lang="html">
  <aside>
    <ul class="btn">
      <li>
        <font-awesome-icon
          @click="$emit('hide')"
          class="icn"
          icon="times-circle"
        />
      </li>
    </ul>
    <p>
      Room Information Panel
    </p>
  </aside>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
aside {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 10em;
  background: #fff;
  z-index: 100;
  border-bottom: 1px solid $almost-white;
  padding: $spacer;
}

ul {
  width: $sidebar-width - 2 * $spacer;
}
</style>
