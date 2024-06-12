import {Component, Input, OnInit} from '@angular/core';
import csvDownload from 'json-to-csv-export';

@Component({
  selector: 'app-excel-export',
  templateUrl: './excel-export.component.html',
  styleUrls: ['./excel-export.component.css']
})

export class ExcelExportComponent implements OnInit {

  @Input() data: any;

  @Input() file_name: string;

  constructor() {}

  ngOnInit() {}

  public exportAsExcelFile(): void {
    const data = this.data;
    const worksheetColumns = [];
    for (let i = 0; i < Object.keys(data[0]).length; i++) {
      worksheetColumns[i] = Object.keys(data[0])[i];
    }

    const writedata = {data: data, filename: this.file_name, delimiter: ',', headers: worksheetColumns}
    csvDownload(writedata)
  }

}
