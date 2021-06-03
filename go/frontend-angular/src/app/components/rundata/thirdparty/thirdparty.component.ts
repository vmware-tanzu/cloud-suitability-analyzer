import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {ThirdPartyLibsUsage} from "../../../model/thirdpartylibsusage";

@Component({
  selector: 'thirdparty',
  templateUrl: './thirdparty.component.html',
  styleUrls: ['./thirdparty.component.css','../rundatasummary/rundatasummary.component.css']
})
export class ThirdPartyComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'annotations.xlsx';
  thirdPartyLibs :ThirdPartyLibsUsage[]=[];

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService) {

  }


  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      console.log("runid is : "+runId);
      this.fetchApiUsageDetailed(runId);
    });
  }

  resetPage(): void {
    this.thirdPartyLibs=[];
  }

  fetchApiUsageDetailed(runId: number): void{
    this.rundataService.getThirdPartyData(runId).subscribe(thirdPartyLibsReturned => {
      console.log(thirdPartyLibsReturned);
      this.thirdPartyLibs = thirdPartyLibsReturned;
    }, error => {
      console.log(error);
    });
  }

}
