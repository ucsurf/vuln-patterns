export const auth = betterAuth({
  // ruleid: sso-trust-email-verified
  plugins: [sso({ trustEmailVerified: true, provisionUser })],
});
// ruleid: sso-trust-email-verified
const opts = { trustEmailVerified: true };
// ok: sso-trust-email-verified
const safe = sso({ trustEmailVerified: false });
