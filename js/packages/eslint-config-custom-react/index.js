module.exports = {
  // ignore lint errors in lint files
  ignorePatterns: ['**/eslint-config-custom-react/index.js', '.eslintrc.js'],
  extends: ['custom', 'plugin:react/jsx-runtime', 'plugin:react-hooks/recommended'],
  plugins: ['react', 'react-hooks', 'jsx-a11y'],
  env: {
    es2021: true,
  },
  rules: {
    'react/prop-types': 'off',
    'react/function-component-definition': [
      'error',
      { namedComponents: 'arrow-function', unnamedComponents: 'arrow-function' },
    ],
    'react-hooks/rules-of-hooks': 'error',
    'react-hooks/exhaustive-deps': [
      'error',
      {
        enableDangerousAutofixThisMayCauseInfiniteLoops: true,
      },
    ],
  },
};
