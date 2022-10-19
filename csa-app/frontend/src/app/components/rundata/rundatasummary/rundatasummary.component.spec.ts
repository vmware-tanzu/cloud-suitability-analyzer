import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RunDataSummaryComponent } from './rundatasummary.component';

describe('RunDataSummaryComponent', () => {
  let component: RunDataSummaryComponent;
  let fixture: ComponentFixture<RunDataSummaryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RunDataSummaryComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RunDataSummaryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
