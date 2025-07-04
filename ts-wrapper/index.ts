import { GTFSValidatorSummary } from '@tmlmobilidade/types';
import { dirname } from 'path';
import { fileURLToPath } from 'url';

import { runGoBinary } from './src/utils.js';

const BINARY_DISTRIBUTIONS_FILES = {
	'darwin-arm64': 'validator-darwin-arm64',
	'darwin-x64': 'validator-darwin-amd64',
	'linux-arm64': 'validator-linux-arm64',
	'linux-x64': 'validator-linux-amd64',
	'windows-x64': 'validator.exe',
};

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

function getCurrentPlatform() {
	const platform = process.platform;
	const arch = process.arch;
	return `${platform}-${arch}`;
}

interface GTFSValidatorOptions {
	out_file?: string
	rules_path?: string
}

export async function GTFSValidator(input: string, options?: GTFSValidatorOptions) {
	const args = ['-input', input];

	if (options) {
		const { out_file, rules_path } = options;

		if (out_file) {
			args.push('-o', out_file);
		}

		// Prefer rules_path over deprecated rules
		if (rules_path) {
			args.push('-rules', rules_path);
		}
	}

	try {
		const result = await runGoBinary<GTFSValidatorSummary>(`${__dirname}/bin/${BINARY_DISTRIBUTIONS_FILES[getCurrentPlatform()]}`, args);
		return result;
	}
	catch (err) {
		console.error('❌ Error:', (err as Error).message);
		throw err;
	}
}
