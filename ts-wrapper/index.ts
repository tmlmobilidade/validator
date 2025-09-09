import { GTFSValidatorSummary } from '@tmlmobilidade/types';
import { access, constants } from 'fs/promises';
import { dirname, resolve } from 'path';
import { fileURLToPath } from 'url';

import { GoBinaryError, runGoBinary, type RunGoBinaryOptions } from './src/utils.js';

const BINARY_DISTRIBUTIONS: Record<string, string> = {
	'darwin-arm64': 'validator-darwin-arm64',
	'darwin-x64': 'validator-darwin-amd64',
	'linux-arm64': 'validator-linux-arm64',
	'linux-x64': 'validator-linux-amd64',
	'win32-x64': 'validator.exe',
} as const;

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

type SupportedPlatform = keyof typeof BINARY_DISTRIBUTIONS;

export interface GTFSValidatorOptions {
	/** Working directory for the validation process */
	cwd?: string
	/** Additional environment variables */
	env?: Record<string, string>
	/** Language for validation messages (e.g., 'en', 'pt') */
	lang?: string
	/** Output file path for detailed validation results */
	out_file?: string
	/** Path to custom validation rules file */
	rules_path?: string
	/** Timeout in milliseconds (default: 10 minutes) */
	timeout?: number
}

export interface GTFSValidatorResult {
	/** Arguments passed to the validator */
	args: string[]
	/** Execution time in milliseconds */
	executionTime: number
	/** Raw stderr from the validator */
	stderr: string
	/** Raw stdout from the validator */
	stdout: string
	/** Parsed validation summary */
	summary: GTFSValidatorSummary
}

export class GTFSValidatorError extends Error {
	constructor(
		message: string,
		public readonly code: string,
		public readonly originalError?: Error,
		public readonly stdout?: string,
		public readonly stderr?: string,
	) {
		super(message);
		this.name = 'GTFSValidatorError';
	}
}

/**
 * Gets the current platform identifier in the format expected by the binary distributions.
 */
function getCurrentPlatform(): SupportedPlatform {
	const platform = process.platform;
	const arch = process.arch;
	const platformKey = `${platform}-${arch}` as SupportedPlatform;

	if (!BINARY_DISTRIBUTIONS[platformKey]) {
		throw new GTFSValidatorError(
			`Unsupported platform: ${platformKey}. Supported platforms: ${Object.keys(BINARY_DISTRIBUTIONS).join(', ')}`,
			'UNSUPPORTED_PLATFORM',
		);
	}

	return platformKey;
}

/**
 * Gets the path to the validator binary for the current platform.
 */
async function getValidatorBinaryPath(): Promise<string> {
	const platform = getCurrentPlatform();
	const binaryName = BINARY_DISTRIBUTIONS[platform];
	const binaryPath = resolve(__dirname, 'bin', binaryName);

	try {
		await access(binaryPath, constants.F_OK | constants.X_OK);
		return binaryPath;
	}
	catch (err) {
		throw new GTFSValidatorError(
			`GTFS validator binary not found or not executable: ${binaryPath}. Please ensure the binary is installed for platform ${platform}`,
			'BINARY_NOT_FOUND',
			err instanceof Error ? err : new Error(String(err)),
		);
	}
}

/**
 * Validates input parameters before running the validator.
 */
async function validateInput(input: string): Promise<void> {
	if (!input || typeof input !== 'string') {
		throw new GTFSValidatorError(
			'Input path is required and must be a non-empty string',
			'INVALID_INPUT',
		);
	}

	try {
		const inputPath = resolve(input);
		await access(inputPath, constants.F_OK | constants.R_OK);
	}
	catch (err) {
		throw new GTFSValidatorError(
			`Input path does not exist or is not readable: ${input}`,
			'INPUT_NOT_ACCESSIBLE',
			err instanceof Error ? err : new Error(String(err)),
		);
	}
}

/**
 * Builds command line arguments for the GTFS validator.
 */
