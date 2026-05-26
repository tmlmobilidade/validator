/* eslint-disable @typescript-eslint/naming-convention */

import { GtfsValidationSummary } from '@tmlmobilidade/types';
import { access, constants, readFile } from 'fs/promises';
import { dirname, resolve } from 'path';
import { fileURLToPath } from 'url';

import { GoBinaryError, runGoBinary, type RunGoBinaryOptions } from './src/utils.js';

const BINARY_DISTRIBUTIONS = {
	'darwin-arm64': 'validator-darwin-arm64',
	'darwin-x64': 'validator-darwin-amd64',
	'linux-arm64': 'validator-linux-arm64',
	'linux-x64': 'validator-linux-amd64',
	'win32-x64': 'validator.exe',
} as const;

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

type SupportedPlatform = keyof typeof BINARY_DISTRIBUTIONS;

/**
 * Supported language codes for validation messages.
 */
export type SupportedLanguage = 'en' | 'pt';

/**
 * Default timeout for validation operations (30 minutes).
 */
const DEFAULT_TIMEOUT_MS = 30 * 60 * 1000;

export interface GTFSValidatorOptions {
	/** Working directory for the validation process */
	cwd?: string
	/** Additional environment variables */
	env?: Record<string, string>
	/** Language for validation messages (e.g., 'en', 'pt') */
	lang?: SupportedLanguage
	/** Log level for validation messages */
	log_level?: 'debug' | 'error' | 'info'
	/** Output file path for detailed validation results */
	out_file?: string
	/** Path to custom validation rules file */
	rules_path?: string
	/** Timeout in milliseconds (default: 30 minutes) */
	timeout?: number
}

export interface GtfsValidationResult {
	/** Arguments passed to the validator */
	args: string[]
	/** Execution time in milliseconds */
	executionTime: number
	/** Raw stderr from the validator */
	stderr: string
	/** Raw stdout from the validator */
	stdout: string
	/** Parsed validation summary */
	summary: GtfsValidationSummary
}

