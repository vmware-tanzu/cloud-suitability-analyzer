import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SourceCodeComponent } from './sourcecode.component';

describe('SourcecodeComponent', () => {
  let component: SourceCodeComponent;
  let fixture: ComponentFixture<SourceCodeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SourceCodeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SourceCodeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
