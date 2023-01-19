import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, NavigationExtras, ParamMap, Router} from "@angular/router";
import {ApplicationScore} from "../../model/applicationscore";
import {ApplicationSummaryService} from "../../services/applicationsummary.service";
import {ClarityIcons, downloadIcon, pinboardIcon, thermometerIcon} from '@cds/core/icon';
import {ChartElement} from "../../model/chartelement";
import '@cds/core/search/register.js';
import {Language} from 'src/app/model/language';
import {Api} from 'src/app/model/api';
import {TagSummary} from 'src/app/model/tagsummary';
import {SelectedTag} from 'src/app/model/selectedtag';
import {ToastrService} from 'ngx-toastr';
import {pushErrorNotification, pushInfoNotification} from 'src/app/utils/notificationutil';
import {Observable} from "rxjs/index";
import {Findings} from "../../model/findings";
import { Color, ScaleType } from '@swimlane/ngx-charts';

@Component({
  selector: 'csa-application-summary',
  templateUrl: './application-summary.component.html',
  styleUrls: ['./application-summary.component.css']
})
export class ApplicationSummaryComponent implements OnInit {

  public searchCrit: any = '';

  applicationScores: ApplicationScore [] = [];
  allApplicationFindings: any;
  displayApplicationFindings: Findings[] = [];
  runId = 0;
  findings = 0;
  selectedAppId = 0;
  filteredApplicationScores: ApplicationScore[] = [];
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

  selectedFinding: Findings;
  isOpen: boolean;
  findingLoaded: boolean = false;

  view: [number, number] = [1390, 400];

  // options
  gradient = false;
  animations = true;

  colorScheme: Color = {
    name: 'applicationSummary',
    selectable: true,
    group: ScaleType.Ordinal,
    domain: ['#0095D3', '#00BFA9', '#60B515', '#8939AD', '#F57600', '#6870C4']
  };

  // tslint:disable-next-line:max-line-length
  constructor(private router: Router, private route: ActivatedRoute, private applicationSummaryService: ApplicationSummaryService, public toastr: ToastrService) {
    ClarityIcons.addIcons(pinboardIcon);
    ClarityIcons.addIcons(downloadIcon);
    ClarityIcons.addIcons(thermometerIcon);
  }

