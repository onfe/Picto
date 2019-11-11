<template lang="html">
  <main>
    <Navbar class="nav" />
    <MessageHistory class="hist" />
    <ToolBox class="toolbox" />
    <Compose class="compose" />
    <Stamps class="stamps" />
    <Keyboard class="keyboard" />
    {{ $route.params.id }}
  </main>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
.app {
  display: grid;
  grid-template:
    "N H"
    "T C"
    "S K";
}
</style>
