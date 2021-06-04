import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";
import {Annotation} from "../../../model/annotation";

@Component({
  selector: 'annotations',
  templateUrl: './annotations.component.html',
  styleUrls: ['./annotations.component.css','../rundatasummary/rundatasummary.component.css']
})
export class AnnotationsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'annotations.xlsx';
  annotations :Annotation[]=[];

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
    this.annotations=[];
  }

  fetchApiUsageDetailed(runId: number): void{
    this.rundataService.getAnnotationData(runId).subscribe(annotationDataReturned => {
      this.annotations = annotationDataReturned;
    }, error => {
      console.log(error);
    });
  }
}
