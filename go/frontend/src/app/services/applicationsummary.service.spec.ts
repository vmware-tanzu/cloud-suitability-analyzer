import { TestBed } from '@angular/core/testing';

import { ApplicationSummaryService } from './applicationsummary.service';
import {HttpClientTestingModule} from '@angular/common/http/testing';

describe('ApplicationsummaryService', () => {
  let service: ApplicationSummaryService;
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [ApplicationSummaryService]
  }));
  beforeEach(() => {
    service = TestBed.inject(ApplicationSummaryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
