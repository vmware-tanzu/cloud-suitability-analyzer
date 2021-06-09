import {Component, Input, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {Findings} from "../../../model/findings";

@Component({
  selector: 'findings',
  templateUrl: './findings.component.html',
  styleUrls: ['./findings.component.css','../rundatasummary/rundatasummary.component.css']
})
export class FindingsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;
  findings :Findings[]=[];
  selectedRunId: number;
  selectedFinding: Findings;
  isOpen: boolean;
  findingLoaded: boolean = false;

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService) {

  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      console.log("runid is : "+runId);
      this.selectedRunId = runId;
      this.fetchFindings(runId);
    });
  }

  resetPage(): void {
    this.findings=[];
  }

  fetchFindings(runId: number): void{
    this.rundataService.getRunFindings(runId).subscribe(findingsReturned => {
      this.findings = findingsReturned;
      this.findings.map(finding => {
        if(finding.recipes == null){
          finding.recipes = [];
        }
      })
    }, error => {
      console.log(error);
    });
  }

  fetchFindingById(findingId: number): void{
    this.rundataService.getFinding(this.selectedRunId, findingId).subscribe(findingDataReturned => {
      this.selectedFinding = findingDataReturned;
      this.findingLoaded = true;
      this.isOpen = true;
      console.log("findingLoaded: "+ this.findingLoaded);
      console.log("isOpen: "+ this.isOpen);
    }, error => {
      console.log(error);
    });
  }

  doNothing(): boolean{
    return false;
  }
}
