<template>
  <section class="nav">
    <RoomInfo v-on:hide="hideInfo" v-show="showInfo" />
    <ul class="btn">
      <li>
        <font-awesome-icon @click="info" class="icn" icon="info-circle" />
      </li>
      <li>
        <font-awesome-icon @click="link" class="icn" icon="user-plus" />
      </li>
    </ul>
  </section>
</template>

<script>
import RoomInfo from "@/components/RoomInfo.vue";
export default {
  components: {
    RoomInfo
  },
  data() {
    return {
      showInfo: false
    };
  },
  methods: {
    info() {
      this.showInfo = true;
    },
    hideInfo() {
      this.showInfo = false;
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
</style>
