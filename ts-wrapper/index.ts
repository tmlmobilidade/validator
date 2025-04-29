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

export async function GTFSValidator(input: string) {
	try {
		const result = await runGoBinary<GTFSValidatorSummary>(`${__dirname}/bin/${BINARY_DISTRIBUTIONS_FILES[getCurrentPlatform()]}`, [
			'-input',
			input,
		]);
		return result;
	}
	catch (err) {
		console.error('❌ Error:', (err as Error).message);
		throw err;
	}
}
