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
        <a>{{ room.Name }} ({{ room.Unpublished }})</a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: ["token", "selectedRoom"],
  data() {
    return {
      rooms: []
    };
  },
  mounted() {
    this.refresh();
  },
  methods: {
    refresh() {
      const url =
        window.location.origin +
        "/api/?method=get_submission_rooms&token=" +
        this.token;
      const options = {
        method: "GET"
      };

      fetch(url, options)
        .then(resp => {
          return resp.text();
        })
        .then(result => {
          this.rooms = JSON.parse(result) || [];
          this.rooms.map(r => (r.Full = r.Pop >= r.Cap));
        });
    }
  }
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
