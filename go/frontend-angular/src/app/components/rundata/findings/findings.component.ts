import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";

@Component({
  selector: 'findings',
  templateUrl: './findings.component.html',
  styleUrls: ['./findings.component.css']
})
export class FindingsComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'findings.xlsx';

  constructor(private router: Router, private rundataService: RundataService) {

  }

  ngOnInit(): void {
  }

}
