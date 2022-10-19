import { TestBed } from '@angular/core/testing';
import { AnalyzerRunService } from './analyzerrun.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('AnalyzerRunService', () => {
  let service: AnalyzerRunService;
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [AnalyzerRunService]
  }));
  beforeEach(() => {
    service = TestBed.inject(AnalyzerRunService);
  });
  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});


