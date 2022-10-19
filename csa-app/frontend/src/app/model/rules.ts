export interface Tag {
  Value: string;
}

export interface Pattern {
  Value: string;
}

export interface Rule {
  Name: string;
  FileType: string;
  Target: string;
  Type: string;
  DefaultPattern: string;
  Advice: string;
  Effort: number;
  Category: string;
  Tags: Tag[];
  Patterns: Pattern[];
}

export interface CsaRules {
  rules: Rule[];
}
