import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {Annotation} from "../../../model/annotation";
import {ToastrService} from 'ngx-toastr';
import {pushErrorNotification} from '../../../utils/notificationutil';

@Component({
  selector: 'annotations',
  templateUrl: './annotations.component.html',
  styleUrls: ['./annotations.component.css', '../rundatasummary/rundatasummary.component.css']
})
export class AnnotationsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'annotations.csv';
  annotations: Annotation[] = [];

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
    this.annotations = [];
  }

  fetchApiUsageDetailed(runId: number): void {
    this.rundataService.getAnnotationData(runId).subscribe(annotationDataReturned => {
      this.annotations = annotationDataReturned;
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }
}
