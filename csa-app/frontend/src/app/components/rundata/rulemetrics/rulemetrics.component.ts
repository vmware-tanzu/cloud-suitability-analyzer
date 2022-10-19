import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {Metric, RuleMetric} from "../../../model/rulemetric";
import {ToastrService} from 'ngx-toastr';
import {pushErrorNotification, pushInfoNotification} from '../../../utils/notificationutil';

@Component({
  selector: 'rulemetrics',
  templateUrl: './rulemetrics.component.html',
  styleUrls: ['./rulemetrics.component.css','../rundatasummary/rundatasummary.component.css']
})
export class RuleMetricsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;
  ruleMetrics: Metric[] = [];

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService, public toastr: ToastrService) {

  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      this.fetchRuleMetrics(runId);
    });
  }

  resetPage(): void {
    this.ruleMetrics = [];
  }

  fetchRuleMetrics(runId: number): void{
    this.rundataService.getRuleMetrics(runId).subscribe(ruleMetricsReturned => {
      this.ruleMetrics = ruleMetricsReturned.metrics;
      pushInfoNotification('Found [' + this.ruleMetrics.length + '] metrics!', this.toastr);
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

}
