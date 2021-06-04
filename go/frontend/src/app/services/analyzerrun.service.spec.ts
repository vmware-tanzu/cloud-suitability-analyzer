import { TestBed, getTestBed } from '@angular/core/testing';
import { AnalyzerRunService } from './analyzerrun.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

describe('AnalyzerRunService', () => {
  let service: AnalyzerRunService;
  let backendMock: HttpTestingController;
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AnalyzerRunService);
    backendMock = TestBed.inject(HttpTestingController);
  });
  afterEach(() => {
    // After every test, assert that there are no more pending requests.
    backendMock.verify();
  });
  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});


