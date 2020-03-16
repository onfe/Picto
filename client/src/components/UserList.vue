<template lang="html">
  <div class="users">
    <ul>
      <li
        v-for="(userColour, i) in $store.getters['client/userColours']"
        v-bind:key="i"
        :style="`background: ${userColour[1]}`"
      >
        {{ userColour[0] }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
}
</script>

<style lang="scss" scoped>
.users {
  color: $grey;
}

ul {
  margin: 0;
  padding: 0;
  max-height: 6em;
  overflow-y: scroll;
  overflow-x: hidden;
}

li {
  display: inline-block;
  font-family: 'pixel 5x7';
  font-size: 1.5em;
  color: #fff;
  padding: 0.125em 0.25em;
  border-radius: 0.25em;
  margin-right: 0.25em;
  margin-bottom: 0.25em;
}
</style>
