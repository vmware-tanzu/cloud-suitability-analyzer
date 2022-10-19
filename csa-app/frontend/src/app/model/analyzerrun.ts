//http://localhost:3001/api/analyze-runs
export class AnalyzerRuns {
  runs: AnalyzerRun[];
}
export class AnalyzerRun {
  id: number;
  Alias: string;
  User: string;
  Command: string;
  Target: string;
  ReportsRequested: string;
  Files: number;
  Findings: number;
  requestDate: string;
  Runtime: string;
  Homepath: string;
  Exepath: string;
  OutputPath: string;
  RulesDir: string;
  DbPath: string;
  TmpPath: string;

  constructor(id: number, alias: string, user: string, command: string, target: string, reportsRequested: string, files: number, findings: number, requestDate: string, runtime: string, homepath: string, exepath: string, outputPath: string, rulesDir: string, dbPath: string, tmpPath: string) {
    this.id = id;
    this.Alias = alias;
    this.User = user;
    this.Command = command;
    this.Target = target;
    this.ReportsRequested = reportsRequested;
    this.Files = files;
    this.Findings = findings;
    this.requestDate = requestDate;
    this.Runtime = runtime;
    this.Homepath = homepath;
    this.Exepath = exepath;
    this.OutputPath = outputPath;
    this.RulesDir = rulesDir;
    this.DbPath = dbPath;
    this.TmpPath = tmpPath;
  }
}


