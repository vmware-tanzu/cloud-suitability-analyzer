import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {ExecutiveSummaryService} from "../../services/executivesummary.service";
import {ApplicationScore} from "../../model/applicationscore";
import {RunSloc} from "../../model/runsloc";
import {applicationsIcon, ClarityIcons, codeIcon, downloadIcon, fileIcon, pinboardIcon} from '@cds/core/icon';
import '@cds/core/icon/register.js';
import {LanguagesByCodeLines} from "../../model/languagesbycodelines";
import {ApisByScore} from "../../model/apisbyscore";
import {ChartElement} from "../../model/chartelement";
import '@cds/core/search/register.js';
import {pushErrorNotification, pushInfoNotification} from '../../utils/notificationutil';
import {ToastrService} from "ngx-toastr";
import { Color, ScaleType } from '@swimlane/ngx-charts';

@Component({
  selector: 'csa-executive-summary',
  templateUrl: './executive-summary.component.html',
  styleUrls: ['./executive-summary.component.css']
})
export class ExecutiveSummaryComponent implements OnInit {

  public searchCrit: any = '';

  applicationScores: ApplicationScore [] = [];
  filteredApplicationScores: ApplicationScore[] = [];
  numAppsByRun: number = 0;
  locByRun: number = 0;
  numFilesByRun: number = 0;
  findings: number = 0;
  top5LanguagesByLocData: ChartElement[] = [];
  top5ApisByScoreData: ChartElement[] = [];
  cardBackground: string = 'count card-background';
  activeIndex: number = 0;
  filter: string = "";
  loading: boolean = true;
  selectedApp: string = "";
  tableEditable: boolean = true;

  view: [number, number] = [500, 400];

  // chart display options
  showXAxis: boolean = true;
  showYAxis: boolean = true;
  gradient: boolean = true;
  showXAxisLabel: boolean = true;
  yAxisLabelForApiByScore: string = 'Top 5 APIs';
  yAxisLabelForLanguagesByLoc: string = 'Top 5 languages';
  showYAxisLabel: boolean = true;
  xAxisLabelForApiByScore: string = 'USAGES';
  xAxisLabelForLanguagesByLoc: string = 'LINES OF CODE';
  barPadding: number = 40;
  showGridLines: boolean = true;

  colorScheme: Color = {
    name: 'executiveSummary',
    selectable: true,
    group: ScaleType.Ordinal,
    domain: ['#0095D3', '#00BFA9', '#60B515', '#8939AD', '#F57600', '#6870C4']
  };

  constructor(private router: Router, private route: ActivatedRoute, private executiveSummaryService: ExecutiveSummaryService, public toastr: ToastrService) {
    ClarityIcons.addIcons(pinboardIcon);
    ClarityIcons.addIcons(fileIcon);
    ClarityIcons.addIcons(codeIcon);
    ClarityIcons.addIcons(downloadIcon);
    ClarityIcons.addIcons(applicationsIcon);
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      this.fetchAppScoresAndFindings(runId);
      this.fetchTop5ApisByScore(runId);
      this.fetchTop5LanguagesByLoc(runId);
    });
  }

  resetPage(): void {
    this.findings = 0;
    this.applicationScores = [];
    this.numFilesByRun = 0;
    this.locByRun = 0;
    this.numAppsByRun = 0;
    this.filteredApplicationScores = [];
    this.top5LanguagesByLocData = [];
    this.top5ApisByScoreData = [];
  }

  fetchAppScoresAndFindings(runid: number): void {
    let runSlocBlank: RunSloc = {
      blankLines: 0, commentLines: 0, runId: 0,
      applicationCount: 0,
      codeLines: 0,
      totalFiles: 0
    };
    this.executiveSummaryService.getApplicationScoresByRun(runid).subscribe(scores => {
      /* istanbul ignore else */
      if (scores.scores.appScores) {
        this.applicationScores = scores.scores.appScores;
        this.findings = scores.scores.findings;
        scores.scores.appScores.forEach(appScore => {
          this.filteredApplicationScores.push(appScore);
        });
        pushInfoNotification('Found [' + this.applicationScores.length + '] applications!', this.toastr);
      }
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
    if (runid === 0 || runid === null) {
      this.numAppsByRun = runSlocBlank.applicationCount;
      this.locByRun = runSlocBlank.codeLines;
      this.numFilesByRun = runSlocBlank.totalFiles;
    } else {
      this.executiveSummaryService.getSummaryFindingsByRun(runid).subscribe(runSlocReturned => {
        this.numAppsByRun = runSlocReturned.applicationCount;
        this.locByRun = runSlocReturned.codeLines;
        this.numFilesByRun = runSlocReturned.totalFiles;
      }, error => {
        pushErrorNotification(error, this.toastr);
      });
    }
  }

  fetchTop5LanguagesByLoc(runid: number) {
    this.executiveSummaryService.getLanguagesByLoc(runid).subscribe(languagesByLocReturned => {
      languagesByLocReturned.sort((a, b) => b.codeLines - a.codeLines);
      const top5LanguagesByLoc: LanguagesByCodeLines[] = languagesByLocReturned.slice(0, 5);
      top5LanguagesByLoc.forEach(languageByLoc => {
        let chartElement: ChartElement = new ChartElement(languageByLoc.codeLines, languageByLoc.language);
        this.top5LanguagesByLocData.push(chartElement);
      });
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  fetchTop5ApisByScore(runid: number) {
    this.executiveSummaryService.getApisByScore(runid).subscribe(apisByScoreReturned => {
      apisByScoreReturned.sort((a, b) => b.usageCount - a.usageCount);
      const top5ApisByScore: ApisByScore[] = apisByScoreReturned.slice(0, 5);
      top5ApisByScore.forEach(apiScore => {
        let chartElement: ChartElement = new ChartElement(apiScore.usageCount, apiScore.api);
        this.top5ApisByScoreData.push(chartElement);
      });
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  onSelectLanguagesByLoc(top5LanguagesByLoc: LanguagesByCodeLines[]): void {
  }

  onActivateLanguagesByLoc(top5LanguagesByLoc: LanguagesByCodeLines[]): void {
  }

  onDeactivateLanguagesByLoc(top5LanguagesByLoc: LanguagesByCodeLines[]): void {
  }

  onSelectApisByScore(top5ApisByScoreData: ApisByScore[]): void {
  }

  onActivateApisByScore(top5ApisByScoreData: ApisByScore[]): void {
  }

  onDeactivateApisByScore(top5ApisByScoreData: ApisByScore[]): void {
  }
}
