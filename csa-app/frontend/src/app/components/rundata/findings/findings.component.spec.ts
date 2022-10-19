import {ComponentFixture, fakeAsync, flush, TestBed, waitForAsync} from '@angular/core/testing';

import {SearchFilterPipe} from '../../../pipes/search-filter.pipe';
import {RouterTestingModule} from '@angular/router/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {RundataService} from '../../../services/rundata.service';
import {ActivatedRoute, convertToParamMap} from '@angular/router';
import {of} from 'rxjs';
import {FindingsComponent} from './findings.component';
import {Findings} from '../../../model/findings';

describe('FindingsComponent', () => {
  let component: FindingsComponent;
  let fixture: ComponentFixture<FindingsComponent>;
  let mockRunDataService: jasmine.SpyObj<RundataService>;
  const mockData = [
    {
      id: 132,
      run: 1,
      filename: '2.f6873732.chunk.css',
      fqn: '/cloud-suitability-analyzer/go/frontend-react/build/static/css/2.f6873732.chunk.css',
      ext: '.css',
      line: 0,
      rule: '',
      pattern: 'Analyzed File',
      value: 'File read and analyzed fully! [0] findings were documented! Rules Used: docker-dockerFile,docker-non-root-user,docker-sudo,java-system-config',
      advice: '',
      level: '',
      effort: 0,
      readiness: 0,
      category: 'File Finding',
      criticality: '',
      application: 'cloud-suitability-analyzer',
      tags: [
        'info',
        'ff'
      ],
      recipes: null
    }];

  beforeEach(waitForAsync(() => {
    mockRunDataService = jasmine.createSpyObj('RundataService', ['getRunFindings']);
    TestBed.configureTestingModule({
      declarations: [ FindingsComponent, SearchFilterPipe],
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
    fixture = TestBed.createComponent(FindingsComponent);
    component = fixture.componentInstance;
    mockRunDataService.getRunFindings.withArgs(1).and.returnValue(of(mockData as Findings[]));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('when component initializes then getRunFindings called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockRunDataService.getRunFindings).toHaveBeenCalledTimes(1);
    flush();
  }));
});
