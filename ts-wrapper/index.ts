import { runGoBinary } from './src/utils.js';

export interface GTFSValidatorMessage {
	field: string
	fileName: string
	message: string
	row: number
	severity: 'error' | 'info' | 'warning'
	validation_id: string
}

export interface GTFSValidatorSummary {
	messages: GTFSValidatorMessage[]
	total_errors: number
	total_infos: number
	total_warnings: number
}

const BINARY_DISTRIBUTIONS_FILES = {
	'darwin-arm64': 'validator-darwin-arm64',
	'darwin-x64': 'validator-darwin-amd64',
	'linux-arm64': 'validator-linux-arm64',
	'linux-x64': 'validator-linux-amd64',
	'windows-x64': 'validator.exe',
};

function getCurrentPlatform() {
	const platform = process.platform;
	const arch = process.arch;
	return `${platform}-${arch}`;
}

export async function GTFSValidator(input: string) {
	try {
		const result = await runGoBinary<GTFSValidatorSummary>(`./bin/${BINARY_DISTRIBUTIONS_FILES[getCurrentPlatform()]}`, [
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
