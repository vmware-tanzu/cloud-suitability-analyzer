import {ComponentFixture, fakeAsync, flush, TestBed, waitForAsync} from '@angular/core/testing';

import {RouterTestingModule} from '@angular/router/testing';
import {AnalyzerRunService} from '../../services/analyzerrun.service';
import {of} from 'rxjs';
import {AnalyzerRunComponent} from './analyzer-run.component';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {Routes} from '@angular/router';
import {ExecutiveSummaryComponent} from '../executive-summary/executive-summary.component';

describe('AnalyzerRunComponent', () => {
  let component: AnalyzerRunComponent;
  let fixture: ComponentFixture<AnalyzerRunComponent>;
  let mockAnalyzerRunService: jasmine.SpyObj<AnalyzerRunService>;
  const mockData = {
    runs: [
      {
        id: 1,
        Alias: 'cloud-suitability-analyzer',
        User: 'bajajro',
        Command: 'analyze',
        Target: '/cloud-suitability-analyzer',
        ReportsRequested: '1,2,3,4,5',
        Files: 765,
        Findings: 1928,
        requestDate: '06/04/2021 15:24:12',
        Runtime: '5.029099859s',
        Homepath: '',
        Exepath: '',
        OutputPath: '',
        RulesDir: '',
        DbPath: '',
        TmpPath: ''
      }
    ]};
  const routes: Routes = [
    {
      path: 'runs/:id',
      children: [
        {
          path: 'summary',
          component: ExecutiveSummaryComponent
        }
      ]
    }
  ];

  beforeEach(waitForAsync(() => {
    mockAnalyzerRunService = jasmine.createSpyObj('AnalyzerRunService', ['getDistinctRuns']);
    TestBed.configureTestingModule({
      declarations: [ AnalyzerRunComponent],
      imports: [RouterTestingModule.withRoutes(routes), ToastrModule.forRoot(),
        BrowserAnimationsModule
      ],
      providers: [
        { provide: AnalyzerRunService, useValue: mockAnalyzerRunService },
        {provide: ToastrService, useClass: ToastrService},
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AnalyzerRunComponent);
    component = fixture.componentInstance;
    mockAnalyzerRunService.getDistinctRuns.and.returnValue(of(mockData));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('when component initializes then getDistinctRuns called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockAnalyzerRunService.getDistinctRuns).toHaveBeenCalledTimes(1);
    flush();
  }));
});
