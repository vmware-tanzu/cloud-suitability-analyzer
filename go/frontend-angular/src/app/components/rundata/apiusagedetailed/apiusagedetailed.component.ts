import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {RundataService} from "../../../services/rundata.service";

@Component({
  selector: 'apiusagedetailed',
  templateUrl: './apiusagedetailed.component.html',
  styleUrls: ['./apiusagedetailed.component.css']
})
export class ApiUsageDetailedComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string = 'api-usage-detailed.xlsx';

  constructor(private router: Router, private rundataService: RundataService) {

  }

  ngOnInit(): void {
  }

}
