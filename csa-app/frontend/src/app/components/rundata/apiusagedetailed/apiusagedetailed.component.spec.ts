import {ComponentFixture, fakeAsync, TestBed, waitForAsync} from '@angular/core/testing';

import {SearchFilterPipe} from '../../../pipes/search-filter.pipe';
import {RouterTestingModule} from '@angular/router/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {RundataService} from '../../../services/rundata.service';
import {ActivatedRoute, convertToParamMap} from '@angular/router';
import {of} from 'rxjs';
import {ApiUsageDetailedComponent} from './apiusagedetailed.component';

describe('ApiUsageDetailedComponent', () => {
  let component: ApiUsageDetailedComponent;
  let fixture: ComponentFixture<ApiUsageDetailedComponent>;
  let mockRunDataService: jasmine.SpyObj<RundataService>;

  beforeEach(waitForAsync(() => {
    mockRunDataService = jasmine.createSpyObj('RundataService', ['getApiDetailedUsage']);
    TestBed.configureTestingModule({
      declarations: [ ApiUsageDetailedComponent, SearchFilterPipe],
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
    fixture = TestBed.createComponent(ApiUsageDetailedComponent);
    component = fixture.componentInstance;
    mockRunDataService.getApiDetailedUsage.withArgs(1).and.returnValue(of([]));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('when component initializes then getApiDetailedUsage called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockRunDataService.getApiDetailedUsage).toHaveBeenCalledTimes(1);
  }));
});
