<template>
  <section class="nav">
    <ExitButton class="fix" />
    <hr />
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
import ExitButton from "@/components/ExitButton.vue";
import InviteButton from "@/components/InviteButton.vue";
import SidebarHistory from "@/components/SidebarHistory.vue";
export default {
  components: {
    ToggleButton,
    ExitButton,
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
  width: 100%;
  margin: 0 0 $spacer 0;
}

.fix {
  flex-shrink: 0;
}

.hist {
  flex-shrink: 1;
  flex-grow: 1;
}
</style>
