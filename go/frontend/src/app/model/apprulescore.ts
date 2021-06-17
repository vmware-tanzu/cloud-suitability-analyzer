export interface Ruleusage {
  rule: string;
  count: number;
}

export interface Datum {
  application: string;
  ruleusage: Ruleusage[];
}

export interface AppRuleScore {
  cols: string[];
  data: Datum[];
}
