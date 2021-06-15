import {ComponentFixture, TestBed, waitForAsync} from '@angular/core/testing';

import { FindingDetailsComponent } from './finding-details.component';
import {RouterTestingModule} from '@angular/router/testing';
import {RundataService} from '../../../../services/rundata.service';
import {Findings} from '../../../../model/findings';

describe('FindingDetailsComponent', () => {
  let component: FindingDetailsComponent;
  let fixture: ComponentFixture<FindingDetailsComponent>;
  let mockRunDataService: jasmine.SpyObj<RundataService>;

  beforeEach(waitForAsync(() => {
    mockRunDataService = jasmine.createSpyObj('RundataService', ['getFinding']);
    TestBed.configureTestingModule({
      declarations: [FindingDetailsComponent],
      imports: [RouterTestingModule.withRoutes([
      ])
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FindingDetailsComponent);
    component = fixture.componentInstance;
    component.selectedFinding = {} as Findings;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
