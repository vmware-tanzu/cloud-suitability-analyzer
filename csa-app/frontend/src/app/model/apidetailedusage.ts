export interface ApiDetailedUsage {
  application: string;
  api: string;
  filename: string;
  pattern: string;
  line: number;
  value: string;
  effort: number;
  level: string;
  advice: string;
}
