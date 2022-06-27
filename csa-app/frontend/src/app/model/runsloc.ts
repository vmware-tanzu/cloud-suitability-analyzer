//http://localhost:3001/api/runs/1/summary/run_slocs
export interface RunSloc {
  runId: number;
  applicationCount: number;
  totalFiles: number;
  blankLines: number;
  commentLines: number;
  codeLines: number;
}
