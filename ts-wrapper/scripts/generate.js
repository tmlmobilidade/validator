import { dirname, resolve } from "path";
import { spawn } from "child_process";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const validatorPath = resolve(__dirname, "../../validator/main.go");
const outputPath = resolve(__dirname, "../bin/validator");

const child = spawn("go", ["build", "-o", outputPath, validatorPath]);

child.stdout.on("data", (data) => {
  console.log(data.toString());
});

child.stderr.on("data", (data) => {
  console.error(data.toString());
});
