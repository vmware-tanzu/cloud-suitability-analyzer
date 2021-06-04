import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RuleMetricsComponent } from './rulemetrics.component';

describe('RuleMetricsComponent', () => {
  let component: RuleMetricsComponent;
  let fixture: ComponentFixture<RuleMetricsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RuleMetricsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RuleMetricsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
