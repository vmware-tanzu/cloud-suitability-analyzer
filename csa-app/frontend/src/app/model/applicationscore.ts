//http://localhost:3001/api/runs/1/summary/application_scores?rangeMin=0&rangeMax=10&reverse=false
export interface Tag {
  Value: string;
}
export interface BinTag {
  name: string;
}

export interface Bin {
  name: string;
  tags: BinTag[];
  matched: string[];
}

export interface ApplicationScore {
  appId: number;
  runId: number;
  name: string;
  path: string;
  businessdomain: string;
  businessvalue: number;
  findings: number;
  ciFindings: number;
  infoFindings: number;
  rawScore: number;
  numCrits: number;
  model: string;
  score: number;
  originalScore: number;
  scoreModified: boolean;
  recommendation: string;
  slocCnt: number;
  filesCnt: number;
  findingsRatio: number;
  tags: Tag[];
  bins: Bin[];
}
