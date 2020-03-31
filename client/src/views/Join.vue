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
      <div class="item">
        <font-awesome-icon class="icn pad" icon="code-branch" />
        <a href="https://github.com/onfe/Picto/releases">{{ version }}</a>
      </div>
      <div class="item">
        <font-awesome-icon class="icn pad" icon="bug" />
        <a
          href="https://github.com/onfe/Picto/issues?q=is%3Aissue+is%3Aopen+label%3Abug"
          >Found a bug?</a
        >
      </div>
      <div class="item">
        <font-awesome-icon class="icn pad" :icon="['fab', 'twitter']" />
        <a href="https://twitter.com/PictoTweets">Twitter</a>
      </div>
      <div class="item">
        Made with <font-awesome-icon class="icn" icon="heart" /> by
        <a href="https://joshuarainbow.co.uk/">Josh</a>,
        <a href="https://onfe.uk/">Eddie</a> &amp;
        <a href="https://freddyheppell.com/">Freddy</a>
      </div>
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
  },
  computed: {
    version() {
      return process.env.VUE_APP_VERSION;
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
  max-width: 675px;

  padding: 0 1.5rem 1rem 3.5rem;

  @media (min-width: 992px) {
    padding-left: 8rem;
  }

  font-family: monospace;
  font-size: 1.2rem;
  color: $grey-d;
}

footer {
  color: $grey;
  font-size: 0.75rem;
  line-height: 1.2;
  display: flex;
  flex-wrap: wrap;
  .item {
    margin-right: 1rem;
    margin-top: 0.75rem;
    .icn.pad {
      margin-right: 1ch;
    }

    a {
      color: $grey;
      transition: color 200ms ease-in-out;
    }

    a:hover {
      color: $grey-d;
    }
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

a {
  color: $grey-d;
  text-decoration: underline;
  text-decoration-style: dotted;
}

.logo {
  max-width: 100%;
  width: auto;
  margin-bottom: 3rem;

  max-height: 10rem;
}
</style>
