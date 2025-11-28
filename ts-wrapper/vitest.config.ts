import { defineConfig } from 'vitest/config';

export default defineConfig({
	test: {
		coverage: {
			exclude: [
				'node_modules/',
				'dist/',
				'**/*.test.ts',
				'**/*.spec.ts',
				'**/scripts/**',
			],
			provider: 'v8',
			reporter: ['text', 'json', 'html'],
		},
		environment: 'node',
		exclude: ['node_modules', 'dist', '.idea', '.git', '.cache'],
		globals: true,
		include: ['**/*.{test,spec}.{js,mjs,cjs,ts,mts,cts,jsx,tsx}'],
		testTimeout: 60000, // 60 seconds for integration tests
	},
});
