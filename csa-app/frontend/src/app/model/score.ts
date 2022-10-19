import {ApplicationScore} from "./applicationscore";

export class Score {
  appScores: ApplicationScore [];
  findings: number;
  ciFindings: number;
  infoFindings: number;
  rawScore: number;
  recommendation: string;

  constructor(applicationScores: ApplicationScore[], findings: number, ciFindings: number, infoFindings: number, rawScore: number, recommendation: string) {
    this.appScores = applicationScores;
    this.findings = findings;
    this.ciFindings = ciFindings;
    this.infoFindings = infoFindings;
    this.rawScore = rawScore;
    this.recommendation = recommendation;
  }
}


