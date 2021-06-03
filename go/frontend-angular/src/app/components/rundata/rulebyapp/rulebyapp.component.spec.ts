import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RuleByAppComponent } from './rulebyapp.component';

describe('RulebyappComponent', () => {
  let component: RuleByAppComponent;
  let fixture: ComponentFixture<RuleByAppComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RuleByAppComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RuleByAppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
