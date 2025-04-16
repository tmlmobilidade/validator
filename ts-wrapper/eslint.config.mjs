/* * */

import { node } from '@carrismetropolitana/eslint'

/* * */

export default [
  ...node,
  {
    ignorePatterns: ['results/**'],
    rules: {
      '@typescript-eslint/no-extraneous-class': 'off',
    },
  },
]
