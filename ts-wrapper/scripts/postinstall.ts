import { existsSync, copyFileSync, mkdirSync } from "fs";
import { dirname, join } from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

// Lookup table for all platforms and binary distribution files
const BINARY_DISTRIBUTIONS_FILES = {
  'linux-x64': 'validator-linux-amd64',
  'linux-arm64': 'validator-linux-arm64',
  'darwin-x64': 'validator-darwin-amd64',
  'darwin-arm64': 'validator-darwin-arm64',
  'windows-x64': 'validator.exe',
}

const DEV_BIN_PATH = join(__dirname,'..', '..', 'bin');
const REMOTE_BIN_PATH = "SOME REMOTE PATH";
const LOCAL_BIN_PATH = join(__dirname, '..', 'bin');

//function to get the current platform
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

  //check if local bin path exists
  if (!existsSync(LOCAL_BIN_PATH)) {
    mkdirSync(LOCAL_BIN_PATH);
  }

  copyFileSync(join(DEV_BIN_PATH, binaryDistributionFile), join(LOCAL_BIN_PATH, binaryDistributionFile));
}

function main() {
  const platform = getCurrentPlatform();
  const binaryDistributionFile = BINARY_DISTRIBUTIONS_FILES[platform];

  if (!binaryDistributionFile) {
    console.error(`No binary distribution file found for platform: ${platform}`);
    return;
  }

  const binaryDistributionFilePath = join(DEV_BIN_PATH, binaryDistributionFile);
  console.log(binaryDistributionFilePath);

  //check if the file exists
  if (existsSync(binaryDistributionFilePath)) {
    try {
      buildDevEnvironment();
    } catch (error) {
      console.error(`Error building dev environment: ${error}`);
      return;
    }
  } else {
    console.info(`Local file not found: ${binaryDistributionFilePath}`);
    console.info(`Downloading file from remote server...`);
    return;
  }
}


console.log('\n\n----- RUNNING POSTINSTALL SCRIPT -----\n\n');
main();
console.log('\n\n----- END POSTINSTALL SCRIPT -----\n\n');
