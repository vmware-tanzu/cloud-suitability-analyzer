export interface Apiusage {
  api: string;
  usageCount: number;
}

export interface Datum {
  application: string;
  apiusage: Apiusage[];
}

export interface AppApiUsage {
  cols: string[];
  data: Datum[];
}
