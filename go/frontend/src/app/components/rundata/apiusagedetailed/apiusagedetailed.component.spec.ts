import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiUsageDetailedComponent } from './apiusagedetailed.component';

describe('ApiusagedetailedComponent', () => {
  let component: ApiUsageDetailedComponent;
  let fixture: ComponentFixture<ApiUsageDetailedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApiUsageDetailedComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiUsageDetailedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
