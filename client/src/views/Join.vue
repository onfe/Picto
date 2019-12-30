<template>
  <div class="join">
    <div class="container">
      <img class="logo" src="/img/logo.svg" alt="picto" />
      <p>
        Picto is an online blah blah words send photos but none of that naughty
        stuff eh?
      </p>
      <Join />
    </div>
  </div>
</template>

<script>
import Join from "@/components/JoinForm.vue";

export default {
  name: "home",
  components: {
    Join
  }
};
</script>

<style lang="scss" scoped>
.join {
  height: 100%;
  background-color: $almost-white;

  background-image: url("/img/stripe.svg");
  background-repeat: repeat-y;
  background-position-x: 0.8rem;
}

.container {
  max-width: 600px;

  padding: 0 2.5rem 0 3.5rem;

  @media (min-width: 992px) {
    padding-left: 8rem;
  }

  font-family: monospace;
  font-size: 1.2rem;
  color: $grey-d;
}

p {
  margin-bottom: 3rem;
  line-height: 1.2;
}

.logo {
  max-width: 100%;
  width: auto;
  margin-bottom: 3rem;

  max-height: 10rem;
}
</style>
