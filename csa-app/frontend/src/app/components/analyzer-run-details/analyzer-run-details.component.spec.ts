import {ComponentFixture, fakeAsync, TestBed, waitForAsync} from '@angular/core/testing';

import { AnalyzerRunDetailsComponent } from './analyzer-run-details.component';
import {RouterTestingModule} from '@angular/router/testing';
import {AnalyzerRunService} from '../../services/analyzerrun.service';
import {AnalyzerRun, AnalyzerRuns} from '../../model/analyzerrun';

describe('AnalyzerRunDetailsComponent', () => {
  let component: AnalyzerRunDetailsComponent;
  let fixture: ComponentFixture<AnalyzerRunDetailsComponent>;
  let mockAnalyzerRunService: jasmine.SpyObj<AnalyzerRunService>;

  beforeEach(waitForAsync(() => {
    mockAnalyzerRunService = jasmine.createSpyObj('AnalyzerRunService', ['getDistinctRuns']);
    TestBed.configureTestingModule({
      declarations: [ AnalyzerRunDetailsComponent],
      imports: [RouterTestingModule.withRoutes([
      ])
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AnalyzerRunDetailsComponent);
    component = fixture.componentInstance;
    component.analyzerRun = {} as AnalyzerRun;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
