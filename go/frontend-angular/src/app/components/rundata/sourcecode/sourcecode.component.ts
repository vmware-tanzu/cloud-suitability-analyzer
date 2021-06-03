import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";

@Component({
  selector: 'sourcecode',
  templateUrl: './sourcecode.component.html',
  styleUrls: ['./sourcecode.component.css']
})
export class SourceCodeComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'source-code.xlsx';

  constructor(private router: Router, private rundataService: RundataService) {

  }

  ngOnInit(): void {
  }

}
