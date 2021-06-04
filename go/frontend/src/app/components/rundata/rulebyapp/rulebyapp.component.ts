import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {Ruleusage} from "../../../model/apprulescore";

@Component({
  selector: 'rulebyapp',
  templateUrl: './rulebyapp.component.html',
  styleUrls: ['./rulebyapp.component.css','../rundatasummary/rundatasummary.component.css']
})
export class RuleByAppComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;

  public gridData: (string|number)[][]=[];
  public gridColumns: string[]=[];

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService) {

  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      console.log("runid is : "+runId);
      this.fetchRuleByApp(runId);
    });
  }

  resetPage(): void {
    this.gridData= [];
    this.gridColumns = [];
  }

  fetchRuleByApp(runId :number) : void {
    this.rundataService.getRuleByAppUsage(runId).subscribe(ruleByAppReturned => {
      ruleByAppReturned.cols.forEach(col => {
        this.gridColumns.push(col);
      });
      if (ruleByAppReturned.cols && ruleByAppReturned.data) {
        console.log(this.gridData.length);
        ruleByAppReturned.data.forEach(datum => {
          let innerArr :(string|number)[] = [];
          innerArr.push(datum.application);
          ruleByAppReturned.cols.forEach(col => {
            let ruleUsage: Ruleusage;
            ruleUsage = datum.ruleusage.find(indRuleUsage=> indRuleUsage.rule.toLowerCase()==col.toLowerCase());
            if(ruleUsage) {
              innerArr.push(ruleUsage.count);
            }
          });
          this.gridData.push(innerArr);
        });
        console.log(this.gridColumns);
      };
    }, error => {
      console.log(error);
    });
  }

}
