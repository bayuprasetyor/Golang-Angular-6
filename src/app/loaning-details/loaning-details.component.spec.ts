import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { LoaningDetailsComponent } from './loaning-details.component';

describe('LoaningDetailsComponent', () => {
  let component: LoaningDetailsComponent;
  let fixture: ComponentFixture<LoaningDetailsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LoaningDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LoaningDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
