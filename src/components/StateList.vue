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
        <a
          >{{ { submitted: "Unpublished", published: "Published" }[state] }} ({{
            {
              submitted: selectedRoom.Unpublished,
              published: selectedRoom.Published
            }[state]
          }})</a
        >
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: ["selectedRoom", "selectedState"]
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
