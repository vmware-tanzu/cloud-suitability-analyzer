import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {Findings} from "../../../model/findings";
import {ToastrService} from 'ngx-toastr';
import {pushErrorNotification, pushInfoNotification} from '../../../utils/notificationutil';

@Component({
  selector: 'findings',
  templateUrl: './findings.component.html',
  styleUrls: ['./findings.component.css', '../rundatasummary/rundatasummary.component.css']
})
export class FindingsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;
  findings: Findings[] = [];
  selectedRunId: number;
  selectedFinding: Findings;
  isOpen: boolean;
  findingLoaded: boolean = false;

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService, public toastr: ToastrService) {

  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      this.selectedRunId = runId;
      this.fetchFindings(runId);
    });
  }

  resetPage(): void {
    this.findings = [];
  }

  fetchFindings(runId: number): void {
    this.rundataService.getRunFindings(runId).subscribe(findingsReturned => {
      this.findings = findingsReturned;
      pushInfoNotification('Found [' + this.findings.length + '] findings!', this.toastr);
      this.findings.map(finding => {
        /* istanbul ignore else */
        if (finding.recipes == null) {
          finding.recipes = [];
        }
      });
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

  fetchFindingById(findingId: number): void {
    this.rundataService.getFinding(this.selectedRunId, findingId).subscribe(findingDataReturned => {
      this.selectedFinding = findingDataReturned;
      this.findingLoaded = true;
      this.isOpen = true;
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }
}
