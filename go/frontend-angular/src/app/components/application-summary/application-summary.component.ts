import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, ParamMap, Router } from "@angular/router";
import { ApplicationScore } from "../../model/applicationscore";
import { ApplicationSummaryService} from "../../services/applicationsummary.service";
import { ClarityIcons, pinboardIcon, thermometerIcon, downloadIcon } from '@cds/core/icon';
import { ChartElement } from "../../model/chartelement";
import '@cds/core/search/register.js';
import { Language } from 'src/app/model/language';
import { Api } from 'src/app/model/api';
import { TagSummary } from 'src/app/model/tagsummary';

@Component({
  selector: 'csa-application-summary',
  templateUrl: './application-summary.component.html',
  styleUrls: ['./application-summary.component.css']
})
export class ApplicationSummaryComponent implements OnInit {

  public searchCrit: any = '';

  applicationScores: ApplicationScore [] = [];
  applicationFindings: any;
  runId: number = 0;
  findings: number = 0;
  selectedAppId: number | undefined;
  filteredApplicationScores: ApplicationScore[] =[];
  applicationSelected: any;
  tags: TagSummary[] = [];
  binTags: TagSummary[] = [];
  languages: ChartElement[] = [];
  apis: ChartElement[] = [];
  scoreCard: any;
  
  view: [number, number] = [700, 400];

  // options
  gradient: boolean = false;
  animations: boolean = true;

  colorScheme = {
    domain: ['#0095D3', '#00BFA9', '#60B515', '#8939AD', '#F57600', '#6870C4']
  };

  constructor(private router: Router, private route: ActivatedRoute, private applicationSummaryService: ApplicationSummaryService) {
    ClarityIcons.addIcons(pinboardIcon);
    ClarityIcons.addIcons(downloadIcon);
    ClarityIcons.addIcons(thermometerIcon);
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      this.runId = Number(params.get('id'));
      console.log("runid is : "+ this.runId);
      this.fetchAppScoresAndFindings(this.runId);
    });
  }

  resetPage(): void{
    this.findings = 0;
    this.selectedAppId = 0;
    this.applicationScores = [];
    this.filteredApplicationScores = [];
    this.languages = [];
    this.apis = [];
    this.tags = [];
    this.binTags = [];
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
    this.applicationSummaryService.getApplicationAllFindings(runid, applicationSelected.name).subscribe(applicationFindings => {
      this.applicationFindings = applicationFindings;
    }, error => {
      console.log(error);
    })
  }

  fetchTagsForApplication(runid: number, applicationSelected: ApplicationScore): void{
    applicationSelected.tags.forEach(tag => {
      this.applicationSummaryService.getApplicationFindingsByTag(runid, applicationSelected.name, tag).subscribe(tags => {
        let tagSummary : TagSummary = new TagSummary(tag.Value, tags.length);
        this.tags.push(tagSummary);
      }, error => {
        console.log(error);
      })
    })
  }

  fetchBinTagsForApplication(runid: number, applicationSelected: ApplicationScore): void{
    applicationSelected.bins.forEach(bin => {
      this.applicationSummaryService.getApplicationFindingsByTags(runid, applicationSelected.name, bin).subscribe(bintags => {
        let tagSummary : TagSummary = new TagSummary(bin.name, bintags.length);
        this.binTags.push(tagSummary);
      }, error => {
        console.log(error);
      })
    })

  }

  fetchScoreCardForApplication(runid: number, applicationSelected: ApplicationScore): void{
    this.applicationSummaryService.getApplicationScorecard(runid, applicationSelected.name).subscribe(scoreCard => {
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
    this.languages = [];
    this.apis = [];
    this.tags = [];
    this.binTags = [];

    this.applicationSelected = this.filteredApplicationScores.find(applicationScore => applicationScore.appId == this.selectedAppId);
    this.fetchApplicationFindings(this.runId, this.applicationSelected);
    this.fetchScoreCardForApplication(this.runId, this.applicationSelected);
    
    this.fetchBinTagsForApplication(this.runId, this.applicationSelected);
    this.fetchTagsForApplication(this.runId, this.applicationSelected);
    this.fetchLanguagesByApplicationRun(this.runId, this.applicationSelected);
    this.fetchApisByApplicationRun(this.runId, this.applicationSelected);
  }

  onSelectLanguage(language: Language[]): void {
    console.log('Active', JSON.parse(JSON.stringify(language)));
  }

  onSelectApis(api: Api[]): void {
    console.log('Active', JSON.parse(JSON.stringify(api)));
  }
}
