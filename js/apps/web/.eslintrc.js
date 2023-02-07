module.exports = {
  root: true,
  extends: ["eslint-config-custom-react"],
  ignorePatterns: ["next.config.js"],
  parserOptions: {
    project: './apps/web/tsconfig.json',
  },
};
