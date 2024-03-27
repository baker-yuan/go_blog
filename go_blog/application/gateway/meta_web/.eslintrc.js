module.exports = {
  extends: [
    require.resolve('@umijs/fabric/dist/eslint'),
    'plugin:import/errors',
    'plugin:import/warnings',
  ],
  plugins: ['simple-import-sort'],
  globals: {
    ANT_DESIGN_PRO_ONLY_DO_NOT_USE_IN_YOUR_PRODUCTION: true,
    page: true,
    REACT_APP_ENV: true,
  },
  rules: {
    '@typescript-eslint/naming-convention': 'off',
    'import/no-unresolved': [2, { ignore: ['^@/', '^umi/', '^@@/'] }],
    'sort-imports': 'off',
    'import/order': 'off',
    'simple-import-sort/imports': 'error',
    'simple-import-sort/exports': 'error',
    'no-underscore-dangle': 'off',
  },
  settings: {
    'import/resolver': {
      node: {
        path: ['src'],
        extensions: ['.js', '.ts', '.tsx', '.jsx'],
      },
    },
  },
  overrides: [
    {
      files: 'server/**/*.js',
      env: { node: true },
      rules: {
        'import/order': ['error', { 'newlines-between': 'always' }],
      },
    },
  ],
};
