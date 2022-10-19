import {ComponentFixture, fakeAsync, flush, TestBed, waitForAsync} from '@angular/core/testing';

import {SearchFilterPipe} from '../../../pipes/search-filter.pipe';
import {RouterTestingModule} from '@angular/router/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {RundataService} from '../../../services/rundata.service';
import {ActivatedRoute, convertToParamMap} from '@angular/router';
import {of} from 'rxjs';
import {RuleMetricsComponent} from './rulemetrics.component';
import {RuleMetric} from '../../../model/rulemetric';

describe('RuleMetricsComponent', () => {
  let component: RuleMetricsComponent;
  let fixture: ComponentFixture<RuleMetricsComponent>;
  let mockRunDataService: jasmine.SpyObj<RundataService>;

  beforeEach(waitForAsync(() => {
    mockRunDataService = jasmine.createSpyObj('RundataService', ['getRuleMetrics']);
    TestBed.configureTestingModule({
      declarations: [ RuleMetricsComponent, SearchFilterPipe],
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
    fixture = TestBed.createComponent(RuleMetricsComponent);
    component = fixture.componentInstance;
    mockRunDataService.getRuleMetrics.withArgs(1).and.returnValue(of({metrics: [{rule: 'java-system-config'}]} as RuleMetric));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('when component initializes then getRuleMetrics called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockRunDataService.getRuleMetrics).toHaveBeenCalledTimes(1);
    flush();
  }));
});