export class GtfsValidationError extends Error {
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
 *
 * @returns The platform key matching the current system
 * @throws {GtfsValidationError} If the platform is not supported
 *
 * @internal
 */
function getCurrentPlatform(): SupportedPlatform {
	const platform = process.platform;
	const arch = process.arch;
	const platformKey = `${platform}-${arch}` as SupportedPlatform;

	if (!(platformKey in BINARY_DISTRIBUTIONS)) {
		const supportedPlatforms = Object.keys(BINARY_DISTRIBUTIONS).join(', ');
		throw new GtfsValidationError(
			`Unsupported platform: ${platformKey}. Supported platforms: ${supportedPlatforms}`,
			'UNSUPPORTED_PLATFORM',
		);
	}

	return platformKey;
}

/**
 * Gets the path to the validator binary for the current platform.
 *
 * @returns The absolute path to the validator binary
 * @throws {GtfsValidationError} If the binary is not found or not executable
 *
 * @internal
 */
async function getValidatorBinaryPath(): Promise<string> {
	const platform = getCurrentPlatform();
	const binaryName = BINARY_DISTRIBUTIONS[platform];
	const binaryPath = resolve(__dirname, 'bin', binaryName);

	try {
		await access(binaryPath, constants.F_OK | constants.X_OK);
		return binaryPath;
	} catch (err) {
		const error = err instanceof Error ? err : new Error(String(err));
		throw new GtfsValidationError(
			`GTFS validator binary not found or not executable: ${binaryPath}. Please ensure the binary is installed for platform ${platform}`,
			'BINARY_NOT_FOUND',
			error,
		);
	}
}

/**
 * Validates input parameters before running the validator.
 *
 * @param input - The input path to validate
 * @throws {GtfsValidationError} If the input is invalid or not accessible
 *
 * @internal
 */
async function validateInput(input: string): Promise<void> {
	if (typeof input !== 'string' || input.trim().length === 0) {
		throw new GtfsValidationError(
			'Input path is required and must be a non-empty string',
			'INVALID_INPUT',
		);
	}

	try {
		const inputPath = resolve(input);
		await access(inputPath, constants.F_OK | constants.R_OK);
	} catch (err) {
		const error = err instanceof Error ? err : new Error(String(err));
		throw new GtfsValidationError(
			`Input path does not exist or is not readable: ${input}`,
			'INPUT_NOT_ACCESSIBLE',
			error,
		);
	}
}

/**
 * Validates options object and normalizes values.
 *
 * @param options - The options to validate
 * @returns Normalized options
 * @throws {GtfsValidationError} If options are invalid
 *
 * @internal
 */
function validateOptions(options: GTFSValidatorOptions = {}): GTFSValidatorOptions {
	const { cwd, env, lang, out_file, rules_path, timeout } = options;

	if (timeout !== undefined && (typeof timeout !== 'number' || timeout <= 0 || !Number.isFinite(timeout))) {
		throw new GtfsValidationError(
			'Timeout must be a positive finite number',
			'INVALID_OPTIONS',
		);
	}

	if (lang !== undefined && typeof lang !== 'string') {
		throw new GtfsValidationError(
			'Language must be a string',
			'INVALID_OPTIONS',
		);
	}

	if (out_file !== undefined && (typeof out_file !== 'string' || out_file.trim().length === 0)) {
		throw new GtfsValidationError(
			'Output file path must be a non-empty string',
			'INVALID_OPTIONS',
		);
	}

	if (rules_path !== undefined && (typeof rules_path !== 'string' || rules_path.trim().length === 0)) {
		throw new GtfsValidationError(
			'Rules path must be a non-empty string',
			'INVALID_OPTIONS',
		);
	}

	if (cwd !== undefined && typeof cwd !== 'string') {
		throw new GtfsValidationError(
			'Working directory must be a string',
			'INVALID_OPTIONS',
		);
	}

	if (env !== undefined && (typeof env !== 'object' || env === null || Array.isArray(env))) {
		throw new GtfsValidationError(
			'Environment variables must be an object',
			'INVALID_OPTIONS',
		);
	}

	return options;
}

/**
 * Builds command line arguments for the GTFS validator.
 *
 * @param input - The input path
 * @param options - Validation options
 * @returns Array of command line arguments
 *
 * @internal
 */
function buildValidatorArgs(input: string, options: GTFSValidatorOptions = {}): string[] {
	const { lang, log_level, out_file, rules_path } = options;
	const args: string[] = ['-input', input];

	if (out_file) {
		args.push('-out', out_file);
	}

	if (rules_path) {
		args.push('-rules', rules_path);
	}

	if (lang) {
		args.push('-lang', lang);
	}

	if (log_level) {
		args.push('-log', log_level);
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
/**
 * Runs the GTFS validator on the specified input.
 *
 * @param input - Path to the GTFS feed (file or directory)
 * @param options - Validation options
 * @returns Promise resolving to validation results
 *
 * @throws {GtfsValidationError} If validation fails or input is invalid
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
 *   if (err instanceof GtfsValidationError) {
 *     console.error(`Validation failed: ${err.message}`);
 *     console.error(`Error code: ${err.code}`);
 *   }
 * }
 * ```
 */
export async function GtfsValidator(
	input: string,
	options: GTFSValidatorOptions = {},
): Promise<GtfsValidationResult> {
	// Validate and normalize options
	const validatedOptions = validateOptions(options);
	const {
		cwd,
		env,
		timeout = DEFAULT_TIMEOUT_MS,
		...validatorOptions
	} = validatedOptions;

	try {
		// Validate input
		await validateInput(input);

		// Get binary path
		const binaryPath = await getValidatorBinaryPath();

		// Build arguments
		const args = buildValidatorArgs(input, validatorOptions);

		// Determine output file path (resolve relative to cwd if provided)
		const outputFilePath = validatorOptions.out_file
			? resolve(cwd || process.cwd(), validatorOptions.out_file)
			: undefined;

		// Run validator
		const runOptions: RunGoBinaryOptions = {
			args,
			cwd,
			env,
			maxStderrSize: 5 * 1024 * 1024, // 5MB for error messages
			maxStdoutSize: 50 * 1024 * 1024, // 50MB for large validation reports
			timeout,
		};

		const startTime = Date.now();
		let result: Awaited<ReturnType<typeof runGoBinary<GtfsValidationSummary>>>;
		let summary: GtfsValidationSummary;

		try {
			result = await runGoBinary<GtfsValidationSummary>(binaryPath, runOptions);
			summary = result.data;
		} catch (err) {
			// If output file was specified and we got a JSON parse error or no output,
			// try reading from the file instead (the validator writes to file, not stdout)
			if (
				outputFilePath
				&& err instanceof GoBinaryError
				&& (err.code === 'NO_OUTPUT' || err.code === 'JSON_PARSE_ERROR' || err.code === 'NO_VALID_LINES')
			) {
				try {
					// Wait a bit for the file to be written (in case of race condition)
					await new Promise(resolve => setTimeout(resolve, 100));
					const fileContent = await readFile(outputFilePath, 'utf-8');
					summary = JSON.parse(fileContent.trim()) as GtfsValidationSummary;
					// Calculate execution time from when we started
					const executionTime = Date.now() - startTime;
					// Create a result object with the file data, preserving error info for stderr/stdout
					result = {
						data: summary,
						executionTime,
						exitCode: err.exitCode ?? 0,
						stderr: err.stderr || '',
						stdout: err.stdout || '',
					};
				} catch (fileErr) {
					const error = fileErr instanceof Error ? fileErr : new Error(String(fileErr));
					throw new GtfsValidationError(
						`Failed to read or parse output file: ${outputFilePath}. ${error.message}`,
						'OUTPUT_FILE_READ_ERROR',
						error,
						err.stdout,
						err.stderr,
					);
				}
			} else {
				// Re-throw the original error
				throw err;
			}
		}

		return {
			args,
			executionTime: result.executionTime,
			stderr: result.stderr,
			stdout: result.stdout,
			summary,
		};
	} catch (err) {
		// Re-throw GTFSValidatorError as-is
		if (err instanceof GtfsValidationError) {
			throw err;
		}

		// Convert GoBinaryError to GTFSValidatorError with context
		if (err instanceof GoBinaryError) {
			let errorMessage = `GTFS validation failed: ${err.message}`;
			let errorCode = 'VALIDATION_FAILED';

			// Provide more specific error messages based on the binary error
			switch (err.code) {
				case 'JSON_PARSE_ERROR':
					errorMessage = 'Failed to parse validation results. The validator may have crashed or produced invalid output.';
					errorCode = 'PARSE_ERROR';
					break;
				case 'NON_ZERO_EXIT':
					errorMessage = `GTFS validator exited with error code ${err.exitCode ?? 'unknown'}${err.stderr ? `: ${err.stderr}` : ''}`;
					errorCode = 'VALIDATOR_ERROR';
					break;
				case 'STDERR_TOO_LARGE':
					errorMessage = 'Validation error output exceeded maximum size.';
					errorCode = 'ERROR_OUTPUT_TOO_LARGE';
					break;
				case 'STDOUT_TOO_LARGE':
					errorMessage = 'Validation output exceeded maximum size. The GTFS feed may be too large.';
					errorCode = 'OUTPUT_TOO_LARGE';
					break;
				case 'TIMEOUT':
					errorMessage = `GTFS validation timed out after ${timeout}ms. Consider increasing the timeout for large feeds.`;
					errorCode = 'VALIDATION_TIMEOUT';
					break;
			}

			throw new GtfsValidationError(
				errorMessage,
				errorCode,
				err,
				err.stdout,
				err.stderr,
			);
		}

		// Handle unexpected errors
		const errorMessage = err instanceof Error ? err.message : String(err);
		throw new GtfsValidationError(
			`Unexpected error during GTFS validation: ${errorMessage}`,
			'UNEXPECTED_ERROR',
			err instanceof Error ? err : new Error(String(err)),
		);
	}
}

/**
 * Gets information about the available validator binary for the current platform.
 *
 * @returns Information about the validator binary including availability status
 *
 * @example
 * ```ts
 * const info = await getValidatorInfo();
 * if (info.isAvailable) {
 *   console.log(`Binary found at: ${info.binaryPath}`);
 * } else {
 *   console.log(`Binary not found for platform: ${info.platform}`);
 * }
 * ```
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
	} catch {
		// Binary not available - this is expected in some scenarios
	}

	return {
		binaryName,
		binaryPath,
		isAvailable,
		platform,
	};
}
