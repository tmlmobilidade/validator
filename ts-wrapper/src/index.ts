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

export async function GTFSValidator(input: string) {
  try {
    const result = await runGoBinary<GTFSValidatorSummary>("./bin/validator", [
      "-input",
      input
    ]);
    return result;
  } catch (err) {
    console.error("❌ Error:", (err as Error).message);
    throw err;
  }
}
