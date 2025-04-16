import { runGoBinary } from "./utils.js";

export type GTFSValidatorMessage = {
    row: number
    field: string
    fileName: string
    message: string
    validation_id: string
    severity: "error" | "info" | "warning"
}

export type GTFSValidatorSummary = {
    messages: GTFSValidatorMessage[]
    total_errors: number
    total_infos: number
    total_warnings: number
}

const BINARY_DISTRIBUTIONS_FILES = {
  'linux-x64': 'validator-linux-amd64',
  'linux-arm64': 'validator-linux-arm64',
  'darwin-x64': 'validator-darwin-amd64',
  'darwin-arm64': 'validator-darwin-arm64',
  'windows-x64': 'validator.exe',
}

function getCurrentPlatform() {
  const platform = process.platform;
  const arch = process.arch;
  return `${platform}-${arch}`;
}

export async function GTFSValidator(input: string) {
  try {
    const result = await runGoBinary<GTFSValidatorSummary>(`./bin/${BINARY_DISTRIBUTIONS_FILES[getCurrentPlatform()]}`, [
      "-input",
      input
    ]);
    return result;
  } catch (err) {
    console.error("❌ Error:", (err as Error).message);
    throw err;
  }
}
