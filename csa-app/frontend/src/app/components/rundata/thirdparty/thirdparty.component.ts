import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {ThirdPartyLibsUsage} from "../../../model/thirdpartylibsusage";
import {ToastrService} from 'ngx-toastr';
import { pushErrorNotification } from '../../../utils/notificationutil';

@Component({
  selector: 'thirdparty',
  templateUrl: './thirdparty.component.html',
  styleUrls: ['./thirdparty.component.css','../rundatasummary/rundatasummary.component.css']
})
export class ThirdPartyComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'annotations.csv';
  thirdPartyLibs: ThirdPartyLibsUsage[] = [];

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService, public toastr: ToastrService) {

  }


  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      this.fetchThirdPartyLibsUsage(runId);
    });
  }

  resetPage(): void {
    this.thirdPartyLibs = [];
  }

  fetchThirdPartyLibsUsage(runId: number): void{
    this.rundataService.getThirdPartyData(runId).subscribe(thirdPartyLibsReturned => {
      this.thirdPartyLibs = thirdPartyLibsReturned;
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

}
