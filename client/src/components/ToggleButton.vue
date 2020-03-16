<template>
  <ul class="btn">
    <li>
      <font-awesome-icon
        class="icn"
        :class="{ turn: active }"
        :icon="mainIcon"
      />
    </li>
  </ul>
</template>

<script>
export default {
  props: {
    icon: String,
    iconActive: String,
    active: Boolean
  },
  computed: {
    mainIcon() {
      return this.active ? this.iconActive : this.icon;
    }
  }
};
</script>

<style lang="scss" scoped>
.icn {
  transform: rotateZ(0deg);
  transition: transform 200ms ease-in-out;

  &.turn {
    transform: rotateZ(90deg);
  }
}
</style>
