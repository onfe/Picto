<template>
  <section class="nav">
    <ToggleButton
      @click.native="toggleInfo"
      icon="info-circle"
      iconActive="times-circle"
      :active="showInfo"
    />
    <hr />
    <InviteButton />
  </section>
</template>

<script>
import ToggleButton from "@/components/ToggleButton.vue";
import InviteButton from "@/components/InviteButton.vue";
export default {
  components: {
    ToggleButton,
    InviteButton
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
