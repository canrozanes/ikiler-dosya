export const DEV_AUTH0_DOMAIN = "dev-tjs38ggl.eu.auth0.com";
export const DEV_AUTH0_CLIENT_ID = "7a6iPPwRxxHnXDK1gQqpMaN6bW57zrA0";
export const DEV_AUTH0_AUDIENCE = "Ikiler-dosya-api";

export const AUTH0_CALLBACK_URL = "/callback";

export function getAuthOConfig() {
  const auth0callbackUrl = `${window.location.origin}${AUTH0_CALLBACK_URL}`;
  return {
    domain: DEV_AUTH0_DOMAIN,
    clientId: DEV_AUTH0_CLIENT_ID,
    redirectUri: auth0callbackUrl,
    audience: DEV_AUTH0_AUDIENCE,
  };
}
