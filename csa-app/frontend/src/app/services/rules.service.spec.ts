import { TestBed } from '@angular/core/testing';

import { RulesService } from './rules.service';
import {HttpClientTestingModule} from '@angular/common/http/testing';
describe('RulesService', () => {
  let service: RulesService;
  beforeEach(() => TestBed.configureTestingModule({
    imports: [HttpClientTestingModule],
    providers: [RulesService]
  }));
  beforeEach(() => {
    service = TestBed.inject(RulesService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});


