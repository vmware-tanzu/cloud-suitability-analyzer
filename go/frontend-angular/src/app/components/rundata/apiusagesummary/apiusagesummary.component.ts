import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";

@Component({
  selector: 'apiusagesummary',
  templateUrl: './apiusagesummary.component.html',
  styleUrls: ['./apiusagesummary.component.css']
})
export class ApiUsageSummaryComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'api-usage-summary.xlsx';

  constructor(private router: Router, private rundataService: RundataService) {

  }

  ngOnInit(): void {
  }

}
