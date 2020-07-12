<template>
  <div>
    <strong>States:</strong>
    <ul>
      <li
        v-for="state in ['visible', 'invisible']"
        :key="state"
        v-on:click="disabled ? null : $emit('select', state)"
        :class="{ selected: selectedState == state, disabled: disabled }"
      >
        <a
          >{{ { visible: "Visible", invisible: "Invisible" }[state] }} ({{
            {
              visible: selectedRoom.Visible,
              invisible: selectedRoom.Invisible
            }[state]
          }})</a
        >
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: ["selectedRoom", "selectedState", "disabled"]
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
  white-space: nowrap;

  &.selected {
    font-weight: bold;
    text-decoration: underline;
  }

  &.disabled {
    color: $grey;
  }

  > * {
    color: inherit;
    line-height: 1.5;
  }
}
</style>
