import { TestBed } from '@angular/core/testing';

import { RulesService } from './rules.service';
import {of} from 'rxjs';
import * as mockRulesDataJson from '../../../public/mock/rules.json';
import {CsaRules} from '../model/rules';
describe('RulesService', () => {
  let mockRulesService: jasmine.SpyObj<RulesService>;
  mockRulesService = jasmine.createSpyObj('RulesService', ['getRules']);

  /*beforeEach(() => {
    TestBed.configureTestingModule({});
    const csarules: CsaRules = mockRulesDataJson;
    mockRulesService.getRules.and.returnValue(of[csarules]);
  });

  it('should be created', () => {
    expect(mockRulesService).toBeTruthy();
  });
  it('We can check if the consumer called a method on the class instance', () => {
    const mockRules = mockRulesDataJson;
    mockRulesService.getRules();
    expect(mockRulesService.getRules).toHaveBeenCalled();
  });*/
});


