import { TestBed } from '@angular/core/testing';

import { ExecutiveSummaryService } from './executivesummary.service';

describe('ExecutiveSummaryService', () => {
  let service: ExecutiveSummaryService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ExecutiveSummaryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
