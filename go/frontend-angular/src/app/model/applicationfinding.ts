import { Tag } from "./applicationscore";

export class ApplicationFinding {
    id: number;
    run: number;
    filename: string;
    fqn: string;
    ext: string;
    line: number;
    rule: string;
    pattern: string;
    value: string;
    advice: string;
    level: string;
    effort: number;
    readiness: number;
    category: string;
    criticality: string;
    application: string;
    recipes: string;
    tags: Tag[];
  
    constructor(id: number, run: number, filename: string, fqn: string, ext: string, line: number, rule: string, pattern: string, value: string, advice: string, level: string, effort: number, readiness: number, category: string, criticality: string, application: string, recipes: string, tags: Tag[]) {
      this.id = id;
      this.run = run;
      this.filename = filename;
      this.fqn = fqn;
      this.ext = ext;
      this.line = line;
      this.rule = rule;
      this.pattern = pattern;
      this.value = value;
      this.advice = advice;
      this.level = level;
      this.effort = effort;
      this.readiness = readiness;
      this.category = category;
      this.criticality = criticality;
      this.application = application;
      this.recipes = recipes;
      this.tags = tags;
    }
  }