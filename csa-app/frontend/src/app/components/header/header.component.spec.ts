import {ComponentFixture, fakeAsync, flush, TestBed, waitForAsync} from '@angular/core/testing';

import { HeaderComponent } from './header.component';
import {HttpClientTestingModule} from '@angular/common/http/testing';
import {ToastrModule, ToastrService} from 'ngx-toastr';
import {AnalyzerRunService} from '../../services/analyzerrun.service';
import {of} from 'rxjs';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';

describe('HeaderComponent', () => {
  let component: HeaderComponent;
  let fixture: ComponentFixture<HeaderComponent>;
  let mockAnalyzerRunService: jasmine.SpyObj<AnalyzerRunService>;

  beforeEach(waitForAsync(() => {
    mockAnalyzerRunService = jasmine.createSpyObj('AnalyzerRunService', ['getForgeVersion']);
    TestBed.configureTestingModule({
      declarations: [ HeaderComponent ],
      imports: [HttpClientTestingModule, ToastrModule.forRoot(),
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
    fixture = TestBed.createComponent(HeaderComponent);
    component = fixture.componentInstance;
    mockAnalyzerRunService.getForgeVersion.and.returnValue(of('v3.2.5'));
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('when component initializes then getForgeVersion called 1 time', fakeAsync(() => {
    fixture.detectChanges();
    expect(mockAnalyzerRunService.getForgeVersion).toHaveBeenCalledTimes(1);
    flush();
  }));
});
