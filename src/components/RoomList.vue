<template lang="html">
  <div>
    <strong>Rooms:</strong>
    <ul>
      <li
        v-for="(room, index) in this.rooms"
        :key="index"
        v-on:click="disabled ? null : $emit('select', room)"
        :class="{
          selected: selectedRoom ? room.Name == selectedRoom.Name : false,
          disabled: disabled
        }"
      >
        <a>{{ room.Name }} ({{ room.Visible }}:{{ room.Invisible }})</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: ["token", "rooms", "selectedRoom", "disabled"]
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
