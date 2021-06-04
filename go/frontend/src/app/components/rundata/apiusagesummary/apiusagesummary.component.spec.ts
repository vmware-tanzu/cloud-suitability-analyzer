import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiUsageSummaryComponent } from './apiusagesummary.component';

describe('ApiUsageSummaryComponent', () => {
  let component: ApiUsageSummaryComponent;
  let fixture: ComponentFixture<ApiUsageSummaryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApiUsageSummaryComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiUsageSummaryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
