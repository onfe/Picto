<template>
  <div>
    <strong>States:</strong>
    <ul>
      <li
        v-for="state in ['submitted', 'published']"
        :key="state"
        v-on:click="$emit('select', state)"
        :class="{ selected: selectedState == state }"
      >
        <a>{{ state }}</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: ["selectedState"]
};
</script>

<style lang="scss" scoped>
ul {
  width: 100%;
  margin: $spacer 0;
  padding: 0;
}

li {
  margin: 0;
  padding: 0;
  display: flex;
  width: 100%;
  justify-content: space-between;
  color: $grey-d;

  &.selected {
    color: $grey;
    font-weight: bold;
  }

  > * {
    color: inherit;
    line-height: 1.5;
    text-decoration: underline;
  }
}
</style>
