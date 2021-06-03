import {Component, Input, OnChanges, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {ApplicationScore} from "../../model/applicationscore";
import {ApplicationSummaryService} from "../../services/applicationsummary.service";
import {ClarityIcons, pinboardIcon, fileIcon, codeIcon, applicationsIcon, downloadIcon} from '@cds/core/icon';
import {RunSloc} from "../../model/runsloc";
import { ApplicationFindings } from 'src/app/model/applicationfindings';
import { Tags } from 'src/app/model/tags';

@Component({
  selector: 'csa-application-summary',
  templateUrl: './application-summary.component.html',
  styleUrls: ['./application-summary.component.css']
})
export class ApplicationSummaryComponent implements OnInit {

  //@Input()
  analyzerRun: any;

  applicationScores: ApplicationScore [] = [];
  applicationFindings: any;
  findings: number = 0;
  selectedAppId: number | undefined;
  filteredApplicationScores: ApplicationScore[] =[];
  applicationSelected: any;
  tags: any;
  scoreCard: any;

  constructor(private router: Router, private route: ActivatedRoute, private applicationSummaryService: ApplicationSummaryService) {
    ClarityIcons.addIcons(pinboardIcon);
    ClarityIcons.addIcons(fileIcon);
    ClarityIcons.addIcons(codeIcon);
    ClarityIcons.addIcons(downloadIcon);
    ClarityIcons.addIcons(applicationsIcon);
  }

  ngOnInit(): void {
    this.resetPage();
    this.route.paramMap.subscribe((params: ParamMap) => {
      const runId = Number(params.get('id'));
      console.log("runid is : "+runId);
      this.fetchAppScoresAndFindings(runId);
    });
  }

  resetPage(): void{
    this.applicationScores = [];
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

  appChanged(): void {
    this.applicationSelected = this.filteredApplicationScores.find(applicationScore => applicationScore.appId == this.selectedAppId);
    console.log(this.applicationSelected);
    this.fetchApplicationFindings(this.analyzerRun.id, this.applicationSelected)
    this.fetchScoreCardForApplication(this.analyzerRun.id, this.applicationSelected)
  }
}
