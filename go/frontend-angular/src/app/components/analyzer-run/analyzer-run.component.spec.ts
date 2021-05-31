import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AnalyzerRunComponent } from './analyzer-run.component';

describe('AnalyzerRunComponent', () => {
  let component: AnalyzerRunComponent;
  let fixture: ComponentFixture<AnalyzerRunComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AnalyzerRunComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AnalyzerRunComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
