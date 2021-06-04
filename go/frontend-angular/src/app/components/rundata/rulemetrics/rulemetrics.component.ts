import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {Metric, RuleMetric} from "../../../model/rulemetric";

@Component({
  selector: 'rulemetrics',
  templateUrl: './rulemetrics.component.html',
  styleUrls: ['./rulemetrics.component.css','../rundatasummary/rundatasummary.component.css']
})
export class RuleMetricsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;
  ruleMetrics :Metric[]=[];

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService) {

  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      console.log("runid is : "+runId);
      this.fetchRuleMetrics(runId);
    });
  }

  resetPage(): void {
    this.ruleMetrics=[];
  }

  fetchRuleMetrics(runId: number): void{
    this.rundataService.getRuleMetrics(runId).subscribe(ruleMetricsReturned => {
      this.ruleMetrics = ruleMetricsReturned.metrics;
    }, error => {
      console.log(error);
    });
  }

}
