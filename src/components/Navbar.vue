<template>
  <section class="nav">
    <ul class="btn fix">
      <li>
        <router-link
          title="Leave room"
          class="home"
          :to="{ name: 'join', params: { id: $route.params.id } }"
        >
          <font-awesome-icon class="icn" :icon="'door-open'" />
        </router-link>
      </li>
    </ul>
    <hr />
    <ToggleButton
      @click.native="toggleInfo"
      icon="info-circle"
      iconActive="times-circle"
      :active="showInfo"
      class="fix"
      title="Toggle room info"
    />
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
  border-bottom: $border-subtle;
  width: 100%;
  margin: 0 0 $spacer 0;
}

.btn a {
  color: var(--primary)
}

.fix {
  flex-shrink: 0;
}

.hist {
  flex-shrink: 1;
  flex-grow: 1;
}
</style>
