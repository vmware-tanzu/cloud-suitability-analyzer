import {ComponentFixture, fakeAsync, TestBed, waitForAsync} from '@angular/core/testing';

import {RundataService} from '../../../services/rundata.service';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {RouterTestingModule} from '@angular/router/testing';
import {of} from 'rxjs';
import {SearchFilterPipe} from '../../../pipes/search-filter.pipe';
import {ActivatedRoute, convertToParamMap} from '@angular/router';
import {RuleByAppComponent} from './rulebyapp.component';

describe('ApibyappComponent', () => {
  let component: RuleByAppComponent;
  let fixture: ComponentFixture<RuleByAppComponent>;
  let mockRunDataService: jasmine.SpyObj<RundataService>;
  const mockData = {
    cols: [
      'App',
      'SNAP-build-Ant-Maven',
      'bootCDI',
      'csa-internal',
      'docker-dockerFile',
      'docker-non-root-user',
      'docker-sudo',
      'java-3rdPartyImports',
      'java-MBeans',
      'java-ehcache',
      'java-fileIO',
      'java-handles-term',
      'java-hardIP',
      'java-jboss',
      'java-jvm-runtimeConfigProps',
      'java-logging-import',
      'java-metrics',
      'java-netflix-healthcheck',
      'java-processexit',
      'java-security',
      'java-servlet',
      'java-springboot-annotations',
      'java-springframework',
      'java-system-config',
      'java-ws2liberty-import',
      'weblogic-cluster-config',
      '***TOTAL'
    ],
    data: [
      {
        application: 'cloud-suitability-analyzer',
        ruleusage: [
          {
            rule: 'docker-sudo',
            count: 4
          },
          {
            rule: 'docker-dockerFile',
            count: 15
          },
          {
            rule: 'weblogic-cluster-config',
            count: 2
          },
          {
            rule: 'java-ws2liberty-import',
            count: 180
          },
          {
            rule: 'csa-internal',
            count: 0
          },
          {
            rule: 'java-netflix-healthcheck',
            count: 2
          },
          {
            rule: 'SNAP-build-Ant-Maven',
            count: 0
          },
          {
            rule: 'java-jvm-runtimeConfigProps',
            count: 100
          },
          {
            rule: '***TOTAL',
            count: 5320
          },
          {
            rule: 'java-metrics',
            count: 22
          },
          {
            rule: 'java-system-config',
            count: 321
          },
          {
            rule: 'java-ehcache',
            count: 200
          },
          {
            rule: 'bootCDI',
            count: 24
          },
          {
            rule: 'java-springboot-annotations',
            count: -200
          },
          {
            rule: 'java-springframework',
            count: 0
          },
          {
            rule: 'docker-non-root-user',
            count: 6
          },
          {
            rule: 'java-security',
            count: 40
          },
          {
            rule: 'java-logging-import',
            count: 66
          },
          {
            rule: 'java-fileIO',
            count: 528
          },
          {
            rule: 'java-servlet',
            count: 0
          },
          {
            rule: 'java-processexit',
            count: 6
          },
          {
            rule: 'java-MBeans',
            count: 600
          },
          {
            rule: 'java-hardIP',
            count: 2
          },
          {
            rule: 'java-jboss',
            count: 3400
          },
          {
            rule: 'java-3rdPartyImports',
            count: 0
          },
          {
            rule: 'java-handles-term',
            count: 2
          }
        ]
      },
      {
        application: '***TOTAL',
        ruleusage: [
          {
            rule: 'java-logging-import',
            count: 66
          },
          {
            rule: 'java-fileIO',
            count: 528
          },
          {
            rule: 'java-netflix-healthcheck',
            count: 2
          },
          {
            rule: 'java-hardIP',
            count: 2
          },
          {
            rule: 'weblogic-cluster-config',
            count: 2
          },
          {
            rule: 'java-metrics',
            count: 22
          },
          {
            rule: 'docker-non-root-user',
            count: 6
          },
          {
            rule: 'csa-internal',
            count: 0
          },
          {
            rule: 'java-springboot-annotations',
            count: -200
          },
          {
            rule: '***TOTAL',
            count: 5320
          },
          {
            rule: 'java-jboss',
            count: 3400
          },
          {
            rule: 'java-ws2liberty-import',
            count: 180
          },
          {
            rule: 'java-handles-term',
            count: 2
          },
          {
            rule: 'java-security',
            count: 40
          },
          {
            rule: 'java-3rdPartyImports',
            count: 0
          },
          {
            rule: 'java-servlet',
            count: 0
          },
          {
            rule: 'SNAP-build-Ant-Maven',
            count: 0
          },
          {
            rule: 'bootCDI',
            count: 24
          },
          {
            rule: 'docker-sudo',
            count: 4
          },
          {
            rule: 'java-jvm-runtimeConfigProps',
            count: 100
          },
          {
            rule: 'java-MBeans',
            count: 600
          },
          {
            rule: 'java-ehcache',
            count: 200
          },
          {
            rule: 'java-processexit',
            count: 6
          },
          {
            rule: 'docker-dockerFile',
            count: 15
          },
          {
            rule: 'java-springframework',
            count: 0
          },
          {
            rule: 'java-system-config',
            count: 321
          }
        ]
      }
    ]
  };

  beforeEach(waitForAsync(() => {
    mockRunDataService = jasmine.createSpyObj('RundataService', ['getRuleByAppUsage']);
    TestBed.configureTestingModule({
      declarations: [ RuleByAppComponent, SearchFilterPipe],
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
    fixture = TestBed.createComponent(RuleByAppComponent);
    component = fixture.componentInstance;
    mockRunDataService.getRuleByAppUsage.withArgs(1).and.returnValue(of(mockData));
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

  it('when component initializes then fetchRules called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockRunDataService.getRuleByAppUsage).toHaveBeenCalledTimes(1);
  }));

});
