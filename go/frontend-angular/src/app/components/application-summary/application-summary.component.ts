import {Component, Input, OnChanges, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {ApplicationScore} from "../../model/applicationscore";
import {ApplicationSummaryService} from "../../services/applicationsummary.service";
import {ClarityIcons, pinboardIcon, fileIcon, codeIcon, applicationsIcon, downloadIcon} from '@cds/core/icon';
import { ChartElement } from "../../model/chartelement";
import '@cds/core/search/register.js';
import { Language } from 'src/app/model/language';
import { Api } from 'src/app/model/api';

@Component({
  selector: 'csa-application-summary',
  templateUrl: './application-summary.component.html',
  styleUrls: ['./application-summary.component.css']
})
export class ApplicationSummaryComponent implements OnInit {

  //@Input()
  analyzerRun: any;

  public searchCrit: any = '';

  applicationScores: ApplicationScore [] = [];
  applicationFindings: any;
  findings: number = 0;
  selectedAppId: number | undefined;
  filteredApplicationScores: ApplicationScore[] =[];
  applicationSelected: any;
  tags: any;
  scoreCard: any;
  languages: ChartElement[] = [];
  apis: ChartElement[] = [];

  view: [number, number] = [700, 400];

  // options
  gradient: boolean = false;
  animations: boolean = true;

  colorScheme = {
    domain: ['#0095D3', '#00BFA9', '#60B515', '#8939AD', '#F57600', '#6870C4']
  };

  constructor(private router: Router, private route: ActivatedRoute, private applicationSummaryService: ApplicationSummaryService) {
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
      console.log("runid is : "+runId);
      this.fetchAppScoresAndFindings(runId);
    });
  }

  resetPage(): void{
    this.findings = 0;
    this.selectedAppId = 0;
    this.applicationScores = [];
    this.filteredApplicationScores = [];
    this.languages = [];
    this.apis = [];
  }

  fetchAppScoresAndFindings(runid: number): void{
    this.applicationSummaryService.getApplicationByRun(runid).subscribe(scores => {
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
  }

  fetchApplicationFindings(runid: number, applicationSelected: ApplicationScore): void{
    this.applicationSummaryService.getApplicationFindings(runid, applicationSelected.name, applicationSelected.tags).subscribe(applicationFindings => {
      this.applicationFindings = applicationFindings;
    }, error => {
      console.log(error);
    })
  }

  fetchTagsForApplication(runid: number, applicationSelected: ApplicationScore): void{
    this.applicationSummaryService.getTagsForApplication(runid, applicationSelected.name).subscribe(tags => {
      this.tags = tags.tags;
    }, error => {
      console.log(error);
    })
  }

  fetchScoreCardForApplication(runid: number, applicationSelected: ApplicationScore): void{
    this.applicationSummaryService.getApplicationScorecard(runid, applicationSelected.name, applicationSelected.tags).subscribe(scoreCard => {
      this.scoreCard = scoreCard;
    }, error => {
      console.log(error);
    })
  }

  fetchLanguagesByApplicationRun(runid: number, applicationSelected: ApplicationScore): void{
    this.applicationSummaryService.getLanguagesByApplicationRun(runid, applicationSelected.name).subscribe(languagesResponse => {
      languagesResponse.forEach(language => {
        let chartElement: ChartElement = new ChartElement(language.codeLines, language.language);
        this.languages.push(chartElement)
      });
    }, error => {
      console.log(error);
    })
  }

  fetchApisByApplicationRun(runid: number, applicationSelected: ApplicationScore): void{
    this.applicationSummaryService.getApisByApplicationRun(runid, applicationSelected.name).subscribe(apisResponse => {
      apisResponse.forEach(api => {
        let chartElement: ChartElement = new ChartElement(api.usageCount, api.api);
        this.apis.push(chartElement)
      });
    }, error => {
      console.log(error);
    })
  }

  appChanged(): void {
    this.applicationSelected = this.filteredApplicationScores.find(applicationScore => applicationScore.appId == this.selectedAppId);
    this.fetchApplicationFindings(this.analyzerRun.id, this.applicationSelected);
    this.fetchScoreCardForApplication(this.analyzerRun.id, this.applicationSelected);

    this.languages = [];
    this.apis = [];
    this.fetchLanguagesByApplicationRun(this.analyzerRun.id, this.applicationSelected);
    this.fetchApisByApplicationRun(this.analyzerRun.id, this.applicationSelected);
  }

  onSelectLanguage(language: Language[]): void {
    console.log('Active', JSON.parse(JSON.stringify(language)));
  }

  onSelectApis(api: Api[]): void {
    console.log('Active', JSON.parse(JSON.stringify(api)));
  }
}
