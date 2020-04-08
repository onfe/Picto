<template>
  <div class="moderate">
    <div class="container">
      <h1>Moderation dashboard</h1>
      <AuthForm v-if="token === null" @authenticated="setToken" />
      <div v-else>Token: {{ token }}</div>
    </div>
  </div>
</template>

<script>
import AuthForm from "@/components/AuthForm.vue";
export default {
  name: "moderate",
  components: {
    AuthForm
  },
  data() {
    return {
      token: null
    };
  },
  methods: {
    setToken(token) {
      this.token = token;
    }
  }
};
</script>

<style lang="scss" scoped>
.moderate {
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
