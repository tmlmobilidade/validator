import { spawn } from "child_process";
import path from "path";

/**
 * Runs a Go binary and returns its JSON stdout as an object.
 * @param binaryPath Absolute or relative path to the Go binary.
 * @param args Arguments to pass to the binary (optional).
 * @param timeout Timeout in milliseconds (default 1000 * 60 * 5) - 5 minutes.
 * @returns A promise that resolves to a JSON object from the Go binary.
 */
export async function runGoBinary<T = unknown>(
    binaryPath: string,
    args: string[] = [],
    timeout: number = 1000 * 60 * 5
  ): Promise<T> {
    return new Promise<T>((resolve, reject) => {
      const fullPath = path.resolve(binaryPath);
      const proc = spawn(fullPath, args, {
        stdio: ["ignore", "pipe", "pipe"],
      });
  
      const stdoutChunks: Buffer[] = [];
      const stderrChunks: Buffer[] = [];
      let timedOut = false;
  
      const timer = setTimeout(() => {
        timedOut = true;
        proc.kill();
        reject(new Error(`Process timeout after ${timeout}ms`));
      }, timeout);
  
      proc.stdout.on("data", (chunk: Buffer) => {
        stdoutChunks.push(chunk);
      });
  
      proc.stderr.on("data", (chunk: Buffer) => {
        stderrChunks.push(chunk);
      });
  
      proc.on("error", (err: Error) => {
        clearTimeout(timer);
        reject(new Error(`Failed to start binary: ${err.message}`));
      });
  
      proc.on("close", (code: number) => {
        clearTimeout(timer);
        if (timedOut) return;
  
        const stdout = Buffer.concat(stdoutChunks).toString("utf-8").trim();
        const stderr = Buffer.concat(stderrChunks).toString("utf-8").trim();
  
        if (code !== 0) {
          return reject(
            new Error(`Binary exited with code ${code}: ${stderr || stdout}`)
          );
        }
  
        try {
          // Find the last line that looks like JSON
          const lastLine = stdout.split('\n').filter(line => line.trim()).pop() || '';
          const json = JSON.parse(lastLine) as T;
          resolve(json);
        } catch (e: any) {
          throw new Error(
            `Failed to parse JSON output: ${e.message} - Output: ${stdout}`
          );
        }
      });
    });
  }