export interface Metric {
  rule: string;
  RunID: number;
  checks: number;
  patternChecks: number;
  hits: number;
  totalTime: string;
  longest: string;
  shortest: string;
  avgRule: string;
  avgPat: string;
  avgHit: string;
}

export interface RuleMetric {
  metrics: Metric[];
}
