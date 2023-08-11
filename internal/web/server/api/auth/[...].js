import GithubProvider from "next-auth/providers/github";

import { NuxtAuthHandler } from "#auth";

export default NuxtAuthHandler({
  secret: "my-secret",
  providers: [
    // @ts-expect-error You need to use .default here for it to work during SSR. May be fixed via Vite at some point
    GithubProvider.default({
      clientId: "afd22679cd8118504e36",
      clientSecret: "e32fe9cdc9e69367d1e5eb87880be8eb5e637190",
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
      const apiServerUrl = useRuntimeConfig().apiBase;
      await $fetch(`${apiServerUrl}/api/v1/login`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${account.access_token}`,
        },
      });
      return true;
    },
  },
});
