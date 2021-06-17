import {Component, Input, OnInit} from '@angular/core';
import * as XLSX from 'xlsx';

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
      worksheetColumns[i] = {width: Object.keys(data[0])[i].length + 3};
    }
    const worksheet = XLSX.utils.json_to_sheet(data);
    const workbook: XLSX.WorkBook = {
      Sheets: { data: worksheet },
      SheetNames: ['data']
    };

    worksheet['!cols'] = worksheetColumns;
    XLSX.writeFile(workbook, this.file_name, {
      bookType: 'xlsx',
      type: 'buffer'
    });
  }

}
