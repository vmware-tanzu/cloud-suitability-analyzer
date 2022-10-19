import {ComponentFixture, fakeAsync, TestBed, waitForAsync} from '@angular/core/testing';

import {SearchFilterPipe} from '../../../pipes/search-filter.pipe';
import {RouterTestingModule} from '@angular/router/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {RundataService} from '../../../services/rundata.service';
import {ActivatedRoute, convertToParamMap} from '@angular/router';
import {of} from 'rxjs';
import {ThirdPartyComponent} from './thirdparty.component';

describe('ThirdPartyComponent', () => {
  let component: ThirdPartyComponent;
  let fixture: ComponentFixture<ThirdPartyComponent>;
  let mockRunDataService: jasmine.SpyObj<RundataService>;

  beforeEach(waitForAsync(() => {
    mockRunDataService = jasmine.createSpyObj('RundataService', ['getThirdPartyData']);
    TestBed.configureTestingModule({
      declarations: [ ThirdPartyComponent, SearchFilterPipe],
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
    fixture = TestBed.createComponent(ThirdPartyComponent);
    component = fixture.componentInstance;
    mockRunDataService.getThirdPartyData.withArgs(1).and.returnValue(of([]));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('when component initializes then getThirdPartyData called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockRunDataService.getThirdPartyData).toHaveBeenCalledTimes(1);
  }));
});
