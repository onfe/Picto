<template>
  <section class="nav">
    <ToggleButton
      @click.native="toggleInfo"
      icon="info-circle"
      iconActive="times-circle"
      :active="showInfo"
    />
    <hr />
    <ul class="btn" @click="link">
      <li>
        <font-awesome-icon class="icn" icon="user-plus" />
      </li>
    </ul>
  </section>
</template>

<script>
import ToggleButton from "@/components/ToggleButton.vue";
export default {
  components: {
    ToggleButton
  },
  computed: {
    showInfo() {
      return this.$store.state.client.showInfo;
    }
  },
  methods: {
    toggleInfo() {
      this.$store.dispatch("client/toggleInfo");
    },
    link() {
      return;
    }
  }
};
</script>

<style lang="scss" scoped>
section {
  padding: $spacer;
}

hr {
  border: 0;
  border-bottom: 1px solid $almost-white;
  margin: $spacer;
}
</style>
