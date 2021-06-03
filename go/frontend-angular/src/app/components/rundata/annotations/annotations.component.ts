import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";

@Component({
  selector: 'annotations',
  templateUrl: './annotations.component.html',
  styleUrls: ['./annotations.component.css']
})
export class AnnotationsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'annotations.xlsx';

  constructor(private router: Router, private rundataService: RundataService) {

  }

  ngOnInit(): void {
  }

}
