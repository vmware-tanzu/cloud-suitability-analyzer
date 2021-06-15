import {ComponentFixture, fakeAsync, TestBed, waitForAsync} from '@angular/core/testing';

import { ApiByAppComponent } from './apibyapp.component';
import {RundataService} from '../../../services/rundata.service';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {RouterTestingModule} from '@angular/router/testing';
import {of} from 'rxjs';
import {SearchFilterPipe} from '../../../pipes/search-filter.pipe';
import {ActivatedRoute, convertToParamMap} from '@angular/router';

describe('ApibyappComponent', () => {
  let component: ApiByAppComponent;
  let fixture: ComponentFixture<ApiByAppComponent>;
  let mockRunDataService: jasmine.SpyObj<RundataService>;
  const mockData = {
    cols: [
      'App',
      'buildSystem',
      'cache',
      'config',
      'hardip',
      'healthcheck',
      'io',
      'jboss',
      'logging',
      'metrics',
      'process',
      'servlet',
      'spring-boot',
      'springFramework',
      'termsignal',
      'webSphere',
      'websphere',
      'wlcluster',
      '***TOTAL'
    ],
    data: [
      {
        application: 'cloud-suitability-analyzer',
        apiusage: [
          {
            api: 'process',
            usageCount: 6
          },
          {
            api: 'logging',
            usageCount: 66
          },
          {
            api: 'healthcheck',
            usageCount: 2
          },
          {
            api: 'hardip',
            usageCount: 2
          },
          {
            api: 'metrics',
            usageCount: 22
          },
          {
            api: 'websphere',
            usageCount: 600
          },
          {
            api: 'io',
            usageCount: 528
          },
          {
            api: '***TOTAL',
            usageCount: 4910
          },
          {
            api: 'servlet',
            usageCount: 0
          },
          {
            api: 'buildSystem',
            usageCount: 0
          },
          {
            api: 'spring-boot',
            usageCount: -200
          },
          {
            api: 'cache',
            usageCount: 200
          },
          {
            api: 'jboss',
            usageCount: 3400
          },
          {
            api: 'config',
            usageCount: 100
          },
          {
            api: 'wlcluster',
            usageCount: 2
          },
          {
            api: 'webSphere',
            usageCount: 180
          },
          {
            api: 'springFramework',
            usageCount: 0
          },
          {
            api: 'termsignal',
            usageCount: 2
          }
        ]
      },
      {
        application: '***TOTAL',
        apiusage: [
          {
            api: 'metrics',
            usageCount: 22
          },
          {
            api: 'webSphere',
            usageCount: 180
          },
          {
            api: 'spring-boot',
            usageCount: -200
          },
          {
            api: 'healthcheck',
            usageCount: 2
          },
          {
            api: 'logging',
            usageCount: 66
          },
          {
            api: 'wlcluster',
            usageCount: 2
          },
          {
            api: 'servlet',
            usageCount: 0
          },
          {
            api: 'jboss',
            usageCount: 3400
          },
          {
            api: 'termsignal',
            usageCount: 2
          },
          {
            api: 'hardip',
            usageCount: 2
          },
          {
            api: 'buildSystem',
            usageCount: 0
          },
          {
            api: 'config',
            usageCount: 100
          },
          {
            api: 'io',
            usageCount: 528
          },
          {
            api: 'springFramework',
            usageCount: 0
          },
          {
            api: 'process',
            usageCount: 6
          },
          {
            api: '***TOTAL',
            usageCount: 4910
          },
          {
            api: 'cache',
            usageCount: 200
          },
          {
            api: 'websphere',
            usageCount: 600
          }
        ]
      }
    ]
  };

  beforeEach(waitForAsync(() => {
    mockRunDataService = jasmine.createSpyObj('RundataService', ['getApiByAppUsage']);
    TestBed.configureTestingModule({
      declarations: [ ApiByAppComponent, SearchFilterPipe],
      imports: [RouterTestingModule.withRoutes([
      ]), ToastrModule.forRoot()
      ],
      providers: [
        { provide: RundataService, useValue: mockRunDataService },
        {provide: ToastrService, useClass: ToastrService},
        {provide: ActivatedRoute,
          useValue: { paramMap: of(convertToParamMap({id: 1})) }}
      ]
    })
      .compileComponents();
  }));
  beforeEach(async () => {
    fixture = TestBed.createComponent(ApiByAppComponent);
    component = fixture.componentInstance;
    mockRunDataService.getApiByAppUsage.withArgs(1).and.returnValue(of(mockData));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('tagFilter by valid partial app name isTruthy', () => {
    fixture.detectChanges();
    const returnVal = component.appNameFilter.accepts(component.gridData[0] as [], 'cloud-suitability');
    expect(returnVal).toBeTrue();
  });

  it('tagFilter by invalid partial app name isFalsey', () => {
    fixture.detectChanges();
    const returnVal = component.appNameFilter.accepts(component.gridData[0] as [], 'csa');
    expect(returnVal).toBeFalse();
  });

  it('when component initializes then getApiByAppUsage called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockRunDataService.getApiByAppUsage).toHaveBeenCalledTimes(1);
  }));

});
