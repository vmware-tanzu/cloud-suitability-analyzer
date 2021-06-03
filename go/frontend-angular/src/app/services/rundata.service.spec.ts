import { TestBed } from '@angular/core/testing';

import { RundataService } from './rundata.service';

describe('RundataService', () => {
  let service: RundataService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RundataService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
