import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Router} from '@angular/router';
import {AnalyzerRunService} from "../../services/analyzerrun.service";
import {AnalyzerRun} from "../../model/analyzerrun";
import {RunDropDownItem} from "../../model/rundropdownitem";
import {calendarIcon, ClarityIcons, clockIcon, targetIcon, userIcon} from '@cds/core/icon';
import '@cds/core/icon/register.js';
import {pushErrorNotification} from '../../utils/notificationutil';
import {ToastrService} from "ngx-toastr";

@Component({
  selector: 'app-analyzer-run',
  templateUrl: './analyzer-run.component.html',
  styleUrls: ['./analyzer-run.component.css']
})
export class AnalyzerRunComponent implements OnInit {

  analyzerRuns: AnalyzerRun[] = [];

  runItemsDropDownItems: RunDropDownItem[] = [];

  @Input()
  selectedRunId: number | undefined;

  @Output()
  selectedRunIdChange = new EventEmitter();


  selectedRunDesc: string = '';
  analyzerRunSelected: any;
  loaded: boolean = false;

  constructor(private router: Router, private analyzerRunService: AnalyzerRunService, public toastr: ToastrService) {
    ClarityIcons.addIcons(calendarIcon);
    ClarityIcons.addIcons(clockIcon);
    ClarityIcons.addIcons(targetIcon);
    ClarityIcons.addIcons(userIcon);
  }

  ngOnInit(): void {
    this.analyzerRunService.getDistinctRuns().subscribe(runList => {
      Object.keys(runList).forEach((key) => {
        this.analyzerRuns = runList[key];
      });
      this.analyzerRuns.forEach((analyzerRun, index) => {
        this.runItemsDropDownItems.push(new RunDropDownItem(analyzerRun.id, analyzerRun.id + " - " + analyzerRun.Alias));
        /* istanbul ignore else */
        if (index === 0) {
          this.selectedRunId = analyzerRun.id;
          this.selectedRunDesc = analyzerRun.id + " - " + analyzerRun.Alias;
          this.analyzerRunSelected = analyzerRun;
          this.loaded = true;
          this.selectedRunIdChange.emit(this.selectedRunId);
          this.router.navigateByUrl(`/runs/${this.selectedRunId}/summary`, {skipLocationChange: true});
        }
      });
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  runIdChanged(): void {
    this.analyzerRunSelected = this.analyzerRuns.find(analyzerRun => analyzerRun.id == this.selectedRunId);
    this.selectedRunIdChange.emit(this.selectedRunId);
    this.router.navigateByUrl(`/runs/${this.selectedRunId}/summary`, {skipLocationChange: true});
  }

}
