import { runGoBinary } from "./utils.js";

interface GoOutput {
  message: string;
  success: boolean;
}

async function main() {
  try {
    const result = await runGoBinary<GoOutput>("./bin/validator", [
      "-input",
      "../data/Bom.zip",
      "-output",
      "./results/output.json",
    ]);
    console.log("✅ Output from Go:", result);
  } catch (err) {
    console.error("❌ Error:", (err as Error).message);
  }
}

main();
