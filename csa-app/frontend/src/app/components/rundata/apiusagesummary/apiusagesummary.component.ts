import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {ApiSummary} from "../../../model/apisummary";
import {ToastrService} from 'ngx-toastr';
import { pushErrorNotification } from '../../../utils/notificationutil';

@Component({
  selector: 'apiusagesummary',
  templateUrl: './apiusagesummary.component.html',
  styleUrls: ['./apiusagesummary.component.css', '../rundatasummary/rundatasummary.component.css']
})
export class ApiUsageSummaryComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;
  apiSummaries: ApiSummary[] = [];

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService, public toastr: ToastrService) {

  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      this.fetchApiUsageSummary(runId);
    });
  }

  resetPage(): void {
    this.apiSummaries = [];
  }

  fetchApiUsageSummary(runId: number): void{
    this.rundataService.getApiSummaryUsage(runId).subscribe(apiSummaryReturned => {
      this.apiSummaries = apiSummaryReturned;
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

}
