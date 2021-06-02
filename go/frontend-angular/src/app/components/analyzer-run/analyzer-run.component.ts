import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Router} from '@angular/router';
import {AnalyzerRunService} from "../../services/analyzerrun.service";
import {AnalyzerRun} from "../../model/analyzerrun";
import {RunDropDownItem} from "../../model/rundropdownitem";
import {ClarityIcons, calendarIcon, clockIcon, targetIcon, userIcon} from '@cds/core/icon';
import '@cds/core/icon/register.js';

@Component({
  selector: 'app-analyzer-run',
  templateUrl: './analyzer-run.component.html',
  styleUrls: ['./analyzer-run.component.css']
})
export class AnalyzerRunComponent implements OnInit {

  analyzerRuns: AnalyzerRun[]=[];
  forgeVersion: string='';
  runItemsDropDownItems: RunDropDownItem[]=[];
  selectedRunId: number | undefined;
  selectedRunDesc: string='';
  analyzerRunSelected: any;
  loaded:boolean = false;

  constructor(private router: Router, private analyzerRunService: AnalyzerRunService) {
    ClarityIcons.addIcons(calendarIcon);
    ClarityIcons.addIcons(clockIcon);
    ClarityIcons.addIcons(targetIcon);
    ClarityIcons.addIcons(userIcon);
  }

  public static createBlankAnalyzerRun() :AnalyzerRun {
    return new AnalyzerRun(0,'','','','','',0,0,'','','','','','','','');
  }

    ngOnInit(): void {
      this.analyzerRunService.getForgeVersion().subscribe(version => {
        this.forgeVersion = version;
      }, error => {
        console.log(error);
      });
    this.analyzerRunService.getDistinctRuns().subscribe(runList => {
      console.log(runList);
      Object.keys(runList).forEach((key) => {
        console.log(runList[key]);
        this.analyzerRuns = runList[key];
      });
      this.analyzerRuns.forEach((analyzerRun, index) => {
        console.log(analyzerRun);
        this.runItemsDropDownItems.push(new RunDropDownItem(analyzerRun.id, analyzerRun.id+" - " +analyzerRun.Alias));
        if(index == 0) {
          this.selectedRunId = analyzerRun.id;
          this.selectedRunDesc = analyzerRun.id+" - " +analyzerRun.Alias;
          this.analyzerRunSelected = analyzerRun;
          console.log("loaded - "+this.selectedRunId);
          this.loaded = true;
          this.router.navigateByUrl(`/runs/${this.selectedRunId}/summary`);
          // this.router.navigateByUrl(`/runs/${this.selectedRunId}/summary`, {state: {data: this.analyzerRunSelected}});
        }
      })
    }, error => {
      console.log(error);
    });
  }

  runIdChanged(): void {
    this.analyzerRunSelected = this.analyzerRuns.find(analyzerRun => analyzerRun.id == this.selectedRunId);
    console.log(this.analyzerRunSelected);
    this.router.navigateByUrl(`/runs/${this.selectedRunId}/summary`);
  }

}
