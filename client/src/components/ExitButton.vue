<template>
  <ul class="btn">
    <li>
      <router-link
        title="Leave room"
        class="home"
        :to="{ name: 'join', params: { id: $route.params.id }}"
      >
        <font-awesome-icon class="icn" :icon="'door-open'" />
      </router-link>
    </li>
  </ul>
</template>

<style lang="scss" scoped>
a {
  color: #000;
}
</style>
