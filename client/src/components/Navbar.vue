<template>
  <section class="nav">
    <ToggleButton
      @click.native="toggleInfo"
      icon="info-circle"
      iconActive="times-circle"
      :active="showInfo"
      class="fix"
    />
    <hr />
    <InviteButton class="fix" />
    <SidebarHistory class="hist" />
  </section>
</template>

<script>
import ToggleButton from "@/components/ToggleButton.vue";
import InviteButton from "@/components/InviteButton.vue";
import SidebarHistory from "@/components/SidebarHistory.vue";
export default {
  components: {
    ToggleButton,
    InviteButton,
    SidebarHistory
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
  display: flex;
  flex-direction: column;
}

hr {
  border: 0;
  border-bottom: 1px solid $almost-white;
  margin: $spacer;
}

.fix {
  flex-shrink: 0;
}

.hist {
  flex-shrink: 1;
  flex-grow: 1;
  margin-top: $spacer;
}
</style>
