import {ComponentFixture, fakeAsync, flush, TestBed, waitForAsync} from '@angular/core/testing';

import {SearchFilterPipe} from '../../pipes/search-filter.pipe';
import {RouterTestingModule} from '@angular/router/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {ActivatedRoute, convertToParamMap} from '@angular/router';
import {of} from 'rxjs';
import {ExecutiveSummaryService} from '../../services/executivesummary.service';
import {ExecutiveSummaryComponent} from './executive-summary.component';
import {Scores} from '../../model/scores';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

describe('ExecutiveSummaryComponent', () => {
  let component: ExecutiveSummaryComponent;
  let fixture: ComponentFixture<ExecutiveSummaryComponent>;
  let mockExecutiveSummaryService: jasmine.SpyObj<ExecutiveSummaryService>;
  const mockData = {
    scores: {
      findings: 1928,
        ciFindings: 354,
        infoFindings: 1574,
        rawScore: 5320,
        appScores: [
        {
          appId: 1,
          runId: 1,
          name: 'cloud-suitability-analyzer',
          path: '/cloud-suitability-analyzer',
          businessdomain: '',
          businessvalue: -1,
          findings: 1928,
          ciFindings: 354,
          infoFindings: 1574,
          rawScore: 5320,
          numCrits: 100,
          model: 'default',
          score: 6.27,
          originalScore: 6.27,
          scoreModified: false,
          recommendation: 'Refactor to TAS (consider modernization)',
          slocCnt: 173526,
          filesCnt: 662,
          findingsRatio: 15.028248587570621,
          tags: [],
          bins: []
        }
      ],
        recommendation: ''
    }
  } as Scores;
  const slocMockData = {
    runId: 1,
    applicationCount: 1,
    totalFiles: 662,
    blankLines: 7966,
    commentLines: 7635,
    codeLines: 173526
  };
  const top5ApiMockData = [
    {
      api: 'jboss',
      usageCount: 68
    },
    {
      api: 'io',
      usageCount: 66
    },
    {
      api: 'metrics',
      usageCount: 22
    },
    {
      api: 'logging',
      usageCount: 22
    },
    {
      api: 'servlet',
      usageCount: 20
    }
  ];
  const top5LanguagesMockData = [
    {
      language: 'JSON',
      codeLines: 95745
    },
    {
      language: 'Go',
      codeLines: 32574
    },
    {
      language: 'Plain Text',
      codeLines: 18729
    },
    {
      language: 'JavaScript',
      codeLines: 6257
    },
    {
      language: 'Java',
      codeLines: 5258
    }
  ];

  beforeEach(waitForAsync(() => {
    mockExecutiveSummaryService = jasmine.createSpyObj('ExecutiveSummaryService', ['getApplicationScoresByRun', 'getSummaryFindingsByRun', 'getLanguagesByLoc', 'getApisByScore']);
    TestBed.configureTestingModule({
      declarations: [ ExecutiveSummaryComponent, SearchFilterPipe],
      imports: [RouterTestingModule.withRoutes([
      ]), ToastrModule.forRoot(),
        BrowserAnimationsModule
      ],
      providers: [
        { provide: ExecutiveSummaryService, useValue: mockExecutiveSummaryService },
        {provide: ToastrService, useClass: ToastrService},
        {provide: ActivatedRoute,
          useValue: { paramMap: of(convertToParamMap({id: 1})) }}
      ]
    })
      .compileComponents();
  }));
  beforeEach(async () => {
    fixture = TestBed.createComponent(ExecutiveSummaryComponent);
    component = fixture.componentInstance;
    mockExecutiveSummaryService.getApplicationScoresByRun.withArgs(1).and.returnValue(of(mockData));
    mockExecutiveSummaryService.getSummaryFindingsByRun.withArgs(1).and.returnValue(of(slocMockData));
    mockExecutiveSummaryService.getApisByScore.withArgs(1).and.returnValue(of(top5ApiMockData));
    mockExecutiveSummaryService.getLanguagesByLoc.withArgs(1).and.returnValue(of(top5LanguagesMockData));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('when component initializes then getApplicationScoresByRun called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockExecutiveSummaryService.getApplicationScoresByRun).toHaveBeenCalledTimes(1);
    flush();
  }));
  it('when component initializes then getSummaryFindingsByRun called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockExecutiveSummaryService.getSummaryFindingsByRun).toHaveBeenCalledTimes(1);
    flush();
  }));
  it('when component initializes then getApisByScore called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockExecutiveSummaryService.getApisByScore).toHaveBeenCalledTimes(1);
    flush();
  }));
  it('when component initializes then getLanguagesByLoc called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockExecutiveSummaryService.getLanguagesByLoc).toHaveBeenCalledTimes(1);
    flush();
  }));
});
