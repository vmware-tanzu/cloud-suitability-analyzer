import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnalyzerRunDetailsComponent } from './analyzer-run-details.component';

describe('AnalyzerRunDetailsComponent', () => {
  let component: AnalyzerRunDetailsComponent;
  let fixture: ComponentFixture<AnalyzerRunDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnalyzerRunDetailsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AnalyzerRunDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
