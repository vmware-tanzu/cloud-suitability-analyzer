import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {ApiDetailedUsage} from "../../../model/apidetailedusage";
import {ToastrService} from 'ngx-toastr';
import { pushErrorNotification } from '../../../utils/notificationutil';

@Component({
  selector: 'apiusagedetailed',
  templateUrl: './apiusagedetailed.component.html',
  styleUrls: ['./apiusagedetailed.component.css', '../rundatasummary/rundatasummary.component.css']
})
export class ApiUsageDetailedComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;
  apiDetailedUsages: ApiDetailedUsage[]=[];

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService, public toastr: ToastrService) {

  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      this.fetchApiUsageDetailed(runId);
    });
  }

  resetPage(): void {
    this.apiDetailedUsages = [];
  }

  fetchApiUsageDetailed(runId: number): void{
    this.rundataService.getApiDetailedUsage(runId).subscribe(apiDetailedUsageReturned => {
      apiDetailedUsageReturned.forEach(apiDetailedUsageItemReturned =>
      {
        this.apiDetailedUsages.push(apiDetailedUsageItemReturned);
      });
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

}
