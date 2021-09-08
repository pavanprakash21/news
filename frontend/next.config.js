/** @type {import('next').NextConfig} */

const ghPages = process.env.DEPLOY_TARGET === 'gh-pages';

module.exports = {
  reactStrictMode: true,
  assetPrefix: ghPages ? '/news/' : '',
  basePrefix: ghPages ? '/news' : ''
}
