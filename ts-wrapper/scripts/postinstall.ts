import { chmodSync, copyFileSync, createWriteStream, existsSync, mkdirSync } from 'fs';
import path, { dirname, join } from 'path';
import { Readable } from 'stream';
import { finished } from 'stream/promises';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// Lookup table for all platforms and binary distribution files
const BINARY_DISTRIBUTIONS_FILES = {
	'darwin-arm64': 'validator-darwin-arm64',
	'darwin-x64': 'validator-darwin-amd64',
	'linux-arm64': 'validator-linux-arm64',
	'linux-x64': 'validator-linux-amd64',
	'windows-x64': 'validator.exe',
};

const DEV_BIN_PATH = join(__dirname, '..', '..', 'bin');
const REMOTE_BIN_PATH = 'https://github.com/tmlmobilidade/validator/raw/refs/heads/production/bin/';
const LOCAL_BIN_PATH = join(__dirname, '..', 'bin');

// function to get the current platform
function getCurrentPlatform() {
	const platform = process.platform;
	const arch = process.arch;
	return `${platform}-${arch}`;
}

function buildDevEnvironment() {
	const platform = getCurrentPlatform();
	const binaryDistributionFile = BINARY_DISTRIBUTIONS_FILES[platform];

	if (!binaryDistributionFile) {
		console.error(`No binary distribution file found for platform: ${platform}`);
		return;
	}

	// check if local bin path exists
	if (!existsSync(LOCAL_BIN_PATH)) {
		mkdirSync(LOCAL_BIN_PATH);
	}

	copyFileSync(join(DEV_BIN_PATH, binaryDistributionFile), join(LOCAL_BIN_PATH, binaryDistributionFile));
}

async function downloadRemoteBinaries() {
	const platform = getCurrentPlatform();
	const binaryDistributionFile = BINARY_DISTRIBUTIONS_FILES[platform];

	if (!binaryDistributionFile) {
		console.error(`No binary distribution file found for platform: ${platform}`);
		return;
	}

	// Download the file
	const res = await fetch(REMOTE_BIN_PATH + binaryDistributionFile);

	if (!res.ok) {
		throw new Error(`Error downloading remote binary: ${res.statusText}`);
	}

	// Create the local bin path if it doesn't exist
	if (!existsSync(LOCAL_BIN_PATH)) mkdirSync(LOCAL_BIN_PATH);

	// Download the file
	const arrayBuffer = await res.arrayBuffer();
	const buffer = Buffer.from(arrayBuffer);

	// Create the file stream
	const fileStream = createWriteStream(path.resolve(LOCAL_BIN_PATH, binaryDistributionFile));

	// Write the file to the local bin path
	return await finished(Readable.from(buffer).pipe(fileStream));
}

async function main() {
	const platform = getCurrentPlatform();
	const binaryDistributionFile = BINARY_DISTRIBUTIONS_FILES[platform];

	if (!binaryDistributionFile) {
		console.error(`No binary distribution file found for platform: ${platform}`);
		return;
	}

	const binaryDistributionFilePath = join(DEV_BIN_PATH, binaryDistributionFile);

	// check if the file exists
	if (existsSync(binaryDistributionFilePath)) {
		try {
			buildDevEnvironment();
		}
		catch (error) {
			console.error(`Error building dev environment: ${error}`);
		}
	}
	else {
		console.info(`Local file not found: ${binaryDistributionFilePath}`);
		console.info(`Downloading file from remote server...`);
		// Download the remote binaries
		await downloadRemoteBinaries();
	}

	// CHMOD the file executable
	// chmodSync(join(LOCAL_BIN_PATH, binaryDistributionFile), 0o755);
}

main();
