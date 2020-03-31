<template>
  <div class="join">
    <a class="sr-only" href="#core">Skip to join form</a>
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
    <footer>
      <font-awesome-icon class="icn" :icon="'bug'" />
      <a
        href="https://github.com/onfe/Picto/issues?q=is%3Aissue+is%3Aopen+label%3Abug"
        >Found a bug?</a
      >
    </footer>
  </div>
</template>

<script>
import JoinForm from "@/components/JoinForm.vue";
import PublicRooms from "@/components/PublicRooms.vue";

export default {
  name: "join",
  components: {
    JoinForm,
    PublicRooms
  },
  metaInfo() {
    if (this.$route.params.id) {
      return {
        title: `Join ${this.$route.params.id} - Picto`
      };
    }
  }
};
</script>

<style lang="scss" scoped>
.join {
  min-height: 100%;
  background-color: $almost-white;

  background-image: url("/img/stripe.svg");
  background-repeat: repeat-y;
  background-position-x: 0.8rem;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.container,
footer {
  max-width: 600px;

  padding: 0 2.5rem 1rem 3.5rem;

  @media (min-width: 992px) {
    padding-left: 8rem;
  }

  font-family: monospace;
  font-size: 1.2rem;
  color: $grey-d;
}

footer {
  font-size: 0.8rem;
  .icn {
    margin-right: 1ch;
  }
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
