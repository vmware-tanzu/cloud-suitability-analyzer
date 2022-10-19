import {ComponentFixture, fakeAsync, TestBed, waitForAsync} from '@angular/core/testing';

import { AnnotationsComponent } from './annotations.component';
import {SearchFilterPipe} from '../../../pipes/search-filter.pipe';
import {RouterTestingModule} from '@angular/router/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {RundataService} from '../../../services/rundata.service';
import {ActivatedRoute, convertToParamMap} from '@angular/router';
import {of} from 'rxjs';

describe('AnnotationsComponent', () => {
  let component: AnnotationsComponent;
  let fixture: ComponentFixture<AnnotationsComponent>;
  let mockRunDataService: jasmine.SpyObj<RundataService>;

  beforeEach(waitForAsync(() => {
    mockRunDataService = jasmine.createSpyObj('RundataService', ['getAnnotationData']);
    TestBed.configureTestingModule({
      declarations: [ AnnotationsComponent, SearchFilterPipe],
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
    fixture = TestBed.createComponent(AnnotationsComponent);
    component = fixture.componentInstance;
    mockRunDataService.getAnnotationData.withArgs(1).and.returnValue(of([]));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('when component initializes then getAnnotationData called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockRunDataService.getAnnotationData).toHaveBeenCalledTimes(1);
  }));
});
