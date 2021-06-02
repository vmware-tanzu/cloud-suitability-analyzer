import { ApplicationFinding } from "./applicationfinding";

export class ApplicationFindings {
  applicationFindings: ApplicationFinding [];

  constructor(applicationFinding: ApplicationFinding []) {
    this.applicationFindings = applicationFinding;
  }
}
