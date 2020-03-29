<template>
  <div class="join">
    <div class="container">
      <img class="logo" src="/img/logo.svg" alt="picto" />
      <p>
        Send doodles to your friends and chat with Picto, the scribble-powered
        online messenger.
      </p>
      <hr />
      <JoinForm />
      <PublicRooms v-if="!$route.params.id" />
    </div>
  </div>
</template>

<script>
import JoinForm from "@/components/JoinForm.vue";
import PublicRooms from "@/components/PublicRooms.vue";

export default {
  name: "home",
  components: {
    JoinForm,
    PublicRooms
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

hr {
  border: 0;
  border-bottom: 1px dashed $grey-l;
}

p {
  margin-bottom: 1.5rem;
  line-height: 1.2;
}

.logo {
  max-width: 100%;
  width: auto;
  margin-bottom: 3rem;

  max-height: 10rem;
}
</style>