  getParams(): Observable<ParamMap> {
    return this.route.paramMap;
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      this.runId = Number(params.get('id'));
      this.fetchAppScoresAndFindings(this.runId);
    });
  }

  appChanged(): void {
    this.languages = [];
    this.apis = [];
    this.tags = [];
    this.binTags = [];
    this.route.queryParamMap.subscribe((queryParams: ParamMap) => {
      const appName = queryParams.get('app');
      if (appName && appName !== undefined) {
        let navigationExtras: NavigationExtras = {
          queryParams: { 'app': appName }
        };
        this.selectedAppId = this.filteredApplicationScores.find(applicationScore => applicationScore.name === appName).appId;
        if (!this.runId) {
          this.router.navigate(['/runs/1/application'], navigationExtras);
        }
      }
    });
    this.applicationSelected = this.filteredApplicationScores.find(applicationScore => applicationScore.appId == this.selectedAppId);
    this.fetchApplicationFindings(this.runId, this.applicationSelected);
    this.fetchScoreCardForApplication(this.runId, this.applicationSelected);

    this.fetchBinTagsForApplication(this.runId, this.applicationSelected);
    this.fetchTagsForApplication(this.runId, this.applicationSelected);
    this.fetchLanguagesByApplicationRun(this.runId, this.applicationSelected);
    this.fetchApisByApplicationRun(this.runId, this.applicationSelected);
  }

  resetPage(): void {
    this.findings = 0;
    this.selectedAppId = 0;
    this.applicationScores = [];
    this.displayApplicationFindings = [];
    this.filteredApplicationScores = [];
    this.languages = [];
    this.apis = [];
    this.tags = [];
    this.binTags = [];
    this.selectedBinTags = [];
    this.selectedTags = [];
  }

  fetchAppScoresAndFindings(runid: number): void {
    this.applicationSummaryService.getApplicationByRun(runid).subscribe(scores => {
      /* istanbul ignore else */
      if (scores.scores.appScores) {
        this.applicationScores = scores.scores.appScores;
        this.findings = scores.scores.findings;
        scores.scores.appScores.forEach(appScore => {
          this.filteredApplicationScores.push(appScore);
          this.selectedAppId = appScore.appId;
        });
        this.appChanged();
      }
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  fetchApplicationFindings(runid: number, applicationSelected: ApplicationScore): void {
    this.applicationSummaryService.getApplicationAllFindings(runid, applicationSelected.name).subscribe(applicationFindings => {
      applicationFindings.map(finding => {
        /* istanbul ignore else */
        if (finding.recipes == null) {
          finding.recipes = [];
        }
      }, error => {
        pushErrorNotification(error, this.toastr);
      });

      this.allApplicationFindings = applicationFindings;
      this.displayApplicationFindings = applicationFindings;

      pushInfoNotification('Found [' + this.displayApplicationFindings.length + '] findings!', this.toastr);
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  fetchTagsForApplication(runid: number, applicationSelected: ApplicationScore): void {
    /* istanbul ignore else */
    if (applicationSelected.tags != null && applicationSelected.tags.length != 0) {
      applicationSelected.tags.forEach(tag => {
        this.applicationSummaryService.getApplicationFindingsByTag(runid, applicationSelected.name, tag).subscribe(tags => {
          let tagSummary: TagSummary = new TagSummary(tag.Value, tags.length, tags);
          this.tags.push(tagSummary);
        }, error => {
          pushErrorNotification(error, this.toastr);
        });
      });
    }
  }

  fetchBinTagsForApplication(runid: number, applicationSelected: ApplicationScore): void {
    /* istanbul ignore else */
    if (applicationSelected.bins != null && applicationSelected.bins.length != 0) {
      applicationSelected.bins.forEach(bin => {
        this.applicationSummaryService.getApplicationFindingsByTags(runid, applicationSelected.name, bin).subscribe(bintags => {
          let tagSummary: TagSummary = new TagSummary(bin.name, bintags.length, bintags);
          this.binTags.push(tagSummary);
        }, error => {
          pushErrorNotification(error, this.toastr);
        });
      });
    }
  }

  fetchScoreCardForApplication(runid: number, applicationSelected: ApplicationScore): void {
    this.applicationSummaryService.getApplicationScorecard(runid, applicationSelected.name).subscribe(scoreCard => {
      this.scoreCard = scoreCard;
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  fetchLanguagesByApplicationRun(runid: number, applicationSelected: ApplicationScore): void {
    this.applicationSummaryService.getLanguagesByApplicationRun(runid, applicationSelected.name).subscribe(languagesResponse => {
      languagesResponse.forEach(language => {
        let chartElement: ChartElement = new ChartElement(language.codeLines, language.language);
        this.languages.push(chartElement)
      });
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  fetchApisByApplicationRun(runid: number, applicationSelected: ApplicationScore): void {
    this.applicationSummaryService.getApisByApplicationRun(runid, applicationSelected.name).subscribe(apisResponse => {
      apisResponse.forEach(api => {
        let chartElement: ChartElement = new ChartElement(api.usageCount, api.api);
        this.apis.push(chartElement)
      });
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  onSelectLanguage(language: Language[]): void {
  }

  onSelectApis(api: Api[]): void {
  }

  resetCssForBinTagsAndTags(binTag: string): void {
    document.getElementById(binTag).classList.replace("label-warning", "label-light-blue");

    this.applicationSelected.bins.forEach(bin => {
      /* istanbul ignore else */
      if (bin.name == binTag) {
        bin.tags.forEach(tag => {
          let element = document.getElementById(tag.name);
          /* istanbul ignore else */
          if (element != null || element != undefined) {
            element.classList.replace("label-success", "label-info");
            this.selectedTags.forEach((selectedTag, index) => {
              /* istanbul ignore else */
              /* istanbul ignore else */
              if (selectedTag.name == tag.name) {
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
    /* istanbul ignore else */
    if (this.selectedBinTags.length != 0) {
      this.selectedBinTags.forEach((selectedBinTag, index) => {
        /* istanbul ignore else */
        if (selectedBinTag == event.target.id) {
          this.selectedBinTags.splice(index, 1);
          this.resetCssForBinTagsAndTags(event.target.id)
          found = true;
        }
      });
    }

    /* istanbul ignore else */
    if (!found) {
      document.getElementById(event.target.id).classList.replace("label-light-blue", "label-warning");
      this.selectedBinTags.push(event.target.id);
    }

    this.showApplicationFindings();
  }

  highlightTags(event): void {
    let found = false;
    let removed = false;

    /* istanbul ignore else */
    if (this.selectedTags.length != 0) {
      this.selectedTags.forEach((selectedTag, index) => {
        /* istanbul ignore else */
        if (selectedTag.name == event.target.id) {
          found = true;
          /* istanbul ignore else */
          if (!selectedTag.isBinTag) {
            document.getElementById(event.target.id).classList.replace("label-success", "label-info");
            this.selectedTags.splice(index, 1);
            removed = true;
          }
        }
      });
    }

    /* istanbul ignore else */
    if (!found) {
      document.getElementById(event.target.id).classList.replace("label-info", "label-success");
      let selectedTag: SelectedTag = new SelectedTag(event.target.id, false);
      this.selectedTags.push(selectedTag);
    }

    this.showApplicationFindings();

  }

  showApplicationFindings(): void {
    if (this.selectedBinTags.length == 0 && this.selectedTags.length == 0) {
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
        /* istanbul ignore else */
        if (data.name == binTag) {
          data.tagData.forEach(tagFinding => {
            /* istanbul ignore else */
            if (!availableFindingIds.includes(tagFinding.id)) {
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
        /* istanbul ignore else */
        if (bin.name == binTag) {
          bin.tags.forEach(tag => {
            let element = document.getElementById(tag.name);
            /* istanbul ignore else */
            if (element != null || element != undefined) {
              element.classList.replace("label-info", "label-success");
              let selectedTag: SelectedTag = new SelectedTag(tag.name, true);
              /* istanbul ignore else */
              if (!selectedTags.includes(selectedTag.name)) {
                this.selectedTags.push(selectedTag);
                selectedTags.push(selectedTag.name)
              }
            }
          });
        }
      });
    });

    this.selectedTags.forEach(selectedTag => {
      /* istanbul ignore else */
      if (!selectedTag.isBinTag) {
        this.tags.forEach(tag => {
          /* istanbul ignore else */
          if (tag.name == selectedTag.name) {
            tag.tagData.forEach((tagFinding, index) => {
              /* istanbul ignore else */
              if (!availableFindingIds.includes(tagFinding.id)) {
                this.displayApplicationFindings.push(tagFinding);
              }
            });
          }
        });
      }
    });

    this.displayApplicationFindings.map(finding => {
      /* istanbul ignore else */
      if (finding.recipes == null) {
        finding.recipes = [];
      }
    }, error => {
      pushErrorNotification(error, this.toastr);
    });

    pushInfoNotification('Found [' + this.displayApplicationFindings.length + '] findings!', this.toastr);
  }

  fetchFindingById(findingId: number): void {
    this.displayApplicationFindings.forEach(finding => {
      // tslint:disable-next-line:triple-equals
      if (finding.id == findingId) {
        this.selectedFinding = finding;
        this.findingLoaded = true;
        this.isOpen = true;
      }
    });
  }
}
