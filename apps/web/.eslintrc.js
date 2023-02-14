module.exports = {
  root: true,
  extends: ['eslint-config-custom-react'],
  ignorePatterns: ['next.config.js'],
  parserOptions: {
    project: 'tsconfig.json',
    tsconfigRootDir: __dirname,
    sourceType: 'module',
  },
};
