import { TestBed } from '@angular/core/testing';

import { ApplicationsummaryService } from './applicationsummary.service';

describe('ApplicationsummaryService', () => {
  let service: ApplicationsummaryService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ApplicationsummaryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
