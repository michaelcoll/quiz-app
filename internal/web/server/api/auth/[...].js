import GithubProvider from "next-auth/providers/github";

import { NuxtAuthHandler } from "#auth";

export default NuxtAuthHandler({
  secret: "my-secret",
  providers: [
    // @ts-expect-error You need to use .default here for it to work during SSR. May be fixed via Vite at some point
    GithubProvider.default({
      clientId: useRuntimeConfig().clientId,
      clientSecret: useRuntimeConfig().clientSecret,
    }),
  ],
  callbacks: {
    jwt({ token, account }) {
      if (account) {
        token = Object.assign({}, token, { access_token: account.access_token });
      }
      return token;
    },
    session({ session, token }) {
      if (session) {
        session = Object.assign({}, session, {
          access_token: token.access_token,
          sub: token.sub,
          exp: token.exp,
        });
      }
      return session;
    },
    async signIn({ account }) {
      await $fetch(`${useRuntimeConfig().apiBase}/api/v1/login`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${account.access_token}`,
        },
      });
      return true;
    },
  },
});
