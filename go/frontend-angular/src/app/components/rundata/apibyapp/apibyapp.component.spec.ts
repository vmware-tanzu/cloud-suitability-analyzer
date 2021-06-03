import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiByAppComponent } from './apibyapp.component';

describe('ApibyappComponent', () => {
  let component: ApiByAppComponent;
  let fixture: ComponentFixture<ApiByAppComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApiByAppComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiByAppComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
