import {ComponentFixture, fakeAsync, flush, TestBed, waitForAsync} from '@angular/core/testing';

import {SearchFilterPipe} from '../../pipes/search-filter.pipe';
import {RouterTestingModule} from '@angular/router/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {ActivatedRoute, convertToParamMap} from '@angular/router';
import {of} from 'rxjs';
import {Scores} from '../../model/scores';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {ApplicationSummaryComponent} from './application-summary.component';
import {ApplicationSummaryService} from '../../services/applicationsummary.service';

describe('ApplicationSummaryComponent', () => {
  let component: ApplicationSummaryComponent;
  let fixture: ComponentFixture<ApplicationSummaryComponent>;
  let mockApplicationSummaryService: jasmine.SpyObj<ApplicationSummaryService>;
  const appName = 'cloud-suitability-analyzer';
  const applicationSelected = {
    appId: 1,
    runId: 1,
    name: 'cloud-suitability-analyzer',
    tags: [
      {
        Value: 'docker'
      }
    ],
    bins: [
      {
        name: 'TKG',
        tags: [
          {
            name: 'docker'
          }
        ],
        matched: [
          'docker'
        ]
      }
    ]
  };
  const tag = {
    Value : 'docker'
  };
  const bin = {
    name: 'TKG',
    tags: [
      {
        name: 'docker'
      }
    ],
    matched: [
      'docker'
    ]
  };
  const mockAppData = {
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
          tags: [
            {
              Value: 'docker'
            }
          ],
          bins: [
            {
              name: 'TKG',
              tags: [
                {
                  name: 'docker'
                }
              ],
              matched: [
                'docker'
              ]
            }
          ]
        }
      ],
      recommendation: ''
    }
  } as Scores;
  const mockLanguagesData = [
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
    }
  ];
  const mockApiData = [
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
  const mockFindingsData = [
      {
        id: 521,
        run: 1,
        filename: 'AppStateChecker.java',
        fqn: '/cloud-suitability-analyzer/go/frontend/samplePortfolio/liberty-arquillian/liberty-managed/src/main/java/io/openliberty/arquillian/managed/AppStateChecker.java',
        ext: '.java',
        line: 3,
        rule: 'java-metrics',
        pattern: 'javax.management',
        value: 'import javax.management.InstanceNotFoundException;',
        note: 'note1',
        advice: 'Indicates use of a metrics collection library, which supports containerization',
        level: 'low',
        effort: 1,
        readiness: 7,
        category: 'metrics',
        criticality: '',
        application: 'cloud-suitability-analyzer',
        tags: [
          'docker'
        ],
        recipes: null
      }
    ];
  const mockScoreCard = {
    application: '',
    info: 811,
    low: 183,
    medium: 3,
    high: 166,
    total: 1163
  };

  beforeEach(waitForAsync(() => {
    mockApplicationSummaryService = jasmine.createSpyObj('ApplicationSummaryService', ['getApplicationByRun', 'getLanguagesByApplicationRun', 'getApisByApplicationRun', 'getApplicationAllFindings', 'getApplicationFindingsByTag', 'getApplicationFindingsByTags', 'getApplicationScorecard']);
    TestBed.configureTestingModule({
      declarations: [ ApplicationSummaryComponent, SearchFilterPipe],
      imports: [RouterTestingModule.withRoutes([
      ]), ToastrModule.forRoot(),
        BrowserAnimationsModule
      ],
      providers: [
        { provide: ApplicationSummaryService, useValue: mockApplicationSummaryService },
        {provide: ToastrService, useClass: ToastrService},
        {provide: ActivatedRoute,
          useValue: { paramMap: of(convertToParamMap({id: 1})) }}
      ]
    })
      .compileComponents();
  }));
  beforeEach(async () => {
    fixture = TestBed.createComponent(ApplicationSummaryComponent);
    component = fixture.componentInstance;
    component.applicationSelected = applicationSelected;
    component.selectedAppId = 1;
    mockApplicationSummaryService.getApplicationByRun.withArgs(1).and.returnValue(of(mockAppData));
    mockApplicationSummaryService.getLanguagesByApplicationRun.withArgs(1, appName).and.returnValue(of(mockLanguagesData));
    mockApplicationSummaryService.getApisByApplicationRun.withArgs(1, appName).and.returnValue(of(mockApiData));
    mockApplicationSummaryService.getApplicationAllFindings.withArgs(1, appName).and.returnValue(of(mockFindingsData));
    mockApplicationSummaryService.getApplicationFindingsByTag.withArgs(1, appName, tag).and.returnValue(of(mockFindingsData));
    mockApplicationSummaryService.getApplicationScorecard.withArgs(1, appName).and.returnValue(of(mockScoreCard));
    mockApplicationSummaryService.getApplicationFindingsByTags.withArgs(1, appName, bin).and.returnValue(of(mockFindingsData));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('when component initializes then getApplicationByRun called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockApplicationSummaryService.getApplicationByRun).toHaveBeenCalledTimes(1);
    flush();
  }));
  it('when component initializes then getLanguagesByApplicationRun called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockApplicationSummaryService.getLanguagesByApplicationRun).toHaveBeenCalledTimes(1);
    flush();
  }));
  it('when component initializes then getApisByApplicationRun called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockApplicationSummaryService.getApisByApplicationRun).toHaveBeenCalledTimes(1);
    flush();
  }));
  it('when component initializes then getApplicationAllFindings called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockApplicationSummaryService.getApplicationAllFindings).toHaveBeenCalledTimes(1);
    flush();
  }));
  it('when component initializes then getApplicationScorecard called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockApplicationSummaryService.getApplicationScorecard).toHaveBeenCalledTimes(1);
    flush();
  }));
  it('when component initializes then getApplicationFindingsByTag called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockApplicationSummaryService.getApplicationFindingsByTag).toHaveBeenCalledTimes(1);
    flush();
  }));
});

