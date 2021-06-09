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
import { ApplicationFinding } from 'src/app/model/applicationfinding';
import { SelectedTag } from 'src/app/model/selectedtag';

@Component({
  selector: 'csa-application-summary',
  templateUrl: './application-summary.component.html',
  styleUrls: ['./application-summary.component.css']
})
export class ApplicationSummaryComponent implements OnInit {

  public searchCrit: any = '';

  applicationScores: ApplicationScore [] = [];
  allApplicationFindings: any;
  displayApplicationFindings: ApplicationFinding[] = [];
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
  tagName: number | undefined;
  binTagName: number | undefined;
  
  selectedBinTags: string[] = [];
  selectedTags: SelectedTag[] = [];

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
      let applicationScores = this.fetchAppScoresAndFindings(this.runId);

      applicationScores.forEach(appScore => {
        if(this.selectedAppId == undefined) {
          this.selectedAppId = 0;
          this.appChanged();
        }
      });
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

  fetchAppScoresAndFindings(runid: number): ApplicationScore[]{
    this.applicationSummaryService.getApplicationByRun(runid).subscribe(scores => {
      if (scores.scores.appScores) {
        this.applicationScores = scores.scores.appScores;
        this.findings = scores.scores.findings;
        scores.scores.appScores.forEach(appScore => {
          this.filteredApplicationScores.push(appScore);
        });
      }
    }, error => {
      console.log(error);
    });

    return this.applicationScores;
  }

  fetchApplicationFindings(runid: number, applicationSelected: ApplicationScore): void{
    this.applicationSummaryService.getApplicationAllFindings(runid, applicationSelected.name).subscribe(applicationFindings => {
      this.allApplicationFindings = applicationFindings;
      this.displayApplicationFindings = applicationFindings;
    }, error => {
      console.log(error);
    })
  }

  fetchTagsForApplication(runid: number, applicationSelected: ApplicationScore): void{
    if(applicationSelected.tags != null && applicationSelected.tags.length != 0) {
      applicationSelected.tags.forEach(tag => {
        this.applicationSummaryService.getApplicationFindingsByTag(runid, applicationSelected.name, tag).subscribe(tags => {
          let tagSummary : TagSummary = new TagSummary(tag.Value, tags.length, tags);
          this.tags.push(tagSummary);
        }, error => {
          console.log(error);
        })
      })
    }
  }

  fetchBinTagsForApplication(runid: number, applicationSelected: ApplicationScore): void{
    if(applicationSelected.bins != null && applicationSelected.bins.length != 0) {
      applicationSelected.bins.forEach(bin => {
        this.applicationSummaryService.getApplicationFindingsByTags(runid, applicationSelected.name, bin).subscribe(bintags => {
          let tagSummary : TagSummary = new TagSummary(bin.name, bintags.length, bintags);
          this.binTags.push(tagSummary);
        }, error => {
          console.log(error);
        })
      })
    }
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

  resetCssForBinTagsAndTags(binTag: string): void {
    document.getElementById(binTag).classList.replace("label-warning", "label-light-blue");

    this.applicationSelected.bins.forEach(bin => {
      if(bin.name == binTag) {
        bin.tags.forEach(tag => {
          let element = document.getElementById(tag.name);
          if (element != null || element != undefined) {
            element.classList.replace("label-success", "label-info");
            this.selectedTags.forEach((selectedTag, index) => {
              if(selectedTag.name == tag.name) {
                this.selectedTags.splice(index, 1);
              }
            });
          }
        });
      }
    });
  }

  highlightBinTags(event): void {
    
    let found = false;
    if(this.selectedBinTags.length != 0) {
      this.selectedBinTags.forEach((selectedBinTag, index) => {
        if(selectedBinTag == event.target.id) {
          this.selectedBinTags.splice(index, 1);
          this.resetCssForBinTagsAndTags(event.target.id)
          found = true;
        } 
      });
    }

    if(!found) {
      document.getElementById(event.target.id).classList.replace("label-light-blue", "label-warning");
      this.selectedBinTags.push(event.target.id);
    }

    this.showApplicationFindings();
  }

  highlightTags(event): void {
    console.log("Clicked Id", event.target.id)

    let found = false;
    let removed = false;

    if(this.selectedTags.length != 0) {
      this.selectedTags.forEach((selectedTag, index) => {
        if(selectedTag.name == event.target.id) {
          found = true;
          if(!selectedTag.isBinTag) {
            document.getElementById(event.target.id).classList.replace("label-success", "label-info");
            this.selectedTags.splice(index, 1);
            removed = true;
          }
        }
      });
    }

    if(!found) {
      document.getElementById(event.target.id).classList.replace("label-info", "label-success");
      let selectedTag: SelectedTag = new SelectedTag(event.target.id, false);
      this.selectedTags.push(selectedTag);
    }
    
    this.showApplicationFindings();

  }

  showApplicationFindings(): void {
    if(this.selectedBinTags.length == 0 && this.selectedTags.length == 0) {
      this.displayApplicationFindings = this.allApplicationFindings;
    } else {
      this.displayApplicationFindings = [];
    }


    let availableFindingIds: number[] = [];
    this.displayApplicationFindings.forEach(finding => {
      availableFindingIds.push(finding.id)
    });

    this.selectedBinTags.forEach(binTag => {
      this.binTags.forEach(data => {
        if (data.name == binTag) {
          data.tagData.forEach(tagFinding => {
            if(!availableFindingIds.includes(tagFinding.id)) {
              this.displayApplicationFindings.push(tagFinding);
            }
          });
        }
      });

      let selectedTags: string[] = [];
      this.selectedTags.forEach(selectedTag => {
        selectedTags.push(selectedTag.name);
      });
  
      this.applicationSelected.bins.forEach(bin => {
        if(bin.name == binTag) {
          bin.tags.forEach(tag => {
            let element = document.getElementById(tag.name);
            if (element != null || element != undefined) {
              element.classList.replace("label-info", "label-success");
              let selectedTag: SelectedTag = new SelectedTag(tag.name, true);
              if(!selectedTags.includes(selectedTag.name)) {
                this.selectedTags.push(selectedTag);
                selectedTags.push(selectedTag.name)
              }
            }
          });
        }
      });
    });

    this.selectedTags.forEach(selectedTag => {
      if(!selectedTag.isBinTag) {
        this.tags.forEach(tag => {
          if(tag.name == selectedTag.name) {
            tag.tagData.forEach((tagFinding, index) => {
              if(!availableFindingIds.includes(tagFinding.id)) {
                this.displayApplicationFindings.push(tagFinding);
              }
            });
          }
        });
      }
    });
  }
}
