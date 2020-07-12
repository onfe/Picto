<template lang="html">
  <section v-if="this.rooms.length > 0">
    <p>...or join a public room and get drawing</p>
    <ol class="public-rooms">
      <li
        v-for="(room, index) in this.rooms"
        v-bind:key="index"
        :class="{ full: room.Full }"
      >
        <router-link
          class="room-name"
          :to="room.Full ? '#' : 'join/' + room.Name"
        >
          {{ room.Name }}
        </router-link>
        <div v-if="room.Full">Full!</div>
        <div>{{ room.Pop }}/{{ room.Cap }}</div>
      </li>
    </ol>
  </section>
</template>

<script>
export default {
  data() {
    return {
      rooms: []
    };
  },
  mounted() {
    const url = window.location.origin + "/api/?method=get_public_rooms";
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
};
</script>

<style lang="scss" scoped>
p {
  margin-bottom: 1.5rem;
  line-height: 1.2;
}

ol {
  width: 100%;
  margin: 0;
  padding: 0;
}

li {
  margin: 0;
  padding: 0;
  display: flex;
  width: 100%;
  justify-content: space-between;
  color: var(--primary-join);

  &.full {
    color: $grey;
  }

  &.full > a {
    text-decoration: line-through;
  }

  > * {
    color: inherit;
    line-height: 1.5;
  }
}
</style>
