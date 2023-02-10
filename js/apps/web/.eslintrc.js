module.exports = {
  root: true,
  extends: ["eslint-config-custom-react"],
  ignorePatterns: ["next.config.js"],
  parserOptions: {
    project: './js/apps/web/tsconfig.json',
  },
};
