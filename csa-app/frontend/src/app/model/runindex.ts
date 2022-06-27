export interface Index {
  runId: number;
  alias: string;
  path: string;
  exists: boolean;
  error: string;
  numDocs: number;
}

export interface RunIndex {
  index: Index;
}