function buildValidatorArgs(input: string, options: GTFSValidatorOptions = {}): string[] {
	const { lang, out_file, rules_path } = options;
	const args: string[] = ['-input', input];

	if (out_file) {
		args.push('-o', out_file);
	}

	if (rules_path) {
		args.push('-rules', rules_path);
	}

	if (lang) {
		args.push('-lang', lang);
	}

	return args;
}

/**
 * Runs the GTFS validator on the specified input.
 *
 * @param input Path to the GTFS feed (file or directory)
 * @param options Validation options
 * @returns Promise resolving to validation results
 *
 * @example
 * ```ts
 * try {
 *   const result = await GTFSValidator('./gtfs-feed.zip', {
 *     lang: 'en',
 *     timeout: 300000, // 5 minutes
 *     out_file: './validation-report.json'
 *   });
 *
 *   console.log(`Validation completed in ${result.executionTime}ms`);
 *   console.log(`Found ${result.summary.errorCount} errors`);
 * } catch (err) {
 *   if (err instanceof GTFSValidatorError) {
 *     console.error(`Validation failed: ${err.message}`);
 *   }
 * }
 * ```
 */
export async function GTFSValidator(
	input: string,
	options: GTFSValidatorOptions = {},
): Promise<GTFSValidatorResult> {
	const {
		cwd,
		env,
		timeout = 3000 * 60 * 10, // 30 minutes default
		...validatorOptions
	} = options;

	try {
		// Validate input
		await validateInput(input);

		// Get binary path
		const binaryPath = await getValidatorBinaryPath();

		// Build arguments
		const args = buildValidatorArgs(input, validatorOptions);

		// Run validator
		const runOptions: RunGoBinaryOptions = {
			args,
			cwd,
			env,
			maxStderrSize: 5 * 1024 * 1024, // 5MB for error messages
			maxStdoutSize: 50 * 1024 * 1024, // 50MB for large validation reports
			timeout,
		};

		const result = await runGoBinary<GTFSValidatorSummary>(binaryPath, runOptions);

		return {
			args,
			executionTime: result.executionTime,
			stderr: result.stderr,
			stdout: result.stdout,
			summary: result.data,
		};
	}
	catch (err) {
		if (err instanceof GTFSValidatorError) {
			throw err;
		}

		if (err instanceof GoBinaryError) {
			let errorMessage = `GTFS validation failed: ${err.message}`;
			let errorCode = 'VALIDATION_FAILED';

			// Provide more specific error messages based on the binary error
			switch (err.code) {
				case 'JSON_PARSE_ERROR':
					errorMessage = `Failed to parse validation results. The validator may have crashed or produced invalid output.`;
					errorCode = 'PARSE_ERROR';
					break;
				case 'NON_ZERO_EXIT':
					errorMessage = `GTFS validator exited with error code ${err.exitCode}${err.stderr ? `: ${err.stderr}` : ''}`;
					errorCode = 'VALIDATOR_ERROR';
					break;
				case 'TIMEOUT':
					errorMessage = `GTFS validation timed out after ${timeout}ms. Consider increasing the timeout for large feeds.`;
					errorCode = 'VALIDATION_TIMEOUT';
					break;
			}

			throw new GTFSValidatorError(
				errorMessage,
				errorCode,
				err,
				err.stdout,
				err.stderr,
			);
		}

		// Handle unexpected errors
		throw new GTFSValidatorError(
			`Unexpected error during GTFS validation: ${err instanceof Error ? err.message : String(err)}`,
			'UNEXPECTED_ERROR',
			err instanceof Error ? err : new Error(String(err)),
		);
	}
}

/**
 * Gets information about the available validator binary for the current platform.
 */
export async function getValidatorInfo(): Promise<{
	binaryName: string
	binaryPath: string
	isAvailable: boolean
	platform: SupportedPlatform
}> {
	const platform = getCurrentPlatform();
	const binaryName = BINARY_DISTRIBUTIONS[platform];
	const binaryPath = resolve(__dirname, 'bin', binaryName);

	let isAvailable = false;
	try {
		await access(binaryPath, constants.F_OK | constants.X_OK);
		isAvailable = true;
	}
	catch {
		// Binary not available
	}

	return {
		binaryName,
		binaryPath,
		isAvailable,
		platform,
	};
}
