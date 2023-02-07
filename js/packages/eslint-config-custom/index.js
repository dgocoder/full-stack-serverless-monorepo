module.exports = {
  // ignore lint errors in lint files
  ignorePatterns: [
    '**/eslint-config-custom/index.js',
    '**/eslint-config-custom-react/index.js',
    '.eslintrc.js',
  ],
  extends: [
    'next',
    'turbo',
    'plugin:import/recommended',
    'plugin:import/typescript',
    'airbnb',
    'airbnb-typescript',
    'plugin:typescript-sort-keys/recommended',
    'plugin:unicorn/recommended',
    'prettier',
    'plugin:react/jsx-runtime',
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 2022,
  },
  env: {
    es2021: true,
  },
  plugins: ['promise', '@typescript-eslint', 'prefer-arrow', 'typescript-sort-keys', 'unicorn'],
  overrides: [
    {
      files: ['*.test.ts', '*.fixtures.ts'],
      rules: {
        'no-magic-numbers': 'off',
        'max-len': 'off',
      },
    },
    {
      files: ['*.d.ts'],
      rules: {
        'no-magic-numbers': 'off',
        'max-lines': 'off',
      },
    },
  ],
  rules: {
    /* Cognitive complexity linting */
    // Complexity value may be adjusted as needed
    complexity: ['error', 12],
    // Max-lines default is 300 lines, can be overwritten by setting "max" property
    'max-lines': ['error', { skipBlankLines: true, skipComments: true, max: 350 }],
    // Max-params defaults to 3, otherwise pass object
    'max-params': 'error',

    '@typescript-eslint/no-non-null-assertion': 'error',
    'no-param-reassign': ['error', { ignorePropertyModificationsForRegex: ['^draft'] }],
    '@typescript-eslint/no-explicit-any': 'error',
    '@typescript-eslint/no-unnecessary-condition': 'error',
    'no-underscore-dangle': ['error', { allow: ['__typename'] }],
    '@typescript-eslint/array-type': ['error', { default: 'array' }],
    '@typescript-eslint/consistent-type-imports': ['warn', { prefer: 'type-imports' }],
    'no-duplicate-imports': 'error',
    'require-await': 'error',
    'no-continue': 'off',
    'no-plusplus': ['error', { allowForLoopAfterthoughts: true }],
    'prefer-arrow/prefer-arrow-functions': 'error',
    'import/extensions': 'off',
    'react/prop-types': 'off',
    'no-magic-numbers': [
      'error',
      { ignoreArrayIndexes: true, enforceConst: true, ignore: [0, 1, -1] },
    ],
    'object-shorthand': ['error', 'properties'],
    'array-callback-return': 'error',
    'arrow-body-style': 'off',
    'no-shadow': [
      'warn',
      {
        builtinGlobals: false,
        hoist: 'all',
      },
    ],
    // handled by "no-shadow"
    '@typescript-eslint/no-shadow': 'off',
    '@typescript-eslint/no-unused-vars': ['error', { argsIgnorePattern: '^_' }],
    '@typescript-eslint/no-floating-promises': 'error',
    'promise/prefer-await-to-then': 'error',
    'no-void': ['error', { allowAsStatement: true }],
    '@typescript-eslint/naming-convention': [
      'error',
      {
        selector: 'variable',
        filter: '__typename',
        format: null,
      },
      {
        selector: 'variable',
        types: ['function'],
        format: ['camelCase', 'PascalCase'],
        leadingUnderscore: 'allow',
      },
      {
        selector: 'variable',
        types: ['boolean', 'number', 'string', 'array'],
        format: ['camelCase', 'UPPER_CASE'],
        leadingUnderscore: 'allow',
      },
      {
        selector: 'typeLike',
        format: ['PascalCase'],
      },
    ],
    'typescript-sort-keys/interface': [
      'error',
      'asc',
      { caseSensitive: true, natural: true, requiredFirst: true },
    ],
    'dot-notation': 'error',
    'import/prefer-default-export': 'off',
    'import/no-default-export': 'off',
    'no-return-await': 'error',
    'lines-between-class-members': [
      'error',
      'always',
      {
        exceptAfterSingleLine: true,
      },
    ],
    'max-len': [
      'error',
      {
        ignorePattern: '//|^/\\*\\*|^\\*|^import .*',
        ignoreStrings: true,
        ignoreTemplateLiterals: true,
        ignoreRegExpLiterals: true,
        code: 100,
      },
    ],
    /* Restrict imports */
    'no-restricted-imports': [
      'error',
      {
        name: 'lodash',
        message: 'Please use lodash-es instead.',
      },
      {
        name: 'lodash/*',
        message: 'Please use lodash-es instead.',
      },
      {
        name: 'moment',
        message: 'Please use dayjs instead',
      },
    ],
    'prefer-destructuring': ['error', { object: true, array: true }],
    'id-length': ['error', { min: 2, exceptions: ['i', 'e', '_', 'x', 'y'] }],
    'no-implicit-coercion': ['error', { allow: ['!!'], disallowTemplateShorthand: true }],

    // We are using unicorn/recommended, these are the rules we are disabling
    'unicorn/explicit-length-check': 'off',
    'unicorn/no-array-callback-reference': 'off',
    'unicorn/no-array-for-each': 'off',
    'unicorn/prefer-switch': 'off',
    'unicorn/prevent-abbreviations': 'off',
    'unicorn/no-unreadable-array-destructuring': 'off',
    'unicorn/prefer-module': 'off',
    'unicorn/no-thenable': 'off',
    'unicorn/prefer-node-protocol': 'off',
    'unicorn/no-array-reduce': 'off',
    'unicorn/no-null': 'off',
    'unicorn/prefer-object-from-entries': 'off',
    'unicorn/no-useless-undefined': 'off',
    'unicorn/switch-case-braces': 'off',
    'unicorn/prefer-at': 'error',
    'unicorn/filename-case': 'off',
    'react/button-has-type': 'off',
    /* Sorting */
    'sort-imports': ['error', { ignoreDeclarationSort: true }],
    'import/order': [
      'error',
      {
        groups: ['builtin', 'external', 'internal', 'parent', 'sibling', 'index'],
        'newlines-between': 'always',
        alphabetize: {
          order: 'asc',
          caseInsensitive: false,
        },
        pathGroups: [
          {
            pattern: 'react',
            position: 'before',
            group: 'external',
          },
          {
            pattern: '~/**',
            group: 'internal',
          },
        ],
        pathGroupsExcludedImportTypes: [],
      },
    ],
    'import/no-extraneous-dependencies': 'off',
    'no-console': 'warn',
    'no-debugger': 'warn',
    'no-alert': 'warn',
  },
};
