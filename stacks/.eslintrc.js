module.exports = {
  root: true,
  extends: ['custom'],
  parserOptions: {
    project: './stacks/tsconfig.json',
  },
  rules: {
    'no-magic-numbers': 'off',
  },
};
