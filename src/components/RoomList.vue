<template lang="html">
  <div>
    <strong>Rooms:</strong>
    <ul>
      <li
        v-for="(room, index) in this.rooms"
        :key="index"
        v-on:click="$emit('select', room)"
        :class="{
          selected: selectedRoom ? room.Name == selectedRoom.Name : false
        }"
      >
        <a>{{ room.Name }} ({{ room.Unpublished }}:{{ room.Published }})</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: ["token", "rooms", "selectedRoom"]
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
    color: $grey;
    font-weight: bold;
    text-decoration: underline;
  }

  > * {
    color: inherit;
    line-height: 1.5;
  }
}
</style>
