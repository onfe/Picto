<template lang="html">
  <div class="room-title">
    <h1 class="room-title" v-show="!editable">
      {{ $store.getters["room/title"] }}
    </h1>
    <!-- Stop events here, to prevent propogation to compose element. -->
    <input
      @blur="handleBlur"
      @focus="handleFocus"
      ref="input"
      @keypress.stop
      @keydown="handleEnter"
      v-show="editable"
    />
  </div>
</template>

<script>
export default {
  props: {
    edit: Boolean
  },
  watch: {
    edit(val) {
      if (val) {
        this.$refs["input"].value = this.$store.getters["room/title"];
      }
    }
  },
  computed: {
    editable() {
      // static rooms are not renamable!
      return this.edit && !this.$store.state.client.public;
    }
  },
  methods: {
    handleBlur() {
      const newTitle = this.$refs["input"].value;
      if (newTitle != this.$store.getters["room/title"]) {
        this.$store.dispatch(
          "socket/send",
          {
            event: "rename",
            payload: {
              RoomName: newTitle
            }
          },
          {
            root: true
          }
        );
      }
    },
    handleFocus() {
      this.$refs["input"].value = this.$store.getters["room/title"];
    },
    handleEnter(e) {
      e.stopPropagation();
      // if enter, save and blur.
      if (e.key === "Enter") {
        this.$refs["input"].blur();
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.room-title {
  > * {
    height: $sidebar-width - $spacer;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-family: "pixel 5x7";
    font-size: $size-pixel;
    margin: 0;
    font-weight: normal;
  }

  h1 {
    border: 1px solid transparent;
  }

  input {
    display: block;
    margin-bottom: $spacer;
    text-align: center;
    border: 1px solid $grey-l;
    border-radius: $spacer;
    height: $sidebar-width - $spacer;
    background: var(--message-background);
  }
}
</style>
