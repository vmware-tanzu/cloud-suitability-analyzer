import { TestBed } from '@angular/core/testing';

import { ExecutiveSummaryService } from './executivesummary.service';
import {HttpClientTestingModule} from '@angular/common/http/testing';

describe('ExecutiveSummaryService', () => {
  let service: ExecutiveSummaryService;
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [ExecutiveSummaryService]
  }));
  beforeEach(() => {
    service = TestBed.inject(ExecutiveSummaryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
