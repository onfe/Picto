<template>
  <ul class="btn">
    <li>
      <router-link v-if="$route.params.id" class="home" to="/">
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
