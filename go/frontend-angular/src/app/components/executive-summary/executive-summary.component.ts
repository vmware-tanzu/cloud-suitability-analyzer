import {Component, Input, OnChanges, OnInit} from '@angular/core';
import {Score} from "../../model/score";
import {Router} from "@angular/router";
import {ExecutiveSummaryService} from "../../services/executivesummary.service";
import {ApplicationScore} from "../../model/applicationscore";
import {RunSloc} from "../../model/runsloc";
import {AnalyzerRun} from "../../model/analyzerrun";
import {forkJoin} from "rxjs";
import {Scores} from "../../model/scores";
import {ClarityIcons, pinboardIcon, fileIcon, codeIcon, applicationsIcon, downloadIcon} from '@cds/core/icon';
import '@cds/core/icon/register.js';
import {LanguagesByCodeLines} from "../../model/languagesbycodelines";
import {ApisByScore} from "../../model/apisbyscore";

@Component({
  selector: 'csa-executive-summary',
  templateUrl: './executive-summary.component.html',
  styleUrls: ['./executive-summary.component.css']
})
export class ExecutiveSummaryComponent implements OnChanges {

  @Input()
  analyzerRun: any;

  applicationScores: ApplicationScore [] = [];
  filteredApplicationScores: ApplicationScore[] =[];
  numAppsByRun: number = 0;
  locByRun: number = 0;
  numFilesByRun: number = 0;
  findings: number = 0;
  top5LanguagesByLocData: LanguagesByCodeLines[]= [];
  top5ApisByScoreData: ApisByScore[] = [];
  cardBackground: string = 'count card-background';
  activeIndex: number = 0;
  filter: string = "";
  loading: boolean = true;
  selectedApp: string = "";
  tableEditable: boolean = true;

  constructor(private router: Router, private executiveSummaryService: ExecutiveSummaryService) {
    ClarityIcons.addIcons(pinboardIcon);
    ClarityIcons.addIcons(fileIcon);
    ClarityIcons.addIcons(codeIcon);
    ClarityIcons.addIcons(downloadIcon);
    ClarityIcons.addIcons(applicationsIcon);
  }

/*  ngOnInit(): void {
    if(this.analyzerRun) {
      console.log("in ngOnInit analyzerRun.id - "+this.analyzerRun.id);
      this.resetPage();
      this.fetchAppScoresAndFindings(this.analyzerRun.id);
    }
  }*/
  ngOnChanges(): void {
    if(this.analyzerRun) {
      this.resetPage();
      this.fetchAppScoresAndFindings(this.analyzerRun.id);
    }
  }

  resetPage(): void{
    this.findings = 0;
    this.applicationScores = [];
    this.numFilesByRun = 0;
    this.locByRun = 0;
    this.numAppsByRun = 0;
    this.filteredApplicationScores = [];
    this.top5LanguagesByLocData = [];
    this.top5ApisByScoreData = [];
  }
  fetchAppScoresAndFindings(runid: number): void{
    let runSlocBlank: RunSloc= {
      blankLines: 0, commentLines: 0, runId: 0,
      applicationCount: 0,
      codeLines: 0,
      totalFiles: 0
    };
    this.executiveSummaryService.getApplicationScoresByRun(runid).subscribe(scores => {
      if (scores.scores.appScores) {
        this.applicationScores = scores.scores.appScores;
        this.findings = scores.scores.findings;
        console.log(scores.scores.appScores.length);
        scores.scores.appScores.forEach(appScore => {
          this.filteredApplicationScores.push(appScore);
        });
      }
    }, error => {
      console.log(error);
    });
    if (runid === 0 || runid === null) {
      this.numAppsByRun = runSlocBlank.applicationCount;
      this.locByRun = runSlocBlank.codeLines;
      this.numFilesByRun = runSlocBlank.totalFiles;
    }else {
      this.executiveSummaryService.getSummaryFindingsByRun(runid).subscribe(runSlocReturned => {
        this.numAppsByRun = runSlocReturned.applicationCount;
        this.locByRun = runSlocReturned.codeLines;
        this.numFilesByRun = runSlocReturned.totalFiles;
      }, error => {
        console.log(error);
      })
    }
}

  fetchTop5LanguagesByLoc(runid: number) {
    this.executiveSummaryService.getLanguagesByLoc(runid).subscribe(languagesByLocReturned => {
      this.top5LanguagesByLocData = languagesByLocReturned.slice(0, 5);
    }, error => {
      console.log(error);
    });
  }

  fetchTop5ApisByScore(runid: number) {
    this.executiveSummaryService.getApisByScore(runid).subscribe(apisByScoreReturned => {
      this.top5ApisByScoreData = apisByScoreReturned.slice(0, 5);
    }, error => {
      console.log(error);
    });
  }
}
