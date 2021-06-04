import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ExcelExportComponent } from './excel-export.component';

describe('ExcelExportComponent', () => {
  let component: ExcelExportComponent;
  let fixture: ComponentFixture<ExcelExportComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ExcelExportComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ExcelExportComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
