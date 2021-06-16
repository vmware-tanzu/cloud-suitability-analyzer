import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from '@angular/router';
import {RundataService} from '../../../services/rundata.service';
import {Apiusage} from '../../../model/appapiusage';
import {ToastrService} from 'ngx-toastr';
import {pushErrorNotification} from '../../../utils/notificationutil';
import {ClrDatagridStringFilterInterface} from '@clr/angular';

@Component({
  selector: 'apibyapp',
  templateUrl: './apibyapp.component.html',
  styleUrls: ['./apibyapp.component.css', '../rundatasummary/rundatasummary.component.css']
})
export class ApiByAppComponent implements OnInit {

  public searchCrit: any = '';
  public fileName: string;

  public gridData: (string | number)[][] = [];
  public gridColumns: string[] = [];

  public appNameFilter = new AppNameFilter();

  constructor(private router: Router, private route: ActivatedRoute, private rundataService: RundataService, public toastr: ToastrService) {

  }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params: ParamMap) => {
      this.resetPage();
      const runId = Number(params.get('id'));
      this.fetchApiByApps(runId);
    });
  }

  resetPage(): void {
    this.gridData = [];
    this.gridColumns = [];
  }

  fetchApiByApps(runId: number): void {
    this.rundataService.getApiByAppUsage(runId).subscribe(appApiUsageReturned => {
      appApiUsageReturned.cols.forEach(col => {
        /* istanbul ignore else */
        if (col !== 'App') {
          this.gridColumns.push(col);
        }
      });
      /* istanbul ignore else */
      if (appApiUsageReturned.cols && appApiUsageReturned.data) {
        appApiUsageReturned.data.forEach(datum => {
          const innerArr: (string | number)[] = [];
          innerArr.push(datum.application);
          appApiUsageReturned.cols.forEach(col => {
            let apiusage: Apiusage;
            apiusage = datum.apiusage.find(indApiUsage => indApiUsage.api.toLowerCase() === col.toLowerCase());
            /* istanbul ignore else */
            if (apiusage) {
              innerArr.push(apiusage.usageCount);
            }
          });
          this.gridData.push(innerArr);
        });
      }
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

}

class AppNameFilter implements ClrDatagridStringFilterInterface<[]> {
  accepts(apiByAppRow: [], search: string): boolean {
    // @ts-ignore
    return apiByAppRow[0].toLowerCase().indexOf(search) >= 0;
  }
}
