<template lang="html">
  <div class="ratio">
    <div class="inner">
      <span class="author">Pedanticat</span>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
.ratio {
  position: relative;
  padding-bottom: 33.333%;
}

.inner {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: #fff;
  border-radius: 4px;
  border: 2px solid desaturate(darken(#e97777, 20%), 20%);
  overflow: hidden;
}

.author {
  padding: 4px;
  background: #e97777;
  border-bottom: 2px solid transparent;
  border-right: 2px solid transparent;
  border-color: inherit;
  border-bottom-right-radius: 4px;
  color: #fff;
}
</style>
