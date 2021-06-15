import { TestBed } from '@angular/core/testing';

import { RundataService } from './rundata.service';
import {HttpClientTestingModule} from '@angular/common/http/testing';

describe('RundataService', () => {
  let service: RundataService;
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [RundataService]
  }));
  beforeEach(() => {
    service = TestBed.inject(RundataService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
