<template>
  <div class="line font-pixel"></div>
</template>

<style lang="scss" scoped>
.line {
  &::after {
    content: "";
    animation: 1s step-start infinite spinner;
  }
}

@keyframes spinner {
  0%,
  100% {
    content: "|";
  }
  25% {
    content: "/";
  }
  50% {
    content: "—";
  }
  75% {
    content: "\\";
  }
}
</style>
