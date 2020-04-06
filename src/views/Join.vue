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
      <UpdateManager
        link="https://github.com/onfe/Picto/releases"
        :version="version"
        :status="$store.state.client.swStatus"
        class="item"
        @update="appUpdate"
      />
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
import UpdateManager from "@/components/UpdateManager.vue";

export default {
  name: "join",
  components: {
    JoinForm,
    PublicRooms,
    UpdateManager
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
  },
  methods: {
    appUpdate() {
      document.dispatchEvent(new CustomEvent("sw-perform-update"));
    }
  }
};
</script>

<style lang="scss" scoped>
.join {
  min-height: 100%;
  background-color: var(--background-join);
  color: var(--primary-join);

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
  color: var(--primary-join);
}

footer {
  color: var(--secondary-join);
  font-size: 0.75rem;
  line-height: 1.2;
  display: flex;
  flex-wrap: wrap;
  align-items: center;

  .item {
    margin-right: 1rem;
    margin-top: 0.75rem;
    .icn.pad {
      margin-right: 1ch;
    }

    a {
      color: var(--secondary-join);
      transition: color 200ms ease-in-out;
    }

    a:hover {
      color: var(--primary-join);
    }
  }
}

hr {
  border: 0;
  border-bottom: 1px dashed var(--secondary-join);
}

p {
  margin-bottom: 1.5rem;
  line-height: 1.2;
}

a {
  color: var(--primary-join);
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
