import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {Sloc} from "../../../model/sloc";
import {ToastrService} from 'ngx-toastr';
import { pushErrorNotification } from '../../../utils/notificationutil';

@Component({
  selector: 'sourcecode',
  templateUrl: './sourcecode.component.html',
  styleUrls: ['./sourcecode.component.css','../rundatasummary/rundatasummary.component.css']
})
export class SourceCodeComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;
  slocDetails: Sloc[] = [];

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService, public toastr: ToastrService) {

  }
  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      this.fetchSloc(runId);
    });
  }

  resetPage(): void {
    this.slocDetails = [];
  }

  fetchSloc(runId: number): void{
    this.rundataService.getSourceCodeData(runId).subscribe(slocReturned => {
      this.slocDetails = slocReturned;
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

}
